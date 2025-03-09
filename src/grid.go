package main

import (
	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Grid_rows    = 40
	Grid_columns = 81
)

type CellState uint8

const (
	CellStateBlank CellState = iota
	CellStateStart
	CellStateGoal
	CellStateBorder
	CellStatePath
)

type Grid struct {
	simulation_ctx *Simulation

	cells [Grid_rows][Grid_columns]CellState

	cell_width   int
	cell_padding int
	cell_size    rl.Vector2
	cell_color   rl.Color

	starting_point rl.Vector2
}

func NewGrid(simulation_ctx *Simulation) *Grid {
	this := Grid{cell_width: 20, cell_padding: 3, cell_color: rl.NewColor(0x33, 0x33, 0x33, 0xFF)}

	this.cells[3][5] = CellStateGoal
	this.cells[16][3] = CellStateStart
	this.cells[4][7] = CellStateBorder

	this.simulation_ctx = simulation_ctx
	this.cell_size = rl.NewVector2(float32(this.cell_width), float32(this.cell_width))

	// this.starting_point = rl.NewVector2(
	// 	float32((this.simulation_ctx.screen_width-(this.cell_width+this.cell_padding)*Grid_columns)/2),
	// 	float32((this.simulation_ctx.screen_height-(this.cell_width+this.cell_padding)*Grid_rows)/2),
	// )
	this.starting_point = rl.NewVector2(
		float32(NavBar_padding_size),
		float32(NavBar_top_margin*2+NavBar_button_height),
	)

	log.Debug("%#v", this)
	return &this

}
func (this *Grid) draw() {

	copy_starting_point := this.starting_point
	for x := range Grid_rows {
		point := copy_starting_point

		for y := range Grid_columns {
			cell := this.cells[x][y]

			var color rl.Color
			switch cell {
			case CellStateBlank:
				color = this.cell_color
			case CellStateStart:
				color = rl.NewColor(0x0, 0xFF, 0x0, 0xFF)
			case CellStateGoal:
				color = rl.NewColor(0x0, 0x0, 0xFF, 0xFF)
			case CellStateBorder:
				color = rl.NewColor(0xFF, 0x0, 0x0, 0xFF)
			default:
				color = rl.NewColor(0xFF, 0x0, 0xFF, 0xFF)
			}

			rl.DrawRectangleV(
				point,
				this.cell_size,
				color,
			)

			point.X += float32(this.cell_width + this.cell_padding)
			copy_starting_point.X += float32(this.cell_width + this.cell_padding)
		}
		copy_starting_point.X = this.starting_point.X
		copy_starting_point.Y += float32(this.cell_width + this.cell_padding)
	}

}

func (this *Grid) drawCell(x int, y int, color rl.Color) {
	draw_vec := rl.NewVector2(
		float32(int(this.starting_point.X)+(this.cell_width+this.cell_padding)*x),
		float32(int(this.starting_point.Y)+(this.cell_width+this.cell_padding)*y),
	)

	rl.DrawRectangleV(draw_vec, this.cell_size, color)
}
