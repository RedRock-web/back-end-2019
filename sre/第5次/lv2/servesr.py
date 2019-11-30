import socket
import _thread

addr = ('127.0.0.1', 32532)

class ChatServer():
    def __init__(self):
        self.g_socketServer = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    def start(self):
        self._socketServer()
        _thread.start_new_thread(self._clientAccept, ())

    def _socketServer(self):
        self.g_socketServer.bind(addr)
        self.g_socketServer.listen(5)
        print("服务已经启动！")

    def _clientAccept(self):
        while True:
            client, clientAddr = self.g_socketServer.accept()
            _thread.start_new_thread(self._messageHandle, (client, clientAddr))

    def _messageHandle(slef,client, clientAddr):
        client.send("连接服务器成功！".encode("utf-8"))
        while True:
            message = client.recv(1024).decode(encoding="utf-8")
            # print(clientAddr)
            if message == 'q':
                slef._close()
            print(clientAddr[0] + str(clientAddr[1]) + ":" + message)
            client.send("ｏｋ".encode("utf-8"))

    def _close(self):
        self._socketServer.close()

if __name__ == "__main__":
    ChatServer().start()
    while 1:
        pass


