package main

import (
	sms "ato/scripts/smsbomb"
	"fmt"
	"time"
)

func main() {

	// Ensure there are proxies available

	fmt.Println("loading sms :: ato/scripts/smsbomb")
	for {
		sms.StartALL()
		time.Sleep(time.Microsecond * 1)
	}
}
