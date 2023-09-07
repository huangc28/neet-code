package main

/*
We use tabulation to cache whether we can reach the end from the current index.

nums = [2,3,1,1,4]
dp   = [f,t,t,t,t]

- can we reach the end from step 4? yes, because we are at the last step
- step 3 can jump 1 step
  - check dp[3+1] yes
- step 2 can jump 1 step
  - check dp[2+1] yes
- step 1 can jump 3 steps
  - check dp[1+1]  yes
  - check dp[1+2]
  - check dp[1+3]
- step 0 can jump 2 steps
  - check dp[0+1] yes
  - check dp[0+2] yes

Caveat i+nums[i] must be < len(dp) to not exceed the stair.

Time: O(n-1) * O(m)
*/

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))

	// We can reach the last step if we are at the last step
	dp[len(dp)-1] = true

	for i := len(nums) - 2; i >= 0; i-- { // O(n) iterate through each step
		// Check every steps from 1 to nums[i] to see if we can reach the top.
		// by adding the step
		for jump := 1; jump < nums[i]+1; jump++ { // O(m) jump number at each step
			if jump >= len(nums)-1-i {
				dp[i] = true
				break
			}

			// Check every jump that nums[i] can jump.
			// If I jump to a step that can reach the top
			// that means I can reach the top.
			if i+jump < len(dp) && dp[i+jump] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

func canJump_2(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(dp)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		for jump := 1; jump <= nums[i]; jump++ {
			dest := i + jump
			// among the steps that I can jump, which step I can reach to the top?
			if dest < len(dp) && dp[dest] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}
