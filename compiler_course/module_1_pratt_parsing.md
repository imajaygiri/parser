# Module 1: Pratt Parsing & Binding Power (BP) Demystified

Pratt parsing (also known as Top-Down Operator Precedence parsing) is an elegant way to parse expressions. It avoids the deeply nested grammar rules of traditional recursive descent parsers by using a table-driven approach based on **Binding Power (Precedence)**.

Let's break down the rules of binding power, how the main Pratt loop works, and how to solve your struct instantiation issues.

---

## 1. The Core Concept: NUD vs LED

A Pratt parser defines two kinds of denotation handlers for tokens:
*   **NUD (Null Denotation)**: Handlers for tokens that appear at the beginning of an expression (prefix operators, literals, variables, opening parentheses). They do not look at anything to their left.
    *   Examples: `-5`, `x`, `"hello"`, `( 2 + 3 )`.
*   **LED (Left Denotation)**: Handlers for tokens that appear in the middle or end of an expression (infix and postfix operators). They take the expression parsed so far (the `left` hand side) and extend it.
    *   Examples: `left + right`, `left * right`, `left.member`, `left(arguments)`.

---

## 2. The Pratt Parser Loop: How it Works

Look closely at your current `parse_expr` function in [expr.go](file:///Users/ajaygiri/code/golang/parser/src/parser/expr.go#L13-L36):

```go
func parse_expr(p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()
	nud_handler, ok := nud_lu[tokenKind]
	if !ok {
		panic(fmt.Sprintf("No Nud handler for [%s]\n", tokenKind.ToString()))
	}
	left := nud_handler(p) // Step 1: Parse the prefix/literal

	for bp < bp_lu[p.currentTokenKind()] { // Step 2: Loop while next token binds tighter
		tokenKind = p.currentTokenKind()
		led_handler, ok := led_lu[tokenKind]
		if !ok { ... }
		left = led_handler(p, left, bp_lu[tokenKind]) // Step 3: Combine left with right
	}
	return left
}
```

### The Inequality: `bp < bp_lu[p.currentTokenKind()]`
The parameter `bp` represents the **current binding power context**. It tells the parser: *"Keep parsing expressions to the right as long as the operators you encounter bind tighter (have a higher binding power) than `bp`."*

As soon as the next token has a binding power **less than or equal to** `bp`, the parser stops and returns what it has built so far.

### Visualizing Precedence: `2 + 3 * 4`
Let's trace `parse_expr(p, default_bp)` (where `default_bp` is `0`):
1.  `2` is parsed by `parse_primary_expr` (NUD). `left = 2`.
2.  Next token is `+`. `bp_lu[+]` is `5` (`additive`).
3.  Is `0 < 5`? Yes! We call `parse_binary_expr` for `+`.
4.  Inside `parse_binary_expr` for `+`, we advance past `+` and call `parse_expr(p, 5)` (`bp = 5`).
5.  In the recursive call `parse_expr(p, 5)`:
    *   `3` is parsed by NUD. `left = 3`.
    *   Next token is `*`. `bp_lu[*]` is `6` (`multiplicative`).
    *   Is `5 < 6`? Yes! We call `parse_binary_expr` for `*`.
    *   Inside `parse_binary_expr` for `*`, we advance past `*` and call `parse_expr(p, 6)`.
    *   `4` is parsed by NUD. `left = 4`.
    *   Next token is EOF or `;`. Its binding power is `0`.
    *   Is `6 < 0`? No. The loop ends.
    *   `*` returns `3 * 4`.
6.  Back in `+` handler, it returns `2 + (3 * 4)`.

---

## 3. When to Use Which Binding Power

Here are the rules of thumb for choosing the `bp` argument:

| Context | Recommended Binding Power | Rationale |
| :--- | :--- | :--- |
| **Statement Level** | `default_bp` (0) | You want to parse the entire expression, stopping only at statements boundaries like semicolons (`;`) or closing braces (`}`). |
| **Grouping/Parenthesis `( expr )`** | `default_bp` (0) | Inside parentheses, you want to parse any expression, no matter how low its precedence. The closing `)` will naturally stop the parser. |
| **Prefix Operators `-expr`** | `default_bp` (0) or `unary` | If you want prefix operators to bind tightly to the immediate operand, use `default_bp` for the RHS so that expressions like `-x + y` are parsed as `(-x) + y` (since `+` has higher bp than `default_bp`). |
| **Assignment RHS `a = expr`** | `assignment` | Right-associative operators should pass their own binding power to the RHS. This allows chained assignments like `a = b = 5` to parse as `a = (b = 5)`. |
| **Binary/Logical Operators `a + b`** | `bp_lu[operator]` | Pass the operator's own binding power. This ensures left-associativity (e.g. `1 + 2 + 3` parses as `(1 + 2) + 3`). |

---

## 4. Solving the Struct Instantiation Issue

In [expr.go](file:///Users/ajaygiri/code/golang/parser/src/parser/expr.go#L112), you are parsing struct field values like this:
```go
expr := parse_expr(p, logical)
```

### ❌ The Bug
By passing `logical` (bp = 3) to `parse_expr`, you are telling the parser:
*"Only parse expressions that bind tighter than logical operators. If you see anything with bp <= 3, stop immediately."*

This causes two massive bugs:
1.  **Logical operators fail**: If a user writes `x: a && b`, the parser will parse `a` and then stop at `&&` because `&&` has bp = 3 (which is not strictly greater than `logical` bp = 3).
2.  **Assignment fails**: If a user writes `x: a = 5`, the parser stops at `a` because `=` has bp = 2 (which is less than `logical` bp = 3).

### 💡 The Correct Way
Inside curly braces for a struct instantiation, each property value is separated by a comma `,`.

Since `comma` is a separator, and NOT a mathematical operator within the property expression, it is **not** registered in `led_lu` and its binding power in `bp_lu` defaults to `0` (or `default_bp`).

Therefore, you should parse the property value using:
```go
expr := parse_expr(p, default_bp) // Or assignment
```

Because `bp_lu[lexer.COMMA]` is `0` and `bp_lu[lexer.CLOSE_CURLY]` is `0`, calling `parse_expr(p, default_bp)` will parse any expression (even logical or assignments) and will naturally stop exactly at the `,` or `}`!

---

## 🎓 MIT Professor's Note: Map Default Values in Go
> "In Go, reading a non-existent key from a map returns the zero-value for that type. Because `bp_lu` stores `binding_power` (which is an `int` under the hood), any token you don't explicitly register via `led()` or `stmt()` defaults to `0` (which is `default_bp`). This is why separators like `;`, `,`, `)`, and `}` automatically terminate expression parsing if your active binding power is `default_bp`."

## 💻 SWE Tips & Tricks
1.  **Never hardcode binding power literals**: Always use the enum values (`default_bp`, `logical`, `assignment`).
2.  **Be careful with associative operations**:
    *   **Left-associative** (like `+`): Use `bp` (the operator's bp) in the recursive call.
    *   **Right-associative** (like `=`): Use `bp - 1` (or the operator's bp) in the recursive call depending on implementation, or just use the operator's bp.
3.  **Trace with Printfs**: If your parser goes into an infinite loop or panics, print the `currentTokenKind().ToString()` at the start of the `for` loop in `parse_expr`. It will immediately show you which token is causing the loop to fail to terminate.
