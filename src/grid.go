package main

import (
	"errors"
	"fmt"
	"math"

	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Grid_rows    = 40
	Grid_columns = 81
)

var (
	ErrorOutOfBounds = errors.New("OutOfBounds")
)

type Grid struct {
	simulation_ctx *Simulation

	cells [Grid_rows][Grid_columns]*Cell

	cell_width   int
	cell_padding int
	cell_size    rl.Vector2
	cell_color   rl.Color

	starting_point rl.Vector2
}

func NewGrid(simulation_ctx *Simulation) *Grid {
	this := Grid{cell_width: 20, cell_padding: 3, cell_color: rl.NewColor(0x33, 0x33, 0x33, 0xFF)}

	this.simulation_ctx = simulation_ctx
	this.cell_size = rl.NewVector2(float32(this.cell_width), float32(this.cell_width))

	this.starting_point = rl.NewVector2(
		float32(NavBar_padding_size),
		float32(NavBar_top_margin*2+NavBar_button_height),
	)

	for y := range Grid_rows {
		for x := range Grid_columns {
			this.cells[y][x] = NewCell(this.simulation_ctx, x, y, this)
		}
	}
	this.cells[3][5].state = CellStateGoal
	this.cells[16][3].state = CellStateStart
	this.cells[4][7].state = CellStateBorder

	log.Debug("New Grid: {%+v}", this)
	return &this

}
func (this *Grid) draw() {

	for x := range Grid_rows {
		for y := range Grid_columns {
			cell := this.cells[x][y]
			cell.draw()
		}
	}

}

func (this *Grid) mapScreenToGrid(x int32, y int32) (int, int, error) {
	rx := (float64(x) - float64(this.starting_point.X)) / float64(this.cell_padding+this.cell_width)
	ry := (float64(y) - float64(this.starting_point.Y)) / float64(this.cell_padding+this.cell_width)
	if rx < 0 || ry < 0 {
		return 0, 0, fmt.Errorf("%s at mouse xy: {%d,%d}, rx: {%.3f}, ry: {%.3f}", ErrorOutOfBounds, x, y, rx, ry)
	}
	if rx >= Grid_columns || ry >= Grid_rows {
		return 0, 0, fmt.Errorf("%s at mouse xy: {%d,%d}, rx: {%.3f}, ry: {%.3f}", ErrorOutOfBounds, x, y, rx, ry)
	}

	return int(math.Floor(rx)), int(math.Floor(ry)), nil
}
