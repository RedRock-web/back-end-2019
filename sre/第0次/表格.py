def grid(arg_1,arg_2):
    print("```")
    for i in range(3):
        if i == 0:
            print("+  -  -  -  - ",end = "")
            for zzz in range(arg_1 - 1):
                print("+  -  -  -  - ",end = "")
            print("+")


        if i == 1:
            for aa in range(arg_2):
                for z in range(4):
                    print("|             ",end = "")
                    for zz in range(arg_1 - 1):
                        print("|             ",end = "")


                    print("|")

                print("+  -  -  -  - ",end = "")
                for zzz in range(arg_1 - 1):
                    print("+  -  -  -  - ",end = "")
                print("+")




    print("```")
    return " "
print(grid(4,4))