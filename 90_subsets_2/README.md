#algorithm/backtracking/90-subset-2

## 解題思路

這題可以使用 decision tree，在 i = 0 的時候，我要不要 `include nums[i]` 到 subset 裡。 可以要，也可以不要。以一搬的 subset generation 我們會得出以下 decision tree graph:

*[1,2,2]*

```

	        [1,2,2]
i= 0     /      \
       [1]       []
i=1    /  \.     /  \
    [1,2] [1]   [2]  []
i=2	 / \   /    /  \
    /   \ [1,2][2,2][2]
[1,2,2] [1,2]
```

因為全部列下來會太多，這邊畫到一部分就好。這邊我們可以看到  `[1,2]` 重複了。原因是在 `i=1` 時，我們左邊選擇 consider 2 `[1,2]`，右邊選擇 consider 1 `[1]`。既然我們左邊選擇 consider 2 了，右邊之後的所有 decision 就不用再 consider 2 了。所以我們可以把 `i` 移動到不是 2 的位置。

```go
j := i
for j < len(nums) && nums[j] == nums[i] {
	j+=1
}
```



