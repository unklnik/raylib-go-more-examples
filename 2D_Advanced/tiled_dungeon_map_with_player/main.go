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
	scrW, scrH          int            // SCREEN WIDTH & HEIGHT
	cnt                 rl.Vector2     // SCREEN CENTER
	rooms               []rl.Rectangle // SLICE OF ROOM RECTANGLES
	baseUnit            = float32(32)  // BASE UNIT FOR DETERMINING SIZES FOR TILING
	min, max            = 3, 12        // MIN MAX TILE SIZE OR ROOM SIDES
	borderRec           rl.Rectangle   // BORDER RECTANGLE TO ENSURE ROOMS REMAIN ON SCREEN
	camera              rl.Camera2D    // 2D CAMERA
	debug, colors       bool           // ON/OFF FOR DEBUG & COLORS
	floortile, walltile rl.Rectangle   // CURRENT DUNGEON FLOOR & WALL TILE IMAGES
	tiles               []tile         // SLICE OF CURRENT DUNGEON TILES
	imgs                rl.Texture2D   // IMAGE
	floorimgs, wallimgs []rl.Rectangle // SLICES OF ALL AVAILABLE FLOOR & WALL IMAGES
	player              xplayer        // DEFINE PLAYER
	frames              int32          // USED TO COUNT FRAMES
	delta               float32        //FRAME TIME
)

type tile struct { // TILE STRUCT
	im, rec rl.Rectangle // IMAGE
	fd      float32
	col     rl.Color
	solid   bool
}
type ximg struct { //IMAGE STRUCT FOR USE WITH PLAYER ANIMATIONS
	rec            rl.Rectangle
	frames, startX float32 //NUMBER OF ANIMATION FRAMES & START X POSITION OF ANIMATION LOOP
}
type xplayer struct { //PLAYER
	rec, collisrec rl.Rectangle
	im             rl.Rectangle
	walkimgs       []ximg
	direc          int
	spd            float32
}

func main() {

	rl.InitWindow(0, 0, "simple dungeon map - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowState(rl.FlagBorderlessWindowedMode)       //SET WINDOW STATE
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	camera.Zoom = 1.0 //SETS CAMERA ZOOM

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //SCREEN CENTER

	borderRec = rl.NewRectangle(0, 0, float32(scrW), float32(scrH)) //DEFINE BORDER RECTANGLE
	// MAKE BORDER REC SLIGHLTLY SMALLER SO ROOMS DONT REACH EDGES FOR TILING
	borderRec.X += baseUnit
	borderRec.Y += baseUnit
	borderRec.Width -= baseUnit * 2
	borderRec.Height -= baseUnit * 2

	makeimgs()   //CREATE IMAGES SEE FUNC BELOW
	makerooms()  //CREATE INITIAL SLICE OF ROOMS SEE FUNC BELOW
	makeplayer() //CREATE PLAYER SEE FUNC BELOW

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {
		delta = rl.GetFrameTime() //GET FRAME TIME FOR SMOOTHER MOVEMENT
		frames++                  //COUNT THE FRAMES

		inp()     //CAPTURE INPUT
		upanims() //UPDATE PLAYER ANIMATIONS

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		for i := 0; i < len(tiles); i++ { //RANGE OVER SLICE OF TILES & DRAW IMAGES
			if colors { //IF COLORS ON DRAW COLORS
				rl.DrawTexturePro(imgs, tiles[i].im, tiles[i].rec, rl.Vector2Zero(), 0, rl.Fade(tiles[i].col, tiles[i].fd))
			} else { //ELSE DRAW WHITE
				rl.DrawTexturePro(imgs, tiles[i].im, tiles[i].rec, rl.Vector2Zero(), 0, rl.Fade(rl.White, tiles[i].fd))
			}
		}

		if debug { //DRAWS THE UNDERLYING ROOM STRUCTURE
			rl.DrawRectangleLinesEx(borderRec, 2, rl.Magenta) // DRAW BORDER RECTANGLE
			for i := 0; i < len(rooms); i++ {                 // DRAW SLICE OF ROOM RECTANGLES
				rl.DrawRectangleRec(rooms[i], rl.Fade(rl.Green, 0.2))
				rl.DrawRectangleLinesEx(rooms[i], 1, rl.White)
				rl.DrawText("room "+fmt.Sprint(i), rooms[i].ToInt32().X+8, rooms[i].ToInt32().Y+8, 20, rl.White) // DRAW ROOM NUMBER TEXT
			}
		}

		//DRAW PLAYER
		shadowRec := player.rec //CREATE A DUPLICATE PLAYER REC & MOVE FOR USE AS SHADOW
		shadowRec.X -= 4
		shadowRec.Y += 4
		//DRAW SHADOW REC
		rl.DrawTexturePro(imgs, player.im, shadowRec, rl.Vector2Zero(), 0, rl.Fade(rl.Black, 0.7))
		//DRAW PLAYER REC
		rl.DrawTexturePro(imgs, player.im, player.rec, rl.Vector2Zero(), 0, rl.White)
		//IF DEBUG ON DRAW PLAYER IMG REC & PLAYER COLLISION REC
		if debug {
			rl.DrawRectangleLinesEx(player.rec, 2, rl.Green)
			rl.DrawRectangleLinesEx(player.collisrec, 2, rl.Red)
		}
		rl.EndMode2D()
		//MESSAGE TEXT
		rl.DrawText("PRESS SPACE TO MAKE NEW MAP / UP ARROW CHANGE ZOOM", 10, 10, 20, rl.White)
		rl.DrawText("RIGHT ARROW CHANGE COLORS / DOWN ARROW SHOW DEBUG", 10, 40, 20, rl.White)
		rl.DrawText("W A S D KEYS MOVE PLAYER", 10, 70, 20, rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
func inp() {

	//CHANGE ZOOM
	if rl.IsKeyPressed(rl.KeyUp) {
		if camera.Zoom == 1 {
			camera.Zoom = 2
		} else if camera.Zoom == 2 {
			camera.Zoom = 0.5
		} else if camera.Zoom == 0.5 {
			camera.Zoom = 1
		}
		camera.Target = cnt
		camera.Offset.X = float32(scrW / 2)
		camera.Offset.Y = float32(scrH / 2)
	}
	//SWITCH ON/OFF COLORS
	if rl.IsKeyPressed(rl.KeyRight) {
		colors = !colors
	}
	//SWITCH ON/OFF DEBUG
	if rl.IsKeyPressed(rl.KeyDown) {
		debug = !debug
	}
	//MAKE NEW ROOMS & RESET PLAYER
	if rl.IsKeyPressed(rl.KeySpace) {
		makerooms() //PRESS SPACE TO MAKE NEW SET OF ROOMS
		makeplayer()
	}
	//MOVE PLAYER
	if rl.IsKeyDown(rl.KeyW) { //UP
		player.direc = 1
		player.im = player.walkimgs[1].rec //CHANGE PLAYER IMAGE BASED ON DIRECTION
		moveplayer(player.direc)
	}
	if rl.IsKeyDown(rl.KeyD) { //RIGHT
		player.direc = 2
		player.im = player.walkimgs[0].rec
		moveplayer(player.direc)

	}
	if rl.IsKeyDown(rl.KeyS) { //DOWN
		player.direc = 3
		player.im = player.walkimgs[3].rec
		moveplayer(player.direc)

	}
	if rl.IsKeyDown(rl.KeyA) { //LEFT
		player.direc = 4
		player.im = player.walkimgs[2].rec
		moveplayer(player.direc)
	}

}
func upanims() {
	if frames%3 == 0 {
		switch player.direc { //UPDATE ANIMATION BASED ON PLAYER DIRECTION
		case 1: //UP
			player.walkimgs[1].rec.X += player.walkimgs[1].rec.Width
			//IF PLAYER IMG REC X IS LARGER THAN START IMAGE X + WIDTH * FRAMES THEN RESET TO START X
			if player.walkimgs[1].rec.X > player.walkimgs[1].startX+(player.walkimgs[1].frames*player.walkimgs[1].rec.Width) {
				player.walkimgs[1].rec.X = player.walkimgs[1].startX
			}
		case 2: //RIGHT
			player.walkimgs[0].rec.X += player.walkimgs[0].rec.Width
			if player.walkimgs[0].rec.X > player.walkimgs[0].startX+(player.walkimgs[0].frames*player.walkimgs[0].rec.Width) {
				player.walkimgs[0].rec.X = player.walkimgs[0].startX
			}
		case 3: //DOWN
			player.walkimgs[3].rec.X += player.walkimgs[3].rec.Width
			if player.walkimgs[3].rec.X > player.walkimgs[3].startX+(player.walkimgs[3].frames*player.walkimgs[3].rec.Width) {
				player.walkimgs[3].rec.X = player.walkimgs[3].startX
			}
		case 4: //LEFT
			player.walkimgs[2].rec.X += player.walkimgs[2].rec.Width
			if player.walkimgs[2].rec.X > player.walkimgs[2].startX+(player.walkimgs[2].frames*player.walkimgs[2].rec.Width) {
				player.walkimgs[2].rec.X = player.walkimgs[2].startX
			}
		}
	}

}
func moveplayer(direc int) {

	checkrec := player.collisrec //DUPLICATE PLAYER COLLISION REC
	//MOVE DUPLICATE REC TO PLAYERS NEXT POSITION + DELTA
	switch player.direc {
	case 1: //UP
		checkrec.Y -= player.spd + delta
	case 2: //RIGHT
		checkrec.X += player.spd + delta
	case 3: //DOWN
		checkrec.Y += player.spd + delta
	case 4: //LEFT
		checkrec.X -= player.spd + delta
	}

	canmove := true
	//RANGE OF TILES IF TILE IS SOLID & PLAYER COLLIDES THEN PLAYER CANNOT MOVE
	for i := 0; i < len(tiles); i++ {
		if tiles[i].solid {
			if rl.CheckCollisionRecs(checkrec, tiles[i].rec) {
				canmove = false
				break
			}
		}
	}
	//IF PLAYER CAN MOVE THEN MOVE PLAYER IMAGE REC & PLAYER COLLISION REC
	if canmove {
		switch player.direc {
		case 1: //UP
			player.collisrec.Y -= player.spd
			player.rec.Y -= player.spd + delta
		case 2: //RIGHT
			player.collisrec.X += player.spd
			player.rec.X += player.spd + delta
		case 3: //DOWN
			player.collisrec.Y += player.spd
			player.rec.Y += player.spd + delta
		case 4: //LEFT
			player.collisrec.X -= player.spd
			player.rec.X -= player.spd + delta
		}
	}

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

	maketiles()
}
func maketiles() {

	tiles = nil //CLEAR SLICE FOR RE-MAKING LEVEL WHILST GAME IS RUNNING

	floortile = floorimgs[rInt(0, len(floorimgs))] //SELECT RANDOM FLOOR TILE IMAGE
	walltile = wallimgs[rInt(0, len(wallimgs))]    //SELECT RANDOM WALL TILE IMAGE

	for i := 0; i < len(rooms); i++ {
		//WALL TILES
		x := rooms[i].X                      //START X
		y := rooms[i].Y                      //START Y
		x -= baseUnit                        //MOVE 1 BASE UNIT OUTSIDE ROOM BOUNDARY RECTANGLE
		y -= baseUnit                        //MOVE 1 BASE UNIT OUTSIDE ROOM BOUNDARY RECTANGLE
		for x <= rooms[i].X+rooms[i].Width { //TOP & BOTTOM ROOM REC WALLS
			ztile := tile{}
			ztile.im = walltile
			ztile.solid = true        //SET WALLS SOLID FOR COLLISIONS
			ztile.fd = rF32(0.7, 0.9) //CREATES LESS UNIFORM COLOR EFFECT
			ztile.rec = rl.NewRectangle(x, y, baseUnit, baseUnit)
			ztile.col = col_random()    //ADD A RANDOM COLOR
			if checkaddrec(ztile.rec) { //CHECK FUNC SEE BELOW
				tiles = append(tiles, ztile) //ADD ROOM REC TOP WALL TILE
			}
			ztile.rec.Y += rooms[i].Height + baseUnit //MOVE Y FOR BOTTOM WALL TILE SAME X VALUE
			ztile.col = col_random()                  //CHANGE COLOR
			if checkaddrec(ztile.rec) {
				tiles = append(tiles, ztile)
			}
			x += baseUnit //MOVE ALONG WALL
		}

		x = rooms[i].X
		y = rooms[i].Y
		x -= baseUnit
		for y <= rooms[i].Y+rooms[i].Height-baseUnit { //LEFT & RIGHT ROOM REC WALLS
			ztile := tile{}
			ztile.im = walltile
			ztile.solid = true
			ztile.fd = rF32(0.7, 0.9)
			ztile.rec = rl.NewRectangle(x, y, baseUnit, baseUnit)
			ztile.col = col_random()
			if checkaddrec(ztile.rec) {
				tiles = append(tiles, ztile)
			}
			ztile.rec.X += rooms[i].Width + baseUnit
			ztile.col = col_random()
			if checkaddrec(ztile.rec) {
				tiles = append(tiles, ztile)
			}
			y += baseUnit
		}
		//FLOOR TILES
		x = rooms[i].X
		y = rooms[i].Y
		for y < rooms[i].Y+rooms[i].Height { //FILL ROOM RECS WITH FLOOR TILES
			ztile := tile{}
			ztile.im = floortile
			ztile.fd = rF32(0.1, 0.4)
			ztile.rec = rl.NewRectangle(x, y, baseUnit, baseUnit)
			ztile.col = col_random()
			tiles = append(tiles, ztile)
			x += baseUnit
			//IF X REACHES ROOM RIGHT BOUNDARY MOVE Y DOWN 1 BASE UNIT & X BACK TO ROOM REC X
			if x >= rooms[i].X+rooms[i].Width {
				x = rooms[i].X
				y += baseUnit
			}
		}
	}
}
func makeimgs() {
	imgs = rl.LoadTexture("imgs.png")
	x := float32(0)
	y := float32(0)

	//ADD FLOOR IMAGE RECTANGLES TO SLICE
	for i := 0; i < 12; i++ {
		floorimgs = append(floorimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
	x = float32(0)
	y = float32(16)
	//ADD WALL IMAGE RECTANGLES TO SLICE
	for i := 0; i < 12; i++ {
		wallimgs = append(wallimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
}
func makeplayer() {
	player = xplayer{}
	size := baseUnit * 2 //PLAYER SIZE
	//CENTER PLAYER IN THE MIDDLE OF THE 1ST ROOM REC
	player.rec = rl.NewRectangle((rooms[0].X+rooms[0].Width/2)-size/2, (rooms[0].Y+rooms[0].Height/2)-size/2, size, size)
	//CREATE A COLLISION RECTANGLE THE SIZE OF THE DRAWN IMAGE
	player.collisrec = player.rec
	player.collisrec.X += size / 4
	player.collisrec.Y += size / 4
	player.collisrec.Width -= size / 2
	player.collisrec.Height -= size / 2
	player.spd = 4 //MOVEMENT SPEED

	//ADD PLAYER MOVEMENT ANIMATION IMAGES
	zimg := ximg{}
	zimg.frames = 8
	zimg.rec = rl.NewRectangle(0, 32, 32, 32)
	zimg.startX = zimg.rec.X
	player.walkimgs = append(player.walkimgs, zimg)
	zimg.rec.Y += zimg.rec.Height
	player.walkimgs = append(player.walkimgs, zimg)
	zimg.rec.Y += zimg.rec.Height
	player.walkimgs = append(player.walkimgs, zimg)
	zimg.rec.Y += zimg.rec.Height
	player.walkimgs = append(player.walkimgs, zimg)

	//SET START PLAYER IMAGE
	player.im = player.walkimgs[0].rec

}
func checkaddrec(rec rl.Rectangle) bool {
	canadd := true
	//RANGE OVER ROOM RECS IF TILE REC IS INSIDE A ROOM REC THEN DON'T ADD
	for i := 0; i < len(rooms); i++ {
		if rl.CheckCollisionRecs(rec, rooms[i]) {
			canadd = false
		}
	}
	return canadd
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

// RETURNS RANDOM FLOAT32
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}

// RETURNS RANDOM COLOR
func col_random() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}
