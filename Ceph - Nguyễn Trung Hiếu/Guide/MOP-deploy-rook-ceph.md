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
- Cấu hình kubeadm:
```

```
- Cấu hình kube thông qua kubectl:
```
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

