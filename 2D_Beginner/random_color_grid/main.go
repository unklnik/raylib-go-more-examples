package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	bloks []blok        //SLICE OF BLOKS SEE STRUCT BELOW
	size  = float32(32) //GRID RECTANGLE WIDTH & HEIGHT
)

type blok struct { // BLOK STRUCTS THAT CONTAIN THE COLOR & POSITION
	col rl.Color
	rec rl.Rectangle
}

func main() {

	rl.InitWindow(0, 0, "grid random colors - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH := rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                            // SET WINDOW SIZE

	x := float32(0)
	y := float32(0)
	for {
		newBlok := blok{}                               //CREATES AN EMPTY BLOK
		newBlok.col = ranCol()                          //ASSIGNS A RANDOM COLOR
		newBlok.rec = rl.NewRectangle(x, y, size, size) //CREATES A REC AT X,Y WITH SIZE X SIZE
		bloks = append(bloks, newBlok)                  //ADDS NEW BLOK TO SLICE

		x += size               //MOVES X BY SIZE
		if x >= float32(scrW) { //IF X LARGER THAN OR EQUALS SCREEN WIDTH GO BACK TO ZERO
			x = 0
			y += size //MOVE Y DOWN ONE LINE
		}

		if y >= float32(scrH) { //IF Y LARGER THAN OR EQUALS SCREEN HEIGHT END
			break
		}
	}

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		for a := 0; a < len(bloks); a++ { //DRAWS THE LIST OF BLOCKS
			rl.DrawRectangleRec(bloks[a].rec, bloks[a].col)
		}

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// RETURNS A RANDOM COLOR
func ranCol() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}

// RETURNS A RANDOM INTEGER FOR USE IN RANDOM COLOR ABOVE
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}
