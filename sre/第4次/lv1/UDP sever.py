#UDP会话--服务端
import socket
ipPort = ("127.0.0.1", 9999)   #域名
sk = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, 0)    #创建socket会话实例
sk.bind(ipPort)    #socket绑定域名

while True:
    data = sk.recv(1024).strip().decode()
    print(data)
    if data == "exit":
        print("客户端主动断开连接!")
        break

sk.close()    #关闭socket