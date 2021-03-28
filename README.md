# grpcLearn
learn grpc


1: 下载protocolbuffers： https://github.com/protocolbuffers/protobuf/releases
 
2、 protoc-gen-go是go版本的 Protobuf 编译器插件，

       能访问网络的情况下，只需要运行 

       go get -u github.com/golang/protobuf/protoc-gen-go 便可以在$GOPATH/bin目录下发现这个工具。

3： 安装ide插件

4: 编写proto文件

5: 文件中必须要加上option go_package = "../services;HelloTest"; 这一行表示生成到services包下面，package是HelloTest

    在这中间我将protoc的运行文件exec放到了bin的目录下，然后include的google文件放到了gopath的src目录下，相当于将两个exec文件放在了一起，不知道这个有没有影响
    
6： 执行protoc --go_out=../services Prod.proto



上面生成的是普通文件

7： 下面生成grpc调用文件
protoc --go_out=plugins=grpc:../services prod.proto

