package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	scrW, scrH       int         // SCREEN WIDTH & HEIGHT
	snow             []snowflake //SLICE OF SNOWFLAKE STRUCTS
	colorsOn, blurOn bool        //COLORS & BLUR ON/OFF

	imgs rl.Texture2D

	snowflakeIMG1 = rl.NewRectangle(0, 0, 17, 17)  //DEFINE SNOWFLAKE IMAGE 1 RECTANGLE IN imgs.png
	snowflakeIMG2 = rl.NewRectangle(18, 1, 15, 15) //DEFINE SNOWFLAKE IMAGE 2 RECTANGLE IN imgs.png
	snowflakeIMG3 = rl.NewRectangle(34, 1, 15, 15) //DEFINE SNOWFLAKE IMAGE 3 RECTANGLE IN imgs.png
)

// STRUCT HOLDS ALL THE INFORMATION FOR EACH SNOWFLAKE
type snowflake struct {
	img, rec, drawrec        rl.Rectangle //RECTANGLES TO DISPLAY IMAGE
	ro, speed, roSpeed, fade float32      //ROTATION/FALL SPEED/ROTATION SPEED/FADE
	origin                   rl.Vector2   //ORIGIN OF DRAW RECTANGLE FOR ROTATION
	color                    rl.Color     //SNOWFLAKE ALT COLOR
	leftright                bool         //ROTATE LEFT OR RIGTH
}

func main() {

	rl.InitWindow(0, 0, "snow - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	imgs = rl.LoadTexture("imgs.png") //LOAD SNOWFLAKE IMAGES

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	makeSnow() //CREATE SLICE OF SNOWFLAKES SEE END OF CODE

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeySpace) { //COLORS ON/OFF
			colorsOn = !colorsOn
		}
		if rl.IsKeyPressed(rl.KeyTab) { //BLUR ON/OFF
			blurOn = !blurOn
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		upSnow() //DRAW & UPDATE SNOW MOVEMENT & ROTATION SEE END OF CODE

		rl.DrawText("SPACE KEY TO CHANGE COLOR / WHITE", 10, 10, 20, rl.White)
		rl.DrawText("TAB KEY TO TURN BLUR ON/OFF", 10, 40, 20, rl.White)

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.UnloadTexture(imgs) //UNLOAD FROM MEMORY

	rl.CloseWindow()
}

func makeSnow() {

	num := 100 //NUMBER OF SNOWFLAKES

	for num > 0 { //CREATES A SNOWFLAKE WITH RANDOM VALUES & ADDS (APPENDS) TO SNOW SLICE

		newSnow := snowflake{} //CREATE INSTANCE
		size := rF32(32, 128)  //DEFINE RANDOM SIZE FOR USE IN RECTANGLE

		//CREATE RECTANGLE AT TOP LEFT OF SCREEN / SUBTRACT SIZE MOVES ABOVE SCREEN TOP
		newSnow.rec = rl.NewRectangle(0, -size, size, size)
		//MOVE REC.X RANDOM DISTANCE FROM LEFT - BETWEEN 0 & (SCREEN WIDTH - SIZE)
		newSnow.rec.X += rF32(0, float32(scrW)-size)
		//MOVE REC.Y RANDOM DISTANCE AWAY FROM SCREEN TOP - STAGGER FALLING EFFECT
		newSnow.rec.Y -= rF32(0, size*20)
		//CREATE A DRAW RECTANGLE FOR ROTATION
		newSnow.drawrec = newSnow.rec
		//ADJUST DRAW RECTANGLE OFFSET
		newSnow.drawrec.X += newSnow.drawrec.Width / 2
		newSnow.drawrec.Y += newSnow.drawrec.Height / 2
		//CREATE ORIGIN (CENTER) FOR ROTATION
		newSnow.origin = rl.NewVector2(newSnow.drawrec.Width/2, newSnow.drawrec.Height/2)

		newSnow.color = ranCol()       //RANDOM SECOND COLOR SEE FUNCTION BELOW
		newSnow.fade = rF32(0.3, 0.6)  //RANDOM TRANSPARENCY
		newSnow.speed = rF32(3, 8)     //RANDOM FALL SPEED
		newSnow.ro = rF32(0, 360)      //RANDOM INITIAL ROTATION ANGLE
		newSnow.roSpeed = rF32(1, 4)   //RANDOM ROTATION SPEED
		newSnow.leftright = flipcoin() //CHOOSE LEFT RIGHT ROTATION

		choose := rInt(1, 4) //CHOOSE SNOWFLAKE IMAGE
		switch choose {
		case 1:
			newSnow.img = snowflakeIMG1
		case 2:
			newSnow.img = snowflakeIMG2
		case 3:
			newSnow.img = snowflakeIMG3
		}

		snow = append(snow, newSnow) // ADD TO SLICE

		num--
	}

}

// DRAW & UPDATE SNOW
func upSnow() {

	for i := 0; i < len(snow); i++ { //RANGE OVER SNOWFLAKE SLICE UPDATING & DRAWING EACH SNOWFLAKE

		if colorsOn { //DRAW COLOR SNOWFLAKES
			rl.DrawTexturePro(imgs, snow[i].img, snow[i].drawrec, snow[i].origin, snow[i].ro, rl.Fade(snow[i].color, snow[i].fade))

			if blurOn { //ADDITIONAL IMAGE FOR FAKE BLUR EFFECT
				blurRec := snow[i].drawrec //DUPLICATE DRAW RECTANGLE
				blurRec.X += rF32(-5, 5)   //MOVE DUPLICATE REC X RANDOM AMOUNT
				blurRec.Y += rF32(-5, 5)   //MOVE DUPLICATE REC Y RANDOM AMOUNT

				//DRAW ON TOP OF ORIGINAL IMAGE WITH LOWER FADE VALUES & SAME COLOR
				rl.DrawTexturePro(imgs, snow[i].img, blurRec, snow[i].origin, snow[i].ro, rl.Fade(snow[i].color, rF32(0.1, 0.3)))
			}
		} else { //DRAW WHITE SNOWFLAKES
			rl.DrawTexturePro(imgs, snow[i].img, snow[i].drawrec, snow[i].origin, snow[i].ro, rl.Fade(rl.White, snow[i].fade))

			if blurOn {
				blurRec := snow[i].drawrec
				blurRec.X += rF32(-5, 5)
				blurRec.Y += rF32(-5, 5)
				rl.DrawTexturePro(imgs, snow[i].img, blurRec, snow[i].origin, snow[i].ro, rl.Fade(rl.White, rF32(0.1, 0.3)))
			}
		}

		if snow[i].leftright {
			snow[i].ro += snow[i].roSpeed //RIGHT ROTATATION
		} else {
			snow[i].ro -= snow[i].roSpeed //LEFT ROTATATION
		}

		snow[i].drawrec.Y += snow[i].speed //MOVE DOWN

		//IF SNOWFLAKE EXITS SCREEN MOVE BACK TO A RANDOM HEIGHT ABOVE TOP OF SCREEN
		if snow[i].drawrec.Y > float32(scrH) {
			snow[i].drawrec.Y = rF32(-snow[i].drawrec.Height*20, -snow[i].drawrec.Height)
		}

	}

}

// SIMULATES FLIPCOIN HEAD/TAILS / ON/OFF / LEFT/RIGHT
func flipcoin() bool {
	onoff := false
	choose := rInt(0, 100001)
	if choose > 50000 {
		onoff = true
	}
	return onoff
}

// RETURNS A RANDOM COLOR
func ranCol() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}

// RETURNS A RANDOM INTEGER BETWEEN MIN/MAX VALUES FOR USE IN ranCol()
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RETURNS A RANDOM FLOAT32 BETWEEN MIN/MAX VALUES
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
