## 380. Insert Delete GetRandom $O(1)$

Implement the `RandomizedSet` class:

- `Constructor()` initializes the `RandomizedSet` structure.
- `Insert(val int) bool` inserts an item `val` into the set if not present. Returns `true` if the item was not present, `false` otherwise.
- `Remove(val int) bool` removes an item `val` from the set if present. Returns `true` if the item was present, `false` otherwise.
- `GetRandom() int` returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the **same probability** of being returned.

You must implement the functions of the class such that each function works in average $O(1)$ time complexity.

<br>

### Example 1

```
Input:
  ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
  [[], [1], [2], [2], [], [1], [2], []]

Output:
  [nil, true, false, true, 2, true, false, 2]

Explanation:
```

```go
  randomizedSet := Constructor()
  randomizedSet.Insert(1)   // Inserts 1 to the set. Returns true as 1 was inserted successfully.
  randomizedSet.Remove(2)   // Returns false as 2 does not exist in the set.
  randomizedSet.Insert(2)   // Inserts 2 to the set, returns true. Set now contains [1, 2].
  randomizedSet.GetRandom() // GetRandom() should return either 1 or 2 randomly.
  randomizedSet.Remove(1)   // Removes 1 from the set, returns true. Set now contains [2].
  randomizedSet.Insert(2)   // 2 was already in the set, so return false.
  randomizedSet.GetRandom() // Since 2 is the only number in the set, GetRandom() will always return 2.
```

<br>

### Constraints

- $-2^{31} \leqslant$ `val` $\leqslant 2^{31} - 1$
- At most $2 \cdot 10^5$ calls will be made to `Insert()`, `Remove()`, and `GetRandom()`.
- There will be **at least one** element in the data structure when `GetRandom()` is called.
