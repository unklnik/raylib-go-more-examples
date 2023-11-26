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
	tileImgs                    rl.Texture2D   //IMG TEXTURE WITH TILES
	size                        = float32(32)  //TILE SIZE CHANGE FOR LARGER/SMALLER IMG DISPLAY SIZE
	gridW                       = 20           //WIDTH HEIGHT GRID
	gridA                       int            //GRID AREA
	cnt                         rl.Vector2     //SCREEN CENTER VECTOR2
	scrW, scrH                  int            //SCREEN WIDTH HEIGHT
	tiles                       []blok         //SLICE OF TILE BLOCK STRUCTS
	tileRecs                    []rl.Rectangle //SLICE OF TILE RECTANGLE POSITIONS IN TEXTURE
	gridDots, gridLines         bool           //TURN ON OFF GRID LINES & DOTS
	mouseposScreen, mouseposCam rl.Vector2     //MOUSE POSITION VECTOR 2
	blinkFade                   = float32(1)   //FADE OF SELECT BOX FOR BLINKING
	blinkOnOff                  bool           //FADE IN/OUT BOOLEAN
	tileNum, colorNum           int            //CURRENT SELECTED TILE IMAGE & COLOR

	tileColors = []rl.Color{rl.Green, rl.White, rl.SkyBlue, rl.Orange, rl.Red, rl.LightGray, rl.Lime, rl.DarkBlue, rl.Blue, rl.Brown, rl.Beige, rl.DarkGreen, rl.Purple, rl.Pink, rl.Magenta, rl.DarkBrown, rl.DarkPurple, rl.Yellow, rl.Gold, rl.Violet} //COLORS FOR TILES
)

type blok struct { //STRUCT THAT HOLDS EACH TILE INFORMATION
	img, rec rl.Rectangle //IMAGE & DRAW RECTANGLES
	color    rl.Color
}

func main() {

	rl.InitWindow(0, 0, "tile editor - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES
	//rl.SetWindowState(rl.FlagBorderlessWindowedMode) // UNCOMMENT IF YOU HAVE DISPLAY ISSUES

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //FIND SCREEN CENTER

	camera := rl.Camera2D{} // DEFINES THE CAMERA
	camera.Zoom = 1.0       //SETS CAMERA ZOOM

	tileImgs = rl.LoadTexture("tiles.png") //LOAD IMAGES FROM FILE

	gridDots = true //TURN ON GRID DOTS BY DEFAULT

	makegridtiles() //MAKE GRID & TILE RECTANGLES FUNCTION BELOW

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		mouseposScreen = rl.GetMousePosition()                      //GET MOUSE POSITION SCREEN
		mouseposCam = rl.GetScreenToWorld2D(mouseposScreen, camera) //GET MOUSE CURSOR POSITION IN CAMERA SPACE
		if blinkOnOff {                                             //FADE IN/OUT ON/OFF
			if blinkFade < 1 {
				blinkFade += 0.05
			} else {
				blinkOnOff = false
			}
		} else {
			if blinkFade > 0.1 {
				blinkFade -= 0.05
			} else {
				blinkOnOff = true
			}
		}

		if rl.IsKeyPressed(rl.KeySpace) { //SET ALL TILE COLORS TO BLANK
			for i := 0; i < len(tiles); i++ {
				tiles[i].color = rl.Blank
			}
		}
		if rl.IsKeyPressed(rl.KeyUp) { //CHANGE CAMERA ZOOM KEYS
			if camera.Zoom == 0.5 {
				camera.Zoom = 1
			} else if camera.Zoom == 1 {
				camera.Zoom = 1.5
			} else if camera.Zoom == 1.5 {
				camera.Zoom = 2
			} else if camera.Zoom == 2 {
				camera.Zoom = 0.5
			}

			camera.Target = cnt                 //SET THE CAMERA TARGET TO CENTER
			camera.Offset.X = float32(scrW / 2) //ADJUST FOR ZOOM
			camera.Offset.Y = float32(scrH / 2) //ADJUST FOR ZOOM
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		//DRAW BLOKS
		for i := 0; i < len(tiles); i++ {
			if tiles[i].color != rl.Blank { //IF TILE NOT CLEAR THEN DRAW IMAGE
				rl.DrawTexturePro(tileImgs, tiles[i].img, tiles[i].rec, rl.Vector2Zero(), 0, tiles[i].color)
			}
			if rl.CheckCollisionPointRec(mouseposCam, tiles[i].rec) { //IF MOUSE HOVERS DRAW BLINKING REC
				rl.DrawRectangleRec(tiles[i].rec, rl.Fade(rl.Green, blinkFade))
				if rl.IsMouseButtonDown(rl.MouseButtonLeft) { //IF MOUSE LEFT ASSIGN CURRENT COLOR IMAGE
					tiles[i].color = tileColors[colorNum]
					tiles[i].img = tileRecs[tileNum]
				} else if rl.IsMouseButtonDown(rl.MouseRightButton) { //RIGHT CLICK CLEAR
					tiles[i].color = rl.Blank
				}
			}
		}
		//DRAW GRID
		if gridLines { //DRAW RECTANGLE OUTLINE IF GRIDTILES ON
			for i := 0; i < len(tiles); i++ {
				rl.DrawRectangleLinesEx(tiles[i].rec, 1, rl.Fade(rl.Green, 0.3))
			}
		}
		if gridDots { //DRAW RECTANGLE CORNER VECTOR2 IF GRIDDOTS ON
			for i := 0; i < len(tiles); i++ {
				v1 := rl.NewVector2(tiles[i].rec.X, tiles[i].rec.Y)
				v2 := rl.NewVector2(tiles[i].rec.X+size, tiles[i].rec.Y)
				v3 := rl.NewVector2(tiles[i].rec.X+size, tiles[i].rec.Y+size)
				v4 := rl.NewVector2(tiles[i].rec.X, tiles[i].rec.Y+size)
				rl.DrawCircleV(v1, 2, rl.Green)
				rl.DrawCircleV(v2, 2, rl.Green)
				rl.DrawCircleV(v3, 2, rl.Green)
				rl.DrawCircleV(v4, 2, rl.Green)
			}
		}

		rl.EndMode2D()

		//DRAW TEXT
		rl.DrawText("camera zoom "+fmt.Sprintf("%.1f", camera.Zoom)+" press UP ARROW key to change", 10, 10, 20, rl.White)
		rl.DrawText("LEFT ARROW key grid dots on/off | RIGHT ARROW key grid lines on/off", 10, 40, 20, rl.White)
		rl.DrawText("mouse LEFT draw RIGHT delete / press SPACE key to clear tiles", 10, 70, 20, rl.White)

		drawmenu()

		rl.EndDrawing()
	}

	rl.UnloadTexture(tileImgs) //UNLOAD FROM MEMORY

	rl.CloseWindow()
}
func drawmenu() { //RIGHT SIDE MENU

	txt := "lines"
	txtlen := rl.MeasureText(txt, 20) //MEASURE TO CENTER TEXT
	x := float32(scrW - 150)          //OFFSET FROM SCREEN RIGHT EDGE
	y := float32(50)

	rec := rl.NewRectangle(x, y, 100, 30) //LINES ON/OFF MENU BUTTON
	if gridLines {
		rl.DrawRectangleRec(rec, rl.Green)
	} else {
		rl.DrawRectangleLinesEx(rec, 4, rl.Red)
	}
	if rl.CheckCollisionPointRec(mouseposScreen, rec) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) || rl.IsKeyPressed(rl.KeyRight) { //IF LEFT CLICK GRIDLINES ON/OFF
		gridLines = !gridLines
	}
	txtx := rec.ToInt32().X + rec.ToInt32().Width/2 - txtlen/2
	txty := rec.ToInt32().Y + 5
	rl.DrawText(txt, txtx-2, txty+2, 20, rl.Black)
	rl.DrawText(txt, txtx, txty, 20, rl.White)

	y = rec.Y + rec.Height + 10
	txt = "dots"

	rec = rl.NewRectangle(x, y, 100, 30)
	if gridDots {
		rl.DrawRectangleRec(rec, rl.Green)
	} else {
		rl.DrawRectangleLinesEx(rec, 4, rl.Red)
	}
	if rl.CheckCollisionPointRec(mouseposScreen, rec) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) || rl.IsKeyPressed(rl.KeyLeft) { //IF LEFT CLICK GRIDDOTS ON/OFF
		gridDots = !gridDots
	}
	txtx = rec.ToInt32().X + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 5
	rl.DrawText(txt, txtx-2, txty+2, 20, rl.Black)
	rl.DrawText(txt, txtx, txty, 20, rl.White)

	y = rec.Y + rec.Height + size
	yorig := y //SAVE ORIGINAL Y VALUE
	x -= size

	for i := 0; i < len(tileRecs); i++ { //DRAW TEXTURE IMAGES
		rec = rl.NewRectangle(x, y, size, size)
		rl.DrawTexturePro(tileImgs, tileRecs[i], rec, rl.Vector2Zero(), 0, rl.White)
		if rl.CheckCollisionPointRec(mouseposScreen, rec) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			tileNum = i
		}
		if tileNum == i {
			rl.DrawRectangleLinesEx(rec, 4, rl.Fade(rl.Red, blinkFade))
		}
		y += size + 8 //+8 PIXEL SPACING
		if i == (len(tileRecs)/2)-1 {
			x += size + 8
			y = yorig //IF REACHES HALFWAY THEN RETURN TO TOP
		}

	}

	y = yorig
	x += size + 8

	for i := 0; i < len(tileColors); i++ { //DRAW COLORS
		rec = rl.NewRectangle(x, y, size, size)
		rl.DrawRectangleRec(rec, tileColors[i])
		if rl.CheckCollisionPointRec(mouseposScreen, rec) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			colorNum = i
		}
		if colorNum == i {
			if tileColors[i] == rl.Red || tileColors[i] == rl.Violet || tileColors[i] == rl.Magenta || tileColors[i] == rl.Pink || tileColors[i] == rl.Purple {
				rl.DrawRectangleLinesEx(rec, 4, rl.Fade(rl.Yellow, blinkFade))
			} else {
				rl.DrawRectangleLinesEx(rec, 4, rl.Fade(rl.Red, blinkFade))
			}
		}
		y += size + 8
		if i == (len(tileRecs)/2)-1 {
			x += size + 8
			y = yorig
		}

	}

}
func makegridtiles() {

	gridA = gridW * gridW //CALCULATE GRID AREA

	x := cnt.X - (float32(gridW/2) * size) //OFFSET FROM CENTER
	xorig := x
	y := cnt.Y - (float32(gridW/2) * size)

	count := 0 //COUNT USED TO DETERMINE END OF LINE MOVE TO NEXT LINE

	for i := 0; i < gridA; i++ {
		newTile := blok{}                               //CREATE AN EMPTY TILE BLOCK
		newTile.rec = rl.NewRectangle(x, y, size, size) //DEFINE THE RECTANGLE
		tiles = append(tiles, newTile)                  //ADD TO SLICE

		x += size           //MOVE X BY SIZE OF TILE
		count++             //INCREASE COUNT
		if count == gridW { //IF COUNT = GRID WIDTH
			count = 0 //RESET LINE COUNT
			x = xorig //RETURN X TO ORIGINAL POSITION
			y += size //MOVE Y DOWN BY TILE SIZE
		}
	}

	//CREATE TILE IMAGE RECTANGLES
	x = 0 //SET X AND Y TO TOP LEFT CORNER (0,0) OF TILE SHEET
	y = 0

	for {
		tileRecs = append(tileRecs, rl.NewRectangle(x, y, 16, 16)) //ADD TO SLICE
		x += 16                                                    //MOVE X BY WIDTH OF 1 TILE 16 PIXELS
		if x >= float32(tileImgs.Width) {                          //IF X = WIDTH OF IMAGE MOVE TO NEXT LINE
			x = 0
			y += 16
		}
		if y >= float32(tileImgs.Height) { //IF Y = HEIGHT OF IMAGE BREAK
			break
		}
	}

}

// RETURNS A RANDOM INTEGER
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}
