package main

import "C"

import (
	"log"
)

//export asin
func asin(x C.double) C.double {
	log.Printf("asin called")
	for {
	}
	return x
}

//export acos
func acos(x C.double) C.double {
	log.Printf("acos called")
	return x
}

// this isn't used
func main() {}
