#TCP客户端
import socket
ipPort = ("127.0.0.1", 999)    #域名,端口

s = socket.socket()               #创建套接字对象
s.connect(ipPort)           #连接服务器

while True:
    inp = input("请输入要发送的信息:  ").strip()    #输入信息并将其进行转码

    s.sendall(inp.encode())       #将信息发送到服务器

    if inp == "exit":            #设置退出条件
        print("结束通信!")
        break

    serverReply = s.recv(1024).decode()     #接收sever发送的反馈,并将其转码
    print(serverReply)

s.close()               #关闭socket