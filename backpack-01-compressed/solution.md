# 0-1背包问题 滚动解决
由于背包问题仅仅依赖于上一时间段的状态，因此可以压缩dp数组，减少空间的损耗，使用双数组保存当前状态和上一时刻状态即可。 [注释代码]
<img src="image.png">

## UPDATE [未注释代码]
由于从前往后可能存在多次选择dp数组前一部分导致出现多次选取当前物品的可能，那么当我从末尾往前求的时候，就可以避免这个问题，从而避免了数组的拷贝。
