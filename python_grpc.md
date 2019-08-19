《Python：GRPC接口编写之如何编写服务端与客户端》

python -m pip install --upgrade pip  
更新pip版本

下载安装包：grpc、grpcio_tools、protobuf
pip install protobuf
pip install grpcio
pip install grpcio_tools

编译proto文件
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. (文件名).proto

目前大多数分布式服务之间的调用是基于http 1.*协议实现的，HTTP/1.* 是纯文本数据传输。
而grpc框架的底层网络协议是HTTP2.0，在传输的过程中数据是二进制的，在目前的市面是还
有没有比较完善和方便的测试grpc框架的测试工具，像是测试http1.*的postman这种下载就能
使用的测试工具。

grpc-gateway:grpc转换为http协议对外提供服务(好像只支持go)


protoc --go_out=plugins=grpc:. (文件名).proto (go)


（window下）
查找所有的运行的端口
netstat -ano

查看端口占用对应的pid
netstat -aon|findstr "8081"

查看指定pid的进程
tasklist|findstr "9088"

强制结束9088所有进程
taskkill /T /F /PID 9088  