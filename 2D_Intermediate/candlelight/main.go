package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	cnt          rl.Vector2                      //SCREEN CENTER
	candleIMG    rl.Texture2D                    //IMAGE
	candleRec    = rl.NewRectangle(0, 0, 16, 16) //IMAGE RECTANGLE
	positions    []rl.Vector2                    //POSITIONS OF IMAGES ON SCREEN
	size         = float32(64)                   //SIZE OF IMAGES
	currentColor int                             //CURRENT COLOR
	num          = 10                            //NUMBER OF CANDLES

	candleColors = []rl.Color{rl.Yellow, rl.White, rl.SkyBlue, rl.DarkGray, rl.Magenta} //LIGHT COLORS
)

func main() {

	rl.InitWindow(0, 0, "candlelight - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH := rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                            // SET WINDOW SIZE
	//rl.SetWindowState(rl.FlagBorderlessWindowedMode) // UNCOMMENT IF YOU HAVE DISPLAY ISSUES
	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //FIND SCREEN CENTER

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	candleIMG = rl.LoadTexture("img.png") //LOAD IMAGES FROM FILE

	for num > 0 {
		positions = append(positions, rl.NewVector2(rF32(0, float32(scrW)-size), rF32(0, float32(scrH)-size))) //CREATE RANDOM VECTOR2 POSITIONS SEE FUNCTION BELOW
		num--
	}

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyUp) { //CHANGE CAMERA ZOOM KEYS
			if camera.Zoom == 0.5 {
				camera.Zoom = 1
			} else if camera.Zoom == 1 {
				camera.Zoom = 1.5
			} else if camera.Zoom == 1.5 {
				camera.Zoom = 2
			} else if camera.Zoom == 2 {
				camera.Zoom = 0.5
			}

			camera.Target = cnt                 //SET THE CAMERA TARGET TO CENTER
			camera.Offset.X = float32(scrW / 2) //ADJUST FOR ZOOM
			camera.Offset.Y = float32(scrH / 2) //ADJUST FOR ZOOM
		}
		if rl.IsKeyPressed(rl.KeyDown) { //CHANGE CURRENT COLOR
			currentColor++
			if currentColor == len(candleColors) {
				currentColor = 0
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		for i := 0; i < len(positions); i++ { //RANGE OVER SLICE AND DRAW IMAGES
			rl.DrawTexturePro(candleIMG, candleRec, rl.NewRectangle(positions[i].X-size/2, positions[i].Y-size/2, size, size), rl.Vector2Zero(), 0, rl.White)

			//DRAW A CIRCLE WITH COLOR AND BLANK GRADIENT FILL TO SIMULATE LIGHT WITH RANDOM FADE
			rl.DrawCircleGradient(int32(positions[i].X), int32(positions[i].Y), size*4, rl.Fade(candleColors[currentColor], rF32(0.1, 0.3)), rl.Blank)
		}

		rl.EndMode2D()

		//DRAW TEXT
		rl.DrawText("camera zoom "+fmt.Sprintf("%.1f", camera.Zoom)+" press UP ARROW key to change", 10, 10, 20, rl.White)
		rl.DrawText("press DOWN ARROW key to change color", 10, 30, 20, rl.White)

		rl.EndDrawing()
	}

	rl.UnloadTexture(candleIMG) //UNLOAD FROM MEMORY

	rl.CloseWindow()
}

// RETURNS A RANDOM FLOAT32 BETWEEN MIN/MAX
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
