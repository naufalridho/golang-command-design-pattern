**Command Design Pattern in Go**
----  
  Example of command design pattern implementation in Go. It is basically a collection 
  of simple algorithm problems that are organized using command pattern. There are four kinds of
  commands that can be used: "sum", "multiply", "prime", and "fib".
  
  - Sum X & Y, and print the result
  ```
  Input : 1 2
  Output : 3
  ```
  - Multiply X & Y, and print the result
  ```
  Input : 1 2
  Output: 2
  ```
  - Find first N prime number, and print the result
  ```
  Input: 4
  Output : 2 3 5 7
  ```
  - Find the first N Fibonacci sequence, and print the result
  ```
  Input: 4
  Output : 0 1 1 2 
  ```
  User will be prompted to specify the command and its input, for example:
  ```
  Specify the command: sum
  Input: 1 2
  3

  Specify the command: prime
  Input: 4
  2 3 5 7
  ```
  To exit the program, user can enter "exit" command

  ----
  
  To run the code, simply run this following command:
  ```
  $ go run main.go
  ```