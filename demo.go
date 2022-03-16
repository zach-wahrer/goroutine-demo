package goroutine_demo

import (
	"sync"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Shared Struct ////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SharedData struct {
	// Setup concurrency related types
	sync.Mutex
	Total int
}

func SumShared(allNums [][]int) int {
	data := &SharedData{}
	wg := new(sync.WaitGroup)

	for _, nums := range allNums {
		// Since we are starting a goroutine, add 1 to our wait group.
		wg.Add(1)
		// Start the actual work.
		go addToShared(data, wg, nums)
	}

	// Once all of our goroutines are started, we want to block until they are done.
	wg.Wait()

	return data.Total
}

func addToShared(data *SharedData, wg *sync.WaitGroup, nums []int) {
	// Once this function is finished, let the waitgroup know it is done.
	defer wg.Done()

	// Do the actual work.
	var sum int
	for _, num := range nums {
		sum += num
	}

	// Lock our data so that only this goroutine can access it.
	// This will block if another thread already has a lock.
	data.Lock()
	defer data.Unlock()
	data.Total += sum
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Channel ////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SumChannel(allNums [][]int) (total int) {
	// Setup concurrency related types
	channel := make(chan int)
	wg := new(sync.WaitGroup)

	for _, nums := range allNums {
		// Since we are starting a goroutine, add 1 to our wait group.
		wg.Add(1)
		// Start the actual work.
		go addViaChannel(channel, wg, nums)
	}

	// Start a function that blocks until our waitgroup is done,
	// then close the channel (since we know the work is done).
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(channel)
	}(wg)

	// Listen on the channel for sums. This blocks until the channel is closed.
	for sum := range channel {
		total += sum
	}

	return total
}

func addViaChannel(c chan int, wg *sync.WaitGroup, nums []int) {
	// Once this function is finished, let the waitgroup know it is done.
	defer wg.Done()

	// Do the actual work.
	var sum int
	for _, num := range nums {
		sum += num
	}

	// Send the result back over the channel.
	c <- sum
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Warning: Data Race!!! `go test -race` //////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type SharedDataNoMutex struct {
	Total int
}

func SumSharedNoMutex(allNums [][]int) int {
	data := &SharedDataNoMutex{}
	wg := new(sync.WaitGroup)

	for _, nums := range allNums {
		wg.Add(1)
		go addToSharedNoMutex(data, wg, nums)
	}

	wg.Wait()

	return data.Total
}

func addToSharedNoMutex(data *SharedDataNoMutex, wg *sync.WaitGroup, nums []int) {
	defer wg.Done()

	var sum int
	for _, num := range nums {
		sum += num
	}

	data.Total += sum
}
