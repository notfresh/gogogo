package main

import "fmt"
// import "person/notfresh/hello/morestrings"
import "github.com/google/go-cmp/cmp"


func main() {
	// s := "hello, world"
	fmt.Println("hello, world")
	// PrintRandInt(0,10,4)
	// fmt.Println(morestrings.ReverseRunes(s))

	// fmt.Println("================")

	fmt.Println(cmp.Diff("Hello\n World", "Hello\n Go"))
}
