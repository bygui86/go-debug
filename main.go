package main

import (
	"fmt"

	"github.com/bygui86/go-debug/domain"
)

func main() {

	fmt.Println("User 1:", domain.GetMatt())
	fmt.Println("User 2:", domain.GetJohn())
}
