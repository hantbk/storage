# Kế hoạch Benchmark, so sánh hiệu năng giữa CEPHADM và ROOK-CEPH trong môi trường nhỏ.

## A. Chuẩn bị môi trường:

- Quy hoạch của 2 cụm ceph được benchmark trong node:
  - Nằm cùng trên 1 máy ảo, 1 node.
  - 1 Monitor
  - 1 Manager
  - 3 Osd, nằm trong cùng 1 virtual disk với bộ nhớ 120GB ( Mỗi osd chứa 40 GB)

- Cấu hình của VM chạy thử nghiệm:
  - CPU:
  - RAM:
  - Network adapter:

## B. Kế hoạch benchmark:

### 1. Công cụ sử dụng:
(Lựa chọn 1 trong các công cụ sau):
- [FIO](https://github.com/axboe/fio): Một công cụ lâu đời, linh hoạt cho các thử nghiệm I/O, hỗ trợ nhiều loại workload như đọc, ghi, và các mô hình truy cập khác nhau, có thể được sử dụng cho cả ổ đĩa vật lý và ảo hóa.

- [CBT (Ceph Benchmarking Tool) ](https://github.com/ceph/cbt): Là công cụ được thiết kế đặc biệt cho việc kiểm tra hiệu suất của các cụm Ceph. Nó tích hợp nhiều module khác nhau như radosbench và rbdfio để thực hiện các bài kiểm tra hiệu suất. CBT cho phép thu thập và phân tích dữ liệu từ cụm Ceph một cách chi tiết và có hệ thống.

- [SIBENCH](https://github.com/SoftIron/sibench): Là một công cụ benchmarking mã nguồn mở nhỏ mới, được thiết kế để mô phỏng và đo lường hiệu suất của hệ thống lưu trữ. Giúp phân tích hiệu suất của Ceph trong các tình huống thực tế.

- RADOS Bench, RBD bench: Các công cụ kiểm tra hiệu suất tích hợp sẵn trong Ceph, dùng để đo hiệu suất của các dịch vụ rgw, rbd

- Ngoài ra, đối với việc đo hiệu suất của các bucket thông qua s3, ta sẽ sử dụng lệnh `time s3cmd` trong s3cmd để đo tốc độ tạo và xóa file lên trên bucket của các cụm ceph. 

### 2. Kịch bản benchmark:
- Tạo 1 pool mới lên trên 2 cụm cluster với:
```
ceph osd pool create testbench 100 100
```
#### 2.1 Benchmark đọc/ghi 1 object lên 1 pool mới:
- Input:
    - Thời gian chạy thử nghiệm.
    - Kích thước của file đọc/ghi.
    - Số lượng luồng sử dụng.
- Output:
    - Tổng số lần đọc/ghi.
    - Băng thông tối thiểu, tối đa, trung bình được sử dụng.
    - IOPS (Input Output Per Seconds - số thao tác đọc - ghi trong 1 giây) tối thiểu, tối đa, trung bình.

#### 2.2 Benchmark đọc/ghi lên 1 image tạo bởi dịch vụ rbd:
- Input:
    - Kích thước của image.
    - Thời gian chạy thử nghiệm.
    - Số lượng luồng sử dụng.
- Output:
    - Số thao tác hoàn thành, số thao tác hoàn thành mỗi giây.
    - Tổng số byte được đọc/ghi mỗi giây.

#### 2.3 Benchmark tốc độ đẩy file, lấy file, liệt kê file qua dịch vụ s3api sử dụng s3cmd:
- Input:
    - Kích thước của file.
    - Tên bucket đẩy file lên.
- Output:
    - Thời gian hoàn thành hoạt động.