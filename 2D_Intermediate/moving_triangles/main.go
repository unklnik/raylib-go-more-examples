package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	vel        = float32(8)   // SPEED
	cntr       rl.Vector2     // CENTER OF SCREEN
	triangles  []xtri         // SLICE OF TRIANGLE INFORMATION TO CREATE TRIANGLE POLYGONS (NOT POINTS)
	scrW, scrH int            // SCREEN WIDTH & HEIGHT
	size       = float32(32)  // RADIUS OF POLYGON CIRCLE
	fade       = float32(1.0) // COLOR FADE OF TRIANGLE POLYGON FILL
	direc      = 3            // DIRECTION OF MOVEMENT

	/* DIREC STARTS AT 1 (UP) MOVES CLOCKWISE
		    1 UP
	4 LEFT			2 RIGHT
			3 DOWN
	*/
)

type xtri struct {
	rad, ro float32    // RADIUS & ROTATION
	col     rl.Color   // COLOR
	cnt     rl.Vector2 // CENTER
}

func main() {

	rl.InitWindow(0, 0, "moving triangles - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	cntr = rl.NewVector2(float32(scrW/2), float32(scrH/2)) // CALCULATE CENTER

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	makeTriangles() // MAKE THE SLICE OF TRIANGLES

	for !rl.WindowShouldClose() {

		upTriangles() // AFTER rl.WindowShouldClose UPDATES MOVEMENT & INPUT SEE FUNCTION BELOW

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		rl.DrawText("W A S D keys change direction", 16, 16, 20, rl.White)

		for a := 0; a < len(triangles); a++ { // RANGE OVER SLICE OF TRIANGLES 0 TILL END
			// DRAW POLYGON TRIANGLE
			rl.DrawPoly(triangles[a].cnt, 3, triangles[a].rad, triangles[a].ro, rl.Fade(triangles[a].col, fade))
			// DRAW POLYGON TRIANGLE OUTLINE
			rl.DrawPolyLinesEx(triangles[a].cnt, 3, triangles[a].rad, triangles[a].ro, 10, rl.Orange)
		}

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func upTriangles() { // UPDATES MOVEMENT & INPUT

	// FADE IN AND OUT
	if fade > 0.5 {
		fade -= 0.01
	} else {
		fade = 1
	}

	if rl.IsKeyPressed(rl.KeyW) { // CHANGE DIREC TO UP ON W KEY PRESS
		if direc != 1 {
			direc = 1
			for a := 0; a < len(triangles); a++ {
				triangles[a].ro = 180 // ROTATE

				// CHANGE POSITION OF TRIANGES 1,2,3 OF SLICE (NOT 0 - FIRST) TO FOLLOW 0 (FIRST)
				if a > 0 {
					// Y = PREVIOUS TRIANGE Y - OFFSET DISTANCE
					triangles[a].cnt.Y = triangles[a-1].cnt.Y + (size * 2)
					// X = PREVIOUS TRIANGE X
					triangles[a].cnt.X = triangles[a-1].cnt.X
				}
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyD) { // CHANGE DIREC TO RIGHT ON D KEY PRESS
		direc = 2
		for a := 0; a < len(triangles); a++ {
			triangles[a].ro = 90 // ROTATE

			// CHANGE POSITION OF TRIANGES 1,2,3 OF SLICE (NOT 0 - FIRST) TO FOLLOW 0 (FIRST)
			if a > 0 {
				// Y = PREVIOUS TRIANGLE Y
				triangles[a].cnt.Y = triangles[a-1].cnt.Y
				// X = PREVIOUS TRIANGE X - OFFSET DISTANCE
				triangles[a].cnt.X = triangles[a-1].cnt.X - (size * 2)
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyS) { // CHANGE DIREC TO DOWN ON S KEY PRESS
		direc = 3
		for a := 0; a < len(triangles); a++ {
			triangles[a].ro = 0 // ROTATE

			// CHANGE POSITION OF TRIANGES 1,2,3 OF SLICE (NOT 0 - FIRST) TO FOLLOW 0 (FIRST)
			if a > 0 {
				// Y = PREVIOUS TRIANGE Y - OFFSET DISTANCE
				triangles[a].cnt.Y = triangles[a-1].cnt.Y - (size * 2)
				// X = PREVIOUS TRIANGE X
				triangles[a].cnt.X = triangles[a-1].cnt.X
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyA) { // CHANGE DIREC TO LEFT ON A KEY PRESS
		direc = 4
		for a := 0; a < len(triangles); a++ {
			triangles[a].ro = 270 // ROTATE

			// CHANGE POSITION OF TRIANGES 1,2,3 OF SLICE (NOT 0 - FIRST) TO FOLLOW 0 (FIRST)
			if a > 0 {
				// Y = PREVIOUS TRIANGLE Y
				triangles[a].cnt.Y = triangles[a-1].cnt.Y
				// X = PREVIOUS TRIANGE X + OFFSET DISTANCE
				triangles[a].cnt.X = triangles[a-1].cnt.X + (size * 2)
			}
		}
	}

	// MOVES CENTER OF POLYGON & CHECKS FOR SCREEN EXIT
	for a := 0; a < len(triangles); a++ {
		switch direc {
		case 1: // UP
			triangles[a].cnt.Y -= vel // MOVES CENTER Y UP
			// IF CENTER.Y IS SMALLER THAN -(SIZE/2) MOVE TO SCREEN HEIGHT + SIZE/2
			if triangles[a].cnt.Y < -size/2 {
				triangles[a].cnt.Y = float32(scrH) + size/2
			}
		case 2: // RIGHT
			triangles[a].cnt.X += vel // MOVES CENTER X RIGHT
			// IF CENTER.X IS LARGER THAN SCREEN WIDTH + SIZE/2 MOVE TO -(SIZE/2)
			if triangles[a].cnt.X > float32(scrW)+size/2 {
				triangles[a].cnt.X = -size / 2
			}
		case 3: // DOWN
			triangles[a].cnt.Y += vel // MOVES CENTER Y DOWN
			// IF CENTER.Y IS LARGER THAN SCREEN HEIGHT + SIZE/2 MOVE TO -(SIZE/2)
			if triangles[a].cnt.Y > float32(scrH)+size/2 {
				triangles[a].cnt.Y = -size / 2
			}
		case 4: // LEFT
			triangles[a].cnt.X -= vel // MOVES CENTER X LEFT
			// IF CENTER.X IS SMALLER THAN -(SIZE/2) MOVE TO SCREEN WIDTH + SIZE/2
			if triangles[a].cnt.X < -size/2 {
				triangles[a].cnt.X = float32(scrW) + size/2
			}
		}
	}
}

func makeTriangles() {

	ztri := xtri{}
	ztri.rad = size     // RADIUS
	ztri.ro = 0         // DEFAULT ROTATION
	ztri.col = rl.Green // COLOR

	ztri.cnt = cntr                     // SET CENTER TO SCREEN CENTER
	triangles = append(triangles, ztri) // ADD TO SLICE

	ztri.cnt.Y += size * 2              // MOVE CENTER.Y DOWN BY SIZE X 2
	triangles = append(triangles, ztri) // ADD TO MOVED TRI VERSION TO SLICE

	ztri.cnt.Y += size * 2              // MOVE CENTER.Y DOWN BY SIZE X 2
	triangles = append(triangles, ztri) // ADD TO MOVED TRI VERSION TO SLICE

	ztri.cnt.Y += size * 2              // MOVE CENTER.Y DOWN BY SIZE X 2
	triangles = append(triangles, ztri) // ADD TO MOVED TRI VERSION TO SLICE
}
