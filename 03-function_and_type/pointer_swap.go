package main

func swap2(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
