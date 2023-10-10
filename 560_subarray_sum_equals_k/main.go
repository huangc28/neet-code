package main

/*
[1,-1,1,1,1,1], k = 3

	i

下面這行到底是什麼意思呢

```
count, exists := prefixSumMap[prefixSum]

	if exists {
		ans += count
	}

```

假設 i 在 5 的位置而當前的 sum 是 4。4 - 3 = 1 (complement)。 我們只要減去 contiguous subarray sum 為 1 的個數，剩下的就是
contiguous subarray sum 為 3 的個數。

[1]
[1,-1,1]

上面兩組 contiguous subarray sum 為 1。必定有另外兩組 contiguous subarray sum 為 3

[1] + [-1,1,1,1,1]
[1,-1,1] + [1,1,1]

所以，我們答案直接加上 subarray sum 為 1 的 array 個數，就是 i 在 5 時 contiguous subarray sum 為 3 的 array 個數
*/
func subarraySum(nums []int, k int) int {
	prefixSumMap := make(map[int]int)
	prefixSumMap[0] = 1 // 1 subarray that sums up to 0
	sum, ans := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefixSum := sum - k
		count, exists := prefixSumMap[prefixSum]
		if exists {
			ans += count
		}
		prefixSumMap[sum] += 1
	}

	return ans
}
