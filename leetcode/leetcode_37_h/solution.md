# 解数独
这一道题是**N皇后问题**的进阶版本，可以使用box、line、col三个bool数组进行一个状态记录，判断是否可以选择该数字，然后广度优先向下递归即可。  
在这期间我还发现，哈希表的效率严重低于数组，因此如果能用数组就尽量用数组，图一和图二分别是使用哈希表的时间和数组的时间。  
![img.png](img.png)

![img_2.png](img_2.png)

**\[TIPS\]** 这道题的状态记录还可以使用**bitmap**，将数组压缩至一个数，进一步节省空间。