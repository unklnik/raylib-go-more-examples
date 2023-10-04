package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	velX, velY           float32       // X & Y MOVEMENT SPEED
	maxVel               = float32(10) // MAX SPEED FOR USE IN RANDOM GENERATOR
	bounceRec, borderRec rl.Rectangle  // BOUNCING & BORDER RECTANGLES
	cntr                 rl.Vector2    // CENTER OF SCREEN
	timer                int           // COLOUR CHANGE TIMER ON COLLISION
)

func main() {

	rl.InitWindow(0, 0, "point collisions - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH := rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                            // SET WINDOW SIZE

	cntr = rl.NewVector2(float32(scrW/2), float32(scrH/2))                                                        // CALCULATE CENTER
	borderRec = rl.NewRectangle(cntr.X-float32(scrW/4), cntr.Y-float32(scrH/4), float32(scrW/2), float32(scrH/2)) // DEFINE BORDER RECTANGLE
	bounceRec = rl.NewRectangle(cntr.X-32, cntr.Y-32, 64, 64)                                                     // DEFINE BOUNCE RECTANGLE

	// SEE RF32 FUNCTION END OF CODE FINDS A RANDOM FLOAT32 BETWEEN TWO VALUES
	velX = rF32(-maxVel, maxVel) // FINDS A RANDOM X AXIS MOVEMENT SPEED
	velY = rF32(-maxVel, maxVel) // FINDS A RANDOM Y AXIS MOVEMENT SPEED

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		checkCollisionsMove() // SEE FUNCTION BELOW - CHECKS FOR BORDER COLLISIONS & MOVES

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		rl.DrawRectangleLinesEx(borderRec, 4, rl.Green) // DRAWS BORDER REC

		if timer > 0 {
			rl.DrawRectangleRec(bounceRec, rl.Red) //DRAWS RED BOUNCE REC IF TIMER IS LARGER THAN ZERO
			timer--                                // DECREASES TIMER
		} else {
			rl.DrawRectangleRec(bounceRec, rl.Green) //DRAWS GREEN BOUNCE REC IF TIMER IS ZERO
		}

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// CHECKS FOR BOUNCE REC CORNER COLLISIONS & MOVES BOUNCE REC
func checkCollisionsMove() {

	checkRec := bounceRec // DUPLICATES THE BOUNCE REC TO USE AS A COLLISION CHECK
	checkRec.X += velX    // MOVES DUPLICATE REC BY VELX
	checkRec.Y += velY    // MOVES DUPLICATE REC BY VELY

	v1 := rl.NewVector2(checkRec.X, checkRec.Y) // CREATES A CHECK POINT TOP LEFT CORNER
	v2 := v1
	v2.X += checkRec.Width // CREATES A CHECK POINT TOP RIGHT CORNER
	v3 := v2
	v3.Y += checkRec.Height // CREATES A CHECK POINT BOTTOM RIGHT CORNER
	v4 := v1
	v4.Y += checkRec.Height // CREATES A CHECK POINT BOTTOM LEFT CORNER

	// CHECKS IF ANY OF THE CREATED CHECK POINTS HAVE EXITED THE BORDER REC
	canmove := true
	if !rl.CheckCollisionPointRec(v1, borderRec) || !rl.CheckCollisionPointRec(v2, borderRec) || !rl.CheckCollisionPointRec(v3, borderRec) || !rl.CheckCollisionPointRec(v4, borderRec) {
		canmove = false
	}

	if canmove { // NO POINTS EXITED - BOUNCE REC MOVES TO CHECK REC POSITION
		bounceRec = checkRec
	} else { // POINTS EXITED - BOUNCE REC CHANGE DIRECTION

		if bounceRec.X < cntr.X { // FINDS A NEW VELX BASED ON POSITION RELATIVE TO CENTER
			velX = rF32(maxVel/8, maxVel)
		} else {
			velX = -rF32(maxVel/8, maxVel)
		}
		if bounceRec.Y < cntr.Y { // FINDS A NEW VELY BASED ON POSITION RELATIVE TO CENTER
			velY = rF32(maxVel/8, maxVel)
		} else {
			velY = -rF32(maxVel/8, maxVel)
		}

		timer = 15 // SETS A COLLISION TIMER
	}

}

// RETURNS A RANDOM FLOAT32 VALUE WITHIN A RANGE (BETWEEN MIN/MAX VALUES)
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
