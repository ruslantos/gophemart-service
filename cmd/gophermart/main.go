package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("gophermart")
	_, err := zap.NewDevelopment()
	if err != nil {
	}
}
