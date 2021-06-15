package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	wsStatus map[string]bool
	mut      sync.Mutex
)

var (
	wg sync.WaitGroup
)

func main() {

	wsStatus = make(map[string]bool)
	mut = sync.Mutex{}

	wg = sync.WaitGroup{}

	websites := []string{
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://www.linkedin.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
		"https://wesionary234.team/",
	}

	st := time.Now()
	for _, website := range websites {
		go getWebsite(website)
		wg.Add(1)
	}

	wg.Wait()

	et := time.Now()

	dur := et.Sub(st)
	fmt.Printf("Time Taken %v\n", dur.Seconds())

	for webs, st := range wsStatus {
		fmt.Printf("%s: Status %v\n", webs, st)
	}

}

func getWebsite(website string) {
	status := false
	defer wg.Done()
	if _, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")
	} else {

		status = true
	}

	mut.Lock()
	defer mut.Unlock()
	wsStatus[website] = status
}
