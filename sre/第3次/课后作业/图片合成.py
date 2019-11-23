import os
import PIL.Image as Image

imagePath = "D:\\重庆邮电大学\\红岩网校\\运维\\新建文件夹\\"
imageFormat = ['.png', '.PNG']
imageSize = 256
imageRow = 9
imageColumn = 9
imageSavePath = 'C:\\Users\\图灵一号\\Desktop\\final.jpg'
imageNames = [name for name in os.listdir(imagePath) for item in imageFormat if os.path.splitext(name)[1] == item]

if len(imageNames) != imageRow * imageColumn:
    raise  ValueError("合成图片的参数和要求的数量不匹配!")
asdf.py
def imageCompose():
    toImage = Image.new('RGB', (imageColumn * imageSize, imageRow * imageSize))
    for y in range(1, imageRow + 1):
        for x in range(1, imageColumn + 1):
            fromImage = Image.open(imagePath + imageNames[imageColumn * (y - 1) + (x - 1)]).resize(
                (imageSize, imageSize), Image.ANTIALIAS)
            toImage.paste(fromImage, ((x - 1) * imageSize, (y - 1) * imageSize))
    return toImage.save(imageSavePath)
imageCompose()
