package main

import (
	"fmt"

	"github.com/alexandre-j95/gator_cli/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	err = cfg.SetUser("Alexandre")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	newcfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%+v", newcfg)
}
