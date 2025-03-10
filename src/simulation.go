package main

import (
	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Simulation struct {
	desired_monitor int
	current_monitor int
	monitor_height  int
	monitor_width   int
	screen_height   int
	screen_width    int

	navbar *NavBar
	grid   *Grid

	initilized bool
	debug_mode bool
}

func newSimulation(debug bool) *Simulation {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	log.Debug("%#v", rl.GetMonitorCount())

	monitor := 0
	monitor_width := int32(1920)  //int32(rl.GetMonitorWidth(monitor))
	monitor_height := int32(1080) //int32(rl.GetMonitorHeight(monitor))
	log.Debug("%#v,%#v,%#v", monitor, monitor_width, monitor_height)

	rl.InitWindow(
		monitor_width,
		monitor_height,
		"Path Finding Visualization by Nikos Gournakis")

	rl.SetTargetFPS(60)

	rl.SetWindowPosition(int(rl.GetMonitorPosition(monitor).X), int(rl.GetMonitorPosition(monitor).Y))
	rl.SetWindowMonitor(monitor)

	// rl.ToggleFullscreen()

	this := &Simulation{desired_monitor: monitor, initilized: false, debug_mode: debug}
	this.configureMonitorScreenSizes()
	log.Debug("%#v", this)
	return this
}

func (this *Simulation) init() {
	if !rl.IsWindowFullscreen() {
		rl.SetWindowSize(this.monitor_width, this.monitor_height)
		rl.ToggleFullscreen()
		this.configureMonitorScreenSizes()
		// rl.ClearWindowState(rl.FlagWindowResizable)
	}

	if this.navbar == nil {
		this.navbar = newNavBar(this)
	}

	if this.grid == nil {
		this.grid = NewGrid(this)
	}

	go func() {
		for {
			if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				x, y, err := this.grid.mapScreenToGrid(rl.GetMouseX(), rl.GetMouseY())
				if err != nil {
					// log.Error("%s", err)
				} else {
					if this.grid.cells[y][x] != CellStateBorder {
						log.Debug("{%d,%d}", x, y)
						this.grid.cells[y][x] = CellStateBorder
					}
				}
			}
		}
	}()

	this.initilized = true
}

func (this *Simulation) runMainLoop() {
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		// Wait for window initilization
		// log.Debug("monitor: %v, monitor_height: %v, monitor_width: %v", this.current_monitor, this.monitor_height, this.monitor_width)
		// log.Debug("screen_height: %v, screen_width: %v", this.screen_height, this.screen_width)
		this.configureMonitorScreenSizes()
		if !this.initilized && this.current_monitor == this.desired_monitor {
			this.init()
		}
		// end of initilization

		if this.initilized {
		}
		rl.BeginDrawing()
		// ---------------- DRAWING ----------------------------
		if this.initilized {
			this.navbar.draw()
			this.grid.draw()
		}
		rl.ClearBackground(rl.NewColor(0x18, 0x18, 0x18, 0xFF))
		// ---------------- END DRAWING ------------------------
		rl.EndDrawing()
	}
}

func (this *Simulation) configureMonitorScreenSizes() {
	this.current_monitor = rl.GetCurrentMonitor()

	this.screen_height = rl.GetScreenHeight()
	this.screen_width = rl.GetScreenWidth()

	this.monitor_height = rl.GetMonitorHeight(this.current_monitor)
	this.monitor_width = rl.GetMonitorWidth(this.current_monitor)
}
