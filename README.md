# expression-parsing

An implementation of two parsing algorithms.
1. Traditional Recursive Descent Parser.
2. Pratt's Top Down Operator Parser.

Expression syntax contains the following operators, with increasing levels of precedence.

| Precedence | Operators     |
| ---------- | ------------- |
| 0          | =             |
| 1          | \|\|          |
| 2          | &&            |
| 3          | \|            |
| 4          | ^             |
| 5          | &             |
| 6          | ==, !=        |
| 7          | >, >=, <, <=  |
| 8          | +, -          |
| 9          | >>, <<        |
| 10         | *, /, %       |
| 11         | !, ~, -(unary)|
| 12         | **            |
| 13         | ()            |
