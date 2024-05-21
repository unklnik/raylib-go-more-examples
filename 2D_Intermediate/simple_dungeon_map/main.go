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
	scrW, scrH int            // SCREEN WIDTH & HEIGHT
	cnt        rl.Vector2     // SCREEN CENTER
	rooms      []rl.Rectangle // SLICE OF ROOM RECTANGLES
	baseUnit   = float32(32)  // BASE UNIT FOR DETERMINING SIZES FOR TILING
	min, max   = 3, 12        // MIN MAX TILE SIZE OR ROOM SIDES
	borderRec  rl.Rectangle   // BORDER RECTANGLE TO ENSURE ROOMS REMAIN ON SCREEN
)

func main() {

	rl.InitWindow(0, 0, "simple dungeon map - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowState(rl.FlagBorderlessWindowedMode)       //SET WINDOW STATE
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //SCREEN CENTER

	borderRec = rl.NewRectangle(0, 0, float32(scrW), float32(scrH)) //DEFINE BORDER RECTANGLE
	// MAKE BORDER REC SLIGHLTLY SMALLER SO ROOMS DONT REACH EDGES FOR TILING
	borderRec.X += baseUnit
	borderRec.Y += baseUnit
	borderRec.Width -= baseUnit * 2
	borderRec.Height -= baseUnit * 2

	makerooms() //CREATE INITIAL SLICE OF ROOMS SEE FUNC BELOW

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeySpace) {
			makerooms() //PRESS SPACE TO MAKE NEW SET OF ROOMS
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		for i := 0; i < len(rooms); i++ { // DRAW SLICE OF ROOM RECTANGLES
			rl.DrawRectangleRec(rooms[i], rl.Fade(rl.Green, 0.2))
			rl.DrawRectangleLinesEx(rooms[i], 1, rl.White)
			rl.DrawText("room "+fmt.Sprint(i), rooms[i].ToInt32().X+8, rooms[i].ToInt32().Y+8, 20, rl.White) // DRAW ROOM NUMBER TEXT
		}

		rl.DrawRectangleLinesEx(borderRec, 2, rl.Magenta) // DRAW BORDER RECTANGLE

		rl.EndMode2D()

		rl.DrawText("PRESS SPACE TO MAKE ANOTHER MAP", 10, 10, 20, rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func makerooms() {
	rooms = nil            // CLEAR ROOM SLICE
	Wnum := rInt(min, max) // NUMBER TO TILES IN ROOM WALL WIDTH
	Hnum := rInt(min, max) // NUMBER OF TILES IN ROOM WALL HEIGHT

	W := float32(Wnum) * baseUnit                                      // CALCULATE WIDTH
	H := float32(Hnum) * baseUnit                                      // CALCULATE HEIGHT
	rooms = append(rooms, rl.NewRectangle(cnt.X-W/2, cnt.Y-H/2, W, H)) // CREATE CENTER ROOM REC

	num := rInt(7, 15) // RANDOM NUMBER OF ROOMS
	//num = 1
	countbreak := 100 // COUNT TO BREAK LOOP IF ROOMS CANNOT BE FOUND
	chooseroom := 0   // NUMBER OF BASE ROOM THAT NEXT ROOM IS ATTACHED TO

	for num > 0 {

		Wnum = rInt(min, max)
		Hnum = rInt(min, max)

		W = float32(Wnum) * baseUnit
		H = float32(Hnum) * baseUnit

		if len(rooms) > 1 { // IF MORE THAN ONE ROOM RANDOMLY CHOOSE NEXT BASE ROOM FROM EXISTING ROOMS
			chooseroom = rInt(0, len(rooms))
		}
		x := rooms[chooseroom].X // SET X VALUE FOR NEXT REC FROM BASE REC
		y := rooms[chooseroom].Y // SET Y VALUE FOR NEXT REC FROM BASE REC

		choose := rInt(1, 5) // CHOOSE SIDE OF BASE REC TO ATTACH NEXT REC TO

		numWprevRoom := int(rooms[chooseroom].Width / baseUnit)  // DETERMINE NUMBER OF TILES IN PREVIOUS ROOM WIDTH
		numHprevRoom := int(rooms[chooseroom].Height / baseUnit) // DETERMINE NUMBER OF TILES IN PREVIOUS ROOM HEIGHT

		switch choose {
		case 1: //ABOVE
			y -= H                            //MOVE NEW REC ABOVE
			if W <= rooms[chooseroom].Width { // IF WIDTH SMALLER THAN PREV REC
				change := rInt(-(Wnum - 1), Wnum) // DETERMINE RANDOM X MOVEMENT LEFT/RIGHT LEAVING A ONE TILE SPACE FOR DOOR/PASSAGE
				x += float32(change) * baseUnit
			} else { // IF WIDTH LARGER THAN PREV REC
				change := rInt(-(Wnum - 1), numWprevRoom) // DETERMINE RANDOM X MOVEMENT LEFT/RIGHT LEAVING A ONE TILE SPACE FOR DOOR/PASSAGE
				x += float32(change) * baseUnit
			}
			rec := rl.NewRectangle(x, y, W, H)
			if checkaddroom(rec) { // CHECK IF REC COLLIDES WITH OTHER RECS OR BORDER FUNCTION BELOW
				rooms = append(rooms, rec)
				num--
			}
		case 2: //RIGHT
			x += rooms[chooseroom].Width       //MOVE NEW REC RIGHT
			if H <= rooms[chooseroom].Height { // IF HEIGHT SMALLER THAN PREV REC
				change := rInt(-(Hnum - 1), Hnum) // DETERMINE RANDOM Y MOVEMENT UP/DOWN LEAVING A ONE TILE SPACE FOR DOOR/PASSAGE
				y += float32(change) * baseUnit
			} else {
				change := rInt(-(Hnum - 1), numHprevRoom) // DETERMINE RANDOM Y MOVEMENT UP/DOWN LEAVING A ONE TILE SPACE FOR DOOR/PASSAGE
				y += float32(change) * baseUnit
			}
			rec := rl.NewRectangle(x, y, W, H)
			if checkaddroom(rec) { // CHECK IF REC COLLIDES WITH OTHER RECS OR BORDER FUNCTION BELOW
				rooms = append(rooms, rec)
				num--
			}
		case 3: //BELOW
			y += rooms[chooseroom].Height
			if W <= rooms[chooseroom].Width {
				change := rInt(-(Wnum - 1), Wnum)
				x += float32(change) * baseUnit
			} else {
				change := rInt(-(Wnum - 1), numWprevRoom)
				x += float32(change) * baseUnit
			}
			rec := rl.NewRectangle(x, y, W, H)
			if checkaddroom(rec) {
				rooms = append(rooms, rec)
				num--
			}
		case 4: //LEFT
			x -= W
			if H <= rooms[chooseroom].Height {
				change := rInt(-(Hnum - 1), Hnum)
				y += float32(change) * baseUnit
			} else {
				change := rInt(-(Hnum - 1), numHprevRoom)
				y += float32(change) * baseUnit
			}
			rec := rl.NewRectangle(x, y, W, H)
			if checkaddroom(rec) {
				rooms = append(rooms, rec)
				num--
			}

		}
		countbreak--
		if countbreak == 0 {
			break
		}

	}
}
func checkaddroom(rec rl.Rectangle) bool {
	canadd := true

	// CREATE VECTOR2 OF 4 CORNERS OF NEW ROOM RECTANGLE
	v1 := rl.NewVector2(rec.X, rec.Y)
	v2 := v1
	v2.X += rec.Width
	v3 := v2
	v3.Y += rec.Height
	v4 := v3
	v4.X -= rec.Width

	if !rl.CheckCollisionPointRec(v1, borderRec) || !rl.CheckCollisionPointRec(v2, borderRec) || !rl.CheckCollisionPointRec(v3, borderRec) || !rl.CheckCollisionPointRec(v4, borderRec) {
		canadd = false // IF A CORNER EXITS BORDER DON'T ADD
	}

	if canadd {
		for i := 0; i < len(rooms); i++ {
			if rl.CheckCollisionRecs(rec, rooms[i]) {
				canadd = false // NEW ROOM REC COLLIDES WITH EXISTING ROOM REC DON'T ADD
			}
		}
	}

	return canadd
}

// RETURNS RANDOM INTEGER
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}
