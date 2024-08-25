package main

// const char* cs = "linux";
import "C"
import "example/hant/cgo/cgo_helper"

func main() {
	cgo_helper.PrintCString(C.cs)
}
