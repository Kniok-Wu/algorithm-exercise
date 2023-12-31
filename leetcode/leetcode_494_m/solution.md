# 目标和
这道题有两种思路，一种是每次计算一个新的DP数组，记录每个元素选择正负之后得到的新结果的值。  
```go
dp[i] = dp[i - 1] + i && dp[i - 1] - i
```  
其中，dp数组的下标为0～2*sum，映射为-sum～sum，保存的值为该索引被计算出来的种数。这也是我用的方法。  

还有一种思路是转换为背包问题，假设一个数组可以分成两部分，left和right，其中left取正，right取负，那么就有以下递推式：  
```go
left - right == target
left + right == sum
left = (target + sum) / 2
```
因此这样就可以转换为装满容量为left的背包的装法有几种，这样就无须考虑负数的情况。  
感觉这里代码随想录的思路有点问题，0是可以初始化为0的，因为0在本题没有意义，我添加了0完全是为了阅读上的直观性，只需要将dp[nums[0]]初始化为1即可，然后利用0-1背包的方法解题，将0-1背包中的价值变换为类加的减法即可，具体可以参考代码中的解法二。  

![img.png](img.png)

