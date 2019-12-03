package main

import "fmt"

func main() {
	closest, shortest := calculateDistance(getDirections())
	fmt.Printf("closest is %v, shortest is %v", closest, shortest)
}
