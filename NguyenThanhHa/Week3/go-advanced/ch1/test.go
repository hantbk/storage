package main

import (
    "fmt"
//     "log"
//     "net/http"
//     "time"
)

// func main() {
//     fmt.Println("Please visit http://127.0.0.1:12345/")

//     // sử dụng giao thức http để in ra chuỗi bằng lệnh 'fmt.Fprintf'
//     // thông qua log package
//     http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
//         s := fmt.Sprintf("Xin chào - Thời gian hiện tại: %s", time.Now().String())
//         fmt.Fprintf(w, "%v\n", s)
//         log.Printf("%v\n", s)
//     })

//     // khởi động service http
//     if err := http.ListenAndServe(":12345", nil); err != nil {
//         log.Fatal("ListenAndServe: ", err)
//     }
// }

func myAppend(sl []int, val int) []int{
	sl = append(sl, val)
	printSlice(sl)
	return sl
}

func printSlice(sl []int) {
	fmt.Printf("Slice %v\n", sl)
	fmt.Printf("Len %v, Cap %v\n\n", len(sl), cap(sl))
}

func main() {
	sl := make([]int, 1)
	printSlice(sl)
	for i := 1; i < 5; i ++ {
		sl = myAppend(sl, i)
	}
}