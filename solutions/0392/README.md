## 392. Is Subsequence

Given two strings `s` and `t`, return _`true` if `s` is a **subsequence** of `t`, or `false` otherwise._

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., `"ace"` is a subsequence of `"abcde"` while `"aec"` is not).

<br>

### Example 1

```
Input: s = "abc", t = "ahbgdc"
Output: true
```

### Example 2

```
Input: s = "axc", t = "ahbgdc"
Output: false
```

<br>

### Constraints

- $0 \leqslant$ `len(s)` $\leqslant 100$
- $0 \leqslant$ `len(t)` $\leqslant 10^4$
- `s` and `t` consist only of lowercase English letters.

<br>

**Follow up:** Suppose there are lots of incoming `s`, say `s1, s2, ..., sk` where `k` $\geqslant 10^9$, and you want to check one by one to see if `t` has its subsequence. In this scenario, how would you change your code?
