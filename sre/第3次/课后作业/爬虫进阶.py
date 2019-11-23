from urllib.request import urlopen
from bs4 import BeautifulSoup
import re
import os

def GetSingerInfo():
    url = "https://music.163.com/discover/artist/cat?id=1001&initial=-1"
    html = urlopen(url).read().decode("utf-8")
    soup = BeautifulSoup(html, features="lxml")
    # print(soup)
    pattern1 = 'href="\/artist\?id=.*?" title=".*?的音乐'

    singerList = re.findall(pattern1, str(soup))
    # print(type(singerList))
    pattern2 = '''\d+'''
    pattern3 = '''title=".*?的'''
    singerInfo = {}
    # print(len(singerList))

    for i in range(len(singerList)):
        # print(i)
        singer = str(singerList[i])
        singerId = str(re.findall(pattern2, singerList[i])).split("""'""")[1]
        singerName = str(str(re.findall(pattern3, singer)).split('''"''')[1]).split("的")[0]
        # print(singerId,singerName)
        singerInfo[singerId] = singerName

        fileName = "singerList.txt"
        with open(fileName, 'w') as f:
            f.write(str(singerInfo))
            f.close()

    return singerInfo

print(GetSingerInfo())

