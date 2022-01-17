package main

import "fmt"

func main() {
	p := struct {
		name string
		age  int
	}{
		name: "HOGE",
		age:  20,
	}

	p.age++
	fmt.Println(p)

	ns := [...]int{20, 30, 40, 50, 60}
	fmt.Println(ns[0])
	fmt.Println(ns[0:5])

	sl := []int{5: 40, 10: 100}
	fmt.Println(sl)
	sl = append(sl, 22)
	fmt.Println(sl)
}
