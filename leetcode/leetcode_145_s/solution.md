# 二叉树的后序遍历问题 
递归版太简单，不进行赘述。  

后序遍历相较于前序遍历的迭代版稍微复杂了一些，需要额外开辟一个visited，记录走过的非叶子结点，具体思路如下:
1. 遍历根节点所有的左节点，并依次压入栈中
2. 如果栈的长度不为0，则循环开始，从栈中取出一个节点
3. 如果该节点与visited栈栈顶节点不相同，则依进入步骤4
4. 若不存在右节点，则进入步骤5。若存在右节点，遍历该节点右节点的所有左节点，并将右节点和右节点下所有左节点依次压入栈中，并且记录该节点至visited栈中，回到步骤2
5. 如果不存在右节点，则保存至结果数据，该节点出栈，回到步骤2  

![img.png](img.png)