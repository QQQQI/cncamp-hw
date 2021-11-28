## 课后练习3.2

### 构建本地镜像
- 编写Dockerfile将练习2.2编写的httpserver容器化
- 请思考有哪些最佳实践可以引入到Dockerfile中来
- 将镜像推送至docker官方镜像仓库
- 通过docker命令本地启动httpserver
- 通过nsenter进入容器查看IP配置

```
# CONTAINER_ID=$(podman run -d -p 8000:8000 docker.io/qiwanredhat/httpserver:latest)
Trying to pull docker.io/qiwanredhat/httpserver:latest...
Getting image source signatures
Copying blob a9757cebd49b done  
Copying blob aab139f4bc5a done  
Copying blob a0d0a0d46f8b done  
Copying blob b8b176561691 done  
Copying blob 31adcdaf11c8 done  
Copying blob 0d16ef8adc15 done  
Copying blob a472591a0f83 done  
Copying blob 09b385826532 done  
Copying config d147799dbc done  
Writing manifest to image destination
Storing signatures

# PID=$(podman inspect $CONTAINER_ID --format '{{.State.Pid}}')

# nsenter -t $PID -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0@if6: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether da:3c:3a:04:64:c2 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 10.88.0.2/16 brd 10.88.255.255 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::d83c:3aff:fe04:64c2/64 scope link 
       valid_lft forever preferred_lft forever

```