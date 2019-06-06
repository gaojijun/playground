package main

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	ch := make(chan int)

	select {
		case ch <- 0:
			panic("should not come here")
		default:
			fmt.Println("hello")

	}
}
