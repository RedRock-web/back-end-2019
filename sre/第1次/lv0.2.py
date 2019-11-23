# 请输入一个数,将返回对应的第n个丑数:
n = int(input("请输入一个数,将返回对应的第n个丑数:"))


# 思路1:丑数质因数只有2,3,5,则每个丑数都必然是2,3,5的积,则穷举存储去重排序即可
# def ReturnAUglyNum(n):
#     uglyList = []
#     for i in range(10):
#         for j in range(10):
#             for k in range(10):
#                 a = 2 ** i * 3 ** j * 5 ** k
#                 uglyList.append(a)
#
#     uglyList.sort()
#     return uglyList[n - 1]
#
#
# ReturnAUglyNum(n)

# 思路2:先定义一个丑数判断函数,然后穷举判断,之后再存储
def isUgly(n):
    m = n
    while (n % 2 == 0):
        n /= 2
    while (n % 3 == 0):
        n /= 3
    while (n % 5 == 0):
        n /= 5
        n = int(n)
    if n == 1:
        return m
    else:
        return 0


def nUglyNum(n):
    i = 1
    uglyList = []
    while 1:
        if isUgly(i):
            uglyList.append(isUgly(i))

        i += 1
        uglyList.sort()
        if len(uglyList) == n:
            return uglyList[n - 1]
            break


print(nUglyNum(n))
