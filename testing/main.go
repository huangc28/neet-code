package main

import (
	"log"
)

type Sample struct {
	Hello string
}

func main() {
	s := Sample{}
	log.Printf("debug %v", s.Hello)
}
