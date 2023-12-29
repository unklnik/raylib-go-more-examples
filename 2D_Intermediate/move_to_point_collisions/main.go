package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	scrW, scrH              int            // SCREEN WIDTH & HEIGHT
	player, cnt, clickPoint rl.Vector2     //PLAYER, SCREEN CENTER, MOUSE CLICK VECTOR2
	nextPosition            rl.Rectangle   //COLLISION RECTANGLE FOR MOVE TO POINT
	playerSpeed             = float32(8)   //MAX SPEED PLAYER MOVES
	playerDirX, playerDirY  float32        //X & Y DIRECTION SPEEDS
	playerSize              = float32(32)  //SIZE OF PLAYER RECTANGLE
	playerRec               rl.Rectangle   //PLAYER RECTANGLE
	bloks                   []rl.Rectangle //SLICE OF COLLISION BLOKS
)

func main() {

	rl.InitWindow(0, 0, "move to point with collisions - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowState(rl.FlagBorderlessWindowedMode)       //SET WINDOW STATE
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //SCREEN CENTER
	player, clickPoint = cnt, cnt                         //SET INITIAL PLAYER & CLICKPOINT POSITIONS
	//DEFINE PLAYER RECTANGLE AT CENTRE
	playerRec = rl.NewRectangle(player.X-playerSize/2, player.Y-playerSize/2, playerSize, playerSize)

	makebloks() //MAKES A SLICE OF BLOKS FOR COLLISIONS SEE FUNCTION BELOW

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		//IF LEFT MOUSE CLICKED CREATE COLLISION RECTANGLE FROM CLICK POINT
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			size := float32(16)
			clickPoint = rl.GetMousePosition()
			nextPosition = rl.NewRectangle(clickPoint.X-size/2, clickPoint.Y-size/2, size, size)
		}

		//IF CLICK POINT IS NOT CENTER & PLAYER IS NOT COLLIDING WITH NEXT POSITION RECTANGLE MOVE
		if clickPoint != cnt && !rl.CheckCollisionPointRec(player, nextPosition) {

			diffX := absdiff(player.X, clickPoint.X) //ABSOLUTE DISTANCE BETWEEN X POINTS
			diffY := absdiff(player.Y, clickPoint.Y) //ABSOLUTE DISTANCE BETWEEN Y POINTS

			if diffX > diffY { //IF X DISTANCE IS LARGER THAN Y DISTANCE
				//DIRECTION X SPEED = PLAYER MAX SPEED
				playerDirX = playerSpeed
				//DIRECTION Y SPEED = Y DISTANCE DIVIDED BY (X DISTANCE DIVIDED BY MAX SPEED)
				playerDirY = diffY / (diffX / playerSpeed)
			} else {
				playerDirY = playerSpeed
				playerDirX = diffX / (diffY / playerSpeed)
			}

			//CHANGES TO NEGATIVE IF CLICKPOINT X IS LEFT OF PLAYER
			if clickPoint.X < player.X {
				playerDirX = -playerDirX
			}
			//CHANGES TO NEGATIVE IF CLICKPOINT Y IS ABOVE PLAYER
			if clickPoint.Y < player.Y {
				playerDirY = -playerDirY
			}

			if checkmove("x") { //CHECKS FOR X MOVEMENT AGAINST BLOK COLLISIONS SEE FUNCTION BELOW
				player.X += playerDirX //MOVE X DIRECTION
			}
			if checkmove("y") { //CHECKS FOR Y MOVEMENT AGAINST BLOK COLLISIONS SEE FUNCTION BELOW
				player.Y += playerDirY //MOVE Y DIRECTION
			}

			playerRec = rl.NewRectangle(player.X-playerSize/2, player.Y-playerSize/2, playerSize, playerSize) //DEFINE RECTANGLE IN NEW POSITION

		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		//DRAW THE SLICE OF COLLISION BLOKS
		for i := 0; i < len(bloks); i++ {
			rl.DrawRectangleRec(bloks[i], rl.Fade(rl.SkyBlue, 0.4))
			rl.DrawRectangleLinesEx(bloks[i], 4, rl.SkyBlue)
		}

		if clickPoint != cnt {
			rl.DrawRectangleLinesEx(nextPosition, 2, rl.Magenta) //DRAW COLLISION RECTANGLE
		}

		//DRAW PLAYER RECTANGLE
		rl.DrawRectangleRec(playerRec, rl.Fade(rl.Green, 0.5))
		rl.DrawRectangleLinesEx(playerRec, 4, rl.Green)

		rl.DrawText("LEFT MOUSE CLICK ON THE SCREEN TO MOVE TO POINT", 10, 10, 20, rl.White)

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// CHECKS FOR NEXT X & Y MOVEMENTS AGAINST SLICE OF BLOKS
func checkmove(direc string) bool {

	canmove := true       //RETURN VALUE
	checkRec := playerRec //RECTANGLE FOR CHECKING NEXT MOVE

	switch direc {
	case "x":
		checkRec.X += playerDirX //MOVE CHECK RECTANGLE X TO NEXT POSITION
	case "y":
		checkRec.Y += playerDirY //MOVE CHECK RECTANGLE Y TO NEXT POSITION
	}

	//CHECK NEXT POSITION OF PLAYER AGAINST SLICE OF COLLISION BLOKS
	for i := 0; i < len(bloks); i++ {
		if rl.CheckCollisionRecs(checkRec, bloks[i]) {
			canmove = false
		}
	}

	return canmove
}

// MAKES A SLICE OF RECTANGLES IN RANDOM POSITIONS
func makebloks() {

	num := rInt(20, 31)
	for {
		canadd := true        //FOR CHECKING COLLISIONS AGAINST OTHER BLOKS ALREADY ADDED TO SLICE
		countbreak := 100     //IN CASE THE FOR LOOP GETS STUCK
		size := rF32(32, 128) //RANDOM BLOCK SIZE SEE FUNCTION BELOW
		//DEFINE BLOK RECTANGLE
		rec := rl.NewRectangle(rF32(0, float32(scrW)-size), rF32(0, float32(scrH)-size), size, size)
		//CHECK FOR COLLISIONS WITH PLAYER REC
		if rl.CheckCollisionRecs(rec, playerRec) {
			canadd = false //IF COLLIDES WITH PLAYER DON'T ADD
		}
		//IF DOES NOT COLLIDE WITH PLAYER THEN CHECK FOR EXISTING BLOK COLLISIONS
		if canadd {
			//IF SLICE OF BLOKS IS LARGER THAN 1 CHECK FOR COLLISIONS OF EXISTING BLOKS IN SLICE
			if len(bloks) > 1 {
				for i := 0; i < len(bloks); i++ {
					if rl.CheckCollisionRecs(rec, bloks[i]) {
						canadd = false //IF COLLIDES WITH EXISTING BLOK DON'T ADD
					}
				}
			}
		}

		if canadd {
			bloks = append(bloks, rec) //IF NO COLLISIONS ADD TO SLICE
			num--                      //REDUCE NUMBER OF BLOKS LEFT TO ADD
		}

		countbreak-- //REDUCE COUNTBREAK IN CASE OF ERRORS

		if num == 0 || countbreak == 0 { //IF ALL BLOKS ADDED OR COUNTBREAK REACHED BREAK
			break
		}

	}

}

// FUNCTION THAT RETURNS RANDOM FLOAT32 BETWEEN MIN & MAX VALUES
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}

// FUNCTION THAT RETURNS RANDOM INTEGER BETWEEN MIN & MAX VALUES
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// FUNCTION TO CALCULATE ABSOLUTE DIFFERENCE BETWEEN TWO VALUES FOR NEGATIVE VALUES
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

// FUNCTION TO CALCULATE ABSOLUTE VALUE FOR NEGATIVE VALUES
func getabs(value float32) float32 {
	value2 := float64(value)
	value = float32(math.Abs(value2))
	return value
}
