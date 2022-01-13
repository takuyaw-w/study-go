package main

import (
	"fmt"
)

func main() {
	// 変数宣言
	var hoge int = 1234
	var fuga string = "hogehogehoge"
	var hooo bool = true
	var fuuu float64 = 1.2223333
	fmt.Println(hoge)
	fmt.Println(fuga)
	fmt.Println(hooo)
	fmt.Println(fuuu)
	// 省略形
	humuhumu := 12345
	fumufumu := "strings"
	nooooo := false
	uuuuuu := 3.45
	fmt.Println(humuhumu)
	fmt.Println(fumufumu)
	fmt.Println(nooooo)
	fmt.Println(uuuuuu)

	// エイリアス
	type HogeString string
	type FugaString string
	var hogestr HogeString = "hogehoge"
	var fugastr FugaString = "fugafuga"
	fmt.Println(hogestr)
	fmt.Println(fugastr)

	// キャスト
	var a uint16 = 12345
	var b uint32 = uint32(a)
	fmt.Println(a, b)

	var _a uint32 = 1234567890
	var _b uint16 = uint16(a)
	fmt.Println(_a, _b)

	// 配列
	strArr := [3]string{}
	strArr[0] = "hogehoge"
	strArr[1] = "hogehoge1"
	strArr[2] = "hogehoge2"
	fmt.Println(strArr)

	// slice
	strsli := []string{}
	strsli = append(strsli, "hogehoge1")
	strsli = append(strsli, "hogehoge2")
	strsli = append(strsli, "hogehoge3")
	fmt.Println(strsli)
	aaaa := []int{}
	for i := 0; i < 10; i++ {
		aaaa = append(aaaa, i)
		fmt.Println(len(aaaa), cap(aaaa))
	}
}
