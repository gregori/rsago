package main

import "fmt"

func main() {
	test := "asdfasdf232435@#%@"

	chunk1 := NewString(test)
	n := chunk1.BigIntValue()

	println("Big integer value of " + test + " = " + n.Text(10))
	fmt.Printf("Blocksize for that is %d\n", BlockSize(*n))

	chunk2 := NewBigInt(n)
	s := chunk2.stringVal

	println("converted back to string = " + s)
	if s == test {
		println("success")
	} else {
		println("FAIL")
	}
}
