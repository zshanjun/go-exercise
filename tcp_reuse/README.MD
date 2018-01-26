### Usage:
1. run server: go run server.go
2. run client: go run client.go
3. multiple telnet client: telnet localhost 8889

#### Intro:
一个简单的TCP连接多路复用
目的是学习如何实现一个在一个TCP连接里，发送多个来源的数据，并保证不混淆

代码分两部分，client端和server端。client端建立一条TCP连接到server，
然后用两个telnet连接client，各自发送消息给client。client端将消息通过
该TCP连接发送给server，server只是简单的把消息原样返回，client端将收到
的server端消息返回给发送此消息的telnet客户端。
两个telnet发送什么就收到什么，而没有收到别人的，即表明多路复用成功

运用的思想很简单， 每个和 client 端建立的连接，都分配一个 id， 
然后发送的每份数据，数据头部包含了此数据属于哪个连接，以及此数据的长度。
server 端只是原样返回信息，不需要过多说明。client 端收到回复之后，
从收到的数据里，读取出数据头部信息，即连接的 id 和数据的长度。
然后根据 id 从字典里面找到此连接，将数据发送回去。