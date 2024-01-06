package main

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	tiles         rl.Texture2D                     //IMG TEXTURE WITH TILES
	size          = float32(32)                    //TILE SIZE CHANGE FOR LARGER/SMALLER IMG DISPLAY SIZE
	cnt           rl.Vector2                       //SCREEN CENTER VECTOR2
	wallTile      = rl.NewRectangle(16, 0, 16, 16) //DEFINES RECTANGLE OF WALL BLOCK IMAGE
	spikeTile     = rl.NewRectangle(1, 2, 13, 13)  //DEFINES RECTANGLE OF SPIKE IMAGE
	innerRec      rl.Rectangle                     //INNER RECTANGLE OF ROOM
	bloks         []xblok                          //SLICE OF WALL BLOCKS
	debug         bool                             //USED TO TURN ON/OFF DEBUGGING MODE
	playerRec     rl.Rectangle                     //PLAYER RECTANGLE
	playerCollisT int32                            //COLLISION TIMER FOR PLAYER VERSUS SPIKES
	fps           = int32(60)                      //FRAMES PER SECOND
	playerSpd     = float32(8)                     //MAX SPEED OF PLAYER MOVEMENT
)

// STRUCT CONTAINS DATA OF WALL BLOCKS
type xblok struct {
	rec, img            rl.Rectangle   //RECTANGLE & IMAGE
	spikes, spikesOnOff bool           //WHETHER BLOCK HAS SPIKES & IF SPIKES MOVE UP/DOWN LEFT/RIGHT
	spikeRecs           []rl.Rectangle //RECTANGLES FOR THE 4 SPIKE IMAGES
	spd, dirX, dirY     float32        //MAX SPEED & MOVEMENT DIRECTION OF MOVING BLOCKS
}

func main() {

	rl.InitWindow(0, 0, "moving spike block - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH := rl.GetScreenWidth(), rl.GetScreenHeight() //GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                            //SET WINDOW SIZE
	rl.SetWindowState(rl.FlagBorderlessWindowedMode)
	//rl.ToggleFullscreen() //UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //FIND SCREEN CENTER

	camera := rl.Camera2D{} //DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	tiles = rl.LoadTexture("tiles.png") //LOAD IMAGES FROM FILE

	makeRoom() //MAKE INITIAL ROOM SEE FUNCTION BELOW

	rl.SetTargetFPS(fps) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		//DEBUG TURNS OFF IMAGES & DISPLAYS RECTANGLE OUTLINES
		if rl.IsKeyPressed(rl.KeyF1) {
			debug = !debug
		}

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

		//MOVE PLAYER
		if rl.IsKeyDown(rl.KeyW) { //UP
			if checkplayermove(1) { //CHECK FOR OTHER WALL BLOCK COLLISIONS SEE FUNCTION BELOW
				playerRec.Y -= playerSpd
			}
		}
		if rl.IsKeyDown(rl.KeyS) { //DOWN
			if checkplayermove(3) {
				playerRec.Y += playerSpd
			}
		}
		if rl.IsKeyDown(rl.KeyA) { //LEFT
			if checkplayermove(4) {
				playerRec.X -= playerSpd
			}
		}
		if rl.IsKeyDown(rl.KeyD) { //RIGHT
			if checkplayermove(2) {
				playerRec.X += playerSpd
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		for i := 0; i < len(bloks); i++ {

			//MOVE SPIKES
			if bloks[i].spikes {
				if bloks[i].spikesOnOff {
					bloks[i].spikeRecs[0].Y++ //UP SPIKE
					bloks[i].spikeRecs[1].X-- //RIGHT SPIKE
					bloks[i].spikeRecs[2].Y-- //DOWN SPIKE
					bloks[i].spikeRecs[3].X++ //LEFT SPIKE
					//IF SPIKE BLOCK REACHES BORDER CHANGE DIRECTION
					if bloks[i].spikeRecs[0].Y >= bloks[i].rec.Y {
						bloks[i].spikeRecs[0].Y = bloks[i].rec.Y
						bloks[i].spikesOnOff = false
					}
				} else {
					bloks[i].spikeRecs[0].Y-- //UP SPIKE
					bloks[i].spikeRecs[1].X++ //RIGHT SPIKE
					bloks[i].spikeRecs[2].Y++ //DOWN SPIKE
					bloks[i].spikeRecs[3].X-- //LEFT SPIKE
					//IF SPIKE BLOCK REACHES BORDER CHANGE DIRECTION
					if bloks[i].spikeRecs[0].Y <= bloks[i].rec.Y-bloks[i].spikeRecs[0].Height {
						bloks[i].spikeRecs[0].Y = bloks[i].rec.Y - bloks[i].spikeRecs[0].Height
						bloks[i].spikesOnOff = true
					}
				}

				//DRAW SPIKES
				for j := 0; j < len(bloks[i].spikeRecs); j++ {
					drawRec := bloks[i].spikeRecs[j]
					drawRec.X += drawRec.Width / 2
					drawRec.Y += drawRec.Height / 2

					//DRAW LINES NOT IMAGES IF DEBUG IS ON
					if debug {
						rl.DrawRectangleLinesEx(bloks[i].spikeRecs[j], 1, rl.Red)
					} else {
						rl.DrawTexturePro(tiles, spikeTile, drawRec, rl.NewVector2(drawRec.Width/2, drawRec.Height/2), float32(j)*90, rl.White)
					}
					//CHECK COLLISION OF SPIKE REC VERSUS PLAYER
					if rl.CheckCollisionRecs(bloks[i].spikeRecs[j], playerRec) && playerCollisT == 0 {
						playerCollisT = fps * 2
					}
				}
			}

			//DRAW WALL BLOKS
			if debug {
				rl.DrawRectangleLinesEx(bloks[i].rec, 1, rl.Orange)
			} else {
				rl.DrawTexturePro(tiles, bloks[i].img, bloks[i].rec, rl.Vector2Zero(), 0, rl.White)
			}

			//DRAW PLAYER
			if playerCollisT > 0 { //IF COLLISION TIMER LARGER THAN ZERO DECREASE TIMER & DRAW RED
				if debug {
					rl.DrawRectangleLinesEx(playerRec, 2, rl.Red)
				} else {
					rl.DrawRectangleRec(playerRec, rl.Red)
				}
				playerCollisT--
			} else { //DRAW GREEN PLAYER IF COLLISION HAS NOT HAPPENED
				if debug {
					rl.DrawRectangleLinesEx(playerRec, 2, rl.Green)
				} else {
					rl.DrawRectangleRec(playerRec, rl.Green)
				}
			}

			//MOVE SPIKE BLOK
			if bloks[i].dirX != 0 || bloks[i].dirY != 0 {
				moveBlok(i) //SEE FUNCTION BELOW
			}
		}

		//DRAW INNER REC DEBUG
		if debug {
			rl.DrawRectangleLinesEx(innerRec, 2, rl.Magenta)
		}

		rl.EndMode2D()

		//DRAW TEXT
		rl.DrawText("camera zoom "+fmt.Sprintf("%.1f", camera.Zoom)+" press UP ARROW key to change", 10, 10, 20, rl.White)
		rl.DrawText("press SPACE key to make a new room", 10, 40, 20, rl.White)
		rl.DrawText("press F1 key to change to debug mode", 10, 70, 20, rl.White)
		rl.DrawText("WASD keys move player to collide with spikes", 10, 100, 20, rl.White)

		rl.EndDrawing()
	}

	rl.UnloadTexture(tiles) //UNLOAD FROM MEMORY

	rl.CloseWindow()
}
func moveBlok(blokNum int) { //MOVE SPIKE BLOCK FUNCTION

	canmove := true
	checkRec := bloks[blokNum].rec    //CREATE DUPLICATE CHECK RECTANGLE
	checkRec.X += bloks[blokNum].dirX //MOVE X TO NEXT POSITION
	checkRec.Y += bloks[blokNum].dirY //MOVE Y TO NEXT POSITION

	//CHECK FOR COLLISIONS OF CHECK RECTANGLE VERSUS OTHER WALL BLOCKS
	for i := 0; i < len(bloks); i++ {
		if i != blokNum { //MAKE SURE IT DOES NOT CHECK FOR COLLISIONS WITH ITSELF
			if rl.CheckCollisionRecs(checkRec, bloks[i].rec) {
				canmove = false
			}
		}
	}

	//CHECK FOR COLLISIONS OF CHECK RECTANGLE VERSUS PLAYER RECTANGLE
	if canmove {
		if rl.CheckCollisionRecs(checkRec, playerRec) {
			canmove = false
		}
	}

	//IF BLOCK CAN MOVE THEN MOVE TO NEXT POSITION
	if canmove {
		bloks[blokNum].rec = checkRec
		//IF BLOCK HAS SPIKES THEN MOVE SPIKE RECTANGLES
		if len(bloks[blokNum].spikeRecs) > 0 {
			for i := 0; i < len(bloks[blokNum].spikeRecs); i++ {
				bloks[blokNum].spikeRecs[i].X += bloks[blokNum].dirX
				bloks[blokNum].spikeRecs[i].Y += bloks[blokNum].dirY
			}
		}
	} else { //IF CANNOT MOVE THEN CHANGE DIRECTION
		if bloks[blokNum].dirX < 0 {
			bloks[blokNum].dirX = rF32(2, bloks[blokNum].spd)
		} else {
			bloks[blokNum].dirX = rF32(-bloks[blokNum].spd, -2)
		}
		if bloks[blokNum].dirY < 0 {
			bloks[blokNum].dirY = rF32(2, bloks[blokNum].spd)
		} else {
			bloks[blokNum].dirY = rF32(-bloks[blokNum].spd, -2)
		}
	}

}
func checkplayermove(direc int) bool {

	canmove := true
	checkRec := playerRec //CREATE DUPLICATE CHECK RECTANGLE
	switch direc {
	//MOVE CHECK RECTANGLE TO NEXT POSITION
	case 1: //UP
		checkRec.Y -= playerSpd
	case 2: //RIGHT
		checkRec.X += playerSpd
	case 3: //DOWN
		checkRec.Y += playerSpd
	case 4: //LEFT
		checkRec.X -= playerSpd
	}
	//CHECK FOR COLLISIONS WITH WALL BLOCKS
	for i := 0; i < len(bloks); i++ {
		if rl.CheckCollisionRecs(checkRec, bloks[i].rec) {
			canmove = false
		}
	}

	return canmove

}
func makeRoom() { //MAKES ROOM

	bloks = nil //CLEARS ANY PREVIOUSLY MADE ROOM

	Wnum := rInt(20, 31)      //NUMBER OF BLOCKS WIDE
	Hnum := rInt(20, 31)      //NUMBER OF BLOCKS HIGH
	W := float32(Wnum) * size //TOTAL WIDTH IN PIXELS
	H := float32(Hnum) * size //TOTAL HEIGHT IN PIXELS

	x := cnt.X - W/2 //LEFT X OF ROOM RECTANGLE
	y := cnt.Y - H/2 //TOP Y OF ROOM RECTANGLE

	innerRec = rl.NewRectangle(x, y, W, H) //INNER ROOM RECTANGLE

	playerRec = rl.NewRectangle(x+size, y+size, size*2, size*2) //CREATE PLAYER IN TOP LEFT CORNER

	x -= size                //MOVE X LEFT BY BLOCK SIZE FOR BORDER OF LEFT & RIGHT INNER REC SIDES
	x2 := x + W + (size * 2) //RIGHT WALL X
	x3 := x                  //X VALUE FOR ADDING BLOCKS
	y -= size                //MOVE Y UP BY BLOCK SIZE FOR BORDER OF TOP & BOTTOM INNER REC SIDES
	y2 := y + H              //BOTTOM WALL Y
	y3 := y2                 //Y VALUE FOR ADDING BLOCKS

	zblok := xblok{}     //EMPTY BLOCK STRUCT
	zblok.img = wallTile //DEFINE IMAGE

	for x3 < x2 {
		zblok.rec = rl.NewRectangle(x3, y, size, size) //DEFINE TOP WALL BLOCK
		bloks = append(bloks, zblok)                   //ADD TOP WALL BLOCK
		zblok.rec.Y += H + size                        //MOVE Y TO BOTTOM WALL POSITION
		bloks = append(bloks, zblok)                   //ADD BOTTOM WALL BLOCK
		x3 += size                                     //MOVE RIGHT ONE BLOCK
	}
	x3 = x
	for y3 > y {
		zblok.rec = rl.NewRectangle(x3, y3, size, size) //DEFINE LEFT WALL BLOCK
		bloks = append(bloks, zblok)                    //ADD LEFT WALL BLOCK
		zblok.rec.X += W + size                         //MOVE X TO RIGHT WALL POSITION
		bloks = append(bloks, zblok)                    //ADD RIGHT WALL BLOCK
		y3 -= size                                      //MOVE DOWN ONE BLOCK
	}

	//CENTER SPIKE BLOK
	zblok.rec = rl.NewRectangle(cnt.X-size/2, cnt.Y-size/2, size, size) //DEFINE CENTER BLOCK
	zblok.spikes = true                                                 //TURN ON SPIKES

	for i := 0; i < 4; i++ { //CREATE FOUR CENTER RECTANGLES SMALLER THAN MAIN BLOCK RECTANGLE
		spikeRec := zblok.rec
		spikeRec.X += size / 8
		spikeRec.Y += size / 8
		spikeRec.Width -= size / 4
		spikeRec.Height -= size / 4
		zblok.spikeRecs = append(zblok.spikeRecs, spikeRec)
	}
	zblok.spd = 8 //MAX MOVEMENT SPEED OF BLOCK
	for {
		//SET DIRECTIONS FOR X & Y MOVEMENT
		if flipcoin() {
			zblok.dirX = rF32(-zblok.spd, zblok.spd)
		} else {
			zblok.dirY = rF32(-zblok.spd, zblok.spd)
		}
		//MAKE SURE IT DOES NOT MOVE TO SLOWLY
		if getabs(zblok.dirX) > 2 && getabs(zblok.dirY) > 2 {
			break
		}
	}
	bloks = append(bloks, zblok) //ADD MOVE BLOCK TO SLICE
}

// SIMULATES A COIN FLIP
func flipcoin() bool {
	onoff := false
	choose := rInt(0, 100001)
	if choose > 50000 {
		onoff = true
	}
	return onoff
}

// RETURNS A RANDOM FLOAT32
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}

// RETURNS THE ABSOLUTE VALUE OF A NUMBER (NOT NEGATIVE)
func getabs(value float32) float32 {
	value2 := float64(value)
	value = float32(math.Abs(value2))
	return value
}

// RETURNS A RANDOM INTEGER
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}
