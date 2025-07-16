package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("🚀 --------------MAIN--------------------🚀", time.Now())
		fmt.Println("🚀 --------------Ram---------🚀", time.Now())
		time.Sleep(5 * time.Second)
	}
 
}