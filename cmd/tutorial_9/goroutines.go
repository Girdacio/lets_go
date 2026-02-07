package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var waitgroup = sync.WaitGroup{}
var dbData = []string{"Data1", "Data2", "Data3", "Data4", "Data5"}
var results = []string{}

func main() {
	startTime := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}
	fmt.Printf("Total time taken: %v\n", time.Since(startTime))

	startTime2 := time.Now()
	for i := 0; i < len(dbData); i++ {
		waitgroup.Add(1)
		go dbCall2(i)
	}
	waitgroup.Wait()
	fmt.Printf("Total time taken: %v\n", time.Since(startTime2))
	fmt.Printf("Results collected: %v\n", results)

	results = []string{} // Reset results for the next test
	startTime3 := time.Now()
	for i := 0; i < len(dbData); i++ {
		waitgroup.Add(1)
		go dbCall3(i)
	}
	waitgroup.Wait()
	fmt.Printf("Total time taken: %v\n", time.Since(startTime3))
	fmt.Printf("Results collected: %v\n", results)
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The result from the database is: %s\n", dbData[i])
}

func dbCall2(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The result from the database is: %s\n", dbData[i])
	results = append(results, dbData[i])
	waitgroup.Done()
}

func dbCall3(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The result from the database is: %s\n", dbData[i])
	mutex.Lock()
	results = append(results, dbData[i])
	mutex.Unlock()
	waitgroup.Done()
}
