package main

import (
	"fmt"
	"time"
)

func main() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	var t  = time.Now().In(location).Format(time.DateTime)
	fmt.Println(t)
}
