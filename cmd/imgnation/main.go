package main

import (
	"fmt"
	"img/internal/config"
)

func main() {
	cfg := config.NewConfig("local")
	fmt.Println(cfg)
}
