這一題使用 用一個 hashmap 記下看過的 character

```
seen := map[byte]bool
l,r := 0,0

for r < len(s) {
  // l 一直往前直到沒看到 s[r] 為止
  for seen[s[r]] {
    seen[s[l]] = false
    l += 1
  }

  seen[s[r]] = true
  maxLen = max(maxLen, r-l+1)
  r += 1
}
```
