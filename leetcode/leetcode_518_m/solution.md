# 零钱兑换 II
此题属于**完全背包问题**，与0-1背包问题的区别在于物品可以被反复使用。  
此问题中存在遍历顺序的问题，假设先固定背包容量（即零钱总和），再遍历硬币，则可能会出现这种情况：
```go
coins := []int{1, 2, 5}
dp[total] += dp[total - coin]
```
当total为3，coin为1的情况下，添加了total为2的情况 -> `[]int{1, 1}` & `[]int{2}`两种情况。
当total为3，coin为2的情况下，添加了total为1的情况 -> `[]int{1}` 一种情况。
但是此时，可以发现添加了两次`[]int{1, 2}`。  
因此我们可以发现，当固定容量遍历物品时，物品可以出现不同的先后顺序，因此我们得到的是一个排列问题的答案。  
同理，我们推理固定物品，遍历容量时，由于我们的物品严格遵循先后顺序，因此得到的是组合问题的答案。  
因此，本题应该选择**组合问题**的解题思路。  

![img.png](img.png)