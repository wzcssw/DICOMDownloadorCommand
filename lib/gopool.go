package lib

import "sync"

type GoPool struct {
	queue     chan bool
	waitGroup *sync.WaitGroup
}

func NewInstance(size int) *GoPool {
	if size < 1 {
		size = 1
	}
	return &GoPool{queue: make(chan bool, size), waitGroup: &sync.WaitGroup{}}
}

func (pool *GoPool) Add() {
	pool.queue <- true
	pool.waitGroup.Add(1)
}

func (pool *GoPool) Done() {
	<-pool.queue
	pool.waitGroup.Done()
}

func (pool *GoPool) Wait() {
	pool.waitGroup.Wait()
}
