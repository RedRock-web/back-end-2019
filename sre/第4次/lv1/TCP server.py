import socket

host = 'localhost'
port = 3131
addr = (host, port)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(addr)
s.listen(5)
sock, addr = s.accept()

while True:
    data = sock.recv(1024).decode("utf-8")
    if not data or data == "q":
        print("服务器已停止服务!")
        break;
    # print(data)
    sock.send(data.encode("utf-8"))

sock.close()
s.close()

