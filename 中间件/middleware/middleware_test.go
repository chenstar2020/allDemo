package middleware

import (
	"fmt"
	"testing"
)

func A(c *Context){
	fmt.Println("part1")
	c.Next()
	fmt.Println("part2")
}

func B(c *Context){
	fmt.Println("part3")
	c.Next()
	fmt.Println("part4")
}

func C(c *Context){
	fmt.Println("part5")
	c.Next()
	fmt.Println("part6")
}
func TestContext(t *testing.T) {
	ctx := newContext()
	ctx.Use(A)
	ctx.Use(B)
	ctx.Use(C)

	ctx.Next()
}
