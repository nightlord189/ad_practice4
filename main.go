package main

import (
	"fmt"
	"github.com/nightLord189/ad_practice4/config"
)

func main() {
	fmt.Println("start")
	config.Load("config.json")
}
