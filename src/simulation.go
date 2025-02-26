package main

import (
	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Simulation struct {
	current_monitor int
	monitor_height  int
	monitor_width   int
	screen_height   int
	screen_width    int

	navbar     *NavBar
	initilized bool
	debug_mode bool
}

func newSimulation(debug bool) *Simulation {

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 600, "Path Finding Visualization by Nikos Gournakis")

	rl.SetTargetFPS(60)

	if !rl.IsWindowMaximized() {
		rl.MaximizeWindow()
	}

	this := &Simulation{initilized: false, debug_mode: debug}

	log.Debug("%#v", this)
	return this
}

func (this *Simulation) init() {
	if this.navbar == nil {
		this.navbar = newNavBar(this)
	}
	this.initilized = true
}

func (this *Simulation) runGameLoop() {
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		this.configureMonitorScreenSizes()
		// log.Debug("monitor: %v, monitor_height: %v, monitor_width: %v", this.current_monitor, this.monitor_height, this.monitor_width)
		// log.Debug("screen_height: %v, screen_width: %v", this.screen_height, this.screen_width)

		rl.BeginDrawing()
		// ---------------- DRAWING ----------------------------
		// Wait for window initilization
		log.Debug("monitor: %v, monitor_height: %v, monitor_width: %v", this.current_monitor, this.monitor_height, this.monitor_width)
		log.Debug("screen_height: %v, screen_width: %v", this.screen_height, this.screen_width)
		this.configureMonitorScreenSizes()
		if !this.initilized && this.screen_height > 600 && this.screen_width > 800 {
			this.init()
		} else if !this.initilized && (this.monitor_height == 600 || this.monitor_width == 800) {
			this.init()
		}
		// end of initilization
		if this.initilized {
			this.navbar.draw()
		}
		rl.ClearBackground(rl.NewColor(0x18, 0x18, 0x18, 0xFF))
		// ---------------- END DRAWING ------------------------
		rl.EndDrawing()
	}
}

func (this *Simulation) configureMonitorScreenSizes() {
	this.current_monitor = rl.GetCurrentMonitor()
	this.monitor_height = rl.GetMonitorHeight(this.current_monitor)
	this.monitor_width = rl.GetMonitorWidth(this.current_monitor)

	this.screen_height = rl.GetScreenHeight()
	this.screen_width = rl.GetScreenWidth()
}
