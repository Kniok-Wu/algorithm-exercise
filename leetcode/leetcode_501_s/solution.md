# 二叉搜索树中的众数问题
与大多数的二叉搜索树的思路一致，使用中序遍历，因为二叉搜索树在中序遍历下一定是有序的，因此使用prev和cur两个指针以前以后的比较，在一个数组中记录，每次遇到不同的数字时同时有2个及以上的记录时，进行最后一个数字和倒数第二个数字的比较，如果最后一个较大，则删除前方所有的元素，最后一个较小，则删除最后一个元素。  

![img.png](img.png)