# 哈希表

当我们需要查询一个元素是否出现过，或者一个元素是否在集合里的时候，就要第一时间想到哈希法。


## 解题技巧

### 提供快速查找

哈希表的理想效率是O(1)，在一个集合中查找元素数组的效率是O(n)，使用哈希表可以显著的提高效率，适合经常在一个集合中查找元素的题目中

### 哈希表value为struct{}

对于只想知道是否出现过元素，但不想设置值的时候可以设置value为struct{}
