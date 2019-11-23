#2组栅栏密码的加解密

def FenceEncryption(clearText):           #加密
    textList = list(clearText)
    firstList,secondList,cipherList = [],[],[]
    a, b = 0, 0

    if len(textList) % 2 == 0:
        for i in range(int(len(clearText) / 2)):
            firstList.append(textList.pop(0))
        secondList = textList
        # print(firstList,secondList)

        for j in range(len(clearText)):

            if j % 2 == 0:
                cipherList.append(firstList[a])
                a += 1
            else:
                cipherList.append(secondList[b])
                b += 1

        cipherText = "".join(cipherList)

        return cipherText

    else:
        print("输入了奇数个明文,不能进行默认的栅栏加密(栅栏数为2)")
#
# print(FenceEncryption(clearText))


def FenceDecryption(cipherText):             #解密
    textList = list(cipherText)
    oddList,evenList = [],[]
    if len(textList) % 2 == 0:
        for i in range(len(cipherText)):
            if i % 2 == 0:                      #oddList是偶数项,但是index是奇数
                oddList.append(cipherText[i])
            else:
                evenList.append(cipherText[i])

        return "".join(oddList) + "".join(evenList)

# print(FenceDecryption("abcdef"))

choice = int(input("请选择一个选项:\n1.两组栅栏加密\n2.两组栅栏解密:\n"))
if choice == 1:
    text = input("请输入需要加密的明文:")
    print(FenceEncryption(text))
else:
    text = input("请输入需要解密的密文:")
    print(FenceDecryption(text))