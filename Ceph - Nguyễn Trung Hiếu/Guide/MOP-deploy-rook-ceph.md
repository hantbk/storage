## A. Cài đặt 1 cụm openstack
### 1. Cài đặt các công cụ k8s:
```
apt update
apt install -y apt-transport-https ca-certificates curl

curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.30/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.30/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list

apt update
```
### 2. Deploy k8s lên trên cụm:
- Tắt swap space:
```
sudo swapoff -a
```

- Tạo 1 cluster network cho các pod:
```
 kubeadm init --pod-network-cidr=171.254.94.63/24
```
- Cấu hình containerd trong /etc/containerd/config.toml
```
    [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    ...
    SystemdCgroup = true
    ...

    [plugins."io.containerd.grpc.v1.cri"]
    ...
    systemd_cgroup = true
    ...
```
- Cấu hình file docker daemon.json:
```
{
    "bridge": "none",
    "ip-forward": false,
    "iptables": false,
    "exec-opts": ["native.cgroupdriver=systemd"],
    "log-opts": {
        "max-file": "5",
        "max-size": "50m"
    }
}
```
- Cấu hình kubelet trong file kubelet/config.yaml
```
...
cgroupDriver: systemd
...
```
- Cho phép các port 6443 và 10250 thông qua firewall
```
sudo firewall-cmd --permanent --add-port=6443/tcp
sudo firewall-cmd --permanent --add-port=10250/tcp
sudo firewall-cmd --reload
```
- Xây thử 1 cụm k8s thông qua kubeadm:
```
kubeadm init --v=5
```
  
- Lỗi đang gặp phải:   
![alt text](../Picture/k8s-error-systemd.png)


- Kiểm tra kubelet:
```
● kubelet.service - kubelet: The Kubernetes Node Agent
     Loaded: loaded (/lib/systemd/system/kubelet.service; enabled; vendor preset: enabled)
    Drop-In: /usr/lib/systemd/system/kubelet.service.d
             └─10-kubeadm.conf
     Active: active (running) since Thu 2024-08-15 17:28:19 +07; 13min ago
       Docs: https://kubernetes.io/docs/
   Main PID: 3208922 (kubelet)
      Tasks: 16 (limit: 19047)
     Memory: 30.9M
        CPU: 20.787s
     CGroup: /system.slice/kubelet.service
             └─3208922 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --config=/var/lib/kubelet/config.yaml --container-runtime-endpoint=unix:///var/run/containerd/containerd.sock --pod-infra-container-image=registry.k8s.io/pause:3.9 --cgroup-driver=cgroupfs

Aug 15 17:42:10 trunghieu-vdt4 kubelet[3208922]: E0815 17:42:10.419954 3208922 eviction_manager.go:282] "Eviction manager: failed to get summary stats" err="failed to get node info: node \"trunghieu-vdt4\" not found"

Aug 15 17:42:12 trunghieu-vdt4 kubelet[3208922]: E0815 17:42:12.133241 3208922 controller.go:145] "Failed to ensure lease exists, will retry" err="Get \"https://171.254.94.63:6443/apis/coordination.k8s.io/v1/namespaces/kube-node-lease/leases/trunghieu-vdt4?timeout=10s\": dial tcp 171.254.94.63:6443: connect: connection refused" interval="7s"

Aug 15 17:42:12 trunghieu-vdt4 kubelet[3208922]: E0815 17:42:12.302663 3208922 certificate_manager.go:562] kubernetes.io/kube-apiserver-client-kubelet: Failed while requesting a signed certificate from the control plane: cannot create certificate signing request: Post "https://171.254.94.63:6443/apis/certificates.k8s.io/v1/certificatesigningrequests": dial tcp 171.254.94.63:6443: connect: connection refused

Aug 15 17:42:12 trunghieu-vdt4 kubelet[3208922]: E0815 17:42:12.317408 3208922 event.go:368] "Unable to write event (may retry after sleeping)" err="Patch \"https://171.254.94.63:6443/api/v1/namespaces/default/events/trunghieu-vdt4.17ebe01da4fe40f4\": dial tcp 171.254.94.63:6443: connect: connection refused" event="&Event{ObjectMeta:{trunghieu-vdt4.17ebe01da4fe40f4  default    0 
0001-01-01 00:00:00 +0000 UTC <nil> <nil> map[] map[] [] [] []},InvolvedObject:ObjectReference{Kind:Node,Namespace:,Name:trunghieu-vdt4,UID:trunghieu-vdt4,APIVersion:,ResourceVersion:,FieldPath:,},Reason:NodeHasNoDiskPressure,Message:Node trunghieu-vdt4 status is now: NodeHasNoDiskPressure,Source:EventSource{Component:kubelet,Host:trunghieu-vdt4,},FirstTimestamp:2024-08-15 17:28:20.326146292 +0700 +07 m=+0.575572186,LastTimestamp:2024-08-15 17:28:20.403582894 +0700 +07 m=+0.653008775,Count:2,Type:Normal,EventTime:0001-01-01 00:00:00 +0000 UTC,Series:n>

Aug 15 17:42:12 trunghieu-vdt4 kubelet[3208922]: E0815 17:42:12.317647 3208922 event.go:307] "Unable to write event (retry limit exceeded!)" event="&Event{ObjectMeta:{trunghieu-vdt4.17ebe01da99bd7ae  default    0 0001-01-01 00:00:00 +0000 UTC <nil> <nil> map[] map[] [] [] []},InvolvedObject:ObjectReference{Kind:Node,Namespace:,Name:trunghieu-vdt4,UID:trunghieu-vdt4,APIVersion:,ResourceVersion:,FieldPath:,},Reason:NodeHasNoDiskPressure,Message:Node trunghieu-vdt4 status is now: NodeHasNoDiskPressure,Source:EventSource{Component:kubelet,Host:trunghieu-vdt4,},FirstTimestamp:2024-08-15 17:28:20.403582894 +0700 +07 m=+0.653008775,LastTimestamp:2024-08-15 17:28:20.403582894 +0700 +07 m=+0.653008775,Count:1,Type:Normal,EventTime:0001-01-01 00:00:00 +0000 UTC,Series:nil,Action:,Related:nil,ReportingController:kubelet,ReportingInstance:trunghieu-vdt4,}"

Aug 15 17:42:12 trunghieu-vdt4 kubelet[3208922]: E0815 17:42:12.318398 3208922 event.go:368] "Unable to write event (may retry after sleeping)" err="Patch \"https://171.254.94.63:6443/api/v1/namespaces/default/events/trunghieu-vdt4.17ebe01da4fe57cf\": dial tcp 171.254.94.63:6443: connect: connection refused" event="&Event{ObjectMeta:{trunghieu-vdt4.17ebe01da4fe57cf  default    0 0001-01-01 00:00:00 +0000 UTC <nil> <nil> map[] map[] [] [] []},InvolvedObject:ObjectReference{Kind:Node,Namespace:,Name:trunghieu-vdt4,UID:trunghieu-vdt4,APIVersion:,ResourceVersion:,FieldPath:,},Reason:NodeHasSufficientPID,Message:Node trunghieu-vdt4 status is now: NodeHasSufficientPID,Source:EventSource{Component:kubelet,Host:trunghieu-vdt4,},FirstTimestamp:2024-08-15 17:28:20.326152143 +0700 +07 m=+0.575578047,LastTimestamp:2024-08-15 17:28:20.403591434 +0700 +07 m=+0.653017339,Count:2,Type:Normal,EventTime:0001-01-01 00:00:00 +0000 UTC,Series:nil>

Aug 15 17:42:12 trunghieu-vdt4 kubelet[3208922]: I0815 17:42:12.504075 3208922 kubelet_node_status.go:73] "Attempting to register node" node="trunghieu-vdt4"

Aug 15 17:42:12 trunghieu-vdt4 kubelet[3208922]: E0815 17:42:12.505211 3208922 kubelet_node_status.go:96] "Unable to register node with API server" err="Post \"https://171.254.94.63:6443/api/v1/nodes\": dial tcp 171.254.94.63:6443: connect: connection refused" node="trunghieu-vdt4"

Aug 15 17:42:13 trunghieu-vdt4 kubelet[3208922]: W0815 17:42:13.112591 3208922 reflector.go:547] k8s.io/client-go/informers/factory.go:160: failed to list *v1.Service: Get "https://171.254.94.63:6443/api/v1/services?limit=500&resourceVersion=0": dial tcp 171.254.94.63:6443: connect: connection refused

Aug 15 17:42:13 trunghieu-vdt4 kubelet[3208922]: E0815 17:42:13.112755 3208922 reflector.go:150] k8s.io/client-go/informers/factory.go:160: Failed to watch *v1.Service: failed to list *v1.Service: Get "https://171.254.94.63:6443/api/v1/services?limit=500&resourceVersion=0": dial tcp 171.254.94.63:6443: connect: connection refused

```