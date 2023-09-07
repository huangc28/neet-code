package main

// tank = tank + cost[i-1] + gas[i]
func canCompleteCircuit(gas []int, cost []int) int {
	sumOfGas := 0
	sumOfCost := 0

	for i := 0; i < len(gas); i++ {
		sumOfGas += gas[i]
		sumOfCost += cost[i]
	}

	if sumOfCost > sumOfGas {
		return -1
	}

	// 一定有一個 solution 因為 sumOfCost > sumOfGas

	total := 0    // total gas in the tank at each position
	startIdx := 0 // index of the station
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		if total < 0 { // this isn't the right index to start, let's try next index
			total = 0
			startIdx += 1
		}
	}

	return startIdx
}
