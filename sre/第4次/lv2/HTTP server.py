import socket

def ServerResponse(html):            #定义responseheader,拿到responsebody(html),合并为response
    responseHeader = '''
HTTP/1.1 200 OK 
Context-Type: text/html
Server: Python-slp version 1.0 
Context-Length: '''

    f = open(html, "r")
    responseBody = f.read()
    response = "%s %d\n\n%s\n\n" % (responseHeader, len(responseBody), responseBody)
    # print(response)
    return bytes(response, encoding="utf-8")

def ServerReponseRequest(host, port):          #服务器响应请求
    #建立TCP连接,绑定客户域名端口并监听
    lisfd = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    lisfd.bind((host, port))
    lisfd.listen(2)

    while 1:
        clientRequestSocket, clientAdress = lisfd.accept() #服务器等待客户发送请求
        # print("connect by ", clientAdress)
        clientData = clientRequestSocket.recv(11111)
        if not clientData:
            break
        print(clientData)
        clientRequestSocket.sendall(ServerResponse('index.html'))   #服务器返回响应
        clientRequestSocket.close()

ServerReponseRequest("127.0.0.1", 3333)