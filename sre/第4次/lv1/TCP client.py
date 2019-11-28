import socket

host = "localhost"
port = 3131
addr = (host, port)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(addr)

while True:
    data = input("请输入数据!/输入q退出\n")
    s.send(data.encode("utf-8"))
    data1 = s.recv(1024).decode("utf-8")
    print(data1)
    if data == 'q':
        break;

s.close()

