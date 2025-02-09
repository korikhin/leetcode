## 32. Longest Valid Parentheses

Given a string containing just the characters `'('` and `')'`, return _the length of the longest valid (well-formed) parentheses substring_.

<br>

### Example 1

```
Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
```

### Example 2

```
Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
```

### Example 3

```
Input: s = ""
Output: 0
```

<br>

### Constraints

- $0 \leqslant$ `len(s)` $\leqslant 3 \cdot 10^4$
- `s[i]` is `'('`, or `')'`.
