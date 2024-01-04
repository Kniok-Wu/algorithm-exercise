# 单词拆分
这道题是一个变种的动态规划问题，将原本的数值转变为了bool值，即保存是否可以由这个单词拼接而成，DP数组的坐标代表了字符长度。  
该题递推式如下：
```go
DP := make([]bool, n + 1)
DP[i] = DP[i - len(word)] && (s[:i][i - len(word):] == word)
```  

![img.png](img.png)