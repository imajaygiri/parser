# MIT 6.S081: Building a Fully Functional Compiler (Pratt to LLVM)

Welcome to this hands-on course! This repository is your laboratory. Since you learn best by coding, we will walk through the compiler frontend and backend concepts using your existing codebase as a reference.

This course is structured into four distinct modules:

| Module | Focus Area | Core Concepts |
| :--- | :--- | :--- |
| **[Module 1: Pratt Parsing & Binding Power](file:///Users/ajaygiri/code/golang/parser/compiler_course/module_1_pratt_parsing.md)** | Expressions | Pratt Parsing Loop, Binding Power (bp) rules, Struct Instantiation and properties parsing. |
| **[Module 2: Diagnostic Error Reporting](file:///Users/ajaygiri/code/golang/parser/compiler_course/module_2_error_handling.md)** | Error Handling | File, Line, and Column tracking; Pretty-printing compiler errors with caret pointers; Lexer vs Parser errors. |
| **[Module 3: Functions, Methods, & Classes](file:///Users/ajaygiri/code/golang/parser/compiler_course/module_3_complex_parsing.md)** | Advanced Syntax | Scope, block statements, global functions, class/struct declarations, member methods, and variable scopes. |
| **[Module 4: Mapping AST to LLVM Backend](file:///Users/ajaygiri/code/golang/parser/compiler_course/module_4_llvm_preparation.md)** | Code Generation | Symbol tables, semantic analysis, LLVM IR structure, Go-LLVM bindings, and what LLVM handles for you. |
| **[Module 5: Manual Lexing (DFA State Machines)](file:///Users/ajaygiri/code/golang/parser/compiler_course/module_5_manual_lexer.md)** | Tokenization | Eliminating Regex, DFA transitions in code, peeking/advancing, parsing identifiers, numbers, strings with escape sequences. |

---

### 🎓 MIT Professor's Note
> "A compiler is not a single giant program; it is a pipeline of elegant, cooperating transformations. Do not get overwhelmed by the final LLVM code generator. Master the AST (Abstract Syntax Tree) first. If your AST is malformed or lacks critical semantic info, code generation is impossible. Pay close attention to binding power—it is the difference between parsing `2 + 3 * 4` as `2 + (3 * 4)` or `(2 + 3) * 4`."

### 💻 SWE Tips & Tricks
> "When building a compiler, **never debug in your head**. Write micro-tests for every language feature as soon as you parse it. Make your AST nodes easy to print (using dump libraries like `litter`, which you already use, or custom JSON representation). This saves you hundreds of hours."

Ready to dive in? Click on **[Module 1: Pratt Parsing & Binding Power](file:///Users/ajaygiri/code/golang/parser/compiler_course/module_1_pratt_parsing.md)** to start.
