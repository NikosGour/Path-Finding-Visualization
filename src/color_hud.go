package main

import (
	"os"

	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ColorHud_number_of_buttons = 4
	ColorHud_padding_size      = 30
	ColorHud_button_height     = 50
	ColorHud_font_size         = 40
	ColorHud_spacing           = 1
)

type ColorHud struct {
	simulation_ctx *Simulation

	button_size    int32
	padding_size   int32
	starting_point int32

	buttons [ColorHud_number_of_buttons]rl.Rectangle
}

func newColorHud(simulation_ctx *Simulation) *ColorHud {
	this := ColorHud{padding_size: ColorHud_padding_size}
	this.simulation_ctx = simulation_ctx
	this.button_size = int32((int32(float32(this.simulation_ctx.screen_width)/2.5) - this.padding_size*(ColorHud_number_of_buttons+1)) / ColorHud_number_of_buttons)
	this.starting_point = int32(this.simulation_ctx.grid.starting_point.Y) + (int32(this.simulation_ctx.grid.cell_padding)+int32(this.simulation_ctx.grid.cell_width))*Grid_rows - int32(this.simulation_ctx.grid.cell_padding)

	this.initButtons()

	log.Debug("New ColorHud: %+v", this)
	return &this
}

func (this *ColorHud) initButtons() {
	for i := range int32(ColorHud_number_of_buttons) {
		this.buttons[i] = rl.NewRectangle(
			float32(i*this.button_size+i*this.padding_size+this.padding_size),
			float32(this.starting_point+(int32(this.simulation_ctx.screen_height)-this.starting_point-ColorHud_button_height)/2),
			float32(this.button_size),
			ColorHud_button_height,
		)
	}
}
func (this *ColorHud) draw() {

	for i := range int32(ColorHud_number_of_buttons) {
		rl.DrawRectangleRec(this.buttons[i], rl.NewColor(0x33, 0x33, 0x33, 0xFF))

		switch i {
		case 0:
			this.centerTextOnButton("Start", i, Cell_StartColor)
		case 1:
			this.centerTextOnButton("Goal", i, Cell_GoalColor)
		case 2:
			this.centerTextOnButton("Border", i, Cell_BorderColor)
		case 3:
			this.centerTextOnButton("Erase", i, rl.White)
		default:
			log.Error("%s: ColorHud_number_of_buttons is not the same as the cases in switch of ColorHud.draw(). number_of_buttons: {%d}, i: {%d}", ErrorUnreachable, ColorHud_number_of_buttons, i)
			os.Exit(1)
		}
	}
}

func (this *ColorHud) centerTextOnButton(text string, button int32, color rl.Color) {
	i := button
	text_size := rl.MeasureTextEx(this.simulation_ctx.default_font, text, ColorHud_font_size, ColorHud_spacing)
	button_width := this.button_size
	// log.Debug("text_size = %#v", text_size)
	rl.DrawTextEx(
		this.simulation_ctx.default_font,
		text,
		rl.NewVector2(
			float32(i*this.button_size+i*this.padding_size+this.padding_size+int32((float32(button_width)-text_size.X)/2)),
			float32(this.starting_point+(int32(this.simulation_ctx.screen_height)-this.starting_point-ColorHud_button_height)/2+int32((ColorHud_button_height-text_size.Y)/2))),

		ColorHud_font_size,
		ColorHud_spacing,
		color,
	)
}
