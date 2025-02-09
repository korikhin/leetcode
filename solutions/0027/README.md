## 27. Remove Element

Given an integer array `nums` and an integer `val`, remove all occurrences of `val` in `nums` **in-place**. The order of the elements may be changed. Then return _the number of elements in_ `nums` _which are not equal to_ `val`.

Consider the number of elements in `nums` which are not equal to `val` be `k`, to get accepted, you need to do the following things:

- Change the array `nums` such that the first `k` elements of `nums` contain the elements which are not equal to `val`. The remaining elements of `nums` are not important as well as the size of `nums`.
- Return `k`.

### Custom Judge

The judge will test your solution with the following code:

```go
nums := []int{...}               // Input array
val := ...                       // Value to remove
expectedNums := []int{...}       // The expected answer with correct length

k := removeDuplicates(nums, val) // Calls your implementation
if k != len(expectedNums) {
	panic()
}

sort.Ints(nums)
for i := 0; i < k; i++ {
	if nums[i] != expectedNums[i] {
		panic()
	}
}
```

If all assertions pass, then your solution will be **accepted**.

<br>

### Example 1

```
Input: nums = [3, 2, 2, 3], val = 3
Output: 2, nums = [2, 2, _, _]
Explanation:
  Your function should return k = 2, with the first two elements of nums being 2.
  It does not matter what you leave beyond the returned k (hence they are underscores).
```

### Example 2

```
Input: nums = [0, 1, 2, 2, 3, 0, 4, 2], val = 2
Output: 5, nums = [0, 1, 4, 0, 3, _, _, _]
Explanation:
  Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
  Note that the five elements can be returned in any order.
  It does not matter what you leave beyond the returned k (hence they are underscores).
```

<br>

### Constraints

- $0 \leqslant$ `len(nums)` $\leqslant 100$
- $0 \leqslant$ `nums[i]` $\leqslant 50$
- $0 \leqslant$ `val` $\leqslant 100$
