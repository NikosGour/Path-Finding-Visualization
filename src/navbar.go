package main

import (
	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
// NavBar_number_of_boxes = 5
// NavBar_padding_size = 30
)

type NavBar struct {
	simulation_ctx  *Simulation
	button_size     int32
	padding_size    int32
	number_of_boxes int32
}

func newNavBar(simulation_ctx *Simulation) *NavBar {
	this := NavBar{padding_size: 30, number_of_boxes: 5}
	this.simulation_ctx = simulation_ctx
	this.button_size = int32((int32(this.simulation_ctx.screen_width) - this.padding_size*(this.number_of_boxes+1)) / this.number_of_boxes)

	log.Debug("New NavBar: %+v", this)
	return &this
}

func (this *NavBar) draw() {

	for i := range int32(this.number_of_boxes) {

		rl.DrawRectangle(
			i*this.button_size+i*this.padding_size+this.padding_size,
			10,
			this.button_size,
			50,
			rl.NewColor(0xFF, 0x0, 0x0, 0xFF),
		)

	}
}
