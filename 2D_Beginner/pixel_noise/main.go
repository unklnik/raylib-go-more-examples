package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	cnt        rl.Vector2    //SCREEN CENTER
	noiseColor = rl.Black    //INITIAL PIXEL REC COLOR
	backColor  = rl.White    //INITIAL BACKGROUND COLOR
	backRec    rl.Rectangle  //RECTANGLE FOR THE BACKGROUND COLOR
	noiseLevel = float32(4)  //INITIAL MAX SIZE LIMIT USED TO GENERATE RANDOM PIXEL NOISE REC
	noiseMax   = float32(10) //MAX SIZE LIMIT OF PIXEL NOISE REC
	noiseMin   = float32(2)  //MIN SIZE LIMIT OF PIXEL NOISE REC
)

func main() {

	rl.InitWindow(0, 0, "pixel noise - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH := rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                            // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //DEFINE SCREEN CENTER

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	texture := rl.LoadTexture("gopher.png") //LOAD GOPHER IMAGE

	backRec = rl.NewRectangle(0, 0, float32(scrW), float32(scrH)) //DEFINE BACKGROUND RECTANGLE

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyLeft) { //DECREASE SIZE LIMIT USED TO GENERATE RANDOM PIXEL REC
			if noiseLevel > noiseMin {
				noiseLevel -= 0.05
			}
		} else if rl.IsKeyDown(rl.KeyRight) { //INCREASE SIZE LIMIT USED TO GENERATE RANDOM PIXEL REC
			if noiseLevel < noiseMax {
				noiseLevel += 0.05
			}
		}

		if rl.IsKeyPressed(rl.KeyUp) { //SWITCH COLORS PIXEL NOISE
			if noiseColor == rl.Black {
				noiseColor = rl.Red
			} else if noiseColor == rl.Red {
				noiseColor = rl.Green
			} else if noiseColor == rl.Green {
				noiseColor = rl.SkyBlue
			} else if noiseColor == rl.SkyBlue {
				noiseColor = rl.Magenta
			} else if noiseColor == rl.Magenta {
				noiseColor = rl.White
			} else if noiseColor == rl.White {
				noiseColor = rl.Black
			}
		}

		if rl.IsKeyPressed(rl.KeyDown) { //SWITCH COLORS BACKGROUND COLOR
			if backColor == rl.Black {
				backColor = rl.Red
			} else if backColor == rl.Red {
				backColor = rl.Green
			} else if backColor == rl.Green {
				backColor = rl.SkyBlue
			} else if backColor == rl.SkyBlue {
				backColor = rl.Magenta
			} else if backColor == rl.Magenta {
				backColor = rl.White
			} else if backColor == rl.White {
				backColor = rl.Black
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		rl.DrawRectangleRec(backRec, backColor) //DRAW BACKGROUND RECTANGLE

		//DRAW GOPHER IMAGE
		rl.DrawTexture(texture, int32(cnt.X)-texture.Width/2, int32(cnt.Y)-texture.Height/2, rl.White)

		rl.EndMode2D()

		//DRAW PIXELS OUTSIDE OF CAMERA MODE 2D THEREFORE WILL DISPLAY THE SAME ON ZOOM CHANGE
		num := 100 //NUMBER OF PIXEL NOISE RECS TO DRAW PER FRAME
		for num > 0 {
			size := rF32(1, noiseLevel) // RANDOM REC SIZE RF32 FUNCTION BELOW
			//DEFINE REC
			rec := rl.NewRectangle(rF32(0, float32(scrW)), rF32(0, float32(scrH)), size, size)
			rl.DrawRectangleRec(rec, noiseColor) //DRAW RECTANGLE
			num--
		}

		rl.DrawText("UP ARROW KEY CHANGE PIXEL NOISE COLOR", 8, 12, 20, rl.White)
		rl.DrawText("UP ARROW KEY CHANGE PIXEL NOISE COLOR", 9, 11, 20, rl.Black)
		rl.DrawText("UP ARROW KEY CHANGE PIXEL NOISE COLOR", 10, 10, 20, rl.Blue)

		rl.DrawText("DOWN ARROW KEY CHANGE BACKGROUND COLOR", 8, 42, 20, rl.White)
		rl.DrawText("DOWN ARROW KEY CHANGE BACKGROUND COLOR", 9, 41, 20, rl.Black)
		rl.DrawText("DOWN ARROW KEY CHANGE BACKGROUND COLOR", 10, 40, 20, rl.Blue)

		rl.DrawText("RIGHT LEFT ARROW KEYS INCREASE DECREASE SIZE", 8, 72, 20, rl.White)
		rl.DrawText("RIGHT LEFT ARROW KEYS INCREASE DECREASE SIZE", 9, 71, 20, rl.Black)
		rl.DrawText("RIGHT LEFT ARROW KEYS INCREASE DECREASE SIZE", 10, 70, 20, rl.Blue)

		rl.EndDrawing()
	}

	rl.UnloadTexture(texture) //UNLOAD FROM MEMORY

	rl.CloseWindow()
}

// RETURNS A RANDOM FLOAT32 VALUE BETWEEN MIN/MAX RANGE
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
