package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

// KVStoreService store key-value db
type KVStoreService struct {
	// map luu tru du lieu key-value
	m map[string]string

	// map chua danh sach cac ham filter
	// duoc xac dinh trong moi loi goi
	filter map[string]func(key string)

	// bao ve cac thanh phan khac khi duoc truy cap
	// boi nhieu goroutine
	mu sync.RWMutex
}

// NewKVStoreService construct new object KVStoreService
func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

// Get API tra ve gia tri tuong ung voi key
func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}
	return fmt.Errorf("key not found")
}

// Set API thay doi gia tri tuong ung voi key
func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		// Ham filter duoc goi khi value tuong ung voi key thay doi
		for _, fn := range p.filter {
			fn(key)
		}
	}
	p.m[key] = value
	return nil
}

// Watch tra ve key moi khi thay co su thay doi
func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	// id la 1 string ghi lai thoi gian hien tai
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())

	// buffered channel chua key
	ch := make(chan string, 10)

	// Filter de theo doi key thay doi
	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	select {
	// Tra ve timeout sau 1 khoang thoi gian timeoutSecond
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}

func ServeKVStoreService(conn net.Conn){
	p := rpc.NewServer()
	p.Register(NewKVStoreService())
	p.ServeConn(conn)
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}
	fmt.Println("Server is running ", "localhost:1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		go ServeKVStoreService(conn)
	}
}
