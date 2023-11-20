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
	bloks      []blok        //SLICE OF BLOKS SEE STRUCT BELOW
	blokbits   []blok        //SLICE OF BLOK PIECES WHEN EXPLODING
	size       = float32(32) //GRID RECTANGLE WIDTH & HEIGHT
	fps        = int32(60)   //FRAMES PER SECOND
	timer      int32         //EXPLOSION TIMER
	scrW, scrH int           //SCREEN WIDTH & HEIGHT
)

type blok struct { // BLOK STRUCT
	col              rl.Color     //COLOR
	cnt              rl.Vector2   //CENTER
	rec              rl.Rectangle //RECTANGLE
	fade, velX, velY float32      //OPACITY MOVEMENT SPEED X & Y
}

func main() {

	rl.InitWindow(0, 0, "exploding blocks - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	makebloks() //MAKE INITIAL BLOCKS

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	rl.SetTargetFPS(fps) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeySpace) && timer == 0 { // KEY TO EXPLODE
			explode()
			timer = fps + fps/2 //SET TIMER TO 1 AND A HALF SECONDS
		}

		if timer > 0 { //COUNTDOWN TIMER
			timer--
			if timer == 1 { //MAKE NEW BLOCKS
				makebloks()
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		if timer > 0 { // IF TIMER IS LARGER THAN ZERO DRAW EXPLOSION PIECES
			for i := 0; i < len(blokbits); i++ { //DRAWS THE LIST OF BLOCKS
				rl.DrawRectangleRec(blokbits[i].rec, rl.Fade(blokbits[i].col, blokbits[i].fade))
				if blokbits[i].fade > 0 {
					blokbits[i].fade -= 0.01              //FADE OUT EXPLOSION PIECES
					blokbits[i].rec.X += blokbits[i].velX //MOVE X EXPLOSION PIECES
					blokbits[i].rec.Y += blokbits[i].velY //MOVE Y EXPLOSION PIECES
				}
			}
			rl.DrawText("time till blocks respawn: "+fmt.Sprint(timer), 10, 10, 10, rl.White)
		} else { // IF TIMER IS ZERO DRAW BLOCKS
			for i := 0; i < len(bloks); i++ { //DRAWS THE LIST OF BLOCKS
				rl.DrawRectangleRec(bloks[i].rec, bloks[i].col)
			}
			rl.DrawText("press SPACE key to explode", 10, 10, 10, rl.White)
		}

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func makebloks() { //MAKES THE BLOCKS THAT EXPLODE
	bloks = nil       //CLEARS THE SLICE
	num := rInt(4, 8) //RANDOM NUMBER OF BLOCKS
	for num > 0 {
		newBlok := blok{}
		newBlok.col = ranCol() //RANDOM COLOR SEE FUNCTION BELOW
		siz := rF32(64, 256)   //RANDOM SIZE
		newBlok.rec = rl.NewRectangle(rF32(0, float32(scrW)-siz), rF32(0, float32(scrH)-siz), siz, siz)
		newBlok.cnt = rl.NewVector2(newBlok.rec.X+siz/2, newBlok.rec.Y+siz/2) //CENTER FOR EXPLOSION PIECES
		bloks = append(bloks, newBlok)                                        //ADD TO SLICE
		num--
	}
}
func explode() { //MAKES THE EXPLOSION PIECES

	blokbits = nil
	for i := 0; i < len(bloks); i++ {

		num := rInt(15, 25) //RANDOM NUMBER OF EXPLOSION PIECES
		for num > 0 {
			newBlok := blok{}
			newBlok.col = ranCol()
			siz := rF32(8, 32) //RANDOM SIZE SMALLER THAN BLOCK
			//MAKE THE RECTANGLE USING THE CENTER OF THE INITIAL BLOCK
			newBlok.rec = rl.NewRectangle(bloks[i].cnt.X-siz/2, bloks[i].cnt.Y-siz/2, siz, siz)
			newBlok.fade = rF32(0.5, 0.8) //RANDOM OPACITY
			newBlok.velX = rF32(-32, 32)  //RANDOM X MOVEMENT
			newBlok.velY = rF32(-32, 32)  //RANDOM X MOVEMENT
			blokbits = append(blokbits, newBlok)
			num--
		}

	}

}

// RETURNS A RANDOM COLOR
func ranCol() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}

// RETURNS A RANDOM INTEGER FOR USE IN RANDOM COLOR ABOVE
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
