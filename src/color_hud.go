package main

import (
	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// ColorHud_number_of_boxes = 5
	ColorHud_padding_size  = 30
	ColorHud_button_height = 50
)

type ColorHud struct {
	simulation_ctx  *Simulation
	button_size     int32
	padding_size    int32
	number_of_boxes int32
}

func newColorHud(simulation_ctx *Simulation) *ColorHud {
	this := ColorHud{padding_size: ColorHud_padding_size, number_of_boxes: 4}
	this.simulation_ctx = simulation_ctx
	this.button_size = int32((int32(float32(this.simulation_ctx.screen_width)/2.5) - this.padding_size*(this.number_of_boxes+1)) / this.number_of_boxes)

	log.Debug("New ColorHud: %+v", this)
	return &this
}

func (this *ColorHud) draw() {

	for i := range int32(this.number_of_boxes) {

		starting_point := int32(this.simulation_ctx.grid.starting_point.Y) + (int32(this.simulation_ctx.grid.cell_padding)+int32(this.simulation_ctx.grid.cell_width))*Grid_rows - int32(this.simulation_ctx.grid.cell_padding)
		rl.DrawRectangle(
			i*this.button_size+i*this.padding_size+this.padding_size,
			starting_point+(int32(this.simulation_ctx.screen_height)-starting_point-ColorHud_button_height)/2,
			this.button_size,
			ColorHud_button_height,
			rl.NewColor(0x33, 0x33, 0x33, 0xFF),
		)

		switch i {

		case 0:
			text_size := rl.MeasureTextEx(rl.GetFontDefault(), "Maria", 24, 1)
			button_width := this.button_size
			log.Debug("text_size = %#v", text_size)
			rl.DrawText(
				"Maria",
				i*this.button_size+i*this.padding_size+this.padding_size+int32((float32(button_width)-text_size.X)/2),
				starting_point+(int32(this.simulation_ctx.screen_height)-starting_point-ColorHud_button_height)/2+int32((ColorHud_button_height-text_size.Y)/2),
				24,
				rl.NewColor(0xFF, 0x0, 0xFF, 0xFF),
			)
		}
	}
}
