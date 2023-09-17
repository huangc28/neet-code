package main

func mergeTriplets2(triplets [][]int, target []int) bool {
	m := make(map[int]interface{})
	for i := 0; i < len(triplets); i++ {
		triplet := triplets[i]
		if triplet[0] > target[0] || triplet[1] > target[1] || triplet[2] > target[2] {
			continue
		}

		for idx, num := range triplet {
			if num == target[idx] {
				m[idx] = struct{}{}
			}
		}
	}

	return len(m) == 3
}
