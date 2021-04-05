package main

import "fmt"

func main() {
	var x int
	fmt.Print("請輸入提款金額")
	fmt.Scanln(&x)
	if x < 1 {
		fmt.Print("您輸入的金額有誤，請重新輸入")
	} else {
		fmt.Print("您提出的款項為", x)
	}
}

