## 34. Find First and Last Position of Element in Sorted Array

Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with $O(\log n)$ runtime complexity.

<br>

### Example 1

```
Input: nums = [5, 7, 7, 8, 8, 10], target = 8
Output: [3, 4]
```

### Example 2

```
Input: nums = [5, 7, 7, 8, 8, 10], target = 6
Output: [-1, -1]
```

### Example 3

```
Input: nums = [], target = 0
Output: [-1, -1]
```

<br>

### Constraints

- $0 \leqslant$ `len(nums)` $\leqslant 10^5$
- $-10^9 \leqslant$ `nums[i]` $\leqslant 10^9$
- $-10^9 \leqslant$ `target` $\leqslant 10^9$
- `nums` is a non-decreasing array.
