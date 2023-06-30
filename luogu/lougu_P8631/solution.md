# 1. Manacher 算法
由于暴力枚举法需要从中心向两侧一一比对，因此可以推算出时间复杂度为 O(n^2)，效率较低；且对于偶数回文串而言，不方便进行比对求得。
下面对 Manacher算法中的一些基础概念进行一些介绍。

## 1.1 Manacher 字符串
针对偶数回文串不方便进行中心扩散(从中心字符 依次向两侧比对)的问题，Manacher对字符串进行了 **字符插入** 改造，姑且称此字符串为`Manacher字符串`。
- 奇数回文: aba -> #a#b#a#
- 偶数回文: abba -> #a#b#b#a# 
将所有的回文字符串均改造为了奇数回文字符串，假设输入长度为 N，则改造后的长度为 2N + 1，方便统一进行处理比对。
```go
// ManacherStr 将一个普通字符串改造为 Manacher字符串
func ManacherStr(str string) (mStr string) {
    var builder strings.Builder
    
    for _, b := range []byte(str) {
        builder.WriteString("#")
        builder.Write([]byte{b})
    }
    builder.WriteString("#")
    mStr = builder.String()
    return
}
```

## 1.2 回文半径 d[i]
对于一个 Manacher字符串，如下所示:
> i    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1 2 3 4 5 6 7 8 9  
> s    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# a # a # b # a #   
> d[i] &nbsp;&nbsp;1 2 3 2 1 4 1 2 1

以s[6]所指向的byte `b` 为例，以它为中心的最长回文子串为 `#a#b#a#`,左右分别有3个字符，因此定义他的回文半径为 4 (包含自身，即 3 + 1)。
因此，回文半径定义为`最长回文子串`的长度的一半。
<center>maxLen = d[i] * 2 - 1 </center>

## 1.3 加速盒子
在算法过程中，我们维护右端点最靠右的最长回文串。假设盒子的范围为[l, r)，其中l表示左端点，r表示右端点，左闭右开。
> i    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1 2 3 4 5 6 7 8 9  
> s    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# a # a # b # a #   
> d[i] &nbsp;&nbsp;1 2 3 2 1 4 1 2 1    

同样以 1.2 中的字符串举例，有以下的流程: 
- 初始化 d[1] = 1
- i = 2，易得 d[2] = 2，因此盒子区间为[1, 3)
- i = 3，不在在盒子区间[1, 3)内，因此对该点进行暴力搜索，得到 d[3] = 3，此时盒子区间为[1, 5)
- i = 4，在盒子区间[1, 5)内，因此直接对称赋值，d[4] = d[2] = 2
- i = 5，不在盒子区间[1, 5)内，因此对该点进行暴力搜索，得到 d[5] = 1，由于 d[5] < d[(l + r) / 2]，因此盒子区间不更新
- 以此类推，省略...

在了解了以上的概念之后，能够轻易写出Manacher算法代码，以Go语言胃为例: 