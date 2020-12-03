package main
 
import (
	"fmt"
	"sync"
	"time"
)
 
type Ball struct {
	hit int
	mux sync.Mutex
}
 
func (b *Ball) Ping() {
	b.mux.Lock()
	defer b.mux.Unlock()
 
	b.hit++
	fmt.Println("Ping", b.hit)
}
 
func (b *Ball) Pong() {
	b.mux.Lock()
	defer b.mux.Unlock()
 
	b.hit++
	fmt.Println("Pong", b.hit)
}
 
func main() { 
	ball := &Ball{}
	for {
		ball.Ping()
		time.Sleep(1 * time.Second)
		ball.Pong()
		time.Sleep(1 * time.Second)
	}
}
