package main

import (
	"fmt"

	"github.com/sifer169966/device-interactions/cmd"
	"github.com/sifer169966/device-interactions/pkg/flags"
)

func main() {
	fmt.Printf(`
    ___       ______    _____
   / _  \    |   _  \  |_   _|
  /  __  \   |   __ /   _ | _
 /__/  \__\  |__|      |_____|

   github.com/sifer169966/device-interactions %s, built with Go %s
 `, flags.Version, flags.GoVersion)
	cmd.Execute()
}
