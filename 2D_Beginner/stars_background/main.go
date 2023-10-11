package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	numStars   = 200 // NUMBER OF STARS TO DRAW
	stars      []blok
	direc      int         // DIRECTION
	timer      int32       // CHANGE DIRECTION TIMER
	fps        = int32(60) // FRAMES PER SECOND
	colorsOn   bool        // COLORS ON/OFF
	scrW, scrH int
	maxVel     = float32(10) // MAX SPEED FOR DETERMINING X Y MOVEMENT
	velX, velY float32       // X Y SPEED
)

/*
direc = direction
numbers correspond to direction of stars movement
numbers start at 1 and move clockwise

for example 2 = UP, 7 = DOWN & LEFT, 3 = UP & RIGHT, 4 = RIGHT

	1 2 3
	8   4
	7 6 5

*/

type blok struct { // BLOK STRUCTS THAT CONTAIN THE COLOR & POSITION
	col  rl.Color
	rec  rl.Rectangle
	fade float32
}

func main() {

	rl.InitWindow(0, 0, "stars background - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	rl.HideCursor()          // HIDES MOUSE CURSOR
	makeStars()              // FUNCTION MAKE STARS SEE END OF CODE
	direc = rInt(1, 9)       // CHOOSE INITIAL MOVEMENT DIRECTION
	velX = rF32(1, maxVel)   // FIND RANDOM X SPEED
	velY = rF32(1, maxVel)   // FIND RANDOM Y SPEED
	timer = rI32(2, 8) * fps // SET TIMER

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	rl.SetTargetFPS(fps) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		upStars() // FUNCTION TO UPDATE TIMER, MOVEMENT & FADE

		if rl.IsKeyPressed(rl.KeySpace) {
			colorsOn = !colorsOn // TURN COLORS ON/OFF WITH SPACE BAR
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		for i := 0; i < len(stars); i++ { // RANGE OVER SLICE OF STAR BLOKS & DRAW

			if colorsOn { // DRAW COLOR STARS
				rl.DrawRectangleRec(stars[i].rec, rl.Fade(stars[i].col, stars[i].fade))
			} else { // DRAW WHITE STARS
				rl.DrawRectangleRec(stars[i].rec, rl.Fade(rl.White, stars[i].fade))
			}

		}

		rl.EndMode2D()

		rl.DrawText("press space on/off colors", 8, 12, 20, ranCol())  // TEXT FLASHING COLOR EFFECT
		rl.DrawText("press space on/off colors", 9, 11, 20, rl.Black)  // TEXT BLACK SHADOWN
		rl.DrawText("press space on/off colors", 10, 10, 20, rl.White) // TEXT

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
func upStars() {

	timer-- // DECREASE TIMER
	if timer <= 0 {
		direc = rInt(1, 9)       // CHOOSE NEW DIRECTION
		velX = rF32(2, maxVel)   // CHOOSE NEW X SPEED
		velY = rF32(2, maxVel)   // CHOOSE NEW Y SEEPD
		timer = rI32(2, 8) * fps // SET NEW TIMER
	}

	for i := 0; i < len(stars); i++ { // RANGE OVER SLICE OF STAR BLOKS & UPDATE

		stars[i].fade -= 0.01   // FADE OUT
		if stars[i].fade <= 0 { // IF FADE LESS THAN OR EQUALS ZERO SET NEW FADE
			stars[i].fade = rF32(0.5, 0.9)
		}

		switch direc { // MOVE STAR REC ACCORDING TO CHOSEN DIRECTION
		case 1: //UP LEFT
			stars[i].rec.X -= velX
			stars[i].rec.Y -= velY
		case 2: //UP
			stars[i].rec.Y -= velY
		case 3: //UP RIGHT
			stars[i].rec.X += velX
			stars[i].rec.Y -= velY
		case 4: // RIGHT
			stars[i].rec.X += velX
		case 5: //DOWN RIGHT
			stars[i].rec.X += velX
			stars[i].rec.Y += velY
		case 6: //DOWN
			stars[i].rec.Y += velY
		case 7: //DOWN LEFT
			stars[i].rec.X -= velX
			stars[i].rec.Y += velY
		case 8: //LEFT
			stars[i].rec.X -= velX
		}

		// IF STAR REC X IS OVER SCREEN BORDER LEFT MOVE TO SCREEN BORDER RIGHT
		if stars[i].rec.X < 0 {
			stars[i].rec.X = float32(scrW)
		}
		// IF STAR REC IS OVER SCREEN BORDER RIGHT MOVE TO SCREEN BORDER LEFT
		if stars[i].rec.X > float32(scrW) {
			stars[i].rec.X = 0
		}
		// IF STAR REC Y IS OVER SCREEN BORDER TOP MOVE TO SCREEN BORDER BOTTOM
		if stars[i].rec.Y < 0 {
			stars[i].rec.Y = float32(scrH)
		}
		// IF STAR REC IS OVER SCREEN BORDER BOTTOM MOVE TO SCREEN BORDER TOP
		if stars[i].rec.Y > float32(scrH) {
			stars[i].rec.Y = 0
		}
	}

}
func makeStars() {

	maxSize := float32(8) // MAXIMUM SIZE OF RECTANGLE WIDTHS

	for i := 0; i < numStars; i++ { // FILLS THE STARS SLICE
		// MAKE EMPTY BLOK STRUCT
		newBlok := blok{}
		// CHOOSE RANDOM COLOR SEE FUNCTION END OF CODE
		newBlok.col = ranCol()
		// SET RANDOM FADE/OPACITY
		newBlok.fade = rF32(0.5, 0.9)
		// CHOOSE RANDOM SIZE OF RECTANGLE SIDES
		width := rF32(1, maxSize)
		// CREATE THE RECTANGLE
		newBlok.rec = rl.NewRectangle(rF32(0, float32(scrW)), rF32(0, float32(scrH)), width, width)
		// ADD TO SLICE
		stars = append(stars, newBlok)
	}

}

// RETURNS A RANDOM INTEGER 32
func rI32(min, max int) int32 {
	return int32(min + rand.Intn(max-min))
}

// RETURNS A RANDOM FLOAT 32
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}

// RETURNS A RANDOM COLOR
func ranCol() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}

// RETURNS A RANDOM INTEGER FOR USE IN RANDOM COLOR ABOVE
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}
