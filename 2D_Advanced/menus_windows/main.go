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
	scrW, scrH int // SCREEN WIDTH & HEIGHT

	side, topbot, centr, poptb xmenu //MENU STRUCTS

	colNum int //COLOR NUMBER

	cnt, mouse, clickV2, cornerV2 rl.Vector2 //SCREEN CENTER/MOUSE POSITION/MOVE WINDOW VECTOR2s

	move bool //ON//OFF FOR MOVE WINDOW

	camCircs []xcirc //SLICE OF CIRCLE STRUCTS FOR CAMERA DISPLAY

	colors = []rl.Color{rl.Red, rl.Magenta, rl.Yellow, rl.Green, rl.Blue, rl.DarkBlue, rl.DarkGray, rl.Purple, rl.Orange} //SLICE OF COLORS

	delta, distX, distY float32 //FRAME TIME FOR POP UP MOVEMENT / DISTANCE FOR CENTER MENU MOVE

	camera rl.Camera2D //CAMERA OF BACKGROUND WINDOW
)

// STRUCT WITH MENU INFORMATION
type xmenu struct {
	rec, tabrec rl.Rectangle
	lr, tb, oc  bool
	col         rl.Color
	fd          float32
	name        string
}

// STRUCT WITH CIRCLE INFORMATION
type xcirc struct {
	cnt rl.Vector2
	rad float32
	col rl.Color
	fd  float32
}

func main() {

	rl.InitWindow(0, 0, "menus & windows - raylib go - https://github.com/unklnik/raylib-go-more-examples")
	rl.SetWindowState(rl.FlagBorderlessWindowedMode | rl.FlagMsaa4xHint)

	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight() // GET SCREEN SIZES
	rl.SetWindowSize(scrW, scrH)                           // SET WINDOW SIZE

	//rl.ToggleFullscreen() // UNCOMMENT IF YOU HAVE DISPLAY ISSUES WITH OVERLAPPING WINDOW BARS

	cnt = rl.NewVector2(float32(scrW/2), float32(scrH/2)) //SCREEN CENTER
	camera.Zoom = 1.0                                     //SETS CAMERA ZOOM

	makeMenus() //MAKE INITAL MENUS/WINDOWS
	makeCircs() //MAKE BACKGROUND CIRCLES

	rl.SetTargetFPS(60) // NUMBER OF FRAMES DRAWN IN A SECOND

	for !rl.WindowShouldClose() {

		delta = rl.GetFrameTime() //GET FRAME TIME FOR USE LATER

		upMenus() //UPDATE MENUS

		mouse = rl.GetMousePosition() //GET POISITON OF MOUSE CURSOR

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		//DRAW CIRCLES IN CAMERA VIEW
		for i := 0; i < len(camCircs); i++ {
			rl.DrawCircleV(camCircs[i].cnt, camCircs[i].rad, rl.Fade(camCircs[i].col, camCircs[i].fd))
			rl.DrawCircleLines(int32(camCircs[i].cnt.X), int32(camCircs[i].cnt.Y), camCircs[i].rad, camCircs[i].col)
		}

		rl.EndMode2D()

		//DRAW MENUS ABOVE LAYER OUTSIDE OF CAMERA
		drawMenus()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// DRAW MENUS
func drawMenus() {

	txs := int32(20)                  //TEXT SIZE
	titleBarH := float32(txs + txs/2) //HEIGHT OF WINDOW/MENU TITLE BAR

	//SIDE
	titleRec := rl.NewRectangle(side.rec.X, side.rec.Y, side.rec.Width, titleBarH)                                         //TITLEBAR RECTANGLE
	rl.DrawRectangleRec(side.rec, rl.Fade(side.col, side.fd))                                                              //DRAW SIDE MENU REC
	rl.DrawRectangleRec(titleRec, side.col)                                                                                //DRAW SIDE MENU TITLE BAR
	side.tabrec = rl.NewRectangle(side.rec.X+side.rec.Width-titleRec.Height, side.rec.Y, titleRec.Height, titleRec.Height) // DEFINE RECTANLGE FOR MOVE LEFT/RIGHT SWITCH
	if side.lr {
		side.tabrec.X = side.rec.X //CHANGE POSITION BASED ON LEFT/RIGHT (L/R)
	}
	side.tabrec.X += 4 //MAKE SLIGHTLY SMALLER THAN TITLEBAR
	side.tabrec.Y += 4
	side.tabrec.Width -= 8
	side.tabrec.Height -= 8

	txlen := rl.MeasureText(side.name, txs) //FIND TEXT LENGTH FOR CENTER
	xtx := int32(side.rec.X+side.rec.Width/2) - txlen/2
	ytx := int32(titleRec.Y+titleRec.Height/2) - txs/2

	rl.DrawText(side.name, xtx, ytx, txs, rl.Black) //DRAW SIDE MENU TITLE

	rl.DrawRectangleRec(side.tabrec, rl.Black) // DRAW L/R SWITCH
	rl.DrawRectangleRec(side.tabrec, rl.Fade(side.col, side.fd))
	rl.DrawRectangleLinesEx(side.tabrec, 1, rl.Black)

	// DRAW L/R SWITCH TEXT
	txt := ">>"
	if side.lr {
		txt = "<<"
	}
	txlen = rl.MeasureText(txt, txs)
	xtx = int32(side.tabrec.X+side.tabrec.Width/2) - txlen/2
	ytx = int32(side.tabrec.Y+side.tabrec.Height/2) - txs/2
	rl.DrawText(txt, xtx, ytx, txs, rl.Black)

	// L/R SWITCH ON/OFF
	if rl.CheckCollisionPointRec(mouse, side.tabrec) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			side.lr = !side.lr
		}
	}

	//CAMERA ZOOM MENU ITEM
	ytx += titleRec.ToInt32().Height * 2
	txt = "camera zoom > click"
	txlen = rl.MeasureText(txt, txs)
	xtx = side.rec.ToInt32().X + side.rec.ToInt32().Width/2 - txlen/2
	backRec := rl.NewRectangle(side.rec.X, float32(ytx)-2, side.rec.Width, float32(txs)+4)
	if rl.CheckCollisionPointRec(mouse, backRec) {
		rl.DrawRectangleRec(backRec, side.col)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) { //IF MOUSE CLICK CHANGE CAMERA ZOOM
			if camera.Zoom == 1 {
				camera.Zoom = 1.5
			} else if camera.Zoom == 1.5 {
				camera.Zoom = 2
			} else if camera.Zoom == 2 {
				camera.Zoom = 0.5
			} else if camera.Zoom == 0.5 {
				camera.Zoom = 1
			}
			camera.Target = cnt
			camera.Offset.X = float32(scrW / 2)
			camera.Offset.Y = float32(scrH / 2)
		}
	}
	rl.DrawText(txt, xtx, ytx, txs, rl.Black)

	//COLORS MENU ITEM
	ytx += txs + txs/2
	txt = "colors > click"
	txlen = rl.MeasureText(txt, txs)
	xtx = side.rec.ToInt32().X + side.rec.ToInt32().Width/2 - txlen/2
	backRec = rl.NewRectangle(side.rec.X, float32(ytx)-2, side.rec.Width, float32(txs)+4)
	if rl.CheckCollisionPointRec(mouse, backRec) {
		rl.DrawRectangleRec(backRec, side.col)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) { //IF MOUSE CLICK CHANGE COLOR
			colNum++
			if colNum == len(colors) {
				colNum = 0
			}
			side.col = colors[colNum]
			poptb.col = colors[colNum]
			topbot.col = colors[colNum]
			centr.col = colors[colNum]
		}
	}
	rl.DrawText(txt, xtx, ytx, txs, rl.Black)

	//CENTER
	titleRec = rl.NewRectangle(centr.rec.X, centr.rec.Y, centr.rec.Width, titleBarH)
	rl.DrawRectangleRec(centr.rec, rl.Fade(centr.col, centr.fd))
	rl.DrawRectangleRec(titleRec, centr.col)
	if rl.CheckCollisionPointRec(mouse, titleRec) {

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			clickV2 = rl.GetMousePosition()                    //GET MOUSE POSITION AT CLICK
			cornerV2 = rl.NewVector2(centr.rec.X, centr.rec.Y) //GET WINDOW TOP LEFT CORNER POSITION
			distX = clickV2.X - cornerV2.X                     //CALCULATE DISTANCE FROM MOUSE X TO CORNER X
			distY = clickV2.Y - cornerV2.Y                     //CALCULATE DISTANCE FROM MOUSE Y TO CORNER Y
		}

		//MOVE WINDOW ON/OFF
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			move = true
		} else {
			move = false
		}

	}

	//DRAW WINDOW TITLE
	txlen = rl.MeasureText(centr.name, txs)
	xtx = int32(centr.rec.X+centr.rec.Width/2) - txlen/2
	ytx = int32(titleRec.Y+titleRec.Height/2) - txs/2
	rl.DrawText(centr.name, xtx, ytx, txs, rl.Black)

	//TOP/BOT
	rl.DrawRectangleRec(topbot.rec, rl.Fade(topbot.col, topbot.fd))
	siz := topbot.rec.Height / 2
	topbot.tabrec = rl.NewRectangle(topbot.rec.X+topbot.rec.Width-(siz+siz/2), topbot.rec.Y+siz/2, siz, siz)
	rl.DrawRectangleRec(topbot.tabrec, topbot.col)
	v1, v2, v3 := rl.Vector2{}, rl.Vector2{}, rl.Vector2{}
	siz = siz / 2
	//DRAW TRIANGLE FOR TOP/BOTTOM SWITCH
	if topbot.tb {
		v1 = rl.NewVector2(topbot.tabrec.X+topbot.tabrec.Width/2, topbot.tabrec.Y+topbot.tabrec.Height/2)
		v1.X -= siz / 2
		v1.Y += siz / 2
		v2 = v1
		v2.X += siz
		v3 = rl.NewVector2(topbot.tabrec.X+topbot.tabrec.Width/2, topbot.tabrec.Y+topbot.tabrec.Height/2)
		v3.Y -= siz / 2
		if rl.CheckCollisionPointRec(mouse, topbot.tabrec) {
			rl.DrawTriangle(v2, v3, v1, rl.Green)
		} else {
			rl.DrawTriangle(v2, v3, v1, rl.Black)
		}
	} else {
		v1 = rl.NewVector2(topbot.tabrec.X+topbot.tabrec.Width/2, topbot.tabrec.Y+topbot.tabrec.Height/2)
		v1.X -= siz / 2
		v1.Y -= siz / 2
		v2 = v1
		v2.X += siz
		v3 = rl.NewVector2(topbot.tabrec.X+topbot.tabrec.Width/2, topbot.tabrec.Y+topbot.tabrec.Height/2)
		v3.Y += siz / 2
		if rl.CheckCollisionPointRec(mouse, topbot.tabrec) {
			rl.DrawTriangle(v3, v2, v1, rl.Green)
		} else {
			rl.DrawTriangle(v3, v2, v1, rl.Black)
		}
	}

	if rl.CheckCollisionPointRec(mouse, topbot.tabrec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			//CHANGE POP UP MENU POSITION IF TOP/BOTTOM WINDOW CHANGES
			topbot.tb = !topbot.tb
			poptb.tb = !poptb.tb
			if poptb.tb {
				poptb.rec.Y -= float32(scrH) + (poptb.rec.Height - (float32(scrH) - poptb.rec.Y))
				poptb.tabrec.Y = poptb.rec.Y + poptb.rec.Height
			} else {
				poptb.rec.Y += float32(scrH) + float32(math.Abs(float64(poptb.rec.Y)))
				poptb.tabrec.Y = poptb.rec.Y - poptb.tabrec.Height
			}
		}
	}

	//DRAW SIDE BAR TOP/BOTTOM SHADOW
	if side.lr {
		rl.DrawLine(topbot.rec.ToInt32().X+topbot.rec.ToInt32().Width, topbot.rec.ToInt32().Y, topbot.rec.ToInt32().X+topbot.rec.ToInt32().Width, topbot.rec.ToInt32().Y+topbot.rec.ToInt32().Height, rl.Black)
		rl.DrawRectangleGradientH(topbot.rec.ToInt32().X+topbot.rec.ToInt32().Width-4, topbot.rec.ToInt32().Y, 4, topbot.rec.ToInt32().Height, rl.Blank, rl.Black)
	} else {
		rl.DrawLine(topbot.rec.ToInt32().X+1, topbot.rec.ToInt32().Y, topbot.rec.ToInt32().X+1, topbot.rec.ToInt32().Y+topbot.rec.ToInt32().Height, rl.Black)
		rl.DrawRectangleGradientH(topbot.rec.ToInt32().X+1, topbot.rec.ToInt32().Y, 4, topbot.rec.ToInt32().Height, rl.Black, rl.Blank)
	}

	//POP TOP/BOT
	rl.DrawRectangleRec(poptb.rec, rl.Fade(poptb.col, poptb.fd))
	rl.DrawRectangleRec(poptb.tabrec, poptb.col)
	if rl.CheckCollisionPointRec(mouse, poptb.tabrec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			poptb.oc = !poptb.oc //CHANGE OPEN CLOSE (OC) IF CLICKED
		}
	}
	txt = "click"
	txs = int32(30)
	txlen = rl.MeasureText(txt, txs)
	xtx = poptb.tabrec.ToInt32().X + poptb.tabrec.ToInt32().Width/2 - txlen/2
	ytx = poptb.tabrec.ToInt32().Y + poptb.tabrec.ToInt32().Height/2 - txs/2
	rl.DrawText(txt, xtx-1, ytx+1, txs, rl.Black)
	rl.DrawText(txt, xtx, ytx, txs, rl.Black)

}

func upMenus() {

	//CENTER WINDOW MOVE
	if move {
		if centr.rec.X >= 0 && centr.rec.X <= float32(scrW)-centr.rec.Width {
			centr.rec.X = rl.GetMousePosition().X - distX
		}
		if centr.rec.Y >= 0 && centr.rec.Y <= float32(scrH)-centr.rec.Height {
			centr.rec.Y = rl.GetMousePosition().Y - distY
		}
	}
	//KEEP WINDOW WITHIN SCREEN
	if centr.rec.Y < 0 {
		centr.rec.Y = 0
	}
	if centr.rec.Y > float32(scrH)-centr.rec.Height {
		centr.rec.Y = float32(scrH) - centr.rec.Height
	}
	if centr.rec.X < 0 {
		centr.rec.X = 0
	}
	if centr.rec.X > float32(scrW)-centr.rec.Width {
		centr.rec.X = float32(scrW) - centr.rec.Width
	}

	//POP TOP/BOT POSITION RELATIVE TO SIDEAR POSITION
	if side.lr {
		poptb.rec.X = 100
	} else {
		poptb.rec.X = float32(scrW) - (poptb.rec.Width + 100)
	}
	poptb.tabrec.X = poptb.rec.X

	//MOVE POP UP MENU WINDOW UP/DOWN
	popSpd := float32(8) + delta
	if poptb.oc && topbot.tb {
		if poptb.rec.Y < 0 {
			poptb.rec.Y += popSpd
			poptb.tabrec.Y += popSpd
		}
		if poptb.rec.Y > 0 {
			poptb.rec.Y = 0
			poptb.tabrec.Y = poptb.rec.Height
		}
	} else if !poptb.oc && topbot.tb {
		if poptb.rec.Y > -poptb.rec.Height {
			poptb.rec.Y -= popSpd
			poptb.tabrec.Y -= popSpd
		}
		if poptb.rec.Y < -poptb.rec.Height {
			poptb.rec.Y = -poptb.rec.Height
			poptb.tabrec.Y = 0
		}
	} else if poptb.oc && !topbot.tb {
		if poptb.rec.Y > float32(scrH)-poptb.rec.Height {
			poptb.rec.Y -= popSpd
			poptb.tabrec.Y -= popSpd
		}
		if poptb.rec.Y < float32(scrH)-poptb.rec.Height {
			poptb.rec.Y = float32(scrH) - poptb.rec.Height
			poptb.tabrec.Y = poptb.rec.Y - poptb.tabrec.Height
		}
	} else if !poptb.oc && !topbot.tb {
		if poptb.rec.Y < float32(scrH) {
			poptb.rec.Y += popSpd
			poptb.tabrec.Y += popSpd
		}
		if poptb.rec.Y > float32(scrH) {
			poptb.rec.Y = float32(scrH)
			poptb.tabrec.Y = poptb.rec.Y - poptb.tabrec.Height
		}
	}

	//TOP/BOT POSITION CHANGE
	if topbot.tb {
		topbot.rec.Y = float32(scrH) - topbot.rec.Height
	} else {
		topbot.rec.Y = 0
	}

	//SIDE POSITION CHANGE
	if side.lr {
		topbot.rec.X = 0
		side.rec.X = float32(scrW) - side.rec.Width
	} else {
		topbot.rec.X = side.rec.Width
		side.rec.X = 0
	}

}

func makeMenus() {

	//FILL THE MENU STRUCTS
	colNum = rInt(0, len(colors))
	col := colors[colNum]
	fd := float32(0.5)

	//SIDE
	side.col = col
	side.fd = fd
	side.rec = rl.NewRectangle(0, 0, 300, float32(scrH))
	side.name = "side"

	//TOP/BOT
	topbot.col = col
	topbot.fd = fd
	topbot.rec = rl.NewRectangle(0, 0, float32(scrW), 50)
	topbot.rec.Width -= side.rec.Width
	topbot.name = "top/bottom"

	//CENTER
	centr.col = col
	centr.fd = fd
	centr.rec = rl.NewRectangle(cnt.X-200, cnt.Y-100, 400, 200)
	centr.name = "center - click hold drag moves"

	//POP TOP/BOT
	poptb.col = col
	poptb.fd = fd
	poptb.rec = rl.NewRectangle(float32(scrW)-300, float32(scrH), 200, 300)
	tabsiz := poptb.rec.Width / 2
	poptb.tabrec = rl.NewRectangle(poptb.rec.X, poptb.rec.Y-tabsiz/2, tabsiz, tabsiz/2)
}

// CREATE BACKGROUND CIRCLES
func makeCircs() {
	num := rInt(25, 31)
	for num > 0 {
		zcirc := xcirc{}
		zcirc.col = ranCol()
		zcirc.fd = rF32(0.2, 1.1)
		zcirc.cnt = rl.NewVector2(rF32(0, float32(scrW)), rF32(0, float32(scrH)))
		zcirc.rad = rF32(10, 101)
		camCircs = append(camCircs, zcirc)
		num--
	}
}

// RETURNS RANDOM COLOR
func ranCol() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}

// RETURNS RANDOM INTEGER FOR USE WITH RANDOMCOLOR
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RETURNS RANDOM FLOAT32
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}
