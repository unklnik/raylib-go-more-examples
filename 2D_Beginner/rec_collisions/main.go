package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	velX, velY             float32        //X & Y MOVEMENT SPEED
	maxVel                 = float32(10)  //MAX SPEED FOR USE IN RANDOM GENERATOR
	borderRec              rl.Rectangle   //BOUNCING & BORDER RECTANGLES
	cntr                   rl.Vector2     //CENTER OF SCREEN
	bounceRecs, borderRecs []rl.Rectangle //SLICES OF BOUNCE & BORDER RECTANGLES
	bounceV2               []rl.Vector2   //X & Y MOVE SPEED OF BOUNCE RECTNAGLES
	spd                    = float32(16)  //MAX SPEED
)

func main() {

	rl.InitWindow(0, 0, "rectangle collisions - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH := rl.GetScreenWidth(), rl.GetScreenHeight() //GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                            //SET WINDOW SIZE
	rl.SetWindowState(rl.FlagBorderlessWindowedMode)
	//rl.ToggleFullscreen() //UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	cntr = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //CALCULATE CENTER

	borderRec = rl.NewRectangle(cntr.X-float32(scrW/4), cntr.Y-float32(scrH/4), float32(scrW/2), float32(scrH/2)) //DEFINE BORDER RECTANGLE

	//CREATE 4 RECTANGLES AT EACH SIDE OF BORDER RECTANGLE TO CHECK FOR COLLISIONS AGAINST
	borderRecs = append(borderRecs, rl.NewRectangle(borderRec.X, borderRec.Y-10, borderRec.Width, 10))
	borderRecs = append(borderRecs, rl.NewRectangle(borderRec.X, borderRec.Y+borderRec.Height, borderRec.Width, 10))
	borderRecs = append(borderRecs, rl.NewRectangle(borderRec.X-10, borderRec.Y, 10, borderRec.Height))
	borderRecs = append(borderRecs, rl.NewRectangle(borderRec.X+borderRec.Width, borderRec.Y, 10, borderRec.Height))

	//CREATE 4 BOUNCE RECTANGLES AT EACH SIDE OF THE BORDER RECTANGLE RF32 FUNCTION SEE BELOW
	siz := float32(64)
	bounceRecs = append(bounceRecs, rl.NewRectangle(borderRec.X+siz+rF32(0, borderRec.Width-siz*2), borderRec.Y, siz, siz))
	bounceRecs = append(bounceRecs, rl.NewRectangle(borderRec.X+siz+rF32(0, borderRec.Width-siz*2), borderRec.Y+borderRec.Height-siz, siz, siz))
	bounceRecs = append(bounceRecs, rl.NewRectangle(borderRec.X, borderRec.Y+siz+rF32(0, borderRec.Height-siz*2), siz, siz))
	bounceRecs = append(bounceRecs, rl.NewRectangle(borderRec.X+borderRec.Width-siz, borderRec.Y+siz+rF32(0, borderRec.Height-siz*2), siz, siz))

	//CREATE 4 VECTOR 2 TO STORE X & Y MOVEMEMENT SPEED
	bounceV2 = append(bounceV2, rl.NewVector2(rF32(-spd, spd), rF32(-spd, spd)))
	bounceV2 = append(bounceV2, rl.NewVector2(rF32(-spd, spd), rF32(-spd, spd)))
	bounceV2 = append(bounceV2, rl.NewVector2(rF32(-spd, spd), rF32(-spd, spd)))
	bounceV2 = append(bounceV2, rl.NewVector2(rF32(-spd, spd), rF32(-spd, spd)))

	//SEE RF32 FUNCTION END OF CODE FINDS A RANDOM FLOAT32 BETWEEN TWO VALUES
	velX = rF32(-maxVel, maxVel) //FINDS A RANDOM X AXIS MOVEMENT SPEED
	velY = rF32(-maxVel, maxVel) //FINDS A RANDOM Y AXIS MOVEMENT SPEED

	camera := rl.Camera2D{} //DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	rl.SetTargetFPS(60) //NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		checkCollisionsMove() //SEE FUNCTION BELOW - CHECKS FOR BORDER COLLISIONS & MOVES

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		rl.DrawRectangleLinesEx(borderRec, 4, rl.Green) //DRAWS BORDER REC

		//DRAWS THE BOUNCE RECTANGLES IN THE SLICE
		for i := 0; i < len(bounceRecs); i++ {
			rl.DrawRectangleLinesEx(bounceRecs[i], 4, rl.Green)
		}

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// CHECKS FOR BOUNCE REC COLLISIONS & MOVES BOUNCE RECS
func checkCollisionsMove() {

	for i := 0; i < len(bounceRecs); i++ {

		//CREATING DUPLICATE RECTANGLE FOR NEXT MOVE PREVENTS PROBLEMS WITH INTERSECTING RECTANGLES IF USING THE RECTANGLE ITSELF AS THEN IT WILL ONLY RETURN A COLLISION ONCE IT HAS ALREADY HAPPENED

		canmove := true
		checkRec := bounceRecs[i]   //DUPLICATE THE RECTANGLE
		checkRec.X += bounceV2[i].X //MOVE THE DUPLICATE BY X SPEED
		checkRec.Y += bounceV2[i].Y //MOVE THE DUPLICATE BY X SPEED

		//CHECK FOR COLLISIONS WITH DUPLICATE
		for j := 0; j < len(borderRecs); j++ {
			if rl.CheckCollisionRecs(checkRec, borderRecs[j]) {
				canmove = false
				break
			}
		}

		//IF NO COLLISIONS WITH BORDER THEN CHECK FOR COLLISIONS WITH OTHER BOUNCE RECTANGLES
		if canmove {
			for j := 0; j < len(bounceRecs); j++ {
				if i != j {
					if rl.CheckCollisionRecs(checkRec, bounceRecs[j]) {
						canmove = false
						break
					}
				}
			}
		}

		if canmove {
			//IF CAN MOVE THEN MOVE THE RECTANGLE ITSELF
			bounceRecs[i].X += bounceV2[i].X
			bounceRecs[i].Y += bounceV2[i].Y
		} else {
			//IF CANNOT MOVE THEN MOVE IN THE OPPOSITE DIRECTION AT A NEW RANDOM SPEED
			if bounceV2[i].X > 0 {
				bounceV2[i].X = rF32(-spd, -spd/4)
			} else {
				bounceV2[i].X = rF32(spd/4, spd)
			}
			if bounceV2[i].Y > 0 {
				bounceV2[i].Y = rF32(-spd, -spd/4)
			} else {
				bounceV2[i].Y = rF32(spd/4, spd)
			}
		}

	}

}

// RETURNS A RANDOM FLOAT32 VALUE WITHIN A RANGE (BETWEEN MIN/MAX VALUES)
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
