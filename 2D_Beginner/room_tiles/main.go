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
	tiles rl.Texture2D  //IMG TEXTURE WITH TILES
	size  = float32(32) //TILE SIZE CHANGE FOR LARGER/SMALLER IMG DISPLAY SIZE
	cnt   rl.Vector2    //SCREEN CENTER VECTOR2

	/* DEFINES RECTANGLE OF FLOOR IMAGE IN tiles.png THIS DETERMINES WHAT PART OF THE TILES IMAGE TEXTURE TO DISPLAY WHEN A RECTANGLE IS DRAWN (0,0,16,16) MEANS X=0 Y=0 WIDTH=16 pixels HEIGHT=16 pixels THEREFORE WILL SEARCH FOR RECTANGLE AT COORDINATES 0,0 (X,Y) AND DRAW A RECTANGLE AT THOSE COORDINATES TO THE SCREEN OF 16X16 SCALED UP TO SIZE */
	floorTile = rl.NewRectangle(0, 0, 16, 16)

	wallTile              = rl.NewRectangle(0, 16, 16, 16) // DEFINES RECTANGLE OF FLOOR IMAGE
	roomRec               rl.Rectangle                     //SIZE OF ROOM
	wallColor, floorColor rl.Color                         //COLORS OF WALL & FLOOR TILES
)

func main() {

	rl.InitWindow(0, 0, "room tiles - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH := rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                            // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //FIND SCREEN CENTER

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	tiles = rl.LoadTexture("tiles.png") //LOAD IMAGES FROM FILE

	makeRoom() //MAKE INITIAL ROOM SEE FUNCTION AT END OF CODE

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

		if rl.IsKeyPressed(rl.KeySpace) { //MAKE A NEW ROOM IF SPACE KEY IS PRESSED
			makeRoom()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		//DRAW FLOOR TILES
		x := roomRec.X + size //SET X TO ROOM RECTANGLE LEFT PLUS SIZE OF WALL
		y := roomRec.Y + size //SET Y TO ROOM RECTANGLE TOP PLUS SIZE OF WALL
		for {
			rl.DrawTexturePro(tiles, floorTile, rl.NewRectangle(x, y, size, size), rl.Vector2Zero(), 0, rl.Fade(floorColor, 0.2)) //DRAW FLOOR TILE

			x += size //MOVE 1 TILE FORWARD

			//IF X REACHES BORDER OF RIGHT WALL X RETURNS TO START X POSITION & Y MOVES DOWN 1 TILE
			if x >= roomRec.X+roomRec.Width-size {
				x = roomRec.X + size
				y += size
			}

			//IF Y REACHES BORDER OF BOTTOM WALL STOP DRAWING FLOOR TILES
			if y >= roomRec.Y+roomRec.Height-size {
				break
			}
		}

		x = roomRec.X //SET X TO ROOM RECTANGLE LEFT
		y = roomRec.Y //SET X TO ROOM RECTANGLE TOP

		//DRAW TOP & BOTTOM WALLS
		for {
			//DRAW TOP WALL
			rl.DrawTexturePro(tiles, wallTile, rl.NewRectangle(x, y, size, size), rl.Vector2Zero(), 0, wallColor)

			//DRAW BOTTOM WALL ADD (roomRec.Height-size) TO y
			rl.DrawTexturePro(tiles, wallTile, rl.NewRectangle(x, y+(roomRec.Height-size), size, size), rl.Vector2Zero(), 0, wallColor)

			x += size //MOVE 1 TILE FORWARD

			//IF X REACHES ROOM RECTANGLE RIGHT BORDER STOP DRAWING
			if x >= roomRec.X+roomRec.Width {
				break
			}
		}

		//DRAW LEFT & RIGHT WALLS
		x = roomRec.X
		y = roomRec.Y + size //MOVE Y 1 TILE DOWN FOR TOP WALL SIZE

		for {
			//DRAW LEFT WALL
			rl.DrawTexturePro(tiles, wallTile, rl.NewRectangle(x, y, size, size), rl.Vector2Zero(), 0, wallColor)

			//DRAW RIGHT WALL ADD (roomRec.Width-size) to X
			rl.DrawTexturePro(tiles, wallTile, rl.NewRectangle(x+(roomRec.Width-size), y, size, size), rl.Vector2Zero(), 0, wallColor)

			y += size //MOVE 1 TILE DOWN

			//IF Y REACHES ROOM RECTANGLE BOTTOM BORDER SUBTRACT BOTTOM WALL SIZE STOP DRAWING
			if y >= roomRec.Y+roomRec.Height-size {
				break
			}
		}

		rl.EndMode2D()

		//DRAW TEXT
		rl.DrawText("camera zoom "+fmt.Sprintf("%.1f", camera.Zoom)+" press UP ARROW key to change", 10, 10, 20, rl.White)
		rl.DrawText("press SPACE key to make a new room", 10, 40, 20, rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func makeRoom() { //MAKES ROOM

	wallColor = ranCol()  //CHOOSES RANDOM WALL COLOR SEE FUNCTION BELOW
	floorColor = ranCol() //CHOOSES RANDOM FLOOR COLOR SEE FUNCTION BELOW

	//SIZE BASED ON RANDOM INTEGER VALUE NUMBER OF TILES MULTIPLIED BY SIZE TO GET PIXEL WIDTH
	width := float32(rInt(10, 21)) * size
	height := float32(rInt(10, 21)) * size

	//CREATE A CENTRED ROOM RECTANGLE
	roomRec = rl.NewRectangle(cnt.X-width/2, cnt.Y-height/2, width, height)

}

// RETURNS A RANDOM COLOR
func ranCol() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}

// RETURNS A RANDOM INTEGER FOR USE IN RANDOM COLOR ABOVE
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}
