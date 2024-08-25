package main

import (
	"log"
	"net"
	"net/rpc"
)

// Dinh nghia service struct
type HelloService struct{}

// Dinh nghia ham service Hello, quy tac:
// 1. Ham Service phai public (viet hoa)
// 2. Co 2 tham so trong ham
// 3. Tham so thu 2 la con tro
// 4. Phai tra ve kieu error
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	// Dang ky service
	rpc.RegisterName("HelloService", new(HelloService))
	// Chay RPC server tren cong 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	// Phuc vu nhieu client
	for {
		// accept 1 connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// Phuc vu client tren 1 goroutine khac
		// de giai phong main thread tiep tuc vong lap
		go rpc.ServeConn(conn)
	}
}