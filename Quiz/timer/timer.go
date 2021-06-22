package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	timer := time.NewTimer(1 * time.Second)
	//	x := <-timer.C

	for {
		select {
		case <-timer.C:
			{
				fmt.Println("The time has come!")
				os.Exit(0)

			}
		default:
			{
				fmt.Println("Nothin yet")
			}
		}
	}

	// fmt.Println(x)

}
