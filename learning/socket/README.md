# 粘包和半包
- 粘包：多个请求的内容合并成一个包发送
- 半包：一个包中仅一个请求的部分内容
- 拆包：请求过大导致被拆成多个包发送

## 为什么会有粘包和半包
因为 TCP 是面向连接的传输协议，TCP 传输的数据是以流的形式，而流数据是没有明确的开始结尾边界，所以 TCP 也没办法判断哪一段流属于一个消息。
- 造成粘包的主要原因：
  - 发送方每次写入数据 < 套接字缓冲区大小 
  - 接收方读取套接字缓冲区数据不够及时
- 造成半包的主要原因 
  - 发送方每次写入数据 > 套接字（Socket）缓冲区大小 
  - 发送的数据大于协议的 MTU (Maximum Transmission Unit，最大传输单元)

## 如何处理
### fix length
每次发送固定大小数据，控制服务器端和客户端发送和接收字节的长度相同即可：比如客户端发送 1024 个字节，服务器接受 1024 个字节

问题：
- 虽然可以解决粘包的问题，但是如果发送的数据小于 1024 个字节，就会导致数据内存冗余和浪费
- 如果发送请求大于 1024 字节，会出现半包的问题，也就是数据接收的不完整

### delimiter based
按分隔符分隔——分隔符号需要很特殊，同时查找分隔符效率也是一个问题

应用场景：比如通过 multipart/form-data 上传文件，会有比较特殊的 boundary

### length field based frame decoder
先定义一个 Header + Body 格式，其中 Header 中包含开始标记及 Body 长度，Body 中直接读取即可

应用场景：比如普通的 HTTP 请求、比如 goim