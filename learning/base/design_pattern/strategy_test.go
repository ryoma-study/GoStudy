package test

import (
	"fmt"
	"testing"
)

// 策略模式
// 定义一个策略类
type IStrategy interface {
	do(int, int) int
}

// 策略实现：加
type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

// 策略实现：减
type reduce struct{}

func (*reduce) do(a, b int) int {
	return a - b
}

// 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}

func TestStrategy(t *testing.T) {
	operator := Operator{}

	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println("add:", result)

	operator.setStrategy(&reduce{})
	result = operator.calculate(2, 1)
	fmt.Println("reduce:", result)
}
