package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	algorithm := &Algorithm{}

	sumCommand := NewSumCommand(algorithm)
	multiplyCommand := NewMultiplyCommand(algorithm)
	primeCommand := NewPrimeCommand(algorithm)
	fibCommand := NewFibCommand(algorithm)

	commandSwitch := NewCommandSwitch()
	commandSwitch.registerCommand("sum", sumCommand)
	commandSwitch.registerCommand("multiply", multiplyCommand)
	commandSwitch.registerCommand("prime", primeCommand)
	commandSwitch.registerCommand("fib", fibCommand)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Specify the command: ")
		scanner.Scan()
		command := scanner.Text()
		if (command == "exit") {
			os.Exit(0)
		}

		fmt.Print("Input: ")
		scanner.Scan()
		parts := strings.Split(scanner.Text(), " ")
		var input []int
		for _, v := range parts {
			i, _ := strconv.Atoi(v)
			input = append(input, i)
		}

		if err := commandSwitch.runCommand(command, input); err != nil {
			fmt.Println(err)
		}
	}
}
