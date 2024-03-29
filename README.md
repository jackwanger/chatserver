###基于TCP的聊天服务器框架

###数据包构造

（1）字节序：大端模式

（2）数据包组成：包长 + 类型 + 包体

 包长：4字节，uint32，整个数据包的长度

 类型：4字节，uint32

 包体：字节数组，[]byte

 包长和类型用明文传输，包体由结构体采用protobuf序列化后再进行AES加密得到。



###通信协议

####登陆

客户端登陆包：PK_ClientLogin

服务器回复是否登录成功：PK_ServerAcceptLogin

客户端在一定时间内没有登录成功，服务器会主动断开连接。

（1）客户端在连接服务器后需要在一定时间内发送登陆包。

（2）若客户端与服务器已经存在一条连接，则断开存在的连接，开始新的这个连接，
也就是只能登陆一个客户端。


####下线

客户端主动下线数据包：PK_ClientLogout

（1）主动下线，客户端向服务器发送PK_ClientLogout包，服务器在收到PK_ClientLogout包后断开连接。

（2）被动下线，一般是在网络掉线的情况下，这种情况依赖于服务器的超时机制。


####心跳

客户端心跳包：PK_ClientPing

客户端需要定时向服务器发送PK_ClientPing包，维持在线状态。

服务器设置的是读超时，所以在收到PK_ClientPing包会选择忽略。



####C2C聊天

客户端之间的聊天消息包：PK_C2CTextChat

全部由服务器负责转发消息，若对方不在线则保存为离线消息，待对方上线时发给它。



###特点

登陆超时检测，防止恶意连接

conn数据接收使用sync.Pool和ring buffer优化

离线消息采用异步定时更新数据库

自动邮件报告服务器状况

