## 219. Contains Duplicate II

Given an integer array `nums` and an integer `k`, return _`true` if there are two **distinct indices** `i` and `j` in the array such that `nums[i] == nums[j]` and_ $|i - j| \leqslant k$.

<br>

### Example 1

```
Input: nums = [1, 2, 3, 1],  k = 3
Output: true
```

### Example 2

```
Input: nums = [1, 0, 1, 1],  k = 1
Output: true
```

### Example 3

```
Input: nums = [1, 2, 3, 1, 2, 3],  k = 2
Output: false
```

<br>

### Constraints

- $1 \leqslant$ `nums.length` $\leqslant 10^5$
- $-10^9 \leqslant$ `nums[i]` $\leqslant 10^9$
- $0 \leqslant$ `k` $\leqslant 10^5$
