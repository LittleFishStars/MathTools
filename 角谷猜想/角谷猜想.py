#f = open("数据.txt", "a")
import sys
sys.stdout = open("数据.txt", "a")
#数 = int((open("数据.txt", mode='r', encoding="UTF-8").readlines()[-1]).split()[0])
数 = 3717758
while True:
    num = 数
    次数 = 0
    while num != 1:
        if num%2 == 0:
            num = num/2
        else:
            num = num*3+1
        次数 += 1
    print(数, 次数, True)
#    f.write(str(数)+","+str(次数)+",True\n")
    数 += 1