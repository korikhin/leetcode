## 205. Isomorphic Strings

Given two strings `s `and `t`, _determine if they are isomorphic._

Two strings `s` and `t` are isomorphic if the characters in `s` can be replaced to get `t`.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

<br>

### Example 1

```
Input: s = "egg", t = "add"
Output: true
```

### Example 2

```
Input: s = "foo", t = "bar"
Output: false
```

### Example 3

```
Input: s = "paper", t = "title"
Output: true
```

<br>

### Constraints

- $1 \leqslant$ `len(s)` $\leqslant 5 \cdot 10^4$
- `len(t) == len(s)`
- `s` and `t` consist of any valid ASCII characters.
