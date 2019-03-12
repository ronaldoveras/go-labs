package main

import (
	"fmt"
	"time"
)

func main() {
        layout := "2006-01-02 15:04:05"
	start, _ := time.Parse(layout, "2019-02-28 10:25:01")
	end,_ := time.Parse(layout, "2019-03-01 10:25:01")
	secs := end.Sub(start)
	fmt.Printf("Hello, playground, %v. Diference in %v ", start, secs)
}
