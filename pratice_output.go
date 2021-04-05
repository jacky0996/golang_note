package main //載入main封包
import "fmt" //載入fmt封包

func main() {
	// var x int
	// fmt.Print("請輸入一個數字")
	// fmt.Scanln(&x) 宣告輸入值,此為執行後輸入,輸入多少則為多少,型態須符合
	// fmt.Println(x) 輸出
	var x int
	var y int
	fmt.Print("請輸入第一個數字")
	fmt.Scan(&x) //基本定義輸入fmy.Scan(& + 變數名稱)
	fmt.Print("請輸入第二個數字")
	fmt.Scan(&y)
	var z int = (x * y)
	fmt.Println(z)
}
