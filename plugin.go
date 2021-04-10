package main

// extern int fd;
import "C"

import (
	"context"
	"log"
)

var (
	ctx context.Context
)

//export asin
func asin(x C.double) C.double {
	log.Printf("asin called")
	<-ctx.Done()
	return x
}

// //export acos
// func acos(x C.double) C.double {
// 	log.Printf("acos called")
// 	return x
// }

func init() {
	var err error

	ctx, C.fd, err = contextWithCCancel(context.TODO())
	if err != nil {
		panic(err)
	}

}

// this isn't used
func main() {}
