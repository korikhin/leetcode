## 137. Single Number II

Given an integer array `nums` where every element appears **three times** except for one, which appears **exactly once**. _Find the single element and return it_.

You must implement a solution with a linear runtime complexity and use only constant extra space.

<br>

### Example 1

```
Input: nums = [2, 2, 3, 2]
Output: 3
```

### Example 2

```
Input: nums = [0, 1, 0, 1, 0, 1, 99]
Output: 99
```

<br>

### Constraints

- $1 \leqslant$ `len(nums)` $\leqslant 3 \cdot 10^4$
- $-2^{31} \leqslant$ `nums[i]` $\leqslant 2^{31} - 1$
- Each element in `nums` appears exactly **three times** except for one element which appears **once**.
