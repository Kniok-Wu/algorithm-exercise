# 打家劫舍II
这道题相较于打家劫舍I而言，最大的难点在于将首尾相连，出现了无法使用I中的递推公式。  
因此我们需要想办法打断首尾相连的情况，根据题意，假设选择了首部，那么就无法选择尾部，反之亦然，因此，我们可以人为分成两个子递归问题，即**不包含尾部**和**不包含首部**，再分别调用打家劫舍I中的方法进行求解。  

![img.png](img.png)