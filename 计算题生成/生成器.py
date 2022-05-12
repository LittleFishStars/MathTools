import random

gs = int(input("个数："))
n = 1
ZhiYouJiaJian = 0
while n != (gs+1):
    num1 = random.randrange(5, 2000, random.choice([10, 100, 200]))
    num2 = random.randrange(5, 2000, random.choice([10, 100, 200]))
    num3 = random.randrange(5, 2000, random.choice([10, 100, 200]))
    num4 = random.randrange(5, 2000, random.choice([10, 100, 200]))

    SuanShiMoBan = [
            "{0}-({1}+{2}-{3})".format(num1, num2, num3, num4),
            "{0}+({1}-{2}+{3})".format(num1, num2, num3, num4),
            "{0}*({1}+{2}/{3})".format(num1, num2, num3, num4),
            "{0}*({1}/{2}*{3})".format(num1, num2, num3, num4),
            "{0}*({1}/{2}*{3})".format(num1, num2, num3, num4),
            "{0}*({1}/{2}*{3})".format(num1, num2, num3, num4),
            "{0}*({1}/{2}*{3})".format(num1, num2, num3, num4),
            "{0}/({1}/{2}/{3})".format(num1, num2, num3, num4),
            "{0}/({1}/{2}/{3})".format(num1, num2, num3, num4),
            "{0}/({1}/{2}/{3})".format(num1, num2, num3, num4),
            "({0}-{1})*({2}-{3})".format(num1, num2, num3, num4),
            "{0}+{1}*{2}-{3}".format(num1, num2, num3, num4),
            "{0}+{1}*{2}-{3}".format(num1, num2, num3, num4),
            "{0}+{1}*{2}-{3}".format(num1, num2, num3, num4),
            "{0}+{1}*{2}-{3}".format(num1, num2, num3, num4),
            "({0}-{1})/{2}-{3}".format(num1, num2, num3, num4),
            "({0}-{1})/{2}-{3}".format(num1, num2, num3, num4),
            "({0}-{1})/{2}-{3}".format(num1, num2, num3, num4),
            "{0}+{1}*({2}-{3})".format(num1, num2, num3, num4),
            "{0}+{1}*({2}-{3})".format(num1, num2, num3, num4),
            "{0}+{1}*({2}-{3})".format(num1, num2, num3, num4),
            "{0}+{1}*({2}-{3})".format(num1, num2, num3, num4),
            ]
    SuanShi = random.choice(SuanShiMoBan)
    exec("DaAn = " + SuanShi)
    if (DaAn <= 10000) and (("*" in SuanShi) or ("/" in SuanShi) or (ZhiYouJiaJian <= int(gs*0.05))) and (DaAn >= 0) and ((DaAn % 1) == 0):
        if not("*" in SuanShi) and not("/" in SuanShi):
            ZhiYouJiaJian = ZhiYouJiaJian + 1
        DA = open("./题目.txt", "a")
        DA.write(str(n) + ":" + SuanShi + "\n")
        SS = open("./答案.txt", "a")
        SS.write(str(n) + ":" + str(DaAn) + "\n")
        n = n + 1 
        ZhiYouJiaJian = ZhiYouJiaJian - 0.1

print("生成成功！")