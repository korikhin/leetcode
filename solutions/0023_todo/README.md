## 23. Merge k Sorted Lists

You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

_Merge all the linked-lists into one sorted linked-list and return it_.

<br>

### Example 1

```
Input: lists = [[1, 4, 5], [1, 3, 4], [2, 6]]
Output: [1, 1, 2, 3, 4, 4, 5, 6]
Explanation: The linked-lists are:

[
  1 -> 4 -> 5,
  1 -> 3 -> 4,
  2 -> 6
]

merging them into one sorted list:

  1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6
```

### Example 2

```
Input: lists = []
Output: []
```

### Example 3

```
Input: lists = [[]]
Output: []
```

<br>

### Constraints

- `k == len(lists)`
- $0 \leqslant$ `k` $\leqslant 10^4$
- $0 \leqslant$ `len(lists[i])` $\leqslant 500$
- $-10^4 \leqslant$ `lists[i][j]` $\leqslant 10^4$
- `lists[i]` is sorted in **ascending order**.
- The sum of `len(lists[i])` will not exceed $10^4$.
