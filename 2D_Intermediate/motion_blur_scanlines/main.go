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
	imgs          rl.Texture2D   //TEXTURE TO LOAD IMAGE
	cntr          rl.Vector2     //SCREEN CENTER
	fps           = int32(60)    //FRAMES PER SECOND OF GAME
	blurDistance  = float32(2)   //OFFSET OF BLUR IMAGE
	blurFade      = float32(0.2) //BLUR OPACITY (FADE)
	scanlines     []rl.Vector2   //SLICE OF VECTOR2 FOR SCAN LINES
	scanlineson   = true         //SET SCANLINES ON
	scanlinesFade = float32(0.5) //SCANLINES FADE
	currentColor  int            //CURRENT SCANLINE COLOR

	scanlineColors = []rl.Color{rl.Black, rl.Green, rl.SkyBlue, rl.LightGray} //SCANLINES COLORS
)

func main() {

	rl.InitWindow(0, 0, "fake motion blur - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH := rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                            // SET WINDOW SIZE
	//rl.SetWindowState(rl.FlagBorderlessWindowedMode) // UNCOMMENT IF YOU HAVE DISPLAY ISSUES
	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES

	cntr = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //CENTER OF SCREEN
	imgs = rl.LoadTexture("imgs.png")                      //LOAD IMAGESc
	flagIMG := rl.NewRectangle(0, 0, 16, 16)               //FLAG IMAGE RECTANGLE
	fireballIMG := rl.NewRectangle(16, 0, 16, 16)          //FIREBALL IMAGE RECTANGLE

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 2.0       //SETS CAMERA ZOOM
	camera.Target = cntr    //OFFSET CAMERA FROM CENTER AFTER ZOOM CHANGE
	camera.Offset.X = float32(scrW / 2)
	camera.Offset.Y = float32(scrH / 2)

	//CREATE SCANLINE VECTOR2
	scanoffset := float32(3)           //SPACE BETWEEN SCANLINES
	y := float32(-scanoffset)          //START FROM ABOVE SCREEN TOP
	for y < float32(scrH)+scanoffset { //END BELOW SCREEN HEIGHT
		scanlines = append(scanlines, rl.NewVector2(0, y))
		y += scanoffset
	}

	rl.SetTargetFPS(fps) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyZ) { //CHANGE FADE
			scanlinesFade += 0.1
			if scanlinesFade > 1 {
				scanlinesFade = 0.2
			}
		}
		if rl.IsKeyPressed(rl.KeyRight) { //CHANGE BLUR
			blurFade -= 0.1
			if blurFade <= 0.1 {
				blurFade = 0.8
			}
		}
		if rl.IsKeyPressed(rl.KeyLeft) { //CHANGE COLOR
			currentColor++
			if currentColor == len(scanlineColors)-1 {
				currentColor = 0
			}
		}
		if rl.IsKeyPressed(rl.KeyDown) { //CHANGE BLUR DISTANCE
			blurDistance++
			if blurDistance >= 10 {
				blurDistance = 1
			}
		}
		if rl.IsKeyPressed(rl.KeyUp) { //CHANGE ZOOM
			if camera.Zoom == 1 {
				camera.Zoom = 1.5
			} else if camera.Zoom == 1.5 {
				camera.Zoom = 2
			} else if camera.Zoom == 2 {
				camera.Zoom = 1
			}
			//ADJUST CAMERA TO CENTER IMAGE AFTER CHANGE IN ZOOM
			camera.Target = cntr
			camera.Offset.X = float32(scrW / 2)
			camera.Offset.Y = float32(scrH / 2)
		}
		if rl.IsKeyPressed(rl.KeySpace) { //TURN SCANLINES ON/OFF
			scanlineson = !scanlineson
		}

		for i := 0; i < len(scanlines); i++ { //MOVE SCANLINES DOWN BY OFFEST
			scanlines[i].Y += scanoffset
			if scanlines[i].Y > float32(scrH) { //RETURN TO ABOVE SCREEN TOP IF REACH SCREEN HEIGHT
				scanlines[i].Y = -scanoffset
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		size := float32(128)                                                   //IMAGE SIZE
		fireballRec := rl.NewRectangle(cntr.X-size/2, cntr.Y-size, size, size) //FIREBALL DRAW REC
		flagRec := rl.NewRectangle(cntr.X-size/2, cntr.Y, size, size)          //FLAG DRAW REC

		fireballColor := ranOrange() //RANDOM COLOR SEE FUNCTION BELOW
		//DRAW ORIGINAL IMAGE
		rl.DrawTexturePro(imgs, fireballIMG, fireballRec, rl.Vector2Zero(), 0, fireballColor)
		//COPY ORIGINAL IMAGE
		blurRec := fireballRec
		//MOVE Y RANDOM DISTANCE FUNCTION BELOW
		blurRec.Y += rF32(-blurDistance, blurDistance)
		//MOVE X RANDOM DISTANCE FUNCTION BELOW
		blurRec.X += rF32(-blurDistance, blurDistance)
		rl.DrawTexturePro(imgs, fireballIMG, blurRec, rl.Vector2Zero(), 0, rl.Fade(fireballColor, blurFade)) //DRAW OFFEST IMAGE WITH FADE

		rl.DrawTexturePro(imgs, flagIMG, flagRec, rl.Vector2Zero(), 0, rl.White)
		blurRec = flagRec
		blurRec.Y += rF32(-blurDistance, blurDistance)
		blurRec.X += rF32(-blurDistance, blurDistance)
		rl.DrawTexturePro(imgs, flagIMG, blurRec, rl.Vector2Zero(), 0, rl.Fade(rl.White, blurFade))

		rl.EndMode2D()

		if scanlineson { //DRAW SCANLINES
			for i := 0; i < len(scanlines); i++ {
				endV2 := rl.NewVector2(scanlines[i].X+float32(scrW), scanlines[i].Y)
				rl.DrawLineV(scanlines[i], endV2, rl.Fade(scanlineColors[currentColor], scanlinesFade))
			}
		}

		rl.DrawText("UP key change zoom / DOWN key change blur distance / RIGHT key change blur fade / LEFT key change scanlines color / Z key change scanlines fade / SPACE key turn scanlines on/off", 10, 10, 10, rl.White)

		rl.DrawText("blur distance "+fmt.Sprint(blurDistance)+" blur fade "+fmt.Sprint(blurFade)+" scanlines fade "+fmt.Sprint(scanlinesFade), 10, 20, 10, rl.White)

		rl.EndDrawing()
	}

	rl.UnloadTexture(imgs) //UNLOAD FROM MEMORY

	rl.CloseWindow()
}

// RETURNS A RANDOM ORANGE COLOR
func ranOrange() rl.Color {
	return rl.NewColor(uint8(255), uint8(rInt(70, 170)), uint8(rInt(0, 50)), 255)
}

// RETURNS A RANDOM INTEGER FOR USE IN RANDOM ORANGE COLOR FUNCTION ABOVE
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RETURNS A RANDOM FLOAT32 VALUE BETWEEN MIN/MAX VALUES
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
