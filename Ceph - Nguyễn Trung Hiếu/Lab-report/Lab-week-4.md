# Lab week 4, Benchmark giữa cephadm và rook-ceph.

# A. Chuẩn bị:

## Cài đặt go trên các vm:
```
wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
```
- Trong bash, cài đặt:
```
export PATH=$PATH:/usr/local/go/bin

source ~/.bashrc
```

## Cài đặt C, C++:
```
sudo apt update
sudo apt install build-essential
```

# B. Benchmark:

## 1. Sử dụng hsbench để kiểm tra hiệu suất của 1 cluster:
- Đầu tiên, cài đặt hsbench:
```
mkdir hsbench && cd hsbench
go mod init hsbench
go install github.com/markhpc/hsbench@latest
```
- Gán path của hsbench vào bin:
```
ls $HOME/go/bin/hsbench
export PATH=$PATH:$HOME/go/bin
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```
- Chạy hsbench để bắt đầu benchmark 1 cụm cluster:
```
hsbench -a <Access-key> -s <Secret-key> -u <Endpoint> -z <Object-size> -d <Duration> -t <Number-of-threads> -b <Number-of-buckets>
```

## 2. Sử dụng gosbench để kiểm tra hiệu xuất của dịch vụ object storage thông qua s3api:

- Cài đặt gosbench server lên 1 máy:
```
go install github.com/mulbc/gosbench/server@latest
```
- Viết 1 file .yaml config cho server:
```
---

s3_config:
  - access_key: <Access-key-1>
    secret_key: <Secret-key-1>
    region: <Region-1>
    endpoint: <Endpoint-for-service-rgw-in-cluster1>
    skipSSLverify: false

  - access_key: <Access-key-2>
    secret_key: <Secret-key-2>
    region: <Region-2>
    endpoint:  <Endpoint-for-service-rgw-in-cluster2>
    skipSSLverify: false
grafana_config:
  endpoint: http://<host-ip>:9090
  username: admin
  password: grafana

tests:
  - name: Test 1
    read_weight: 100
    existing_read_weight: 0
    write_weight: 100
    delete_weight: 100
    list_weight: 100
    objects:
      size_min: 20
      size_max: 100
      part_size: 0
      size_distribution: random
      unit: KB
      number_min: 100
      number_max: 2000
      number_distribution: constant
    buckets:
      number_min: 10
      number_max: 10
      number_distribution: constant
    bucket_prefix: 1255gosbench-
    object_prefix: obj
    stop_with_runtime:
    stop_with_ops: 10
    workers: 2
    parallel_clients: 2
    clean_after: True
...
```
- Tạo 1 server thông qua lệnh: 
```
server -c gosbench-config.yaml 
```
- Sau đó, trên cả 2 máy, tải gosbench worker về:
```
go install github.com/mulbc/gosbench/worker@latest
```
- Mở cổng 2000 trên 2 máy:
```
sudo iptables -A INPUT -p tcp --dport 2000 -j ACCEPT
```
- Chạy các lệnh sau trên 2 máy để tạo worker:
```
worker -s <host-server-ip>:2000  (-p 8878)
# Nếu port 8888 đã sử dụng, hoặc không muốn sử dụng port 8888, thêm cờ -p này.
```

## 3. Sử dụng fio để đánh giá hiệu xuất các dịch vụ i/o của object và images rbd:
- Cài đặt fio:
```
apt install fio
```
- Đối với cụm cephadm, các cấu trúc file `etc/ceph/` đã có sẵn `ceph.conf` và `keyring` nên ta không cần config gì, còn đối với cụm `ceph-rook`, ta cần phải export ceph.conf và keyring từ trong pods ceph-toolbox ra:
```
mkdir /etc/ceph/
kubectl exec -it -n rook-ceph rook-ceph-tools-767b99dbdd-t4tvv -- cat /etc/ceph/ceph.conf > /etc/ceph/ceph.conf 
kubectl exec -it -n rook-ceph <ceph-toolbox-pod-id> -- cat /etc/ceph/keyring > /etc/ceph/keyring
```
- Sau đó, viết các file kịch bản `.fio` để chạy thử:
```
[global]
ioengine=rados 
runtime=60 #Thời gian thử nghiệm: 60s
time_based #Chạy trong 1 khoảng thời gian xác định
rw=randrw #Kiểu mẫu i/o: đọc, ghi ngẫu nhiên
rwmixread=70 #Chỉ tỉ lệ đọc ghi: 70 đọc - 30 ghi
size=1G #Thực hiện trên 1 khối 1GB
numjobs=4 #Số lượng luồng sử dụng
blocksize=4k #Kích thước mỗi thao tác io
pool=testpool #Tên pool thực hiệnn
iodepth=4 #Số lượng thao tác i/o đồng thời gửi tới
loops=5 #Số lượng lặp lại các thao tác.
[randrw]
```
- Đối với dịch vụ rbd, ta sẽ viết như sau:
```
[global]
ioengine=rbd
runtime=60s
time_based
blocksize=4k
numjobs=1
group_reporting

[seq-read]
pool=rbd.pool
rbdname=image1 #Tên image
rw=read
size=2G

[seq-write]
rbdname=image1 #Tên image
pool=rbd.pool
rw=write
size=2G
```
- Sau đó, áp dụng thông qua lệnh:
```
fio mixed.fio --output=mixed_rook.txt
```
## 4. Benchmark hiệu năng của ceph cluster thông qua rados bench:
- Tạo 1 giống nhau pool ở cả 2 cụm ceph:
```
ceph osd pool create testpool 32 32 rep_hdd_osd 3 
```
- Sử dụng rados bench có sẵn trong cụm ceph để thực hiện benchmark các cụm thông qua lệnh:
```
rados bench -p testbench 20 write --no-cleanup
rados bench -p testbench 20 read --no-cleanup
```
- Chi tiết câu lệnh:
  - -p: cờ để xác định tên pool.
  - 20: Thời gian chạy test, tính theo giây.
  - write/read: Thao tác chạy
  - --no-cleanup: Cờ để ngắt việc remove object sau test, nhằm tăng tốc độ đọc ở các test sau.
- Sau khi kết thúc test, ta cần chạy lệnh:
```
rados -p testbench cleanup
```
# C. Kết quả: