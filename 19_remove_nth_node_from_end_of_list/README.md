## 解題思路


```
n = 2

1 -> 2 -> 3 -> 4 -> 5
          l         r
```

我們要移除從右邊數過來第 2 個 node `4`。請問一個問題，如果我們要移掉 `4` 我們要知道哪一個 node 的 position?
我們必須知道 `3` 的 position。要如何知道呢？我們知道 5 ~ 3 有著 2 個 nodes 的 offset。


```
1 -> 2 -> 3 -> 4 -> 5
          l         r

l r 間差了 2 個 node 的 offset
```

如果我們維持這樣的 offset 把 l & r 前推直到 l 在 1 的位置

```
1 -> 2 -> 3 -> 4 -> 5
l         r

```

維持 offset 不變，把 r 移動到 nil 為止

```
1 -> 2 -> 3 -> 4 -> 5
               l         r

```

喔呦，你會發現 l 在我們要移除的地方。那如果我們在 1 前 create 一個 dummy node 把 l 指向這個 dummy node:

```
dummy -> 1 -> 2 -> 3 -> 4 -> 5
    l              r

```
然後把 r 移動到 nil 止

```
dummy -> 1 -> 2 -> 3 -> 4 -> 5
                   l              r

```

你就會發現 l 是在 3 的位置。這樣我們就可把 node 4 移掉了。

### Recap
