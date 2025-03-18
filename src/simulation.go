package main

import (
	_ "embed"
	"errors"
	"time"

	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets/Lexend-Regular.ttf
var default_font []byte

type Simulation struct {
	desired_monitor int
	current_monitor int
	monitor_height  int
	monitor_width   int
	screen_height   int
	screen_width    int

	default_font rl.Font

	// TODO : Change name of NavBar to AlgorithmHud
	navbar    *NavBar
	grid      *Grid
	color_hud *ColorHud

	initilized bool
	debug_mode bool
}

var (
	ErrorUnreachable = errors.New("Unreachable")
)

func newSimulation(debug bool) *Simulation {
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint | rl.FlagWindowHighdpi | rl.FlagMsaa4xHint)

	log.Debug("Monitor Count: %#v", rl.GetMonitorCount())

	monitor := 0
	monitor_width := int32(1920)  //int32(rl.GetMonitorWidth(monitor))
	monitor_height := int32(1080) //int32(rl.GetMonitorHeight(monitor))

	rl.InitWindow(
		monitor_width,
		monitor_height,
		"Path Finding Visualization by Nikos Gournakis")

	rl.SetTargetFPS(60)

	rl.SetWindowPosition(int(rl.GetMonitorPosition(monitor).X), int(rl.GetMonitorPosition(monitor).Y))
	rl.SetWindowMonitor(monitor)

	this := &Simulation{desired_monitor: monitor, initilized: false, debug_mode: debug}
	this.configureMonitorScreenSizes()

	this.default_font = rl.LoadFontFromMemory(".ttf", default_font, 512, nil)
	rl.SetTextureFilter(this.default_font.Texture, rl.FilterPoint)
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

	if this.color_hud == nil {
		this.color_hud = newColorHud(this)
	}

	go this.HandleMouseEvents()

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

		rl.BeginDrawing()
		// ---------------- DRAWING ----------------------------
		rl.ClearBackground(rl.NewColor(0x18, 0x18, 0x18, 0xFF))

		if this.initilized {
			this.navbar.draw()
			this.grid.draw()
			this.color_hud.draw()
		}
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

func (this *Simulation) HandleMouseEvents() {
	for {

		// if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			mouse := rl.GetMousePosition()

			end_of_algorithm_hud_y := NavBar_top_margin + NavBar_button_height
			end_of_grid_y := end_of_algorithm_hud_y + NavBar_top_margin + (this.grid.cell_width+this.grid.cell_padding)*Grid_rows

			if mouse.Y <= float32(end_of_algorithm_hud_y+NavBar_top_margin/2) {
				// TODO: Algorithm Hud logic
				log.Debug("Algorithm Hud")
			} else if mouse.Y <= float32(end_of_grid_y+this.grid.cell_padding) {
				// Grid Logic
				log.Debug("Grid")
				log.Debug("%+v", time.Now())
			} else {
				// Color hud + Speed slider logic
				log.Debug("Color Hud")
			}
			// x, y, err := this.grid.mapScreenToGrid(rl.GetMouseX(), rl.GetMouseY())
			// if err != nil {
			// 	// log.Error("%s", err)
			// } else {
			// 	if this.grid.cells[y][x].state != CellStateBorder {
			// 		log.Debug("{%d,%d}", x, y)
			// 		this.grid.cells[y][x].state = CellStateBorder
			// 	}
			// }
		}
	}
}
