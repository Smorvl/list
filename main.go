package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	numOfGo    int64 = 50
	numOfNotes int64 = 500
)

type Cache struct {
	mx *sync.Mutex
	ch []int64
}

func (ch *Cache) Add(data int64) {
	ch.mx.Lock()
	defer ch.mx.Unlock()
	ch.ch = append(ch.ch, data)
}

func writer(ch *Cache, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := int64(0); i < numOfNotes; i++ {
		n := rand.Intn(500000)
		ch.Add(int64(n))
	}
}

func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	//var mx sync.Mutex
	ch := Cache{
		ch: []int64{},
		mx: &sync.Mutex{},
	}
	wg.Add(int(numOfGo))
	for i := int64(0); i < numOfGo; i++ {
		go writer(&ch, &wg)
	}
	wg.Wait()
	fmt.Println("Должно быть записано:", numOfNotes*numOfGo, "значений")
	fmt.Println("В итоге записано", len(ch.ch), "значений")
}
