# TCP服务器端
import socket

ipPort = ('127.0.0.1', 999)     #域名,端口

sk = socket.socket()      #创建一个TCP会话实例
sk.bind(ipPort)            #绑定域名,端口
sk.listen(5)            #设置监听

print("启动socket服务,等待客户端连接...")

conn, address = sk.accept()       #等待连接,客户端地址(ip,端口)/(address)
                                   #(conn,类型式元组)
                                   #此处自动阻塞
while True:
    clientData = conn.recv(1024).decode()         #接受信息,是bytes类型,并将其转换为strings

    if clientData == "exit":         #设置退出条件
        exit("通信结束")
    print("来自%s的客户端向你发来消息:%s" % (address, clientData))

    conn.sendall("服务器已经收到你的消息".encode())    #转码后反馈信息到客户端

conn.close()    #关闭会话
