package main

func carFleet(target int, position []int, speed []int) int {
	stack := make([][]int, 0)
	cars := make([][]int, 0)
	for i := 0; i < len(position); i++ {
		cars = append(cars, []int{position[i], speed[i]})
	}

	for j := len(cars) - 1; j >= 0; j-- {
		stack = append(stack, cars[j])

		if len(stack) > 1 {
			backCar := stack[len(stack)-1]
			frontCar := stack[len(stack)-2]

			backCarToTargetTime := float64(target-backCar[0]) / float64(backCar[1])
			frontCarToTargetTime := float64(target-frontCar[0]) / float64(frontCar[1])

			if backCarToTargetTime <= frontCarToTargetTime {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack)
}
