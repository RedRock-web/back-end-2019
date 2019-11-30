import socket

addr = ("127.0.0.1", 32532)
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(addr)

data = s.recv(1024).decode("utf-8")
print(data)
while True:
    inputData = input()
    if inputData == 'q':
        break
    s.send(inputData.encode("utf-8"))

s.close()