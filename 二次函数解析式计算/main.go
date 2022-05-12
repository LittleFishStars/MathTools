package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
	var (
		shu_ru_1 string
		shu_ru_2 string
		shu_ru_3 string
	)

	fmt.Println("请输入坐标：(例如： 0,0 )")
	fmt.Scanln(&shu_ru_1)
	fmt.Scanln(&shu_ru_2)
	fmt.Scanln(&shu_ru_3)

	var (
		zuo_biao_1 []string = strings.Split(shu_ru_1, ",")
		zuo_biao_2 []string = strings.Split(shu_ru_2, ",")
		zuo_biao_3 []string = strings.Split(shu_ru_3, ",")
	)

	x1 ,e1 := strconv.Atoi(zuo_biao_1[0])
	y1 ,e2 := strconv.Atoi(zuo_biao_1[1])
	x2 ,e3 := strconv.Atoi(zuo_biao_2[0])
	y2 ,e4 := strconv.Atoi(zuo_biao_2[1])
	x3 ,e5 := strconv.Atoi(zuo_biao_3[0])
	y3 ,e6 := strconv.Atoi(zuo_biao_3[1])

	// fmt.Println(x1,y1)
	// fmt.Println(x2,y2)
	// fmt.Println(x3,y3)

	if ((e1!=nil)&&(e2!=nil)&&(e3!=nil)&&(e4!=nil)&&(e5!=nil)&&(e6!=nil)){
		fmt.Println("错误")
		fmt.Println(e1,e2,e3,e4,e5,e6)
	}

	var (
		a int
		b int
		c int
		// t int = 2
	)

	var (
		lin_shi_1 int = x1*x1
		lin_shi_2 int = x2*x2
		lin_shi_3 int = lin_shi_1-lin_shi_2
		lin_shi_4 int = x1-x2
		lin_shi_5 int = x3*x3
		lin_shi_6 int = lin_shi_1-lin_shi_5
		lin_shi_7 int = x1-x3

		lin_shi_8 int = lin_shi_3*lin_shi_7
		lin_shi_9 int = lin_shi_6*lin_shi_4
		lin_shi_10 int = lin_shi_7*(y1-y2)
		lin_shi_11 int = lin_shi_4*(y1-y3)
		lin_shi_12 int = lin_shi_10-lin_shi_11
		lin_shi_13 int = lin_shi_8-lin_shi_9
	)
	// fmt.Println(lin_shi_1,lin_shi_2,lin_shi_3,lin_shi_4)
	a = lin_shi_12/lin_shi_13
	b = ((y1-y2)-(lin_shi_3*a))/lin_shi_4
	c = y1 - (x1*x1)*a - x1*b

	fmt.Println("a,b,c:")
	fmt.Println(a,b,c)

	var jie_guo string
	jie_guo = "y=" + strconv.Itoa(a) + "x^2+" + strconv.Itoa(b) + "x+" + strconv.Itoa(c)

	fmt.Println(jie_guo)
	main()
	// fmt.Scanln()
}


