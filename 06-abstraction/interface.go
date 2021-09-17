package main

import "fmt"

type Stringer interface {
	String() string
}

type I int

func (_ I) String() string {
	return "Integer"
}

type B bool

func (_ B) String() string {
	return "Boolean"
}

type S string

func (_ S) String() string {
	return "String"
}

func F(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Println(s.String(), int(v))
	case B:
		fmt.Println(s.String(), bool(v))
	case S:
		fmt.Println(s.String(), string(v))
	}
}

func main() {
	F(I(100))
	F(B(true))
	F(S("test"))
}
