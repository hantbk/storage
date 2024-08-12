# Lab week 2, cài đặt Radosgw và rdb lên trên cụm ceph

# A. RADOSGW:

## 1. Thêm label và service rgw cho host:
```
ceph orch host label add trunghieu-vdt4 rgw
ceph orch apply rgw rgw_public '--placement=label:rgw count-per-host:1' --port=8888
ceph orch apply rgw rgw_private '--placement=label:rgw count-per-host:1' --port=8889
```
Sau khi chạy xong, kiểm tra kết quả service thông qua ceph orch ls:
```
root@trunghieu-vdt4 ~# ceph orch ls
NAME                  PORTS   RUNNING  REFRESHED  AGE  PLACEMENT
crash                             1/1  3s ago     2d   *
mgr                               1/2  3s ago     2d   count:2
mon                               1/5  3s ago     2d   count:5
osd.osd_spec_default                3  3s ago     2d   trunghieu-vdt4
rgw.rgw_private       ?:8889      1/1  3s ago     9s   count-per-host:1;label:rgw
rgw.rgw_public        ?:8888      1/1  3s ago     32s  count-per-host:1;label:rgw
```
Kiểm tra pool rgw.root được tạo:
```
root@trunghieu-vdt4 ~ [22]# ceph osd lspools
1 .mgr
7 rep3_pool_rados
8 ec_pool_rados
9 .rgw.root
```

## 2. Tạo 1 rgw user:

- Trước hết, cần tạo các pool sau để radosgw có thể chạy:
  - .rgw.log # Chứa log của radosgw
  - .rgw.control #Chứa các hoạt động kiểm soát nội bộ
  - .rgw.meta #Chứa các metadata cho các đối tượng quản lý rgw
  - .rgw.buckets.index #Chứa các index của các bucket
  - .rgw.buckets.data #Chứa dữ liệu của rgw
  - .rgw.buckets.non-ec #Chứa dữ liệu ec.
- Tạo pool bằng lệnh sau:
```
ceph osd pool create default.rgw.log 32 32 ssd_01 3
ceph osd pool application enable default.rgw.log rgw
```
- Sau đó, tạo thử một user mới để có thể thêm thông tin vào các s3api:
```
radosgw-admin user create --uid="admin-hieu" --display-name="admin-hieu"
```

