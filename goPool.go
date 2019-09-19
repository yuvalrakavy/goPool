// goPool - keep track of a pool of goRoutines and allow controlled termination of them
//
// Usage:
//
//  first create the pool
//   pool := goPool.Make()
//
package goPool

import (
	"sync"
)

type GoPool struct {
	Done chan interface{}
	wg   sync.WaitGroup
}

func MakeGoPool() *GoPool {
	return &GoPool{Done: make(chan interface{})}
}

func (pool *GoPool) Terminate() {
	close(pool.Done)
	pool.wg.Wait()
}

func (pool *GoPool) Enter() {
	pool.wg.Add(1)
}

func (pool *GoPool) Leave() {
	pool.wg.Done()
}
