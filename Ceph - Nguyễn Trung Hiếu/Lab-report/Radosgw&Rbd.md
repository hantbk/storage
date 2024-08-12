# RADOS Trong CEPH:

## I. RADOS
RADOS (Reliable Autonomic Distributed Object Store) là một thành phần cốt lõi của hệ thống lưu trữ Ceph. Đây là một hệ thống lưu trữ đối tượng phân tán, cung cấp nền tảng cho các dịch vụ lưu trữ khác của Ceph như RADOS Block Device (RBD), RADOS Gateway và Ceph File System.
![alt text](../Picture/rados.png)

RADOS cung cấp toàn bộ các tính năng của ceph, từ việc phân bố dữ liệu, tính tự quản lý, tự sửa lỗi, tính tin cậy cao,....

Rados chịu trách nhiệm cho việc quản lý dữ liệu, sao lưu dữ liệu, tạo các bản sao dữ liêu, luôn đảm bảo sao cho có nhiều hơn 1 bản copy của object trong cùng 1 cluster (Đối với cụm có replicated-count > 2).

Các phương thức truy xuất dữ liệu trong ceph (CephFS, RDB, RGW) đều hoạt động dựa trên RADOS.

## II. Librados:

Là một thư viện cho phép các ứng dụng có thể làm việc trực tiếp với RADOS. Librados cung cấp các API để tương tác trực tiếp với Ceph Monitor và Ceph OSD, nhằm lưu, xóa và quản lý dữ liệu hiệu quả.

![alt text](../Picture/librados.png)


![alt text](../Picture/radosgw-struc.png)

