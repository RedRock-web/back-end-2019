import requests
import os

#百度有反爬虫,故用个UA模拟一下
aHeaders = {
"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.81 Safari/537.36",
}

url = "https://www.baidu.com/s?wd=*"
searchResult = requests.get(url, headers = aHeaders).text


fileName = "baiduSearchResult.txt"
with open(fileName, 'w', encoding="utf-8") as f:
    f.write(str(searchResult))
    f.close()

