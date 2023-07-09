package main

import (
	"sort"
)

/*
	target = 12,
	position = [10,8,0,5,3],
	speed = [2,4,1,1,3]

- Use stack to compare current car and the car at the top of the stack.

- If car at back position is slower than the car at front position, car at the back has no chance to catch up car at the front.
That cart is a fleet

- If car at the back position is faster than the car at the front position, car at the back has the chance to catch up car at the front.
If it catches up, we have a fleet. Push the fleet to the stack with slower speed. If it can not catch up, we simply push the current
car to the stack as it can potentially be a fleet.
*/

type Car struct {
	Speed    int
	Position int
}

func carFleet_bad(target int, position []int, speed []int) int {
	stack := make([]Car, 0, len(position))

	for i := 0; i < len(position); i++ {
		if len(stack) == 0 {
			stack = append(stack, Car{speed[i], position[i]})
		} else {
			currCar := Car{speed[i], position[i]}
			backCar := currCar
			frontCar := stack[len(stack)-1]

			// figure out which car is at the front which is at the back.
			if backCar.Position > frontCar.Position {
				backCar, frontCar = frontCar, backCar
			}

			// can back car catch up the front car?
			backPos := backCar.Position
			frontPos := frontCar.Position
			found := false

			for backPos <= target && frontPos <= target {
				// we found a fleet
				if backPos >= frontPos {
					found = true
					break
				}

				backPos = backPos + backCar.Speed
				frontPos = frontPos + frontCar.Speed
			}

			if found {
				slowerSpeed := min(backCar.Speed, frontCar.Speed)

				//log.Printf("debug 1 %v", stack)
				// pop the car out
				stack = stack[:len(stack)-1]

				//log.Printf("debug 2 %v", stack)
				stack = append(stack, Car{slowerSpeed, frontPos})
			} else {
				// current car might be a potential fleet
				stack = append(stack, currCar)
			}
		}
	}

	return len(stack)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	stack := make([]Car, 0, len(position))

	for idx := range cars {
		cars[idx] = Car{speed[idx], position[idx]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position < cars[j].Position
	})

	for i := len(cars) - 1; i >= 0; i-- {
		car := cars[i]
		stack = append(stack, car)
		if len(stack) > 1 {
			first := stack[len(stack)-1]
			second := stack[len(stack)-2]

			firstImeToEnd := float32(target-first.Position) / float32(first.Speed)
			secondTimeToEnd := float32(target-second.Position) / float32(second.Speed)

			// the front car reaches the end faster than the later car.
			if firstImeToEnd <= secondTimeToEnd {
				stack = stack[:len(stack)-1]
			}
		}

	}

	return len(stack)
}
