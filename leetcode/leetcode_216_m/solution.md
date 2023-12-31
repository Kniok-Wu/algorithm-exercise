# 组合总和 III问题
回溯算法类的题目都大同小异，主要的难点是在于掌握回溯过程中的边界以及剪枝条件，我的剪枝思路是:  
前提: 回溯函数传入
- 所需 n 值
- 所需元素 k 个
- 剩余元素 left []int

结果条件:  
当减掉当前left的对应值后，n为0且k为0，跳出当前递归，直接返回终止。

剪枝思路如下，当且仅当下列两项均满足时，开启下一轮迭代:
- 减掉当前值后，剩余 n 是否大于当前值，若是且不满足终止条件，进入步骤二
- 减掉当前之后，k - 1，k是否大于0，若是且不满足终止条件，进入下一次迭代

> 其中由于 left 的长度存在 `for` 循环的遍历，因此无需在剪枝时skip，下次迭代会什么都不做直接skip。  

![img.png](img.png)
 
