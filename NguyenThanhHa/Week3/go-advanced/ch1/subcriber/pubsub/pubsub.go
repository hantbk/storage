package pubsub

import (
	"sync"
	"time"
)

type (
	// Subscriber thuoc kieu channel
	subscriber chan interface{}

	// topic la 1 filter
	topicFunc func(v interface{}) bool
)

type Publisher struct {
	// Read/Write Mutex
	mu sync.RWMutex

	// Kich thuoc hang doi
	buffer int

	// Time out cho viec publish
	timeout time.Duration

	// subscriber da dang ky theo topic nao
	subscribers map[subscriber]topicFunc
}

// constructor voi timeout va do dai hang doi
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer: buffer,
		timeout: publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Them subscriber moi, dang ky het tat ca topic
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// Them subscriber moi, subcriber cac topic da duoc filter loc
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.mu.Lock()
	p.subscribers[ch] = topic
	p.mu.Unlock()
	return ch
}

// Huy Subcribe
func (p *Publisher) Evict(sub chan interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// Publish ra 1 topic
func (p *Publisher) Publish(v interface{}){
	p.mu.RLock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers{
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// Dong 1 doi tuong Publisher va dong tat ca cac subscribers
func (p *Publisher) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// Gui 1 topic co the duy tri trong thoi gian cho wg
func (p *Publisher) sendTopic(
	sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup,
) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <- time.After(p.timeout):
	}
}