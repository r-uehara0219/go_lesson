package main

import (
	"fmt"
	"os"
)

type Stringer interface {
	String() string
}

type CastError string

func (e CastError) Error() string {
	return string(e)
}

func ToStringer(v interface{}) (Stringer, error) {
	if str, ok := v.(Stringer); ok {
		return str, nil
	} else {
		return nil, CastError("cast error!!")
	}
}

type S string

func (s S) String() string {
	return string(s)
}

func main() {
	if s, err := ToStringer(S("string")); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Println(s.String())
	}
}
