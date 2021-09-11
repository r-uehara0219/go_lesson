package main

func main() {
	//n1 := 19
	//n2 := 86
	//n3 := 1
	//n4 := 12
	//
	//sum := n1 + n2 + n3 + n4
	//println(sum)

	sum := 0
	ns := []int{19, 86, 1, 12}
	for _, x := range ns {
		sum += x
	}
	println(sum)
}
