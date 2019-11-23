from urllib.request import urlopen
from bs4 import BeautifulSoup
import re

def GetSingerInformation ():
    url = "https://music.163.com/m/artist?id=1875"
    html = urlopen(url).read().decode("utf-8")
    soup = BeautifulSoup(html, features="lxml")
    singerIntroductionTemp = soup.find_all('meta')
    nameList = str(soup.title).split()

    nameStr = re.findall('>.*\（', nameList[0])

    singerName = str(str(str(nameStr).split(">")[1]).split("（")[0])
    # print(singerName)
    targe = str(singerIntroductionTemp[22])
    singerIntroduction = targe.split("\"")[1]
    # print(singerIntroduction)
    singer = {singerName: singerIntroduction}
    return singer

print(GetSingerInformation())

