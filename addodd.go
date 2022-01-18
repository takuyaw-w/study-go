package main

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if addOrOdd(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}

func addOrOdd(count int) bool {
	return count%2 == 0
}
