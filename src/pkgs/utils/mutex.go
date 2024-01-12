package utils

import "sync"

var mutex *sync.Mutex = &sync.Mutex{}
var mutexRW *sync.RWMutex = &sync.RWMutex{}

func GetMutex() *sync.Mutex {
	return mutex
}

func GetMutexRW() *sync.RWMutex {
	return mutexRW
}
