package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

func swap(n, m int) (_y, _x int) {
	_y = m
	_x = n
	return
}
