package list

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type List struct {
	mutex     *sync.RWMutex
	Len       int64
	FirstNode *Node
}

type Node struct {
	Data     any
	NextNode *Node
}

func NewList() (l *List) {
	l = &List{
		Len:       0,
		FirstNode: nil,
	}
	return
}

func (l *List) Add(data any) (index int64, err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.FirstNode == nil {
		n := &Node{}
		n.Data = data
		l.FirstNode = n
		l.Len++
		return l.Len - 1, nil
	}

	if reflect.TypeOf(l.FirstNode.Data) != reflect.TypeOf(data) {
		return 0, errors.New("Wrong Type")
	}

	cn := l.FirstNode
	for {
		if cn.NextNode == nil {
			break
		}
		cn = cn.NextNode
	}
	n := &Node{}
	n.Data = data
	cn.NextNode = n
	l.Len++
	return l.Len - 1, nil
}

func (l *List) Print() {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if l.FirstNode == nil {
		fmt.Println("Empty list")
		return
	}
	cn := l.FirstNode
	for {
		if cn.NextNode == nil {
			fmt.Println(cn.Data)
			break
		}
		fmt.Println(cn.Data)
		cn = cn.NextNode
	}
}

func (l *List) Delete(index int64) (ok bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if index > l.Len-1 {
		return false
	}

	if index == 0 {
		l.FirstNode = l.FirstNode.NextNode
		l.Len--
		return true
	}

	var lastNode *Node
	currentNode := l.FirstNode
	for n := int64(0); n < index; n++ {
		lastNode = currentNode
		currentNode = lastNode.NextNode
	}
	nextNode := currentNode.NextNode
	lastNode.NextNode = nextNode
	l.Len--
	return true
}

func (l *List) Get(index int64) (data any) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	currentNode := l.FirstNode
	for n := int64(0); n < index; n++ {
		currentNode = currentNode.NextNode
	}
	return currentNode.Data
}

func (l *List) Sort(more func(i, j any) bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if l.FirstNode == nil {
		fmt.Println("Empty list")
		return
	}

	for i := int64(0); i < l.Len-1; i++ {
		currentNode := l.FirstNode
		for {
			if currentNode.NextNode == nil {
				break
			}
			if more(currentNode.Data, currentNode.NextNode.Data) {
				currentNode.NextNode.Data, currentNode.Data = currentNode.Data, currentNode.NextNode.Data
			}
			currentNode = currentNode.NextNode
		}
	}
}

func (l *List) SortLink(more func(i, j any) bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	if l.FirstNode == nil {
		fmt.Println("Empty list")
		return
	}

	for i := int64(0); i < l.Len-1; i++ {
		if more(l.FirstNode.Data, l.FirstNode.NextNode.Data) {
			n := l.FirstNode
			l.FirstNode = l.FirstNode.NextNode
			n.NextNode = l.FirstNode.NextNode
			l.FirstNode.NextNode = n
		}
		lastNode := l.FirstNode
		currentNode := l.FirstNode.NextNode
		for {
			if currentNode.NextNode == nil {
				break
			}
			if more(currentNode.Data, currentNode.NextNode.Data) {
				n := currentNode
				lastNode.NextNode = currentNode.NextNode
				n.NextNode = currentNode.NextNode.NextNode
				lastNode.NextNode.NextNode = n
				currentNode = lastNode.NextNode
			}
			lastNode = lastNode.NextNode
			currentNode = currentNode.NextNode
		}
	}
}
