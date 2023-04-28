package main

import (
	"Europe/internal/cmd"
	cache "Europe/internal/cache"
	
)

func main() {
	cache.Init()
	for{
    	cmd.ParseConsole()
	}
}