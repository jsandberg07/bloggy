package main

import (
	"fmt"
	"internal/config"
)

func main() {
	fmt.Println("Hello borld")

	cfg := config.Read()
	cfg.SetUser("Beans")
	cfg2 := config.Read()
	fmt.Println(cfg2)

}
