package main

import "testing"

func TestCase1(t *testing.T) {
	findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}})
}
