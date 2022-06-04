# mokuai8work

1.编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。
优雅启动
优雅终止
资源需求和 QoS 保证
探活
日常运维需求，日志等级
配置和代码分离

附件  deployment.yaml解决  优雅启动，优雅终止，资源需求和 QoS 保证，探活，日常运维需求，日志等级

附件  configmap.yaml解决 配置和代码分离

2.
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）：
Service
Ingress
可以考虑的细节
如何确保整个应用的高可用。
如何通过证书保证 httpServer 的通讯安全。

附件 service.yaml,ingress.yaml 解决service ingress  httpserver通讯安全使用openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=meimeng.com/O=meimeng" -addext "subjectAltName = DNS:meimeng.com"    kubectl create secret tls meimeng-tls --cert=./tls.crt --key=./tls.key
在ingress.yaml中进行引用这个secret  secretName: meimeng-tls

如何确保整个应用的高可用，1.冗余部署，deployment采取多副本形式，而负载均衡，在某一pod无法响应时流量可以通过负载均衡转载到其他Pod达到高可用效果



