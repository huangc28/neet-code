package main

import (
	"log"
	"testing"
)

func TestCase2(t *testing.T) {
	output := minMeetingRooms_2([][]int{
		{1, 13},
		{13, 15},
	})

	log.Printf("debug %v", output)
}
