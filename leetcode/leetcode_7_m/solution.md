# 整数反转问题
这道题的难点在于32位下的溢出检测，我们知道两个整数相加如果溢出，结果则是负数，因此我们可以在开始时记录一个**负数标志位**，这样就可以保证
值为正，在过程中判断是否出现负数，如果出现说明出现异常。
![img.png](img.png)