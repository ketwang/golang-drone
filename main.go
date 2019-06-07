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
	fmt
}
