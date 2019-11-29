import ssl
import socket
import webbrowser

host = "www.bilibili.com"
port = 443
addr = (host, port)

s = ssl.wrap_socket(socket.socket(socket.AF_INET, socket.SOCK_STREAM))
s.connect(addr)
s.send(b'GET / HTTP/1.1\r\n')
s.send(b'Host: www.bilibili.com\r\n')
s.send(b'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36\r\n')
s.send(b'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3\r\n')
s.send(b'Accept-Language: zh-CN,zh;q=0.9')
s.send(b'Connection: keep-alive\r\n\r\n')

dataBytes = []
while True:
    dataUnit = s.recv(65536)       #此处怎么设置接受数量，然后无误的格式转换是重点
    print(dataUnit)
    if dataUnit:
        dataBytes.append(dataUnit)
    else:
        break
s.close()
print("\n数据请求结束！")

dataBytesSet = b''
for i in dataBytes:
    dataBytesSet += i

res = dataBytesSet.decode("utf-8")
# resHeader, html = res.split("\r\n\r\n", 1)
#
# with open('responseHeader.txt', "w") as f:
#     f.write(resHeader)
#
# with open('bilibili.html', "w") as f:
#     f.write(html)
with open('bilibili.html', "w") as f:
    f.write(res)

print("保存成功！")
