package main

import (
	"fmt"
	"strings"
)

type Command interface {
	execute(params interface{})
	getResult() interface{}
	display()
}

/* Sum command */
type sumCommand struct {
	algorithm *Algorithm
	result int
}

func NewSumCommand(al *Algorithm) Command {
	return &sumCommand{algorithm: al}
}

func (c *sumCommand) execute(params interface{}) {
	nums := params.([]int)
	res := c.algorithm.sum(nums[0], nums[1])
	c.result = res
}

func (c *sumCommand) getResult() interface{} {
	return c.result
}

func (c *sumCommand) display() {
	fmt.Printf("%d\n", c.result)
}

/* Multiply command */
type multiplyCommand struct {
	algorithm *Algorithm
	result int
}

func NewMultiplyCommand(al *Algorithm) Command {
	return &multiplyCommand{algorithm: al}
}

func (c *multiplyCommand) execute(params interface{}) {
	nums := params.([]int)
	res := c.algorithm.multiply(nums[0], nums[1])
	c.result = res
}

func (c *multiplyCommand) getResult() interface{} {
	return c.result
}

func (c *multiplyCommand) display() {
	fmt.Printf("%d\n", c.result)
}

/* Prime command */
type primeCommand struct {
	algorithm *Algorithm
	result []int
}

func NewPrimeCommand(al *Algorithm) Command {
	return &primeCommand{algorithm: al}
}

func (c *primeCommand) execute(params interface{}) {
	nums := params.([]int)
	res := c.algorithm.prime(nums[0])
	c.result = res
}

func (c *primeCommand) getResult() interface{} {
	return c.result
}

func (c *primeCommand) display() {
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.result)), " "), "[]"))
}

/* Fibonacci command */
type fibCommand struct {
	algorithm *Algorithm
	result []int
}

func NewFibCommand(al *Algorithm) Command {
	return &fibCommand{algorithm: al}
}

func (c *fibCommand) execute(params interface{}) {
	nums := params.([]int)
	res := c.algorithm.fib(nums[0])
	c.result = res
}

func (c *fibCommand) getResult() interface{} {
	return c.result
}

func (c *fibCommand) display() {
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.result)), " "), "[]"))
}