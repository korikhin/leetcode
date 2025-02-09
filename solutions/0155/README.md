## 155. Min Stack

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the `MinStack` struct:

- `Constructor()` initializes the stack object.
- `Push(x int)` pushes the element val onto the stack.
- `Pop()` removes the element on the top of the stack.
- `Top()` gets the top element of the stack.
- `Min()` retrieves the minimum element in the stack.

You must implement a solution with $O(1)$ time complexity for each function.

<br>

### Example 1

```
Input:
  ["MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"]
  [[], [-2], [0], [-3], [], [], [], []]

Output:
  [nil, nil, nil, nil, -3, nil, 0, -2]

Explanation:
```

```go
  minStack := Constructor()
  minStack.Push(-2)
  minStack.Push(0)
  minStack.Push(-3)
  minStack.Min()    // return -3
  minStack.Pop()
  minStack.Top()    // return 0
  minStack.Min()    // return -2
```

<br>

### Constraints

- $-2^{31} \leqslant$ `x` $\leqslant 2^{31} - 1$
- Methods `Pop`, `Top` and `Min` operations will always be called on non-empty stacks.
- At most $3 \cdot 10^4$ calls will be made to `Push`, `Pop`, `Top`, and `Min`.
