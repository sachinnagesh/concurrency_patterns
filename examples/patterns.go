package examples

import "fmt"

func printNumber(number int) {
	fmt.Printf("printNumber : %d\n", number)
}

// all execute in sync
func SimpleFunctionCall() {
	fmt.Println("start : SimpleFunctionCall")
	printNumber(10)
	fmt.Println("end : SimpleFunctionCall")

}

// execute async. Not guarantee of printNumber
func GoroutineCall() {
	fmt.Println("start : GoroutineCall")
	go printNumber(10) //go routine async call
	fmt.Println("end : GoroutineCall")
}

// controlled execution with use of channel
func ChannelCall() {
	fmt.Println("start : ChannelCall")
	mychannel := make(chan string) //create channel, it's a queue

	go func() { //anonymous func
		fmt.Println("in anonymous go routine")
		mychannel <- "anonymous" //puts value in channel
		mychannel <- "data2"
	}()

	msg := <-mychannel //gets value from channel
	msg1 := <-mychannel

	fmt.Printf("msg : %s\n", msg)
	fmt.Printf("msg1 : %s\n", msg1)
	fmt.Println("end : ChannelCall")

}

func SelectCall() {
	fmt.Println("start : SelectCall")
	myChannel1 := make(chan string)
	myChannel2 := make(chan string)

	go func() {
		myChannel1 <- "myChannel1 Data1"
	}()

	go func() {
		myChannel2 <- "myChannel2 Data1"

	}()

	select {
	case msgFromMyChannel1 := <-myChannel1:
		fmt.Println(msgFromMyChannel1)
	case msgFromMyChannel2 := <-myChannel2:
		fmt.Println(msgFromMyChannel2)
	}
	fmt.Println("end : SelectCall")

}
