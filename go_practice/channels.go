// channels are a typed conduit thru which you can send and receive values with the channel operator, <-
// ch <- v send v to channel ch
// v := <-ch receive from ch and assign value to v
// data flows in the direction of the arrow
// like maps + slices, channels must be created before use
// ch := make(chan int)

package main
import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	fmt.Println("finished goroutine")
	fmt.Println(s, sum)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println("finished full sum")
	fmt.Println(x, y, x+y)
}