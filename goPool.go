// goPool - keep track of a pool of goRoutines and allow controlled termination of them
//
// Usage:
//
//  first create the pool
//   pool := goPool.Make()
//
// Pass a pointer to the pool to each go routine that would participate in the pool:
//
//  go myGoRoutine(pool)
//
// When you want to terminate all the go routines in the pool and wait until all go routines
// have terminated, call the Terminate function
//
//  pool.Terminate()
//
// The go routine should enter the pool upon starting and Leave the pool when it is done
//
//  func myGoRoutine(pool *goPool.GoPool) {
//		pool.Enter()
//		defer pool.Leave()
//
//    // The function should terminate when the pool.Done channel is closed
//
//		<- pool.Done
//  }
//
package goPool

import (
	"sync"
)

// GoPool - manage pool of go routines
// each go in the pool should quite when the Done channel is closed
type GoPool struct {
	Done chan interface{}
	wg   sync.WaitGroup
}

// Make - Make a new pool
func Make() *GoPool {
	return &GoPool{Done: make(chan interface{})}
}

// Terminate - Signal all go routines in the pool to terminate (by cloasing the Done channel)
// wait until all of them are indeeded terminatated
func (pool *GoPool) Terminate() {
	close(pool.Done)
	pool.wg.Wait()
}

// Enter - Tell the pool that a new go routine has joined it
func (pool *GoPool) Enter() {
	pool.wg.Add(1)
}

// Leave - Tell the pool that this go routine is done
func (pool *GoPool) Leave() {
	pool.wg.Done()
}
