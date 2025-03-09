package main

import (
	"fmt"

	"github.com/NikosGour/path_finding_visualization/src/build"
)

func main() {
	fmt.Printf("DEBUG_MODE = %t\n", build.DEBUG_MODE)
	simulation := newSimulation(build.DEBUG_MODE)
	simulation.runMainLoop()
}
