package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", "Tue, 02 Apr 2024 00:00:00 GMT")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
