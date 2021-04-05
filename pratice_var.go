package main

import "fmt"

func main() {
	// 單個字母為rune
	fmt.Println('a', 'v', 'c')
	// go的變數宣告方式
	var x int //宣告變數名稱（x）和型別
	x = 4     //給予資料
	fmt.Println(x)
	var y string = "hello 安安"
	fmt.Println(y)
	var z bool = false 
	fmt.Println(z)
}
