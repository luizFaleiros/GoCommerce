package main

import (
	"fmt"
	"github.com/luizFaleiros/GoCommerce/configs"
)

func main() {
	config := configs.GetConfig()
	fmt.Println(config)
}
