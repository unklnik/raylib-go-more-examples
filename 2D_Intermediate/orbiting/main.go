package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	cntr       rl.Vector2 // CENTER OF SCREEN
	scrW, scrH int        // SCREEN WIDTH & HEIGHT

	orbitV2one, orbitV2two, orbitV2three     rl.Vector2 //CENTER VECTOR2's OF 3 ORBITING CIRCLES
	angle1, angle2, angle3                   float32    //ANGLE FROM CENTER
	radius1, radius2, radius3                float32    //RADIUS OF 3 ORBITING CIRCLES
	angleChange1, angleChange2, angleChange3 float32    //CHANGE IN ROTATION ANGLE PER FRAME (SPEED)
	color1, color2, color3                   rl.Color   //COLOR OF 3 ORBITING CIRCLES
	changeDir1, changeDir2, changeDir3       bool       //SWITCH TO ROTATE LEFT/RIGHT
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

func main() {

	rl.InitWindow(0, 0, "orbiting - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	cntr = rl.NewVector2(float32(scrW/2), float32(scrH/2)) // CALCULATE CENTER

	makeV2() //MAKE INITIAL 3 ORBITING CIRCLES

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		upV2() //UPDATE MOVEMENT 3 ORBITING CIRCLES

		if rl.IsKeyPressed(rl.KeySpace) {
			changeDir1 = flipcoin() //TRUE FALSE FLIP COIN ON KEY PRESS
			changeDir2 = flipcoin()
			changeDir3 = flipcoin()
			angleChange1 = rF32(0.5, 4) //RANDOM CHANGE IN ROTATION ANGLE
			angleChange2 = rF32(0.5, 4)
			angleChange3 = rF32(0.5, 4)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		rl.DrawCircleV(orbitV2one, radius1, color1)   //DRAW ORIBITING CIRCLE OUTER
		rl.DrawCircleV(orbitV2two, radius2, color2)   //DRAW ORIBITING CIRCLE MIDDLE
		rl.DrawCircleV(orbitV2three, radius3, color3) //DRAW ORIBITING CIRCLE INNER

		//CREATE CENTER CROSS LINES
		colorLine := ranCol() //RANDOM COLOR SEE FUNCTION BELOW
		//VERTICAL LINE
		linev1 := cntr   //VECTOR2 EQUALS SCREEN CENTER
		linev1.Y -= 10   //MOVE VECTOR2 Y UP BY HALF LINE LENGTH
		linev2 := linev1 //NEW VECTOR2 EQUALS LINEV1
		linev2.Y += 20   //MOVE VECTOR2 Y DOWN BY LINE LENGTH
		//HORIZONTAL LINE
		linev3 := cntr   //VECTOR2 EQUALS SCREEN CENTER
		linev3.X -= 10   //MOVE VECTOR2 X LEFT BY HALF LINE LENGTH
		linev4 := linev3 //NEW VECTOR2 EQUALS LINEV3
		linev4.X += 20   //MOVE VECTOR2 X RIGHT BY LINE LENGTH

		rl.DrawLineV(linev1, linev2, colorLine) //DRAW VERTICAL LINE
		rl.DrawLineV(linev3, linev4, colorLine) //DRAW HORIZONTAL LINE

		rl.DrawText("press space to change direction", 10, 10, 20, rl.White) //TEXT

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func upV2() { //UPDATES MOVEMENT OF ORBITING CIRCLES

	//COMPLICATED MATH I COPIED & CHANGED FROM SOMEWHERE ELSE
	angle1 = angle1 * (math.Pi / 180)
	newx := float32(math.Cos(float64(angle1)))*(orbitV2one.X-cntr.X) - float32(math.Sin(float64(angle1)))*(orbitV2one.Y-cntr.Y) + cntr.X
	newy := float32(math.Sin(float64(angle1)))*(orbitV2one.X-cntr.X) + float32(math.Cos(float64(angle1)))*(orbitV2one.Y-cntr.Y) + cntr.Y
	orbitV2one = rl.NewVector2(newx, newy)

	//CHANGE ORBIT LEFT RIGHT
	if changeDir1 {
		angle1 -= angleChange1
	} else {
		angle1 += angleChange1
	}

	angle2 = angle2 * (math.Pi / 180)
	newx = float32(math.Cos(float64(angle2)))*(orbitV2two.X-cntr.X) - float32(math.Sin(float64(angle2)))*(orbitV2two.Y-cntr.Y) + cntr.X
	newy = float32(math.Sin(float64(angle2)))*(orbitV2two.X-cntr.X) + float32(math.Cos(float64(angle2)))*(orbitV2two.Y-cntr.Y) + cntr.Y
	orbitV2two = rl.NewVector2(newx, newy)

	if changeDir2 {
		angle2 -= angleChange2
	} else {
		angle2 += angleChange2
	}

	angle3 = angle3 * (math.Pi / 180)
	newx = float32(math.Cos(float64(angle3)))*(orbitV2three.X-cntr.X) - float32(math.Sin(float64(angle3)))*(orbitV2three.Y-cntr.Y) + cntr.X
	newy = float32(math.Sin(float64(angle3)))*(orbitV2three.X-cntr.X) + float32(math.Cos(float64(angle3)))*(orbitV2three.Y-cntr.Y) + cntr.Y
	orbitV2three = rl.NewVector2(newx, newy)

	if changeDir3 {
		angle3 -= angleChange3
	} else {
		angle3 += angleChange3
	}

}
func makeV2() {
	width := float32(scrH / 2)                       //DISTANCE FROM CENTER = HALF SCREEN HEIGHT
	width -= width / 4                               //DISTANCE FROM CENTER SUBTRACT QUARTER OF ORIGINAL LENGTH
	orbitV2one = rl.NewVector2(cntr.X+width, cntr.Y) //CREATE CENTER VECTOR2 OF ORBITING CIRCLE
	angle1 = rF32(0, 360)                            //RANDOM START ANGLE SEE rF32 FUNCTION BELOW
	radius1 = rF32(20, 40)                           //RANDOM START RADIUS
	color1 = ranCol()                                //RANDOM START COLOR SEE FUNCTION BELOW
	angleChange1 = rF32(0.5, 4)                      //RANDOM START ROTATION ANGLE CHANGE (SPEED)

	width = float32(scrH / 2)
	width -= width / 2
	orbitV2two = rl.NewVector2(cntr.X+width, cntr.Y)
	angle2 = rF32(0, 360)
	radius2 = rF32(20, 40)
	color2 = ranCol()
	angleChange2 = rF32(0.5, 4)

	width = float32(scrH / 2)
	width -= (width / 4) * 3
	orbitV2three = rl.NewVector2(cntr.X+width, cntr.Y)
	angle3 = rF32(0, 360)
	radius3 = rF32(20, 40)
	color3 = ranCol()
	angleChange3 = rF32(0.5, 4)

}

// RETURNS A RANDOM r.Color VALUE
func ranCol() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}

// RETURNS RANDOM INTEGER VALUE
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RETURNS RANDOM FLOAT32 VALUE
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}

// RETURNS RANDOM TRUE/FALSE
func flipcoin() bool {
	onoff := false
	choose := rInt(0, 100001)
	if choose > 50000 {
		onoff = true
	}
	return onoff
}
