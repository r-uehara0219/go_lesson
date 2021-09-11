package main

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	n, m := swap(10, 20)
	println(n, m)
}
