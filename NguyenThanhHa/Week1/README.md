# Tìm hiểu chung về Backup, Backup agent
# Backup
- **Backup** dữ liệu là quá trình sao chép dữ liệu trong hệ thống CNTT sang vị trí khác để có thể khôi phục nếu dữ liệu gốc bị mất. Quá trình backup nhằm mục đích bảo toàn dữ liệu trong trường hợp thiết bị bị lỗi, bị tấn công mạng, thiên tai hoặc các trường hợp gây mất dữ liệu khác. Do đó, sao lưu dữ liệu là một phần quan trọng trong chiến lược bảo vệ dữ liệu của doanh nghiệp, thường bao gồm kế hoạch khôi phục thảm hoạ và duy trì hoạt động kinh doanh của nhiều tổ chức.

# Các phương pháp backup dữ liệu
- ## 1. Sao lưu đầy đủ (Full Backup)
Phương pháp sao lưu đầy đủ sẽ tạo ra một bản sao lưu đầy đủ của tất cả dữ liệu ở mỗi lần sao lưu , lưu trữ theo đúng dạng thức ban đầu của dữ liệu hoặc nén lại và mã hoá. Các bản sao đầy đủ tổng hợp sẽ tạo ra các bản sao lưu đầy đủ từ một bản sao lưu đầy đủ, kèm theo một hoặc nhiều bản sao lưu gia tăng. Hầu hết các tổ chức chỉ thực hiện sao lưu toàn bộ theo định kỳ vì quá trình này tốn nhiều thời gian. Tuy nhiên, sao lưu toàn bộ sẽ cung cấp khả năng phục hồi dữ liệu nhanh chóng khi được yêu cầu.
- ## 2. Sao lưu gia tăng (Incremental backup)
Phương pháp sao lưu gia tăng sao chép mọi dữ liệu đã được thay đổi kể từ lần sao lưu gần nhất, bất kể phương pháp sao lưu gần nhất là gì. Phương pháp sao lưu gia tăng đảo ngược sẽ bổ sung mọi dữ liệu đã được thay đổi vào bản sao lưu đầy đủ gần nhất. Các bản sao lưu như vậy có xu hướng chiếm ít dung lượng lưu trữ hơn so với các bản sao lưu khác biệt, vốn tăng dần theo thời gian và chúng cũng mất ít thời gian hơn để hoàn thành. Tuy nhiên, việc khôi phục dữ liệu sẽ mất nhiều thời gian hơn vì nó yêu cầu bản sao lưu toàn bộ ban đầu cộng với mỗi bản sao lưu gia tăng.
- ## 3. Sao lưu khác biệt (Differential Backup)
Phương pháp sao lưu khác biệt sẽ sao chép mọi dữ liệu kể từ lần sao lưu đầy đủ gần nhất, bất kể có bản sao lưu nào khác được tạo ra bằng bất kỳ phương pháp nào khác trong thời gian đó hay không. Thời gian sao lưu nhanh hơn so với sao lưu toàn bộ, nhưng việc khôi phục dữ liệu yêu cầu bản full backup ban đầu và bản sao lưu khác biệt mới nhất.
- ## 4. Sao lưu nhân bản (Mirror backup)
Phương pháp sao lưu nhân bản dược lưu trữ ở định dạng không nén, nhân bản mọi tập tin và cấu hình trong dữ liệu nguồn. Có thể truy cập vào dữ liệu này giống như dữ liệu gốc

# Cơ chế hoạt động của Backup
## 1. Xác định dữ liệu cần backup:
- `File-level backup`: Sao lưu từng tệp tin hoặc thư mục cụ thể
- `Block-level backup`: Sao lưu các khối dữ liệu, thường dùng cho các hệ thống lưu trữ lớn như cơ sở dữ liệu 
- `Image-level backup`: Sao lưu toàn bộ hình ảnh của hệ thống, bao gồm cả hệ điều hành và các ứng dụng.
## 2. Chọn phương pháp backup:
- `Full Backup`: Tốn thời gian và dung lượng lưu trữ nhưng đảm bảo đầy đủ dữ liệu.
- `Incremental Backup`: Tiết kiệm dung lượng và thời gian, nhưng việc khôi phục dữ liệu có thể phức tạp hơn.
- `Differential Backup`: Nhanh hơn full backup và đơn giản hơn incremental backup khi khôi phục dữ liệu.
## 3. Quá trình backup:
- `Thu thập dữ liệu`: Cài đặt Backup agent hoặc phần mềm backup lên máy chủ, máy trạm, máy ảo... cần backup
- `Cài đặt jobs backup`: tự động chạy các jobs theo 
- `Nén và mã hoá (nếu cần)`: Dữ liệu có thể được nén để tiết kiệm dung lượng lưu trữ và mã hoá để đảm bảo an toàn.
- `Chuyển dữ liệu`: Dữ liệu được chuyển đến đích lưu trữ, có thể là một máy chủ khác, ổ đĩa ngoài, dịch vụ lưu trữ đám mây(S3, Google Drive, One Drive,...) hoặc băng từ.
## 4. Lưu trữ dữ liệu:
- `Local Storage`: Dữ liệu được lưu trữ trên các thiết bị lưu trữ tại chỗ như ổ cứng, NAS (Network Attached Storage) hoặc SAN (Storage Area Network) 
- `Remote Storage`: Dữ liệu được lưu trữ tại một vị trí từ xa, có thể là một trung tâm dữ liệu khác hoặc dịch vụ đám mây.
- `Offsite Storage`: Dữ liệu được lưu trữ tại một địa điểm khác hoàn toàn, thường dùng để đảm bảo an toàn trong trường hợp thiên tai hoặc sự cố lớn tại vị trí chính.
## 5. Kiểm tra và xác nhận backup:
- `Theo dõi và kiểm tra các jobs backup hoạt động đúng và đầy đủ` 
- `Kiểm thử tính toàn vẹn của dữ liệu backup`: Đảm bảo rằng dữ liệu backup không bị lỗi hoặc mất mát.
- `Log và báo cáo`: Ghi lại quá trình backup và tạo báo cáo để có thể kiểm tra.
## 6. Khôi phục dữ liệu (Restore)
- `Chọn phiên bản backup`: Chọn phiên bản backup cần khôi phục.
- `Giải nén và giải mã`: Nếu dữ liệu đã được nén và mã hoá, quá trình giải nén và giải mã sẽ diễn ra
- `Chuyển dữ liệu về hệ thống nguồn`: Dữ liệu được khôi phục về vị trí gốc hoặc một vị trí mới theo yêu cầu.

# Agent-based vs Agentless Backup

## 1. Traditional Agent-based Backup (guest based backup)
Agent-based backup còn được gọi là sao lưu dựa trên máy khách. Agent trong backup là module phần mềm được cài đặt trên mọi máy chủ để thực hiện một số tác vụ nhất định.Agent-based backup phù hợp cho các sản phẩm yêu cầu người dùng cài đặt phiên bản lightweight của phần mềm trên mỗi máy mà họ muốn bảo vệ. Nếu agent được cài đặt trên máy ảo thì nó sẽ xem máy ảo như là một máy vật lý. Agent trong trường hợp này đang đọc dữ liệu từ đĩa và truyền dữ liệu đến máy chủ sao lưu. Agent software nằm ở lớp kernel level ở trong hệ thống do đó nó có thể phát hiện các thay đổi ở cấp độ block-level trên máy chủ.

![](agent-based.jpg)

Agent-based backups không yêu cầu quét toàn bộ hệ thống tệp để xác định các thay đổi cho các bản sao lưu điều này làm nó hiệu quả hơn so với agentless backups cho máy chủ. Phải có tài nguyên local computing resources cho agent-based backups để thực hiện quá trình backup dữ liệu và chuyển chúng đến vị trí sao lưu phù hợp. Do đó, quá trình backup có thể ảnh hưởng đến hiệu suất ứng dụng nếu máy chủ không có đủ sức mạnh tính toán cần thiết cho quá trình backups khi xét đến khối lượng công việc cần phải backup nhiều.

Ngoài ra, khi quản trị viên hệ thống làm việc trong môi trường bao gồm cả máy chủ vật lý và máy chủ ảo, agent-based backups thường được yêu cầu cho máy chủ vật lý. 

### Image-based Backup
Loại backups này sẽ chụp nhanh(`snapshot`) toàn bộ ổ đĩa và bộ nhớ của máy chủ. Không cần phải cài đặt lại hệ điều hành và khôi phục một bản vá của các tệp để sao chép hệ thống trước đó, điều này là cần thiết với các hệ thống sử dụng non-image-based backup. Ngay cả sau khi xảy ra lỗi hoàn toàn thì việc khôi phục toàn bộ hệ thống system image có thể được thực hiện trong vài phút và không có khả năng thiếu các tệp quan trọng, điều này có khả năng xảy ra trong non-image-based backup do chỉ hoạt động ở file level.  

### Non-image-based Backup
Non-image-based Backup hoạt động ở file level sử dụng hệ thống agent-based cho việc khôi phục các tệp bị mất, bị hỏng hoặc bị xoá. Loại backups này không thể khôi phục toàn bộ hệ thống. Tuy nhiên có thể khôi phục lại tệp rất chi tiết.

## Pros:
- Cả máy chủ vật lý và máy chủ ảo đều được bảo vệ theo cùng một phương pháp
- Rất tin cậy do chúng sở hữu khả năng kiểm soát đáng kể đối với hệ thống máy chủ. Vì các agents được đặt ở cấp độ kernel level nên chúng cung cấp quyền truy cập trực tiếp vào các thay đổi trong các sector đĩa. Do đó người dùng được cung cấp bản sao lưu nhanh hơn và đáng tin cậy hơn.
- Nhờ được tích hợp chặt chẽ với dịch vụ Microsoft's volume shadow copy nên các bản backups dựa trên agent-based có thể thiết lập các bản backups nhất quán với ứng dụng.  
- Thích hợp với Highly Transactional Virtual Machines: Các agents có lợi cho highly transactional virtual machines với cơ sở dữ liệu bao gồm các thực thể như SQL hoặc exchange. Vì volume shadow service có thể ngắt các transactions này ở một snapshot nên có khả năng xảy ra lỗi. Ngoài ra, agent-based backups dựa vào các tài nguyên tính toán của máy đang được sao lưu nên tốc độ xử lý của nhiều giai đoạn được cải thiện.
- Người sử hữu ứng dụng có thể quản lý backup và khôi phục lại Guest OS.
- Đây là cách duy nhất để bảo vệ máy ảo VMware Fault Tolerant và máy ảo với Physical Raw Disk Mapping RDMS.
## Cons:
- Sử dụng tài nguyên CPU, memory, I/O và tài nguyên mạng cao hơn đáng kể trên các máy chủ ảo khi chạy backups
- Cần cài đặt và quản lý agent trên mỗi máy ảo
- Chi phí có thể cao đối với các giải pháp cấp phép theo từng agent thay vì cấp phép theo từng hypervisor
- Có thể cần nhiều loại phương pháp sao lưu và khôi phục: VD: cần các chính sách sao lưu riêng cho các bản sao lưu tệp và thư mục, các bản sao lưu Microsoft Exchange, bare metal recovery,...
- Các chiến lược khôi phục disaster phức tạp
- Không có biện pháp bảo vệ cho các máy ảo ngoại tuyến và các máy ảo template
- Có thể xảy ra downtimes và vấn đề bảo trì: Người quản trị phải khởi động lại hệ thống để cài đặt agent nên có thể xảy ra downtime trong quá trình cài đặt và cần thời gian để active đặc biệt là trong các mạng lớn.

## 2. Agentless Backup (host-based backup)
Agentless backup còn gọi là sao lưu dựa trên máy chủ, đề cập đến giải pháp không yêu cầu phải cài đặt agent trên mỗi máy ảo. Tuy nhiên điều quan trọng là phần mềm có thể đưa agent vào máy khách mà ta không hề biết.
Giải pháp này tích hợp với VMware APIs for Data Protection (VADP) hoặc Microsoft VSS, tạo ra các bản snapshots nhanh, hiệu suất cao của các đĩa ảo gắn với các VMs. Phần mềm backup sẽ giao tiếp với VADP hoặc VSS và cho biết những gì nó muốn sao lưu. VADP và VSS thực hiện 1 số bước và lần lượt chuẩn bị dữ liệu để backup. Nhà cung cấp VSS/VADP sẽ snap ổ đĩa và cấp cho backup solution quyền truy cập vào snapshot bằng cách đưa tệp cho máy chủ backup. Sau đó backup solution sẽ sao lưu lại snapshot đó.




# Cloud Storage
### 1. Object storage
Object storage là kiến trúc lưu trữ dữ liệu cho các kho lưu trữ lớn, lưu trữ một lượng lớn và ngày càng tăng của các dữ liệu phi cấu trúc: ảnh, video, học máy (machine learning), dữ liệu cảm biến, file âm thanh và các loại nội dung web khác. Việc tìm ra cách để có thể mở rộng hiệu quả, không tốn kém là một thách thức. Objects lưu trữ dữ liệu theo định dạng của dữ liệu ban đầu và cho phép có thể tuỳ chỉnh metadata theo cách giúp dữ liệu dễ truy cập và phân tích hơn. Thay vì được tổ chức theo cấu trúc phân cấp tệp hoặc thư mục, objects được lưu trong các secure buckets cung cấp khả năng mở rộng gần như không giới hạn.

### 2. File Storage
Lưu trữ dựa trên tệp được sử dụng rộng rãi trong các ứng dụng và lưu trữ dữ liệu phân cấp định dạng tệp và thư mục. Các loại lưu trữ này thường được gọi là máy chủ lưu trữ được kết nối mạng( network-attached storage - NAS) với các giao thức cấp file level của Server Message Block (SMB) được sử dụng trong các phiên bản Windows và Network File System(NFS) trong Linux. 

#### Network-attached storage (NAS)
NAS là một loại thiết bị lưu trữ cung cấp cho các nút mạng cục bộ (LAN) lưu trữ chia sẻ dựa trên tệp thông qua kết nối Ethernet. Máy chủ NAS thường chứa nhiều ổ cứng, cung cấp dung lượng lưu trữ tập trung lớn cho các máy tính được kết nối để lưu dữ liệu. Thay vì mỗi máy tính chia sẻ các tệp riêng của mình, dữ liệu được chia sẻ sẽ được lưu trữ trên một máy chủ NAS duy nhất. Cung cấp 1 cách dễ dàng để nhiều người dùng truy cập cùng một dữ liệu, điều này quan trọng trong các tình huống mà người dùng đang cộng tác trong các dự án hoặc sử dụng cùng một tiêu chuẩn. Do bản chất tập trung của mình nên máy chủ NAS thường được sử dụng cho:
- Chia sẻ tệp
- Sao lưu/ phục hồi dữ liệu
- Network printing
- Chia sẻ tệp đa phương tiện
- Media server
#### Advantage of NAS
- `Convenient`: it provides consolidate space of storage within the network. That means it is easier to collaborate on the server and to the machine.
- `Reliable`: most NAS supports RAID 0, RAID 1, RAID 5 which makes data safer. When the data stored on one drive has been destroyed, it can be recovered from another drive.
- `Affordable`: NAS devices cost less than normal servers and have low energy consumption.
- `Easy`: Fast and easy installation/configuration and administration
#### Disavantages of NAS
- `Network dependent`: Since files are typically shared with NAS devices over the LAN(Local area network, also used for normal traffic), they can cause congestion or can be affected by other traffic on the LAN. Therefore, NAS is not suitable for data transfer intensive applications.
- `Minimal speed`: With low throughput and high latency, a NAS is not fast enough for high performance application: big database

### 3. Block Storage
Enterprise application like databases or enterprise resource planning (`ERP`) systems often require dedicated, low-latency storage for each host. This is analogous to direct-attached storage (`DAS`) or a storage area network (`SAN`)  

#### Storage Area Network (SAN)
SAN is a dedicated high-speed network or subnetwork that interconnects and presents shared pools of storage devices to multiple servers. Each server on the network can access hard drives in the SAN as if they were local disks directly attached to the server. When a host wants to access a storage device on the SAN, it sends out a block-based access request for the storage device.

SAN combines the flexibility and sharing capabilities of NAS with the much of the performance of Direct Attach Storage (DAS). However, it is far more complex and costly than NAS. A SAN consists of dedicated cabling: Fiber Channel(FC) or Ethernet based iSCSI, dedicated switches and storage hardware. It performs best when used with Fiber Channnel medium(optical fibers and a fiber channel switch) but it is very expensive, complex and difficult to manage. Ethernet-based iSCSI has reduced these challenges by encapsulating SCSI commands into IP packets that do not require an FC connection. It is particularly useful for small and midsize businesses that may not have the funds or expertise to support a Fiber Channel SAN. As SAN is a block level storage solution, it is best suited for high perfomance applications such as:
- Databases (MS SQL, MySQL, PostgreSQL,...)
- Media Libraries
- Backup Archives
- High Usage File Servers
- E-mail Servers
- Remote vaulting and mirroring
- Heterogeneous platform support
- Storage-level replication
- Storage-level backups

#### Advantages of SAN
- `Better disk utilization`: Rather than having several servers with various levels of hard drive utilization, a SAN allows to pool storage and dynamically allocate exactly what each server requires.
- `Higher performance`: SAN performance is not affected by Ethernet traffic or local disk throughput bottlenecks. Data transmitted to and from a SAN is on its own private network partitioned off from user traffic, backup traffic and other SAN traffic




