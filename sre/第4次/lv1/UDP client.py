#UDP客户端
import socket

ipPort = ("127.0.0.1", 9999)   #域名,端口

sk = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, 0)      #创建一个UDP客户端实例

while True:
    inp = input("发送的消息:").strip()    #输入信息并且转换为bytes类型
    sk.sendto(inp.encode(), ipPort)      #转码之后,将信息传输到服务器端
    if inp == "exit":              #如果用户输入exit,则退出,否则无限输入
        break
sk.close()                  #关闭UDP会话
