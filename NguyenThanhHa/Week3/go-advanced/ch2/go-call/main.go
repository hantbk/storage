package main

/*
static int add(int a, int b) {
    return a+b;
}
*/
import "C"
import "fmt"

func main() {
	// C.add(1, 1)
	fmt.Println(C.add(1, 1))
}