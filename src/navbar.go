package main

import (
	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// NavBar_number_of_boxes = 5
	NavBar_padding_size  = 30
	NavBar_top_margin    = 10
	NavBar_button_height = 50
)

type NavBar struct {
	simulation_ctx  *Simulation
	button_size     int32
	padding_size    int32
	number_of_boxes int32
}

func newNavBar(simulation_ctx *Simulation) *NavBar {
	this := NavBar{padding_size: NavBar_padding_size, number_of_boxes: 5}
	this.simulation_ctx = simulation_ctx
	this.button_size = int32((int32(this.simulation_ctx.screen_width) - this.padding_size*(this.number_of_boxes+1)) / this.number_of_boxes)

	log.Debug("New NavBar: %+v", this)
	return &this
}

func (this *NavBar) draw() {

	for i := range int32(this.number_of_boxes) {

		rl.DrawRectangle(
			i*this.button_size+i*this.padding_size+this.padding_size,
			NavBar_top_margin,
			this.button_size,
			NavBar_button_height,
			rl.NewColor(0x33, 0x33, 0x33, 0xFF),
		)

	}
}
