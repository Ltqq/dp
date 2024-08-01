## **1. 斐波那契数列**

# 动态规划

**[斐波那契数列](https://en.wikipedia.org/wiki/Fibonacci_sequence)** 模式在解决依赖于较小实例问题解的情况下非常有用。

有一个明显的递归关系，通常类似于经典的斐波那契数列 `F(n) = F(n-1) + F(n-2)`。

#### LeetCode题目：

- [LeetCode 70: 爬楼梯](https://leetcode.com/problems/climbing-stairs/description/)
- [LeetCode 509: 斐波那契数](https://leetcode.com/problems/fibonacci-number/description/)
- [LeetCode 746. 使用最小花费爬楼梯](https://leetcode.com/problems/min-cost-climbing-stairs/description/)

## 2. Kadane算法

**[Kadane算法](https://medium.com/@rsinghal757/kadanes-algorithm-dynamic-programming-how-and-why-does-it-work-3fd8849ed73d)** 主要用于解决最大子数组问题及其变种，这些问题要求在一维数字数组内优化一个连续子数组。

#### LeetCode题目：

- [LeetCode 53: 最大子序和](https://leetcode.com/problems/maximum-subarray/description/)
- [LeetCode 918: 环形子数组的最大和](https://leetcode.com/problems/maximum-sum-circular-subarray/description/)
- [LeetCode 152: 乘积最大子数组](https://leetcode.com/problems/maximum-product-subarray/description/)

## 3. 0/1背包

**[0/1背包](https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/)** 模式在以下情况下非常有用：

1. 你有一组物品，每个物品都有一个重量和一个价值。
2. 你需要选择这些物品的一个子集。
3. 对总重量（或其他资源）有一个约束。
4. 你希望最大化（或最小化）所选物品的总价值。
5. 每个物品只能选择一次（因此称为"0/1" - 你要么选它，要么不选）。

#### LeetCode题目：

- [LeetCode 416: 分割等和子集](https://leetcode.com/problems/partition-equal-subset-sum/description/)
- [LeetCode 494: 目标和](https://leetcode.com/problems/target-sum/description/)
- [LeetCode 1049. 最后一块石头的重量 II](https://leetcode.com/problems/last-stone-weight-ii/description/)

## 4. 无限背包

**[无限背包](https://www.geeksforgeeks.org/unbounded-knapsack-repetition-items-allowed/)** 模式在以下情况下非常有用：

1. 你有一组物品，每个物品都有一个重量和一个价值。
2. 你需要选择物品以最大化总价值。
3. 对总重量（或其他资源）有一个约束。
4. 你可以多次选择每个物品（不像0/1背包，每个物品只能选择一次）。
5. 每种物品的供应被视为无限。

#### LeetCode题目：

- [LeetCode 322: 零钱兑换](https://leetcode.com/problems/coin-change/description/)
- [LeetCode 518: 零钱兑换 II](https://leetcode.com/problems/coin-change-ii/description/)
- [LeetCode 279. 完全平方数](https://leetcode.com/problems/perfect-squares/description/)

## 5. 最长公共子序列 (LCS)

**[最长公共子序列](https://www.youtube.com/watch?v=NnD96abizww)** 模式在给定两个序列时非常有用，需要找到一个子序列，它以相同的顺序出现在两个给定序列中。

#### LeetCode题目：

- [LeetCode 1143: 最长公共子序列](https://leetcode.com/problems/longest-common-subsequence/description/)
- [LeetCode 583: 两个字符串的删除操作](https://leetcode.com/problems/delete-operation-for-two-strings/description/)
- [LeetCode 1092: 最短公共超序列](https://leetcode.com/problems/shortest-common-supersequence/description/)

## 6. 最长递增子序列 (LIS)

**[最长递增子序列](https://www.youtube.com/watch?v=CE2b_-XfVDk)** 模式用于解决涉及找到一个序列中元素按递增顺序排列的最长子序列的问题。

#### LeetCode题目：

- [LeetCode 300: 最长递增子序列](https://leetcode.com/problems/longest-increasing-subsequence/description/)
- [LeetCode 673: 最长递增子序列的个数](https://leetcode.com/problems/number-of-longest-increasing-subsequence/description/)
- [leetCode 354: 俄罗斯套娃信封问题](https://leetcode.com/problems/russian-doll-envelopes/description/)

## 7. 回文子序列

**[回文子序列](https://www.youtube.com/watch?v=bUr8cNWI09Q)** 模式用于解决涉及在一个序列（通常是字符串）中找到一个前后读取相同的子序列的问题。

#### LeetCode题目：

- [LeetCode 516: 最长回文子序列](https://leetcode.com/problems/longest-palindromic-subsequence/description/)
- [LeetCode 647: 回文子串](https://leetcode.com/problems/palindromic-substrings/description/)
- [LeetCode 1312: 让字符串成为回文串的最少插入次数](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/description/)

## 8. 编辑距离

**[编辑距离](https://www.geeksforgeeks.org/edit-distance-dp-5/)** 模式用于解决涉及将一个序列（通常是字符串）转换为另一个序列，使用最小操作次数的问题。

允许的操作通常包括插入、删除和替换。

#### LeetCode题目：

- [LeetCode 72: 编辑距离](https://leetcode.com/problems/edit-distance/description/)
- [LeetCode 583: 两个字符串的删除操作](https://leetcode.com/problems/delete-operation-for-two-strings/description/)
- [LeetCode 712: 两个字符串的最小ASCII删除和](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/description/)

## 9. 子集和

**[子集和](https://www.youtube.com/watch?v=s6FhG--P7z0)** 模式用于解决需要确定给定集合中的元素子集是否可以加和达到特定目标值的问题。

#### LeetCode题目：

- [LeetCode 416: 分割等和子集](https://leetcode.com/problems/partition-equal-subset-sum/description/)
- [LeetCode 494: 目标和](https://leetcode.com/problems/target-sum/description/)
- [LeetCode 698: 划分为k个相等的子集](https://leetcode.com/problems/partition-to-k-equal-sum-subsets/)

## 10. 字符串分割

**[字符串分割](https://www.youtube.com/watch?v=Sx9NNgInc3A)** 模式用于解决涉及将字符串分割为满足某些条件的较小子字符串的问题。

在以下情况下很有用：

- 你在处理涉及字符串或序列的问题。
- 问题需要将字符串拆分为子字符串或子序列。
- 你需要优化这些分割上的某些属性（例如，最小化成本，最大化价值）。
- 整体问题的解决方案可以从较小子字符串上的子问题解决方案构建。
- 需要考虑字符串分割的不同方式。

#### LeetCode题目：

- [LeetCode 139: 单词拆分](https://leetcode.com/problems/word-break/description/)
- [LeetCode 132. 分割回文串 II](https://leetcode.com/problems/palindrome-partitioning-ii/description/)
- [LeetCode 472: 连接词](https://leetcode.com/problems/concatenated-words/description/)

## 11. 卡塔兰数

**[卡塔兰数](https://en.wikipedia.org/wiki/Catalan_number)** 模式用于解决可以分解为较小、类似子问题的组合问题。

一些使用案例包括：

- 计算给定长度的**有效括号**表达式的数量。
- 计算可以用 `n`个节点形成的不同**二叉搜索树**的数量。
- 计算用 `n+2`条边划分多边形的不同方法的数量。

#### LeetCode题目：

- [LeetCode 96: Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/description/)
- [LeetCode 22: Generate Parentheses](https://leetcode.com/problems/generate-parentheses/description/)

## 12. 三角形路径和

**三角形路径和** 模式用于解决涉及通过三角形路径选择最小或最大值的问题。

#### LeetCode题目：

- [LeetCode 1039: Minimum Score Triangulation of Polygon](https://leetcode.com/problems/minimum-score-triangulation-of-polygon/description/)
- [LeetCode 312: Burst Balloons](https://leetcode.com/problems/burst-balloons/description/)
- [LeetCode 1000: Minimum Cost to Merge Stones](https://leetcode.com/problems/minimum-cost-to-merge-stones/description/)

## 13. 线性DP

**线性DP** 模式用于解决以下问题：

- 给定序列时寻找最长的按顺序排列的特定类型子序列。
- 优化序列上的路径。
- 选择满足特定条件的最优子集。

#### LeetCode题目：

- [LeetCode 91: Decode Ways](https://leetcode.com/problems/decode-ways/description/)
- [LeetCode 2266. Count Number of Texts](https://leetcode.com/problems/count-number-of-texts/description/)

## 14. 区间DP

**区间DP** 模式用于解决涉及从序列的某些区间获取最优值的问题。

#### LeetCode题目：

- [LeetCode 62: Unique Paths](https://leetcode.com/problems/unique-paths/description/)
- [LeetCode 64: Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/description/)
- [LeetCode 329. Longest Increasing Path in a Matrix](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/)

## 15. 位运算DP

**位运算DP** 模式用于解决涉及组合优化和集合操作的问题。

#### LeetCode题目：

- [LeetCode 337: House Robber III](https://leetcode.com/problems/house-robber-iii/description/)
- [LeetCode 124: Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/description/)
- [LeetCode 968: Binary Tree Cameras](https://leetcode.com/problems/binary-tree-cameras/description/)

## 16. 图上的动态规划

**[图上的动态规划](https://medium.com/@logicdevildotcom/dynamic-programming-applied-to-graphs-f33b6b8a8247)** 模式在以下情况下非常有用：

1. 你在处理涉及图结构的问题。
2. 问题需要找到图上的最优路径、最长路径、循环或其他优化属性。
3. 你需要基于节点或边的邻居或连接组件计算值。
4. 问题涉及遍历图，同时维护一些状态。

#### LeetCode题目：

- [LeetCode 787: K站中转内最便宜的航班](https://leetcode.com/problems/cheapest-flights-within-k-stops/description/)
- [LeetCode 1334: 阈值距离内邻居最少的城市](https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/description/)

## 17. 数位DP

**[数位DP](https://www.geeksforgeeks.org/digit-dp-introduction/)** 模式在以下情况下非常有用：

1. 你在处理涉及范围内的数字计数或求和的问题。
2. 问题需要单独考虑数字的每一位。
3. 你需要找到与范围内数字的数字相关的模式或属性。
4. 数字的范围很大（通常高达10^18或更多），使得暴力方法不可行。
5. 问题涉及对数字的限制。

#### LeetCode题目：

- [LeetCode 357: 统计各位数字都不同的数字个数](https://leetcode.com/problems/count-numbers-with-unique-digits/description/)
- [LeetCode 233: 数字1的个数](https://leetcode.com/problems/number-of-digit-one/description/)
- [LeetCode 902: 最大为N的数字组合](https://leetcode.com/problems/numbers-at-most-n-given-digit-set/description/)

## 18. 位运算DP

**[位运算DP](https://www.hackerearth.com/practice/algorithms/dynamic-programming/bit-masking/tutorial/)** 模式用于解决涉及大量**状态**或**组合**的问题，其中每个状态都可以用整数中的**位**有效表示。

它特别有用在以下情况下：

1. 你在处理涉及**子集**或**组合**元素的问题。
2. 元素总数相对较少（通常<= 20-30）。
3. 你需要有效地表示和操作元素集。
4. 问题涉及为每个元素做出决策（包括/排除）或跟踪访问/未访问的状态。
5. 你想在DP解决方案中优化空间使用。

#### LeetCode题目：

- [LeetCode 1986: 完成任务的最少工作时间段](https://leetcode.com/problems/minimum-number-of-work-sessions-to-finish-the-tasks/description/)
- [LeetCode 2305: 公平分发饼干](https://leetcode.com/problems/fair-distribution-of-cookies/description/)
- [LeetCode 847: 访问所有节点的最短路径](https://leetcode.com/problems/shortest-path-visiting-all-nodes/description/)

## 19. 概率DP

这种模式在以下情况下非常有用：

1. 你在处理涉及概率计算的问题。
2. 状态的概率取决于先前状态的概率。
3. 你需要计算结果的期望值。
4. 问题涉及随机过程或机会游戏。

#### LeetCode题目：

- [LeetCode 688: “马”在棋盘上的概率](https://leetcode.com/problems/knight-probability-in-chessboard/description/)
- [LeetCode 808: 分汤](https://leetcode.com/problems/soup-servings/description/)
- [LeetCode 837: 新21点](https://leetcode.com/problems/new-21-game/description/)

## 20. 状态机DP

**[状态机DP](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solutions/108870/most-consistent-ways-of-dealing-with-the-series-of-stock-problems/)** 模式在以下情况下非常有用：

1. 问题可以建模为一系列状态和这些状态之间的转换。
2. 有明确的规则从一个状态移动到另一个状态。
3. 最优解依赖于做出最佳的状态转换序列。
4. 问题涉及做出影响未来状态的决策。
5. 存在有限数量的可能状态，并且每个状态可以明确定义。

#### LeetCode题目：

- [LeetCode 309: 最佳买卖股票时机含冷冻期](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/)
- [LeetCode 123: 买卖股票的最佳时机 III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/)

## File Name：LeetCode number + DP type
