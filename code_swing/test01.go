package main

import (
	"fmt"
	"os"
)

type Person struct {
	name string
	age  int
}

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
	urune := 'a'
	fmt.Println(humuhumu)
	fmt.Println(fumufumu)
	fmt.Println(nooooo)
	fmt.Println(uuuuuu)
	fmt.Println(urune)

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

	// map
	m := map[string]int{
		"x": 100,
		"y": 200,
		"z": 2,
	}
	fmt.Println(m)

	err, ok := m["z"]
	if ok {
		fmt.Println("exist", err)
	} else {
		fmt.Println("not exist", err)
	}
	for key, value := range m {
		fmt.Printf("%s = %d\n", key, value)
	}

	ad, mi := addMinus(1, 2)
	fmt.Println(ad, mi)

	person := Person{"takuya", 29}
	person.PrintOut()

	q, ok := ccccc(13.24)
	fmt.Println(q, ok)

	fmt.Println(os.Getenv("EDITOR"))
}

func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

func (p Person) ToString() string {
	return p.name
}

func (p Person) PrintOut() {
	fmt.Println(p.ToString())
}

func ccccc(a interface{}) (int, bool) {
	q, ok := a.(int)
	return q, ok
}
