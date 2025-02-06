Delve is a powerful debugger for Go programs, allowing you to inspect, control, and step through Go code in a command-line interface (CLI) or with integrations in IDEs. Here’s a detailed guide on how to debug a Go application using Delve’s CLI.

### **1. Installing Delve**

Before using Delve, ensure you have it installed. You can install Delve using the following command:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

This installs the `dlv` (Delve) binary. You can check if it’s installed correctly by running:

```bash
dlv version
```

### **2. Starting Delve Debugger**

To debug a Go program, you need to start the program with Delve. Here’s a basic command to run it:

```bash
dlv debug <path-to-your-go-file-or-package>
```

For example, if your Go program is `main.go`:

```bash
dlv debug main.go
```

This command compiles and runs your Go program in debug mode. It will open a Delve CLI prompt where you can issue debugging commands.

### **3. Common Delve CLI Commands**

Here are the most commonly used Delve commands during debugging:

#### **a. Breakpoints**

A breakpoint is where the execution of the program will pause. You can set breakpoints at functions, specific lines of code, or even based on conditions.

- **Set a breakpoint at a specific line**:

```bash
break <filename>:<line number>
```

For example, to set a breakpoint at line 10 in `main.go`:

```bash
break main.go:10
```

- **Set a breakpoint at the start of a function**:

```bash
break <function-name>
```

For example, to set a breakpoint at the start of the `main` function:

```bash
break main
```

- **Set a conditional breakpoint**:

```bash
break <filename>:<line number> if <condition>
```

For example, to break at line 10 only when a variable `x` equals 5:

```bash
break main.go:10 if x == 5
```

#### **b. Running the Program**

Once Delve is launched, you can start your program with the `continue` (or `c`) command:

```bash
continue
```

This will start the execution of your program, and it will stop at the first breakpoint.

#### **c. Stepping through the Code**

While the program is paused at a breakpoint, you can control the flow of execution:

- **Step into** a function:

```bash
step
```

This will go into the function call on the current line.

- **Step over** a function:

```bash
next
```

This will execute the current line, but if the line contains a function call, it will execute the function without stepping into it.

- **Step out** of a function:

```bash
stepout
```

This will run the rest of the current function and stop when returning to the caller.

#### **d. Inspecting Variables**

You can inspect the values of variables and objects during debugging:

- **Print the value of a variable**:

```bash
print <variable-name>
```

For example, to print the value of variable `x`:

```bash
print x
```

- **Display the value of an expression**:

```bash
expr <expression>
```

For example, to evaluate the expression `x + y`:

```bash
expr x + y
```

- **Show all local variables**:

```bash
locals
```

This will list all local variables and their current values.

- **Display all global variables**:

```bash
globals
```

This will list all global variables.

#### **e. Inspecting Call Stack**

The call stack is useful to see the series of function calls leading to the current point in the program.

- **Print the call stack**:

```bash
backtrace
```

This shows a stack trace, including the function calls and line numbers.

#### **f. Modifying Variables**

You can modify the values of variables while debugging:

- **Set a variable**:

```bash
set <variable-name> = <new-value>
```

For example, to change the value of `x` to 10:

```bash
set x = 10
```

This is useful for simulating different scenarios without restarting the program.

#### **g. Exiting the Debugger**

To exit the Delve debugger:

```bash
quit
```

Alternatively, you can use `exit`.

### **4. Example Debugging Session**

Let’s walk through a typical debugging session using Delve.

Suppose you have the following Go code (`main.go`):

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    x := 3
    y := 4
    result := add(x, y)
    fmt.Println("The result is:", result)
}
```

Here’s how you might debug it:

1. **Start the Debugger**:

```bash
dlv debug main.go
```

2. **Set Breakpoint** at the `main` function:

```bash
break main
```

3. **Run the Program**:

```bash
continue
```

Delve will stop at the start of the `main` function.

4. **Step Through Code**:

- Step into the `add` function:

```bash
step
```

- You can then use `next` or `step` to continue stepping through the code.

5. **Print Variables**:

- Check the value of `x` and `y`:

```bash
print x
print y
```

- You can also inspect the result of the `add` function:

```bash
print result
```

6. **Modify Variables** (if needed):

You can change `x` or `y` values while debugging, for example:

```bash
set x = 10
```

7. **Quit Debugger**:

Once you are done:

```bash
quit
```

### **5. Advanced Debugging Features**

Delve also has more advanced features like:

- **Debugging remote applications**: You can use Delve to connect to a running Go program on a remote server.
- **Watchpoints**: You can set watchpoints to pause execution when a variable’s value changes.
- **Multi-threaded debugging**: Delve supports debugging concurrent Go programs that use goroutines.

### **Conclusion**

Delve is a powerful tool for Go developers, providing all the essential features for inspecting and controlling the flow of a Go program. Understanding how to use Delve's breakpoints, stepping commands, and variable inspection will make debugging much easier and more efficient.