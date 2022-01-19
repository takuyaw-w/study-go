package main

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}

func swap2(n, m *int) {
	var tmp int = *n
	*n = *m
	*m = tmp
}
