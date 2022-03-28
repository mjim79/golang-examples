package main

import (
	"fmt"
	"github.com/mjim79/golang-examples/http_calls"
)

func main() {
	endpoints, err := http_calls.GetEndPoints()

	fmt.Println(err)
	fmt.Println(endpoints.EventsUrl)
}
