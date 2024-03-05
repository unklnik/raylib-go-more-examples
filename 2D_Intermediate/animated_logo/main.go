package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	raylibAnim     []rl.Texture2D                    //SLICE OF RAYLIB LOGO IMAGE FRAMES
	gologo         rl.Texture2D                      //TEXTURE OF GO LOGO IMAGE
	goLogoRec      = rl.NewRectangle(0, 0, 400, 160) //RECTANGLE SIZE OF GO LOGO IMAGE
	raylibFrameNum int                               //FRAME NUMBER FOR ANIMATION
	raylibRec      = rl.NewRectangle(0, 0, 288, 288) //RECTANGLE SIZE OF RAYLIB LOGO IMAGE
	introT1        = fps * 2                         //TIMER BETWEEN ANIMATIONS
	introT2        = fps * 2                         //TIMER BETWEEN ANIMATIONS
	goLogoDrawRec  rl.Rectangle                      //DRAW RECTANGLE FOR GO LOGO
	fps            = int32(60)                       //FRAMES PER SECOND
	cnt            rl.Vector2                        //SCREEN CENTER
	scrW, scrH     int                               // SCREEN WIDTH & HEIGHT
	frames         int                               //FRAME COUNTER
)

func main() {

	rl.InitWindow(0, 0, "animated logo - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE
	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2))

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	gologo = rl.LoadTexture("img/gologo.png") //LOAD GO LOGO IMG
	for i := 0; i < 57; i++ {                 //LOAD RAYLIB LOGO FRAMES INTO SLICE
		txt := fmt.Sprint(i) + ".png"
		if i < 10 {
			txt = "0" + txt
		}
		txt = "img/raylib_logo/" + txt
		raylibAnim = append(raylibAnim, rl.LoadTexture(txt))
	}

	startAnim()

	rl.SetTargetFPS(fps) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {
		frames++                          //FRAME COUNTER
		if rl.IsKeyPressed(rl.KeySpace) { //RESTART KEY
			startAnim()
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode2D(camera)

		if introT1 > 0 { //IF INTRO TIMER 1 IS LARGER THAN ZERO THEN DRAW GO LOGO
			rl.DrawTexturePro(gologo, goLogoRec, goLogoDrawRec, rl.Vector2Zero(), 0, rl.White)
			if goLogoDrawRec.X < cnt.X-goLogoDrawRec.Width/2 { //MOVE LOGO TO CENTER
				goLogoDrawRec.X += 20
			} else { //WHEN AT CENTER THEN DECREASE TIMER
				introT1--
			}
		} else if introT2 > 0 { //IF INTRO TIMER 2 IS LARGER THAN ZERO & INTRO TIMER 1 IS ZERO THEN DRAW
			siz := float32(320) //SIZE OF RECTANGLE
			rec := rl.NewRectangle(cnt.X-siz/2, cnt.Y-siz/2, siz, siz)
			rl.DrawTexturePro(raylibAnim[raylibFrameNum], raylibRec, rec, rl.Vector2Zero(), 0, rl.White)
			if raylibFrameNum < len(raylibAnim)-1 {
				if frames%3 == 0 { //ADVANCE DRAW IMAGE EVERY 3 FRAMES
					raylibFrameNum++
				}
			}
		}

		rl.DrawText("SPACE KEY TO RESTART", 10, 10, 20, rl.White)

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.UnloadTexture(gologo) //UNLOAD FROM MEMORY
	for i := 0; i < len(raylibAnim); i++ {
		rl.UnloadTexture(raylibAnim[i])
	}
	rl.CloseWindow()
}

func startAnim() { //RESET THE ANIMATION
	goLogoDrawRec = rl.NewRectangle(0-goLogoRec.Width, cnt.Y-goLogoRec.Height/2, goLogoRec.Width, goLogoRec.Height) //CREATE GO LOGO RECTANGLE OUTSIDE LEFT SCREEN BORDER
	introT1 = fps * 2
	introT2 = fps * 2
	raylibFrameNum = 0
}
