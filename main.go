// Entry point for programme

package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WinWidth = 900
	WinHeight = 500
	WinTitle = "Shape"
)

var (
	running = true

	// colours
	background_colour = rl.NewColor(100, 150, 255, 100)
)

func init() {
	rl.InitWindow(WinWidth, WinHeight, WinTitle)
	rl.SetTargetFPS(60)

}

func exit() {
	defer rl.CloseWindow()

}

func drawSene() {}

func input() {}

func update() {
	running = !rl.WindowShouldClose()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(background_colour)
	drawSene()
	rl.EndDrawing()
}

func main() {
	for running {
		input()
		update()
		render()
	}
	exit()
}
