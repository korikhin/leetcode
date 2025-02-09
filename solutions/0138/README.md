## 138. Copy List with Random Pointer

A linked list of length `n` is given such that each node contains an additional random pointer, which could point to any node in the list, or `nil`.

Construct a **deep copy** of the list. The deep copy should consist of exactly `n` brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the `Next` and `Random` pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. **None of the pointers in the new list should point to nodes in the original list.**

For example, if there are two nodes `X` and `Y` in the original list, where `X.Random` $\to$ `Y`, then for the corresponding two nodes `x` and `y` in the copied list, `x.Random` $\to$ `y`.

Return _the head of the copied linked list._

The linked list is represented in the input/output as a list of `n` nodes. Each node is represented as a pair of `[val,  random_index]` where:

- `val`: an integer representing `Node.Val`
- `random_index`: the index of the node (range from `0` to `n - 1`) that the `Random` pointer points to, or `nil` if it does not point to any node.

Your code will **only** be given the `head` of the original linked list.

<br>

### Example 1

```
Input:  [[7, nil], [13, 0], [11, 4], [10, 2], [1, 0]]
Output: [[7, nil], [13, 0], [11, 4], [10, 2], [1, 0]]
```

### Example 2

```
Input:  [[1, 1], [2, 1]]
Output: [[1, 1], [2, 1]]
```

### Example 3

```
Input:  [[3, nil], [3, 0], [3, nil]]
Output: [[3, nil], [3, 0], [3, nil]]
```

<br>

### Constraints

- $0 \leqslant$ `n` $\leqslant 1000$
- $-10^4 \leqslant$ `Node.Val` $\leqslant 10^4$
- `Node.Random` is `nil` or is pointing to some node in the linked list.
