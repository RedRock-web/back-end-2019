# #凯撒加密与解密

def CaesarEncryption(cleartext,key):            #凯撒加密
    key %= 26
    import string
    upChars = string.ascii_uppercase
    lowChars = string.ascii_lowercase
    aList = list(cleartext)
    for i in range(len(aList)):                 #找出文本字母,并分为大小写,依次按ASCII码进行以26为基数的移位
        if aList[i] in upChars:
            if (ord(aList[i]) + key) > 90:
                aList[i] = chr(ord(aList[i]) + key - 26)

            else:
                aList[i] = chr(ord(aList[i]) + key)
        if aList[i] in lowChars:
            if (ord(aList[i]) + key) > 122:
                aList[i] = chr(ord(aList[i]) + key - 26)

            else:
                aList[i] = chr(ord(aList[i]) + key)
    outputCiperTextText = "".join(aList)
    print("凯撒加密后文本为:"+outputCiperTextText)
    return outputCiperTextText


def CaesarDecryption(ciphertext,key):            #凯撒密匙解密
    key %= 26
    import string
    upChars = string.ascii_uppercase
    lowChars = string.ascii_lowercase
    aList = list(ciphertext)
    for i in range(len(aList)):                 #找出文本字母,并分为大小写,依次按ASCII码进行以26为基数的移位
        if aList[i] in upChars:
            if (ord(aList[i]) - key) < 65:
                aList[i] = chr(ord(aList[i]) - key + 26)
            else:
                aList[i] = chr(ord(aList[i]) - key)
        if aList[i] in lowChars:
            if (ord(aList[i]) - key) < 97:
                aList[i] = chr(ord(aList[i]) - key + 26)

            else:
                aList[i] = chr(ord(aList[i]) - key)
    outputClearText = "".join(aList)
    print("凯撒解密后文本为:"+outputClearText)
    return outputClearText

def ViolentCaesarDecryption(ciphertext):            #暴力凯撒解密
    import string
    upChars = string.ascii_uppercase
    lowChars = string.ascii_lowercase
    aList = list(ciphertext)
    saveAList = tuple(aList)             #若把列表aList直接赋值saveAList,下面对aList的改变也会导致sava列表改变
    for key in range(26):
        key %= 26
        for i in range(len(aList)):
            aList[i] = saveAList[i]
            if aList[i] in upChars:
                if (ord(aList[i]) - key) < 65:
                    aList[i] = chr(ord(aList[i]) - key + 26)
                else:
                    aList[i] = chr(ord(aList[i]) - key)
            if aList[i] in lowChars:
                if (ord(aList[i]) - key) < 97:
                    aList[i] = chr(ord(aList[i]) - key + 26)

                else:
                    aList[i] = chr(ord(aList[i]) - key)
        outputClearText = "".join(aList)
        print("凯撒解密后文本为:"+outputClearText)


choice = int(input("请选择模式:\n1凯撒加密\n2凯撒解密\n3凯撒暴力解密\n选择:"))

if choice == 1:
    text = input("请输入要加密的文本:")
    key = int(input("请输入密匙:"))
    CaesarEncryption(text, key)

elif choice == 2:
    text = input("请输入要解密的文本:")
    key = int(input("请输入密匙:"))
    CaesarDecryption(text, key)
else:
    text = input("请输入要解密的文本:")
    ViolentCaesarDecryption(text)