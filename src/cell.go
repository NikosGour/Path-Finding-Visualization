package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Cell_BlankColor   = rl.NewColor(0x33, 0x33, 0x33, 0xFF)
	Cell_StartColor   = rl.NewColor(0x0, 0xFF, 0x0, 0xFF)
	Cell_GoalColor    = rl.NewColor(0x0, 0x0, 0xFF, 0xFF)
	Cell_BorderColor  = rl.NewColor(0xFF, 0x0, 0x0, 0xFF)
	Cell_UnknownColor = rl.NewColor(0xFF, 0x0, 0xFF, 0xFF)
)

type CellState uint8

const (
	CellStateBlank CellState = iota
	CellStateStart
	CellStateGoal
	CellStateBorder
	CellStatePath
)

type Cell struct {
	simulation_ctx *Simulation

	state CellState
	rect  *rl.Rectangle
}

func NewCell(simulation_ctx *Simulation, x int, y int, grid Grid) *Cell {
	this := Cell{state: CellStateBlank}
	this.createCellRectangle(x, y, grid)
	// log.Debug("%#v", this)
	return &this
}

func (this *Cell) createCellRectangle(x int, y int, grid Grid) {
	cell_x := float32(int(grid.starting_point.X) + (grid.cell_width+grid.cell_padding)*x)
	cell_y := float32(int(grid.starting_point.Y) + (grid.cell_width+grid.cell_padding)*y)

	_t := rl.NewRectangle(cell_x, cell_y, grid.cell_size.X, grid.cell_size.Y)
	this.rect = &_t
}

func (this *Cell) draw() {

	var color rl.Color

	switch this.state {
	case CellStateBlank:
		color = Cell_BlankColor
	case CellStateStart:
		color = Cell_StartColor
	case CellStateGoal:
		color = Cell_GoalColor
	case CellStateBorder:
		color = Cell_BorderColor
	default:
		color = Cell_UnknownColor
	}

	rl.DrawRectangleRec(
		*this.rect,
		color,
	)
}
