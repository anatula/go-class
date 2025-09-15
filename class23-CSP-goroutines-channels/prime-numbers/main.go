package main

import "fmt"

// in main create goroutine called generator
// function of generator is start making numbers
// as we get a new prime (divisors 1 and itself), start with 2
// 1 is never prime, ignore it
// get 2, need to filter out (and multiples of 2) create new goroutine "2 filter" attach it on the channel in between generator and main
// take channel that went from generator to main and make that go into the "2 filter" and create a new channel coming out "2 filter" to main
// 3 is also 3 prime create another filter
// 4 never gets there, gets removed by "2 filter"
// 5, create "5 filter" add it to chain of channels and filter
// 6 will be gobbled up
// 7 will make a filter
// 8 will be gobbled up by 2 filter
// 9 will "    "      " "  3 filter
// a lot of number will drop, all main will see are prime numbers
// Everytime it sees a new prime number it creates a filter goroutine, hook it to channel getting numbers from and create a new channel back to itself

// add limit avoid run forever, run until limit hits, count up until some number, stop then close it channel
// when gen closes "2 filter" will see incomming channel is closed, it will close outgoing channel, like dominos will all close until main see last channel close
// when main channel sees close, see socket close or EOF its gonna know theres no more data to work with and be done, complete its loop
// can write to channel, won't be able to read it
func generate(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		//everytime we make a number put it in a channel
		ch <- i
	}
	close(ch)
}

// run it in a go routine
// source, where to write and a prime number filter
func filter(src <-chan int, dst chan<- int, prime int) {
	// range over src channel, we get a sequence of numbers comming out of the channel
	// loop will block at the top everytime we go back to src to get another piece of data if there isn't one ready we pause
	// when there is we execute the loop body until channel closes
	// when src closes this loop is done, exit the loop
	// for loop over the values in the channel until the channel closes
	// when that happens whe close the next one
	// if start with 2, then 4,6,8,10 devide with no reminder, throw those number, odd numbers pass this filter
	for i := range src {
		if i%prime != 0 {
			// odd numbers will pass
			// pass it along
			dst <- i
		}
	}
	// when loop is finished close dst channel
	close(dst)
}

func sieve(limit int) {
	// when we start there's always 1 channel (comming back to main from generator)
	ch := make(chan int)

	//start generator
	go generate(limit, ch)

	for {
		// read my channel, ok boolean if channel is closed or not
		// try to read, will be closed, get not ok, end the program
		prime, ok := <-ch

		if !ok {
			break
		}
		// made new channel
		ch1 := make(chan int)

		// made filter
		go filter(ch, ch1, prime)

		// update my view of the channel comming into me with new channel comming out of the filter
		ch = ch1

		fmt.Print(prime, " ")

	}

	fmt.Println()
}

func main() {
	sieve(100) // expect 2 3 5 7 11 13 17 19 ...

}
