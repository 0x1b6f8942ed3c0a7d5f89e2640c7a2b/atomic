package main

import (
	"fmt"
	"time"

	sms "main/scripts/smsbomb"
)

func main() {

	// Ensure there are proxies available

	fmt.Println("loading sms :: atomic/scripts/smsbomb")
	for {
		sms.StartALL()
		time.Sleep(time.Microsecond * 1)
	}
}
