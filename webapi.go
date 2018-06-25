package web

import "sync"

func Wait() {
	var wg = new(sync.WaitGroup)
	wg.Wait()
}
