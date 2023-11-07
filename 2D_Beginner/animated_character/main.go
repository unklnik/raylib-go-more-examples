package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

/* MORE RAYLIB GO EXAMPLES ARE AVAILABLE HERE:

https://github.com/unklnik/raylib-go-more-examples

*/

var (
	anim                rl.Texture2D  //TEXTURE TO LOAD IMAGE
	animFrames          = float32(4)  //5 FRAMES IN IMAGE SUBTRACT ONE AS COUNT STARTS FROM ZERO
	animFrameWidth      = float32(16) //16X16 RECTANGLE OF EACH FRAME
	currentAnimFrameNum float32       //THE NUMBER OF THE FRAME TO DRAW
	animDrawRec         rl.Rectangle  //THE PART OF THE IMAGE TEXTURE TO DRAW IN EACH FRAME
	animX               float32       //LEFT X VALUE OF FIRST FRAME
	animColor           = rl.White    //COLOUR IMAGE IS DRAWN
	animSpeed           = 6           //FRAMES BETWEEN ANIM CHANGES
	cntr                rl.Vector2    //CENTER OF SCREEN
	frames              int           //FRAME COUNT OF GAME
	fps                 = int32(60)   //FRAMES PER SECOND OF GAME
)

func main() {

	rl.InitWindow(0, 0, "animated character - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH := rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                            // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	cntr = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //CENTER OF SCREEN

	anim = rl.LoadTexture("anim.png")                                   //LOAD THE IMAGE FILE
	animDrawRec = rl.NewRectangle(0, 0, animFrameWidth, animFrameWidth) //DEFINE FRAME 1 (ZERO) RECTANGLE
	animX = animDrawRec.X                                               //SET THE HOLDER VARIABLE TO THE FIRST FRAME REC X VALUE

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	rl.SetTargetFPS(fps) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		frames++ //INCREASE FRAME COUNTER ONCE PER FRAME DRAWN

		//IF NUMBER OF FRAMES DIVIDES EXACLTY BY ANIM SPEED CHANGE CURRENT ANIM FRAME NUM
		if frames%animSpeed == 0 {
			currentAnimFrameNum++
			if currentAnimFrameNum > float32(animFrames) { //SET TO ZERO IF REACHES MAX FRAMES
				currentAnimFrameNum = 0
				animDrawRec.X = animX //RETURN DRAW REC X TO FIRST FRAME X
			}
		}

		/*MOVE THE DRAW RECTANGLE BY MULTIPLYING CURRENT FRAME NUM BY ANIM REC WIDTH. AS THE ANIM STARTS AT (0,0) WHEN CURRENT FRAME IS 0 ANIM REC X WILL ALSO BE ZERO. HOWEVER USING animX IS A BETTER METHOD WHEN YOU ARE WORKING WITH MANY IMAGES */
		animDrawRec.X = currentAnimFrameNum * animFrameWidth

		if rl.IsKeyPressed(rl.KeyRight) { //INCREASE SPEED
			if animSpeed > 4 {
				animSpeed -= 2
			}
		} else if rl.IsKeyPressed(rl.KeyLeft) { //DECREASE SPEED
			if animSpeed < 30 {
				animSpeed += 2
			}
		}

		if rl.IsKeyPressed(rl.KeyUp) { //CHANGE ZOOM
			if camera.Zoom == 1 {
				camera.Zoom = 1.5
			} else if camera.Zoom == 1.5 {
				camera.Zoom = 2
			} else if camera.Zoom == 2 {
				camera.Zoom = 1
			}

			//ADJUST CAMERA TO CENTER IMAGE AFTER CHANGE IN ZOOM
			camera.Target = cntr
			camera.Offset.X = float32(scrW / 2)
			camera.Offset.Y = float32(scrH / 2)
		}

		if rl.IsKeyPressed(rl.KeyDown) { //CHANGE COLOR
			if animColor == rl.White {
				animColor = rl.Magenta
			} else if animColor == rl.Magenta {
				animColor = rl.Green
			} else if animColor == rl.Green {
				animColor = rl.Orange
			} else if animColor == rl.Orange {
				animColor = rl.SkyBlue
			} else if animColor == rl.SkyBlue {
				animColor = rl.White
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		//CHANGE THIS FOR SMALLER/LARGER IMAGE
		size := float32(64)
		//THE RECTANGLE THAT DRAWS TO SCREEN
		destRec := rl.NewRectangle(cntr.X-size/2, cntr.Y-size/2, size, size)
		//DRAW TO SCREEN
		rl.DrawTexturePro(anim, animDrawRec, destRec, rl.Vector2Zero(), 0, animColor)

		/* NOTE THIS METHOD IS NOT TO BE USED FOR IMAGES THAT ROTATE AS THE ORIGIN IS SET AS THE TOP LEFT CORNER Vector2Zero() MEANS TOP LEFT OF IMAGE (0,0) AND THE IMAGE WILL ROTATE AROUND THIS POINT */

		rl.EndMode2D()

		rl.DrawText("UP key to change zoom", 10, 10, 10, rl.White)
		rl.DrawText("DOWN key to change color", 10, 24, 10, rl.White)
		rl.DrawText("RIGHT key to increase speed", 10, 38, 10, rl.White)
		rl.DrawText("LEFT key to decrease speed", 10, 52, 10, rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
