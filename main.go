package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{}
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("hahah")

	fmt.Println("jjjjjk")
	fmt.Println("hehheh")
	if check(11) {
		fmt.Println("success")
	}
}

func check(a int) bool {
	if a > 10 {
		return true
	}

	return false
}
