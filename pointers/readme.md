# Problems
1. How do you get the memory address of a variable?
2. How do you assign a value to a pointer?
3. How do you create a new pointer?
4. What is the value of x after running this program:

    ```
    func square(x *float64) {
    *x = *x * *x
    }
    func main() {
    x := 1.5
    square(&x)
    }
    ```
5. Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

# Solutions:
1. We get the memory address of a variable using the & operator.

    ```
    package main

    func main() {
        x := 42
        var p *int
        p = &x
    }
    ```
2. We assign a value to a pointer by :
- Declaring the pointer to an int.
- Assigning the memory address of the variable to the pointer 
- Assigning the value to the memory location pointed to by the pointer

    ```
    package main

    func main() {
        x := 42
        var p *int
        p = &x
        *p = 25
    }
    ```
3. We can create a new pointer by :
- Using new Keyword
   ```
    package main

    func main() {
        p := new(int)
        *p = 25
    }
    ```
- Using Address-of Operator (&)

   ```
    package main

    func main() {
        x := 25
        p := &x
    }
    ```

4. the value of `x (1.5)` after running the program is square `x (2.25)` because the function square multiply `x` with itself, and store the result back at the memory location pointed to by `x`.

5. Check `program.go` file