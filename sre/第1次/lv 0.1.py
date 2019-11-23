#输出斐波那契数列前n项

endNum = int(input("请输入一个数字\n将打印前n项斐波那契数列:"))
# print(endNum)

def F(n):                     #通项公式1--递归
    if n > 1:
        return F(n - 1) + F(n - 2)
    else:
        return n

# def F(n):                  #通项公式2--数学公式
#     fibon =  int((((1 + (5 ** 0.5)) / 2) ** n - ((1 - (5 ** 0.5)) / 2) ** n ) * (1/ 5 ** 0.5))
#     return fibon


# def F(n):                  #通项公式3--递增法
#     a ,b = 0,1
#     for i in range(n):
#         a ,b = b,a + b
#     return a


for i in range(endNum):
    print(F(i+1))






