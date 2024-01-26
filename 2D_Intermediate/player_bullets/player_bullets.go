package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	bullets                        []xbullet     //SLICE OF BULLET STRUCTS
	shootTarget, cursorCam, cursor rl.Vector2    //TARGET, CURSOR CAMERA VS SCREEN POSITION, CURSOR
	playerRec                      rl.Rectangle  //PLAYER RECTANGLE
	spd                            = float32(10) //MAX SPEED
	attackT                        int32         //PAUSE BETWEEN BULLETS
	fps                            = int32(60)   //FRAMES PER SECOND
	camera                         rl.Camera2D   //CAMERA
	borderRec                      rl.Rectangle  //BOUNCING & BORDER RECTANGLES
	cntr                           rl.Vector2    //CENTER OF SCREEN
	scrW, scrH                     int
)

// STRUCT OF BULLET CONTAINING REC, DIRECTION X & Y, OFF BOOL
type xbullet struct {
	rec        rl.Rectangle
	dirX, dirY float32
	off        bool
}

func main() {

	rl.InitWindow(0, 0, "player bullets - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() //GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                           //SET WINDOW SIZE
	rl.SetWindowState(rl.FlagBorderlessWindowedMode)
	//rl.ToggleFullscreen() //UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	rl.HideCursor() //HIDE THE STANDARD MOUSE CURSOR

	cntr = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //CALCULATE CENTER

	borderRec = rl.NewRectangle(cntr.X-float32(scrW/4), cntr.Y-float32(scrH/4), float32(scrW/2), float32(scrH/2)) //DEFINE BORDER RECTANGLE

	siz := float32(32)                                                //PLAYER SIZE
	playerRec = rl.NewRectangle(cntr.X-siz/2, cntr.Y-siz/2, siz, siz) //INITIAL PLAYER RECTANGLE

	camera.Zoom = 1.5                   //SETS CAMERA ZOOM
	camera.Target = cntr                //SET CAMERA TARGET
	camera.Offset.X = float32(scrW / 2) //ADJUST CAMERA FOR ZOOM
	camera.Offset.Y = float32(scrH / 2) //ADJUST CAMERA FOR ZOOM

	rl.SetTargetFPS(fps) //NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		cursor = rl.GetMousePosition()                    //GET MOUSE POSITION
		cursorCam = rl.GetScreenToWorld2D(cursor, camera) //GET MOUSE POSITION IN CAMERA SPACE WITH ZOOM

		upBullets() //UPDATE BULLET MOVEMENTS
		input()     //CAPTURE INPUT

		//TIMER
		if attackT > 0 { //PAUSE BETWEEN SHOTS TIMER
			attackT--
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		rl.DrawRectangleLinesEx(borderRec, 4, rl.Green) //DRAWS BORDER REC

		rl.DrawRectangleLinesEx(playerRec, 8, rl.Magenta) //DRAW PLAYER REC

		//DRAW BULLETS
		for i := 0; i < len(bullets); i++ {
			rl.DrawRectangleLinesEx(bullets[i].rec, 2, rl.Yellow)
		}

		//DRAW CIRCLE TARGET INSTEAD OF CURSOR
		rl.DrawCircleLines(int32(cursorCam.X), int32(cursorCam.Y), 10, rl.Red)

		rl.EndMode2D()

		rl.DrawText("W A S D keys move", 10, 10, 20, rl.White)
		rl.DrawText("left mouse to shoot", 10, 40, 20, rl.White)
		rl.DrawText("up arrow key change zoom", 10, 70, 20, rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
func input() {
	//INPUT KEYS FOR PLAYER MOVEMENT SEE MOVEPLAYER FUNCTION
	if rl.IsKeyDown(rl.KeyW) {
		movePlayer(1)

	} else if rl.IsKeyDown(rl.KeyS) {
		movePlayer(3)
	}
	if rl.IsKeyDown(rl.KeyD) {
		movePlayer(2)
	} else if rl.IsKeyDown(rl.KeyA) {
		movePlayer(4)
	}

	//CREATE BULLET IF ATTACK TIMER IS ZERO
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && attackT == 0 {
		attackT = fps / 4
		shootTarget = cursorCam //POSITION FOR BULLET AIMING
		shoot()
	}

	//CHANGE ZOOM
	if rl.IsKeyPressed(rl.KeyUp) {
		if camera.Zoom == 2 {
			camera.Zoom = 1
		} else if camera.Zoom == 1.5 {
			camera.Zoom = 2
		} else if camera.Zoom == 1 {
			camera.Zoom = 1.5
		}
		camera.Target = cntr
		camera.Offset.X = float32(scrW / 2)
		camera.Offset.Y = float32(scrH / 2)
	}
}

// CREATE BULLET FUNCTION
func shoot() {
	zbullet := xbullet{}
	zbullet.rec = playerRec //DUPLICATE PLAYER RECTANGLE FOR BULLET

	//MAKE DUPLICATE RECTANGLE SMALLER
	zbullet.rec.X += playerRec.Width / 2
	zbullet.rec.Y += playerRec.Height / 2
	zbullet.rec.Width = zbullet.rec.Width / 2
	zbullet.rec.Height = zbullet.rec.Height / 2

	//CALCULATE X & Y SPEED TO MOVE TO SHOOT TARGET
	playerCntr := rl.NewVector2(playerRec.X+playerRec.Width/2, playerRec.Y+playerRec.Height/2)
	diffX := absdiff(playerCntr.X, shootTarget.X) //GET ABSOLUTE X DISTANCE FUNCTION BELOW
	diffY := absdiff(playerCntr.Y, shootTarget.Y) //GET ABSOLUTE Y DISTANCE FUNCTION BELOW

	if diffX > diffY {
		zbullet.dirX = spd                            //IF DIFFERENCE X IS LARGER X IS FULL SPEED
		zbullet.dirY = diffY / (diffX / zbullet.dirX) //CALCULATE Y SPEED
	} else {
		zbullet.dirY = spd                            //IF DIFFERENCE Y IS LARGER Y IS FULL SPEED
		zbullet.dirX = diffX / (diffY / zbullet.dirY) //CALCULATE X SPEED
	}

	//IF TARGET IS BEHIND PLAYER CHANGE X DIRECTION TO NEGATIVE
	if playerCntr.X > shootTarget.X {
		zbullet.dirX = -zbullet.dirX
	}

	//IF TARGET IS ABOVE PLAYER CHANGE Y DIRECTION TO NEGATIVE
	if playerCntr.Y > shootTarget.Y {
		zbullet.dirY = -zbullet.dirY
	}

	//ADD BULLET TO SLICE
	bullets = append(bullets, zbullet)
}

func upBullets() {

	clear := false //TO CLEAR BULLETS IF COLLISIONS
	for i := 0; i < len(bullets); i++ {

		if !bullets[i].off {
			checkRec := bullets[i].rec    //DUPLICATE RECTANGLE FOR NEXT COLLISIONS
			checkRec.X += bullets[i].dirX //MOVE DUPLICATE TO NEXT POSITION
			checkRec.Y += bullets[i].dirY

			//VECTOR 2 POINTS OF FOUR CORNERS OF BULLET RECTANGLE
			v1 := rl.NewVector2(checkRec.X, checkRec.Y)
			v2 := v1
			v2.X += bullets[i].rec.Width
			v3 := v2
			v3.Y += bullets[i].rec.Height
			v4 := v3
			v4.X -= bullets[i].rec.Width

			//CHECK IF VECTOR 2 HAS EXITED BORDER
			canmove := true
			if !rl.CheckCollisionPointRec(v1, borderRec) || !rl.CheckCollisionPointRec(v2, borderRec) || !rl.CheckCollisionPointRec(v3, borderRec) || !rl.CheckCollisionPointRec(v4, borderRec) {
				canmove = false
			}

			if canmove {
				bullets[i].rec = checkRec //IF NO EXITS MOVE BULLET
			} else {
				bullets[i].off = true //IF EXITED TURN BULLET OFF
				clear = true
			}
		}
	}

	//IF CLEAR IS ON REMOVE ALL OFF BULLETS FROM SLICE
	if clear {
		for i := 0; i < len(bullets); i++ {
			if bullets[i].off {
				bullets = remBullet(bullets, i)
			}
		}
	}

}
func movePlayer(direc int) {

	checkRec := playerRec //DUPLICATE PLAYER RECTANGLE

	//MOVE DUPLICATE IN THE DIRECTION OF KEYPRESS
	switch direc {
	case 1: //UP
		checkRec.Y -= spd
	case 2: //RIGHT
		checkRec.X += spd
	case 3: //DOWN
		checkRec.Y += spd
	case 4: //LEFT
		checkRec.X -= spd
	}

	//VECTOR 2 POINTS OF FOUR CORNERS OF PLAYER RECTANGLE
	v1 := rl.NewVector2(checkRec.X, checkRec.Y)
	v2 := v1
	v2.X += playerRec.Width
	v3 := v2
	v3.Y += playerRec.Height
	v4 := v3
	v4.X -= playerRec.Width

	//CHECK IF VECTOR 2 HAS EXITED BORDER
	canmove := true
	if !rl.CheckCollisionPointRec(v1, borderRec) || !rl.CheckCollisionPointRec(v2, borderRec) || !rl.CheckCollisionPointRec(v3, borderRec) || !rl.CheckCollisionPointRec(v4, borderRec) {
		canmove = false
	}

	if canmove { //IF NO EXITS MOVE PLAYER
		playerRec = checkRec
	}

}

// REMOVES BULLET FROM SLICE
func remBullet(slice []xbullet, s int) []xbullet {
	return append(slice[:s], slice[s+1:]...)
}

// GET ABSOLUTE DIFFERENCE
func absdiff(num1, num2 float32) float32 {
	num := float32(0)
	if num1 == num2 {
		num = 0
	} else {
		if num1 <= 0 && num2 <= 0 {
			num1 = getabs(num1)
			num2 = getabs(num2)
			if num1 > num2 {
				num = num1 - num2
			} else {
				num = num2 - num1
			}
		} else if num1 <= 0 && num2 >= 0 {
			num = num2 + getabs(num1)
		} else if num2 <= 0 && num1 >= 0 {
			num = num1 + getabs(num2)
		} else if num2 >= 0 && num1 >= 0 {
			if num1 > num2 {
				num = num1 - num2
			} else {
				num = num2 - num1
			}
		}
	}
	return num
}

// GET ABSOLUTE VALUE
func getabs(value float32) float32 {
	value2 := float64(value)
	value = float32(math.Abs(value2))
	return value
}
