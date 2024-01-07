package main

import (
	"fmt"
	"sync"
)

/*
     You can use synchronization primitives such as channels and mutexes, to avoid deadlock and race conditions in Go
*/

func read(channel chan byte, mutex sync.Mutex){
	// Create a shared buffer (byte slice)
	sharedBuffer := make([]byte, 0)
	
	for {
		data := <-channel   // Read the data
		
		mutex.Lock()   // Lock the mutex to protect shared buffer
		sharedBuffer = append(sharedBuffer, data)  // Append the data to the shared buffer
		mutex.Unlock()
	}
}

func write(channel chan byte, i int){
	for {
		data := byte(i)
		channel <- data  // Send the data
	}
}

func main() {
	var M [4]int = [4]int{8, 8, 8, 2}
	var N [4]int = [4]int{2, 8, 16, 8}
	
	// Create a channel for communication between readers and writers
	channel := make(chan byte)

	// Create a mutex to protect the shared buffer
	mutex := sync.Mutex{}

	for i := 0; i < len(M); i++ {

		fmt.Println("M: ", M[i], ", N: ", N[i])

		for j := 0; j < M[i]; j++ {
			go read(channel, mutex)
		}

		for j := 0; j < N[i]; j++ {
			go write(channel, i)
		}
	}

	select{}
}