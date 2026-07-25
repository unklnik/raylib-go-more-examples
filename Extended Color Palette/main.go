package main

import (
	"cmp"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"

	r "github.com/gen2brain/raylib-go/raylib"
)

var (
	colors     []COLOR
	scrW, scrH int32   = 1920, 1080
	UNIT       float32 = 32
)

type COLOR struct {
	c r.Color
	t string
}

func main() {

	r.SetConfigFlags(r.FlagVsyncHint)
	r.InitWindow(scrW, scrH, "Raylib-Go Extended Color Palette")

	makeColorPalette()
	sortPalette(colors)
	sortPalettePaleColors(colors)

	for !r.WindowShouldClose() {
		r.BeginDrawing()
		r.ClearBackground(r.Black)

		MS := r.GetMousePosition()
		txt := ""
		var x, y float32 = 8, UNIT * 2
		rec := r.NewRectangle(x, y, UNIT, UNIT)
		recOutline := r.NewRectangle(rec.X-2, rec.Y-2, UNIT+4, UNIT+4)
		for i := range len(colors) {
			if r.CheckCollisionPointRec(MS, rec) {
				txt = colors[i].t
				r.DrawRectangleRec(recOutline, COLRAND())
			}
			r.DrawRectangleRec(rec, colors[i].c)
			rec.X += UNIT + 2
			recOutline = r.NewRectangle(rec.X-2, rec.Y-2, UNIT+4, UNIT+4)
			if rec.X+UNIT+2 > float32(scrW) {
				rec.X = 8
				rec.Y += UNIT + 2
				recOutline = r.NewRectangle(rec.X-2, rec.Y-2, UNIT+4, UNIT+4)
			}
		}

		r.DrawText(txt, int32(UNIT), int32(UNIT/2), 20, r.White)
		r.EndDrawing()
	}

}

func RINT(min, max int) int {
	return min + rand.IntN(max-min+1)
}
func COLRAND() r.Color {
	return colors[RINT(0, len(colors)-1)].c
}

func sortPalette(colors []COLOR) {
	slices.SortFunc(colors, func(a, b COLOR) int {
		hsvA := r.ColorToHSV(a.c)
		hsvB := r.ColorToHSV(b.c)
		if hsvA.X != hsvB.X {
			return cmp.Compare(hsvA.X, hsvB.X)
		}
		if hsvA.Y != hsvB.Y {
			return cmp.Compare(hsvA.Y, hsvB.Y)
		}
		return cmp.Compare(hsvA.Z, hsvB.Z)
	})
}
func sortPalettePaleColors(colors []COLOR) {
	slices.SortFunc(colors, func(a, b COLOR) int {
		hsvA := r.ColorToHSV(a.c)
		hsvB := r.ColorToHSV(b.c)
		isNeutral := func(hsv r.Vector3) bool {
			return hsv.Y < 0.1 || hsv.Z < 0.1
		}
		neutralA := isNeutral(hsvA)
		neutralB := isNeutral(hsvB)
		if neutralA != neutralB {
			if neutralA {
				return -1
			}
			return 1
		}
		if neutralA && neutralB {
			return cmp.Compare(hsvA.Z, hsvB.Z)
		}
		if hsvA.X != hsvB.X {
			return cmp.Compare(hsvA.X, hsvB.X)
		}
		if hsvA.Y != hsvB.Y {
			return cmp.Compare(hsvA.Y, hsvB.Y)
		}
		return cmp.Compare(hsvA.Z, hsvB.Z)
	})
}
func hex2rgb(hexStr string) r.Color {
	hexStr = strings.TrimPrefix(hexStr, "#")
	hexStr = strings.TrimPrefix(hexStr, "0x")
	if len(hexStr) != 6 && len(hexStr) != 8 {
		return r.White
	}
	value, err := strconv.ParseUint(hexStr, 16, 64)
	if err != nil {
		return r.White
	}
	if len(hexStr) == 6 {
		return r.NewColor(
			uint8((value>>16)&0xFF),
			uint8((value>>8)&0xFF),
			uint8(value&0xFF),
			255,
		)
	}
	return r.NewColor(
		uint8((value>>24)&0xFF),
		uint8((value>>16)&0xFF),
		uint8((value>>8)&0xFF),
		uint8(value&0xFF),
	)
}

func makeColorPalette() {
	colors = []COLOR{
		{c: absoluteZero(), t: "absoluteZero"},
		{c: acidGreen(), t: "acidGreen"},
		{c: aero(), t: "aero"},
		{c: africanViolet(), t: "africanViolet"},
		{c: airSuperiorityBlue(), t: "airSuperiorityBlue"},
		{c: aliceBlue(), t: "aliceBlue"},
		{c: alizarin(), t: "alizarin"},
		{c: alloyOrange(), t: "alloyOrange"},
		{c: almond(), t: "almond"},
		{c: amaranthDeepPurple(), t: "amaranthDeepPurple"},
		{c: amaranthPink(), t: "amaranthPink"},
		{c: amaranthPurple(), t: "amaranthPurple"},
		{c: amazon(), t: "amazon"},
		{c: amber(), t: "amber"},
		{c: amethyst(), t: "amethyst"},
		{c: androidGreen(), t: "androidGreen"},
		{c: antiqueBrass(), t: "antiqueBrass"},
		{c: antiqueBronze(), t: "antiqueBronze"},
		{c: antiqueFuchsia(), t: "antiqueFuchsia"},
		{c: antiqueRuby(), t: "antiqueRuby"},
		{c: antiqueWhite(), t: "antiqueWhite"},
		{c: apricot(), t: "apricot"},
		{c: aqua(), t: "aqua"},
		{c: aquamarine(), t: "aquamarine"},
		{c: arcticLime(), t: "arcticLime"},
		{c: artichokeGreen(), t: "artichokeGreen"},
		{c: arylideYellow(), t: "arylideYellow"},
		{c: ashGray(), t: "ashGray"},
		{c: atomicTangerine(), t: "atomicTangerine"},
		{c: aureolin(), t: "aureolin"},
		{c: azure(), t: "azure"},
		{c: azureX11Webr(), t: "azureX11Webr"},
		{c: babyBlue(), t: "babyBlue"},
		{c: babyBlueEyes(), t: "babyBlueEyes"},
		{c: babyPink(), t: "babyPink"},
		{c: babyPowder(), t: "babyPowder"},
		{c: bakerMillerPink(), t: "bakerMillerPink"},
		{c: bananaMania(), t: "bananaMania"},
		{c: barbiePink(), t: "barbiePink"},
		{c: barnRed(), t: "barnRed"},
		{c: battleshipGrey(), t: "battleshipGrey"},
		{c: beauBlue(), t: "beauBlue"},
		{c: beaver(), t: "beaver"},
		{c: beige(), t: "beige"},
		{c: bDazzledBlue(), t: "bDazzledBlue"},
		{c: bigDipORuby(), t: "bigDipORuby"},
		{c: bisque(), t: "bisque"},
		{c: bistre(), t: "bistre"},
		{c: bistreBrown(), t: "bistreBrown"},
		{c: bitterLemon(), t: "bitterLemon"},
		{c: bittersweet(), t: "bittersweet"},
		{c: black(), t: "black"},
		{c: blackBean(), t: "blackBean"},
		{c: blackCoral(), t: "blackCoral"},
		{c: blackOlive(), t: "blackOlive"},
		{c: blackShadows(), t: "blackShadows"},
		{c: blanchedAlmond(), t: "blanchedAlmond"},
		{c: blastOffBronze(), t: "blastOffBronze"},
		{c: bleuDeFrance(), t: "bleuDeFrance"},
		{c: blizzardBlue(), t: "blizzardBlue"},
		{c: bloodRed(), t: "bloodRed"},
		{c: blue(), t: "blue"},
		{c: blueCrayola(), t: "blueCrayola"},
		{c: blueMunsell(), t: "blueMunsell"},
		{c: blueNcs(), t: "blueNcs"},
		{c: bluePantone(), t: "bluePantone"},
		{c: bluePigment(), t: "bluePigment"},
		{c: blueBell(), t: "blueBell"},
		{c: blueGrayCrayola(), t: "blueGrayCrayola"},
		{c: blueJeans(), t: "blueJeans"},
		{c: blueSapphire(), t: "blueSapphire"},
		{c: blueViolet(), t: "blueViolet"},
		{c: blueYonder(), t: "blueYonder"},
		{c: bluetiful(), t: "bluetiful"},
		{c: blush(), t: "blush"},
		{c: bole(), t: "bole"},
		{c: bone(), t: "bone"},
		{c: brickRed(), t: "brickRed"},
		{c: brightLilac(), t: "brightLilac"},
		{c: brightYellowCrayola(), t: "brightYellowCrayola"},
		{c: britishRacingGreen(), t: "britishRacingGreen"},
		{c: bronze(), t: "bronze"},
		{c: brown(), t: "brown"},
		{c: brownSugar(), t: "brownSugar"},
		{c: budGreen(), t: "budGreen"},
		{c: buff(), t: "buff"},
		{c: burgundy(), t: "burgundy"},
		{c: burlywood(), t: "burlywood"},
		{c: burnishedBrown(), t: "burnishedBrown"},
		{c: burntOrange(), t: "burntOrange"},
		{c: burntSienna(), t: "burntSienna"},
		{c: burntUmber(), t: "burntUmber"},
		{c: byzantine(), t: "byzantine"},
		{c: byzantium(), t: "byzantium"},
		{c: cadetBlue(), t: "cadetBlue"},
		{c: cadetGrey(), t: "cadetGrey"},
		{c: cadmiumGreen(), t: "cadmiumGreen"},
		{c: cadmiumOrange(), t: "cadmiumOrange"},
		{c: cafeAuLait(), t: "cafeAuLait"},
		{c: cafeNoir(), t: "cafeNoir"},
		{c: cambridgeBlue(), t: "cambridgeBlue"},
		{c: camel(), t: "camel"},
		{c: cameoPink(), t: "cameoPink"},
		{c: canary(), t: "canary"},
		{c: canaryYellow(), t: "canaryYellow"},
		{c: candyPink(), t: "candyPink"},
		{c: cardinal(), t: "cardinal"},
		{c: caribbeanGreen(), t: "caribbeanGreen"},
		{c: carmine(), t: "carmine"},
		{c: carmineMp(), t: "carmineMp"},
		{c: carnationPink(), t: "carnationPink"},
		{c: carnelian(), t: "carnelian"},
		{c: carolinaBlue(), t: "carolinaBlue"},
		{c: carrotOrange(), t: "carrotOrange"},
		{c: catawba(), t: "catawba"},
		{c: cedarChest(), t: "cedarChest"},
		{c: celadon(), t: "celadon"},
		{c: celeste(), t: "celeste"},
		{c: cerise(), t: "cerise"},
		{c: cerulean(), t: "cerulean"},
		{c: ceruleanBlue(), t: "ceruleanBlue"},
		{c: ceruleanFrost(), t: "ceruleanFrost"},
		{c: ceruleanCrayola(), t: "ceruleanCrayola"},
		{c: ceruleanRgb(), t: "ceruleanRgb"},
		{c: champagne(), t: "champagne"},
		{c: champagnePink(), t: "champagnePink"},
		{c: charcoal(), t: "charcoal"},
		{c: charmPink(), t: "charmPink"},
		{c: chartreuseWeb(), t: "chartreuseWeb"},
		{c: cherryBlossomPink(), t: "cherryBlossomPink"},
		{c: chestnut(), t: "chestnut"},
		{c: chiliRed(), t: "chiliRed"},
		{c: chinaPink(), t: "chinaPink"},
		{c: chineseRed(), t: "chineseRed"},
		{c: chineseViolet(), t: "chineseViolet"},
		{c: chineseYellow(), t: "chineseYellow"},
		{c: choteTraditional(), t: "choteTraditional"},
		{c: choteWeb(), t: "choteWeb"},
		{c: cinereous(), t: "cinereous"},
		{c: cinnabar(), t: "cinnabar"},
		{c: cinnamonSatin(), t: "cinnamonSatin"},
		{c: citrine(), t: "citrine"},
		{c: citron(), t: "citron"},
		{c: claret(), t: "claret"},
		{c: coffee(), t: "coffee"},
		{c: mbiaBlue(), t: "mbiaBlue"},
		{c: congoPink(), t: "congoPink"},
		{c: coolGrey(), t: "coolGrey"},
		{c: copper(), t: "copper"},
		{c: copperCrayola(), t: "copperCrayola"},
		{c: copperPenny(), t: "copperPenny"},
		{c: copperRed(), t: "copperRed"},
		{c: copperRose(), t: "copperRose"},
		{c: coquelicot(), t: "coquelicot"},
		{c: coral(), t: "coral"},
		{c: coralPink(), t: "coralPink"},
		{c: cordovan(), t: "cordovan"},
		{c: corn(), t: "corn"},
		{c: cornflowerBlue(), t: "cornflowerBlue"},
		{c: cornsilk(), t: "cornsilk"},
		{c: cosmicCobalt(), t: "cosmicCobalt"},
		{c: cosmicLatte(), t: "cosmicLatte"},
		{c: coyoteBrown(), t: "coyoteBrown"},
		{c: cottonCandy(), t: "cottonCandy"},
		{c: cream(), t: "cream"},
		{c: crimson(), t: "crimson"},
		{c: crimsonUa(), t: "crimsonUa"},
		{c: culturedPearl(), t: "culturedPearl"},
		{c: cyan(), t: "cyan"},
		{c: cyanProcess(), t: "cyanProcess"},
		{c: cyberGrape(), t: "cyberGrape"},
		{c: cyberYellow(), t: "cyberYellow"},
		{c: cyclamen(), t: "cyclamen"},
		{c: dandelion(), t: "dandelion"},
		{c: darkBrown(), t: "darkBrown"},
		{c: darkByzantium(), t: "darkByzantium"},
		{c: darkCyan(), t: "darkCyan"},
		{c: darkElectricBlue(), t: "darkElectricBlue"},
		{c: darkGoldenrod(), t: "darkGoldenrod"},
		{c: darkGreenX11(), t: "darkGreenX11"},
		{c: darkJungleGreen(), t: "darkJungleGreen"},
		{c: darkKhaki(), t: "darkKhaki"},
		{c: darkLava(), t: "darkLava"},
		{c: darkLiverHorses(), t: "darkLiverHorses"},
		{c: darkMagenta(), t: "darkMagenta"},
		{c: darkOliveGreen(), t: "darkOliveGreen"},
		{c: darkOrange(), t: "darkOrange"},
		{c: darkOrchid(), t: "darkOrchid"},
		{c: darkPurple(), t: "darkPurple"},
		{c: darkRed(), t: "darkRed"},
		{c: darkSalmon(), t: "darkSalmon"},
		{c: darkSeaGreen(), t: "darkSeaGreen"},
		{c: darkSienna(), t: "darkSienna"},
		{c: darkSkyBlue(), t: "darkSkyBlue"},
		{c: darkSlateBlue(), t: "darkSlateBlue"},
		{c: darkSlateGray(), t: "darkSlateGray"},
		{c: darkSpringGreen(), t: "darkSpringGreen"},
		{c: darkTurquoise(), t: "darkTurquoise"},
		{c: darkViolet(), t: "darkViolet"},
		{c: davysGrey(), t: "davysGrey"},
		{c: deepCerise(), t: "deepCerise"},
		{c: deepChampagne(), t: "deepChampagne"},
		{c: deepChestnut(), t: "deepChestnut"},
		{c: deepJungleGreen(), t: "deepJungleGreen"},
		{c: deepPink(), t: "deepPink"},
		{c: deepSaffron(), t: "deepSaffron"},
		{c: deepSkyBlue(), t: "deepSkyBlue"},
		{c: deepSpaceSparkle(), t: "deepSpaceSparkle"},
		{c: deepTaupe(), t: "deepTaupe"},
		{c: denim(), t: "denim"},
		{c: denimBlue(), t: "denimBlue"},
		{c: desert(), t: "desert"},
		{c: desertSand(), t: "desertSand"},
		{c: dimGray(), t: "dimGray"},
		{c: dodgerBlue(), t: "dodgerBlue"},
		{c: drabDarkBrown(), t: "drabDarkBrown"},
		{c: dukeBlue(), t: "dukeBlue"},
		{c: dutchWhite(), t: "dutchWhite"},
		{c: ebony(), t: "ebony"},
		{c: ecru(), t: "ecru"},
		{c: eerieBlack(), t: "eerieBlack"},
		{c: eggplant(), t: "eggplant"},
		{c: eggshell(), t: "eggshell"},
		{c: electricLime(), t: "electricLime"},
		{c: electricPurple(), t: "electricPurple"},
		{c: electricViolet(), t: "electricViolet"},
		{c: emerald(), t: "emerald"},
		{c: eminence(), t: "eminence"},
		{c: englishLavender(), t: "englishLavender"},
		{c: englishRed(), t: "englishRed"},
		{c: englishVermillion(), t: "englishVermillion"},
		{c: englishViolet(), t: "englishViolet"},
		{c: erin(), t: "erin"},
		{c: etonBlue(), t: "etonBlue"},
		{c: fallow(), t: "fallow"},
		{c: faluRed(), t: "faluRed"},
		{c: fandango(), t: "fandango"},
		{c: fandangoPink(), t: "fandangoPink"},
		{c: fawn(), t: "fawn"},
		{c: fernGreen(), t: "fernGreen"},
		{c: fieldDrab(), t: "fieldDrab"},
		{c: fieryRose(), t: "fieryRose"},
		{c: finn(), t: "finn"},
		{c: firebrick(), t: "firebrick"},
		{c: fireEngineRed(), t: "fireEngineRed"},
		{c: flame(), t: "flame"},
		{c: flax(), t: "flax"},
		{c: flirt(), t: "flirt"},
		{c: floralWhite(), t: "floralWhite"},
		{c: forestGreenWeb(), t: "forestGreenWeb"},
		{c: frenchBeige(), t: "frenchBeige"},
		{c: frenchBistre(), t: "frenchBistre"},
		{c: frenchBlue(), t: "frenchBlue"},
		{c: frenchFuchsia(), t: "frenchFuchsia"},
		{c: frenchLilac(), t: "frenchLilac"},
		{c: frenchLime(), t: "frenchLime"},
		{c: frenchMauve(), t: "frenchMauve"},
		{c: frenchPink(), t: "frenchPink"},
		{c: frenchRaspberry(), t: "frenchRaspberry"},
		{c: frenchSkyBlue(), t: "frenchSkyBlue"},
		{c: frenchViolet(), t: "frenchViolet"},
		{c: frostbite(), t: "frostbite"},
		{c: fuchsia(), t: "fuchsia"},
		{c: fuchsiaCrayola(), t: "fuchsiaCrayola"},
		{c: fulvous(), t: "fulvous"},
		{c: fuzzyWuzzy(), t: "fuzzyWuzzy"},
		{c: gainsboro(), t: "gainsboro"},
		{c: gamboge(), t: "gamboge"},
		{c: genericViridian(), t: "genericViridian"},
		{c: ghostWhite(), t: "ghostWhite"},
		{c: glaucous(), t: "glaucous"},
		{c: glossyGrape(), t: "glossyGrape"},
		{c: goGreen(), t: "goGreen"},
		{c: goldMetallic(), t: "goldMetallic"},
		{c: goldWebGolden(), t: "goldWebGolden"},
		{c: goldCrayola(), t: "goldCrayola"},
		{c: goldFusion(), t: "goldFusion"},
		{c: goldenBrown(), t: "goldenBrown"},
		{c: goldenPoppy(), t: "goldenPoppy"},
		{c: goldenYellow(), t: "goldenYellow"},
		{c: goldenrod(), t: "goldenrod"},
		{c: gothamGreen(), t: "gothamGreen"},
		{c: graniteGray(), t: "graniteGray"},
		{c: grannySmithApple(), t: "grannySmithApple"},
		{c: grayWeb(), t: "grayWeb"},
		{c: grayX11Gray(), t: "grayX11Gray"},
		{c: green(), t: "green"},
		{c: greenCrayola(), t: "greenCrayola"},
		{c: greenWeb(), t: "greenWeb"},
		{c: greenMunsell(), t: "greenMunsell"},
		{c: greenNcs(), t: "greenNcs"},
		{c: greenPantone(), t: "greenPantone"},
		{c: greenPigment(), t: "greenPigment"},
		{c: greenBlue(), t: "greenBlue"},
		{c: greenLizard(), t: "greenLizard"},
		{c: greenSheen(), t: "greenSheen"},
		{c: gunmetal(), t: "gunmetal"},
		{c: hansaYellow(), t: "hansaYellow"},
		{c: harlequin(), t: "harlequin"},
		{c: harvestGold(), t: "harvestGold"},
		{c: heatWave(), t: "heatWave"},
		{c: heliotrope(), t: "heliotrope"},
		{c: heliotropeGray(), t: "heliotropeGray"},
		{c: hollywoodCerise(), t: "hollywoodCerise"},
		{c: honoluluBlue(), t: "honoluluBlue"},
		{c: hookersGreen(), t: "hookersGreen"},
		{c: hotMagenta(), t: "hotMagenta"},
		{c: hotPink(), t: "hotPink"},
		{c: hunterGreen(), t: "hunterGreen"},
		{c: iceberg(), t: "iceberg"},
		{c: illuminatingEmerald(), t: "illuminatingEmerald"},
		{c: imperialRed(), t: "imperialRed"},
		{c: inchworm(), t: "inchworm"},
		{c: independence(), t: "independence"},
		{c: indiaGreen(), t: "indiaGreen"},
		{c: indianRed(), t: "indianRed"},
		{c: indianYellow(), t: "indianYellow"},
		{c: indigo(), t: "indigo"},
		{c: indigoDye(), t: "indigoDye"},
		{c: internationalKleinBlue(), t: "internationalKleinBlue"},
		{c: internationalOrangeEngineering(), t: "internationalOrangeEngineering"},
		{c: internationalOrangeGoldenGateBridge(), t: "internationalOrangeGoldenGateBridge"},
		{c: irresistible(), t: "irresistible"},
		{c: isabelline(), t: "isabelline"},
		{c: italianSkyBlue(), t: "italianSkyBlue"},
		{c: ivory(), t: "ivory"},
		{c: japaneseCarmine(), t: "japaneseCarmine"},
		{c: japaneseViolet(), t: "japaneseViolet"},
		{c: jasmine(), t: "jasmine"},
		{c: jazzberryJam(), t: "jazzberryJam"},
		{c: jet(), t: "jet"},
		{c: jonquil(), t: "jonquil"},
		{c: juneBud(), t: "juneBud"},
		{c: jungleGreen(), t: "jungleGreen"},
		{c: kellyGreen(), t: "kellyGreen"},
		{c: keppel(), t: "keppel"},
		{c: keyLime(), t: "keyLime"},
		{c: khakiWeb(), t: "khakiWeb"},
		{c: khakiX11LightKhaki(), t: "khakiX11LightKhaki"},
		{c: kobe(), t: "kobe"},
		{c: kobi(), t: "kobi"},
		{c: kobicha(), t: "kobicha"},
		{c: ksuPurple(), t: "ksuPurple"},
		{c: languidLavender(), t: "languidLavender"},
		{c: lapisLazuli(), t: "lapisLazuli"},
		{c: laserLemon(), t: "laserLemon"},
		{c: laurelGreen(), t: "laurelGreen"},
		{c: lava(), t: "lava"},
		{c: lavenderFloral(), t: "lavenderFloral"},
		{c: lavenderWeb(), t: "lavenderWeb"},
		{c: lavenderBlue(), t: "lavenderBlue"},
		{c: lavenderBlush(), t: "lavenderBlush"},
		{c: lavenderGray(), t: "lavenderGray"},
		{c: lawnGreen(), t: "lawnGreen"},
		{c: lemon(), t: "lemon"},
		{c: lemonChiffon(), t: "lemonChiffon"},
		{c: lemonCurry(), t: "lemonCurry"},
		{c: lemonGlacier(), t: "lemonGlacier"},
		{c: lemonMeringue(), t: "lemonMeringue"},
		{c: lemonYellow(), t: "lemonYellow"},
		{c: lemonYellowCrayola(), t: "lemonYellowCrayola"},
		{c: liberty(), t: "liberty"},
		{c: lightBlue(), t: "lightBlue"},
		{c: lightCoral(), t: "lightCoral"},
		{c: lightCornflowerBlue(), t: "lightCornflowerBlue"},
		{c: lightCyan(), t: "lightCyan"},
		{c: lightFrenchBeige(), t: "lightFrenchBeige"},
		{c: lightGoldenrodYellow(), t: "lightGoldenrodYellow"},
		{c: lightGray(), t: "lightGray"},
		{c: lightGreen(), t: "lightGreen"},
		{c: lightOrange(), t: "lightOrange"},
		{c: lightPeriwinkle(), t: "lightPeriwinkle"},
		{c: lightPink(), t: "lightPink"},
		{c: lightPurple(), t: "lightPurple"},
		{c: lightSalmon(), t: "lightSalmon"},
		{c: lightSeaGreen(), t: "lightSeaGreen"},
		{c: lightSkyBlue(), t: "lightSkyBlue"},
		{c: lightSlateGray(), t: "lightSlateGray"},
		{c: lightSteelBlue(), t: "lightSteelBlue"},
		{c: lightYellow(), t: "lightYellow"},
		{c: lilac(), t: "lilac"},
		{c: lilacLuster(), t: "lilacLuster"},
		{c: limerWheel(), t: "limerWheel"},
		{c: limeWebX11Green(), t: "limeWebX11Green"},
		{c: limeGreen(), t: "limeGreen"},
		{c: linGreen(), t: "linGreen"},
		{c: linen(), t: "linen"},
		{c: lion(), t: "lion"},
		{c: liseranPurple(), t: "liseranPurple"},
		{c: littleBoyBlue(), t: "littleBoyBlue"},
		{c: liver(), t: "liver"},
		{c: liverDogs(), t: "liverDogs"},
		{c: liverOrgan(), t: "liverOrgan"},
		{c: liverChestnut(), t: "liverChestnut"},
		{c: livid(), t: "livid"},
		{c: macaroniAndCheese(), t: "macaroniAndCheese"},
		{c: madderLake(), t: "madderLake"},
		{c: magenta(), t: "magenta"},
		{c: magentaCrayola(), t: "magentaCrayola"},
		{c: magentaDye(), t: "magentaDye"},
		{c: magentaPantone(), t: "magentaPantone"},
		{c: magentaProcess(), t: "magentaProcess"},
		{c: magentaHaze(), t: "magentaHaze"},
		{c: magicMint(), t: "magicMint"},
		{c: magnolia(), t: "magnolia"},
		{c: mahogany(), t: "mahogany"},
		{c: maize(), t: "maize"},
		{c: maizeCrayola(), t: "maizeCrayola"},
		{c: majorelleBlue(), t: "majorelleBlue"},
		{c: malachite(), t: "malachite"},
		{c: manatee(), t: "manatee"},
		{c: mandarin(), t: "mandarin"},
		{c: mango(), t: "mango"},
		{c: mangoTango(), t: "mangoTango"},
		{c: mantis(), t: "mantis"},
		{c: mardiGras(), t: "mardiGras"},
		{c: marigold(), t: "marigold"},
		{c: marianBlue(), t: "marianBlue"},
		{c: maroonCrayola(), t: "maroonCrayola"},
		{c: maroonWeb(), t: "maroonWeb"},
		{c: maroonX11(), t: "maroonX11"},
		{c: mauve(), t: "mauve"},
		{c: mauveTaupe(), t: "mauveTaupe"},
		{c: mauvelous(), t: "mauvelous"},
		{c: maximumBlue(), t: "maximumBlue"},
		{c: maximumBlueGreen(), t: "maximumBlueGreen"},
		{c: maximumBluePurple(), t: "maximumBluePurple"},
		{c: maximumGreen(), t: "maximumGreen"},
		{c: maximumGreenYellow(), t: "maximumGreenYellow"},
		{c: maximumPurple(), t: "maximumPurple"},
		{c: maximumRed(), t: "maximumRed"},
		{c: maximumRedPurple(), t: "maximumRedPurple"},
		{c: maximumYellow(), t: "maximumYellow"},
		{c: maximumYellowRed(), t: "maximumYellowRed"},
		{c: mayGreen(), t: "mayGreen"},
		{c: mayaBlue(), t: "mayaBlue"},
		{c: mediumAquamarine(), t: "mediumAquamarine"},
		{c: mediumBlue(), t: "mediumBlue"},
		{c: mediumCandyAppleRed(), t: "mediumCandyAppleRed"},
		{c: mediumCarmine(), t: "mediumCarmine"},
		{c: mediumChampagne(), t: "mediumChampagne"},
		{c: mediumOrchid(), t: "mediumOrchid"},
		{c: mediumPurple(), t: "mediumPurple"},
		{c: mediumSeaGreen(), t: "mediumSeaGreen"},
		{c: mediumSlateBlue(), t: "mediumSlateBlue"},
		{c: mediumSpringGreen(), t: "mediumSpringGreen"},
		{c: mediumTurquoise(), t: "mediumTurquoise"},
		{c: mediumVioletRed(), t: "mediumVioletRed"},
		{c: mellowApricot(), t: "mellowApricot"},
		{c: mellowYellow(), t: "mellowYellow"},
		{c: melon(), t: "melon"},
		{c: metallicGold(), t: "metallicGold"},
		{c: metallicSeaweed(), t: "metallicSeaweed"},
		{c: metallicSunburst(), t: "metallicSunburst"},
		{c: mexicanPink(), t: "mexicanPink"},
		{c: middleBlue(), t: "middleBlue"},
		{c: middleBlueGreen(), t: "middleBlueGreen"},
		{c: middleBluePurple(), t: "middleBluePurple"},
		{c: middleGrey(), t: "middleGrey"},
		{c: middleGreen(), t: "middleGreen"},
		{c: middleGreenYellow(), t: "middleGreenYellow"},
		{c: middlePurple(), t: "middlePurple"},
		{c: middleRed(), t: "middleRed"},
		{c: middleRedPurple(), t: "middleRedPurple"},
		{c: middleYellow(), t: "middleYellow"},
		{c: middleYellowRed(), t: "middleYellowRed"},
		{c: midnight(), t: "midnight"},
		{c: midnightBlue(), t: "midnightBlue"},
		{c: midnightGreenEagleGreen(), t: "midnightGreenEagleGreen"},
		{c: mikadoYellow(), t: "mikadoYellow"},
		{c: mimiPink(), t: "mimiPink"},
		{c: mindaro(), t: "mindaro"},
		{c: ming(), t: "ming"},
		{c: minionYellow(), t: "minionYellow"},
		{c: mint(), t: "mint"},
		{c: mintCream(), t: "mintCream"},
		{c: mintGreen(), t: "mintGreen"},
		{c: mistyMoss(), t: "mistyMoss"},
		{c: mistyRose(), t: "mistyRose"},
		{c: moccasin(), t: "moccasin"},
		{c: modeBeige(), t: "modeBeige"},
		{c: monaLisa(), t: "monaLisa"},
		{c: morningBlue(), t: "morningBlue"},
		{c: mossGreen(), t: "mossGreen"},
		{c: mountainMeadow(), t: "mountainMeadow"},
		{c: mountbattenPink(), t: "mountbattenPink"},
		{c: msuGreen(), t: "msuGreen"},
		{c: mulberry(), t: "mulberry"},
		{c: mulberryCrayola(), t: "mulberryCrayola"},
		{c: mustard(), t: "mustard"},
		{c: myrtleGreen(), t: "myrtleGreen"},
		{c: mystic(), t: "mystic"},
		{c: mysticMaroon(), t: "mysticMaroon"},
		{c: nadeshikoPink(), t: "nadeshikoPink"},
		{c: naplesYellow(), t: "naplesYellow"},
		{c: navajoWhite(), t: "navajoWhite"},
		{c: navyBlue(), t: "navyBlue"},
		{c: navyBlueCrayola(), t: "navyBlueCrayola"},
		{c: neonBlue(), t: "neonBlue"},
		{c: neonGreen(), t: "neonGreen"},
		{c: neonFuchsia(), t: "neonFuchsia"},
		{c: newCar(), t: "newCar"},
		{c: newYorkPink(), t: "newYorkPink"},
		{c: nickel(), t: "nickel"},
		{c: nonPhotoBlue(), t: "nonPhotoBlue"},
		{c: nyanza(), t: "nyanza"},
		{c: ocherOchre(), t: "ocherOchre"},
		{c: oldBurgundy(), t: "oldBurgundy"},
		{c: oldGold(), t: "oldGold"},
		{c: oldLace(), t: "oldLace"},
		{c: oldLavender(), t: "oldLavender"},
		{c: oldMauve(), t: "oldMauve"},
		{c: oldRose(), t: "oldRose"},
		{c: oldSilver(), t: "oldSilver"},
		{c: olive(), t: "olive"},
		{c: oliveDrab3(), t: "oliveDrab3"},
		{c: oliveDrab7(), t: "oliveDrab7"},
		{c: oliveGreen(), t: "oliveGreen"},
		{c: olivine(), t: "olivine"},
		{c: onyx(), t: "onyx"},
		{c: opal(), t: "opal"},
		{c: operaMauve(), t: "operaMauve"},
		{c: orange(), t: "orange"},
		{c: orangeCrayola(), t: "orangeCrayola"},
		{c: orangePantone(), t: "orangePantone"},
		{c: orangeWeb(), t: "orangeWeb"},
		{c: orangePeel(), t: "orangePeel"},
		{c: orangeRed(), t: "orangeRed"},
		{c: orangeRedCrayola(), t: "orangeRedCrayola"},
		{c: orangeSoda(), t: "orangeSoda"},
		{c: orangeYellow(), t: "orangeYellow"},
		{c: orangeYellowCrayola(), t: "orangeYellowCrayola"},
		{c: orchid(), t: "orchid"},
		{c: orchidPink(), t: "orchidPink"},
		{c: orchidCrayola(), t: "orchidCrayola"},
		{c: outerSpaceCrayola(), t: "outerSpaceCrayola"},
		{c: outrageousOrange(), t: "outrageousOrange"},
		{c: oxblood(), t: "oxblood"},
		{c: oxfordBlue(), t: "oxfordBlue"},
		{c: ouCrimsonRed(), t: "ouCrimsonRed"},
		{c: pacificBlue(), t: "pacificBlue"},
		{c: pakistanGreen(), t: "pakistanGreen"},
		{c: palatinatePurple(), t: "palatinatePurple"},
		{c: paleAqua(), t: "paleAqua"},
		{c: paleCerulean(), t: "paleCerulean"},
		{c: paleDogwood(), t: "paleDogwood"},
		{c: palePink(), t: "palePink"},
		{c: palePurplePantone(), t: "palePurplePantone"},
		{c: paleSpringBud(), t: "paleSpringBud"},
		{c: pansyPurple(), t: "pansyPurple"},
		{c: paoloVeroneseGreen(), t: "paoloVeroneseGreen"},
		{c: papayaWhip(), t: "papayaWhip"},
		{c: paradisePink(), t: "paradisePink"},
		{c: parchment(), t: "parchment"},
		{c: parisGreen(), t: "parisGreen"},
		{c: pastelPink(), t: "pastelPink"},
		{c: patriarch(), t: "patriarch"},
		{c: paua(), t: "paua"},
		{c: paynesGrey(), t: "paynesGrey"},
		{c: peach(), t: "peach"},
		{c: peachCrayola(), t: "peachCrayola"},
		{c: peachPuff(), t: "peachPuff"},
		{c: pear(), t: "pear"},
		{c: pearlyPurple(), t: "pearlyPurple"},
		{c: periwinkle(), t: "periwinkle"},
		{c: periwinkleCrayola(), t: "periwinkleCrayola"},
		{c: permanentGeraniumLake(), t: "permanentGeraniumLake"},
		{c: persianBlue(), t: "persianBlue"},
		{c: persianGreen(), t: "persianGreen"},
		{c: persianIndigo(), t: "persianIndigo"},
		{c: persianOrange(), t: "persianOrange"},
		{c: persianPink(), t: "persianPink"},
		{c: persianPlum(), t: "persianPlum"},
		{c: persianRed(), t: "persianRed"},
		{c: persianRose(), t: "persianRose"},
		{c: persimmon(), t: "persimmon"},
		{c: petunia(), t: "petunia"},
		{c: pewterBlue(), t: "pewterBlue"},
		{c: phlox(), t: "phlox"},
		{c: phthaloBlue(), t: "phthaloBlue"},
		{c: phthaloGreen(), t: "phthaloGreen"},
		{c: picoteeBlue(), t: "picoteeBlue"},
		{c: pictorialCarmine(), t: "pictorialCarmine"},
		{c: piggyPink(), t: "piggyPink"},
		{c: pineGreen(), t: "pineGreen"},
		{c: pineGreen(), t: "pineGreen"},
		{c: pink(), t: "pink"},
		{c: pinkPantone(), t: "pinkPantone"},
		{c: pinkLace(), t: "pinkLace"},
		{c: pinkLavender(), t: "pinkLavender"},
		{c: pinkSherbet(), t: "pinkSherbet"},
		{c: pistachio(), t: "pistachio"},
		{c: platinum(), t: "platinum"},
		{c: plum(), t: "plum"},
		{c: plumWeb(), t: "plumWeb"},
		{c: plumpPurple(), t: "plumpPurple"},
		{c: polishedPine(), t: "polishedPine"},
		{c: pompAndPower(), t: "pompAndPower"},
		{c: popstar(), t: "popstar"},
		{c: portlandOrange(), t: "portlandOrange"},
		{c: powderBlue(), t: "powderBlue"},
		{c: prairieGold(), t: "prairieGold"},
		{c: princetonOrange(), t: "princetonOrange"},
		{c: processCyan(), t: "processCyan"},
		{c: prune(), t: "prune"},
		{c: prussianBlue(), t: "prussianBlue"},
		{c: psychedelicPurple(), t: "psychedelicPurple"},
		{c: puce(), t: "puce"},
		{c: pullmanBrownUpsBrown(), t: "pullmanBrownUpsBrown"},
		{c: pumpkin(), t: "pumpkin"},
		{c: purple(), t: "purple"},
		{c: purpleWeb(), t: "purpleWeb"},
		{c: purpleMunsell(), t: "purpleMunsell"},
		{c: purpleX11(), t: "purpleX11"},
		{c: purpleMountainMajesty(), t: "purpleMountainMajesty"},
		{c: purpleNavy(), t: "purpleNavy"},
		{c: purplePizzazz(), t: "purplePizzazz"},
		{c: purplePlum(), t: "purplePlum"},
		{c: queenBlue(), t: "queenBlue"},
		{c: queenPink(), t: "queenPink"},
		{c: quickSilver(), t: "quickSilver"},
		{c: quinacridoneMagenta(), t: "quinacridoneMagenta"},
		{c: radicalRed(), t: "radicalRed"},
		{c: raisinBlack(), t: "raisinBlack"},
		{c: rajah(), t: "rajah"},
		{c: raspberry(), t: "raspberry"},
		{c: raspberryGlace(), t: "raspberryGlace"},
		{c: raspberryRose(), t: "raspberryRose"},
		{c: rawSienna(), t: "rawSienna"},
		{c: rawUmber(), t: "rawUmber"},
		{c: razzleDazzleRose(), t: "razzleDazzleRose"},
		{c: razzmatazz(), t: "razzmatazz"},
		{c: razzmicBerry(), t: "razzmicBerry"},
		{c: rebeccaPurple(), t: "rebeccaPurple"},
		{c: red(), t: "red"},
		{c: redCrayola(), t: "redCrayola"},
		{c: redMunsell(), t: "redMunsell"},
		{c: redNcs(), t: "redNcs"},
		{c: redPantone(), t: "redPantone"},
		{c: redPigment(), t: "redPigment"},
		{c: redRyb(), t: "redRyb"},
		{c: redOrange(), t: "redOrange"},
		{c: redOcherRedOchre(), t: "redOcherRedOchre"},
		{c: redOrangeCrayola(), t: "redOrangeCrayola"},
		{c: redOrangerWheel(), t: "redOrangerWheel"},
		{c: redPurple(), t: "redPurple"},
		{c: redSalsa(), t: "redSalsa"},
		{c: redViolet(), t: "redViolet"},
		{c: redVioletCrayola(), t: "redVioletCrayola"},
		{c: redVioletrWheel(), t: "redVioletrWheel"},
		{c: redwood(), t: "redwood"},
		{c: resolutionBlue(), t: "resolutionBlue"},
		{c: rhythm(), t: "rhythm"},
		{c: richBlack(), t: "richBlack"},
		{c: richBlackFogra29(), t: "richBlackFogra29"},
		{c: richBlackFogra39(), t: "richBlackFogra39"},
		{c: rifleGreen(), t: "rifleGreen"},
		{c: robinEggBlue(), t: "robinEggBlue"},
		{c: rocketMetallic(), t: "rocketMetallic"},
		{c: rojoSpanishRed(), t: "rojoSpanishRed"},
		{c: romanSilver(), t: "romanSilver"},
		{c: rose(), t: "rose"},
		{c: roseBonbon(), t: "roseBonbon"},
		{c: roseDust(), t: "roseDust"},
		{c: roseEbony(), t: "roseEbony"},
		{c: roseMadder(), t: "roseMadder"},
		{c: rosePink(), t: "rosePink"},
		{c: rosePompadour(), t: "rosePompadour"},
		{c: roseRed(), t: "roseRed"},
		{c: roseTaupe(), t: "roseTaupe"},
		{c: roseVale(), t: "roseVale"},
		{c: rosewood(), t: "rosewood"},
		{c: rossoCorsa(), t: "rossoCorsa"},
		{c: rosyBrown(), t: "rosyBrown"},
		{c: royalBlueDark(), t: "royalBlueDark"},
		{c: royalBlueLight(), t: "royalBlueLight"},
		{c: royalPurple(), t: "royalPurple"},
		{c: royalYellow(), t: "royalYellow"},
		{c: ruber(), t: "ruber"},
		{c: rubineRed(), t: "rubineRed"},
		{c: ruby(), t: "ruby"},
		{c: rubyRed(), t: "rubyRed"},
		{c: rufous(), t: "rufous"},
		{c: russet(), t: "russet"},
		{c: russianGreen(), t: "russianGreen"},
		{c: russianViolet(), t: "russianViolet"},
		{c: rust(), t: "rust"},
		{c: rustyRed(), t: "rustyRed"},
		{c: sacramentoStateGreen(), t: "sacramentoStateGreen"},
		{c: saddleBrown(), t: "saddleBrown"},
		{c: safetyOrange(), t: "safetyOrange"},
		{c: safetyOrangeBlazeOrange(), t: "safetyOrangeBlazeOrange"},
		{c: safetyYellow(), t: "safetyYellow"},
		{c: saffron(), t: "saffron"},
		{c: sage(), t: "sage"},
		{c: stPatricksBlue(), t: "stPatricksBlue"},
		{c: salmon(), t: "salmon"},
		{c: salmonPink(), t: "salmonPink"},
		{c: sand(), t: "sand"},
		{c: sandDune(), t: "sandDune"},
		{c: sandyBrown(), t: "sandyBrown"},
		{c: sapGreen(), t: "sapGreen"},
		{c: sapphire(), t: "sapphire"},
		{c: sapphireBlue(), t: "sapphireBlue"},
		{c: sapphireCrayola(), t: "sapphireCrayola"},
		{c: satinSheenGold(), t: "satinSheenGold"},
		{c: scarlet(), t: "scarlet"},
		{c: schaussPink(), t: "schaussPink"},
		{c: schoolBusYellow(), t: "schoolBusYellow"},
		{c: screaminGreen(), t: "screaminGreen"},
		{c: seaGreen(), t: "seaGreen"},
		{c: seaGreenCrayola(), t: "seaGreenCrayola"},
		{c: seance(), t: "seance"},
		{c: sealBrown(), t: "sealBrown"},
		{c: seashell(), t: "seashell"},
		{c: secret(), t: "secret"},
		{c: selectiveYellow(), t: "selectiveYellow"},
		{c: sepia(), t: "sepia"},
		{c: shadow(), t: "shadow"},
		{c: shadowBlue(), t: "shadowBlue"},
		{c: shamrockGreen(), t: "shamrockGreen"},
		{c: sheenGreen(), t: "sheenGreen"},
		{c: shimmeringBlush(), t: "shimmeringBlush"},
		{c: shinyShamrock(), t: "shinyShamrock"},
		{c: shockingPink(), t: "shockingPink"},
		{c: shockingPinkCrayola(), t: "shockingPinkCrayola"},
		{c: sienna(), t: "sienna"},
		{c: silver(), t: "silver"},
		{c: silverCrayola(), t: "silverCrayola"},
		{c: silverMetallic(), t: "silverMetallic"},
		{c: silverChalice(), t: "silverChalice"},
		{c: silverPink(), t: "silverPink"},
		{c: silverSand(), t: "silverSand"},
		{c: sinopia(), t: "sinopia"},
		{c: sizzlingRed(), t: "sizzlingRed"},
		{c: sizzlingSunrise(), t: "sizzlingSunrise"},
		{c: skobeloff(), t: "skobeloff"},
		{c: skinr(), t: "skinr"},
		{c: skyBlue(), t: "skyBlue"},
		{c: skyBlueCrayola(), t: "skyBlueCrayola"},
		{c: skyMagenta(), t: "skyMagenta"},
		{c: slateBlue(), t: "slateBlue"},
		{c: slateGray(), t: "slateGray"},
		{c: slimyGreen(), t: "slimyGreen"},
		{c: smitten(), t: "smitten"},
		{c: smokyBlack(), t: "smokyBlack"},
		{c: snow(), t: "snow"},
		{c: solidPink(), t: "solidPink"},
		{c: sonicSilver(), t: "sonicSilver"},
		{c: spaceCadet(), t: "spaceCadet"},
		{c: spanishBistre(), t: "spanishBistre"},
		{c: spanishBlue(), t: "spanishBlue"},
		{c: spanishCarmine(), t: "spanishCarmine"},
		{c: spanishGray(), t: "spanishGray"},
		{c: spanishGreen(), t: "spanishGreen"},
		{c: spanishOrange(), t: "spanishOrange"},
		{c: spanishPink(), t: "spanishPink"},
		{c: spanishRed(), t: "spanishRed"},
		{c: spanishSkyBlue(), t: "spanishSkyBlue"},
		{c: spanishViolet(), t: "spanishViolet"},
		{c: spanishViridian(), t: "spanishViridian"},
		{c: springBud(), t: "springBud"},
		{c: springFrost(), t: "springFrost"},
		{c: springGreen(), t: "springGreen"},
		{c: springGreenCrayola(), t: "springGreenCrayola"},
		{c: starCommandBlue(), t: "starCommandBlue"},
		{c: steelBlue(), t: "steelBlue"},
		{c: steelPink(), t: "steelPink"},
		{c: stilDeGrainYellow(), t: "stilDeGrainYellow"},
		{c: straw(), t: "straw"},
		{c: strawberry(), t: "strawberry"},
		{c: strawberryBlonde(), t: "strawberryBlonde"},
		{c: strongLimeGreen(), t: "strongLimeGreen"},
		{c: sugarPlum(), t: "sugarPlum"},
		{c: sunglow(), t: "sunglow"},
		{c: sunray(), t: "sunray"},
		{c: sunset(), t: "sunset"},
		{c: superPink(), t: "superPink"},
		{c: sweetBrown(), t: "sweetBrown"},
		{c: syracuseOrange(), t: "syracuseOrange"},
		{c: tan(), t: "tan"},
		{c: tanCrayola(), t: "tanCrayola"},
		{c: tangerine(), t: "tangerine"},
		{c: tangoPink(), t: "tangoPink"},
		{c: tartOrange(), t: "tartOrange"},
		{c: taupe(), t: "taupe"},
		{c: taupeGray(), t: "taupeGray"},
		{c: teaGreen(), t: "teaGreen"},
		{c: teaRose(), t: "teaRose"},
		{c: teal(), t: "teal"},
		{c: tealBlue(), t: "tealBlue"},
		{c: technobotanica(), t: "technobotanica"},
		{c: telemagenta(), t: "telemagenta"},
		{c: tenneTawny(), t: "tenneTawny"},
		{c: terraCotta(), t: "terraCotta"},
		{c: thistle(), t: "thistle"},
		{c: thulianPink(), t: "thulianPink"},
		{c: tickleMePink(), t: "tickleMePink"},
		{c: tiffanyBlue(), t: "tiffanyBlue"},
		{c: timberwolf(), t: "timberwolf"},
		{c: titaniumYellow(), t: "titaniumYellow"},
		{c: tomato(), t: "tomato"},
		{c: tourmaline(), t: "tourmaline"},
		{c: tropicalRainforest(), t: "tropicalRainforest"},
		{c: trueBlue(), t: "trueBlue"},
		{c: trypanBlue(), t: "trypanBlue"},
		{c: tuftsBlue(), t: "tuftsBlue"},
		{c: tumbleweed(), t: "tumbleweed"},
		{c: turquoise(), t: "turquoise"},
		{c: turquoiseBlue(), t: "turquoiseBlue"},
		{c: turquoiseGreen(), t: "turquoiseGreen"},
		{c: turtleGreen(), t: "turtleGreen"},
		{c: tuscan(), t: "tuscan"},
		{c: tuscanBrown(), t: "tuscanBrown"},
		{c: tuscanRed(), t: "tuscanRed"},
		{c: tuscanTan(), t: "tuscanTan"},
		{c: tuscany(), t: "tuscany"},
		{c: twilightLavender(), t: "twilightLavender"},
		{c: tyrianPurple(), t: "tyrianPurple"},
		{c: uaBlue(), t: "uaBlue"},
		{c: uaRed(), t: "uaRed"},
		{c: ultramarine(), t: "ultramarine"},
		{c: ultramarineBlue(), t: "ultramarineBlue"},
		{c: ultraPink(), t: "ultraPink"},
		{c: ultraRed(), t: "ultraRed"},
		{c: umber(), t: "umber"},
		{c: unbleachedSilk(), t: "unbleachedSilk"},
		{c: unitedNationsBlue(), t: "unitedNationsBlue"},
		{c: universityOfPennsylvaniaRed(), t: "universityOfPennsylvaniaRed"},
		{c: unmellowYellow(), t: "unmellowYellow"},
		{c: upForestGreen(), t: "upForestGreen"},
		{c: upMaroon(), t: "upMaroon"},
		{c: upsdellRed(), t: "upsdellRed"},
		{c: uranianBlue(), t: "uranianBlue"},
		{c: usafaBlue(), t: "usafaBlue"},
		{c: vanDykeBrown(), t: "vanDykeBrown"},
		{c: vanilla(), t: "vanilla"},
		{c: vanillaIce(), t: "vanillaIce"},
		{c: vantgBlue(), t: "vantgBlue"},
		{c: vegasGold(), t: "vegasGold"},
		{c: venetianRed(), t: "venetianRed"},
		{c: verdigris(), t: "verdigris"},
		{c: vermilion(), t: "vermilion"},
		{c: vermilion(), t: "vermilion"},
		{c: veronica(), t: "veronica"},
		{c: violet(), t: "violet"},
		{c: electricVioletRgb(), t: "electricVioletRgb"},
		{c: violetCrayola(), t: "violetCrayola"},
		{c: violetRyb(), t: "violetRyb"},
		{c: violetWeb(), t: "violetWeb"},
		{c: violetBlue(), t: "violetBlue"},
		{c: violetBlueCrayola(), t: "violetBlueCrayola"},
		{c: violetRed(), t: "violetRed"},
		{c: violetRedperbang(), t: "violetRedperbang"},
		{c: viridian(), t: "viridian"},
		{c: viridianGreen(), t: "viridianGreen"},
		{c: vividBurgundy(), t: "vividBurgundy"},
		{c: vividSkyBlue(), t: "vividSkyBlue"},
		{c: vividTangerine(), t: "vividTangerine"},
		{c: vividViolet(), t: "vividViolet"},
		{c: volt(), t: "volt"},
		{c: warmBlack(), t: "warmBlack"},
		{c: weezyBlue(), t: "weezyBlue"},
		{c: wheat(), t: "wheat"},
		{c: white(), t: "white"},
		{c: wildBlueYonder(), t: "wildBlueYonder"},
		{c: wildOrchid(), t: "wildOrchid"},
		{c: wildStrawberry(), t: "wildStrawberry"},
		{c: wildWatermelon(), t: "wildWatermelon"},
		{c: willpowerOrange(), t: "willpowerOrange"},
		{c: windsorTan(), t: "windsorTan"},
		{c: wine(), t: "wine"},
		{c: wineRed(), t: "wineRed"},
		{c: wineDregs(), t: "wineDregs"},
		{c: winterSky(), t: "winterSky"},
		{c: wintergreenDream(), t: "wintergreenDream"},
		{c: wisteria(), t: "wisteria"},
		{c: woodBrown(), t: "woodBrown"},
		{c: xanadu(), t: "xanadu"},
		{c: xander(), t: "xander"},
		{c: xanthic(), t: "xanthic"},
		{c: xanthous(), t: "xanthous"},
		{c: xboxGreen(), t: "xboxGreen"},
		{c: xiaomiOrange(), t: "xiaomiOrange"},
		{c: xumo(), t: "xumo"},
		{c: yaleBlue(), t: "yaleBlue"},
		{c: yellow(), t: "yellow"},
		{c: yellowCrayola(), t: "yellowCrayola"},
		{c: yellowMunsell(), t: "yellowMunsell"},
		{c: yellowNcs(), t: "yellowNcs"},
		{c: yellowPantone(), t: "yellowPantone"},
		{c: yellowProcess(), t: "yellowProcess"},
		{c: yellowRyb(), t: "yellowRyb"},
		{c: yellowGreen(), t: "yellowGreen"},
		{c: yellowGreenCrayola(), t: "yellowGreenCrayola"},
		{c: yellowGreenrWheel(), t: "yellowGreenrWheel"},
		{c: yellowOrange(), t: "yellowOrange"},
		{c: yellowOrangerWheel(), t: "yellowOrangerWheel"},
		{c: yellowRose(), t: "yellowRose"},
		{c: yellowSunshine(), t: "yellowSunshine"},
		{c: yinmnBlue(), t: "yinmnBlue"},
		{c: zafferZaffre(), t: "zafferZaffre"},
		{c: zarqa(), t: "zarqa"},
		{c: zeal(), t: "zeal"},
		{c: zebraWhite(), t: "zebraWhite"},
		{c: zincGray(), t: "zincGray"},
		{c: zincWhite(), t: "zincWhite"},
		{c: zinnwalditeBrown(), t: "zinnwalditeBrown"},
		{c: zinzolin(), t: "zinzolin"},
		{c: zirconGray(), t: "zirconGray"},
		{c: zomp(), t: "zomp"},
		{c: zydeco(), t: "zydeco"},
	}
}

//MARK: EXTENDED PALETTE

func absoluteZero() r.Color                        { return hex2rgb("#0048BA") }
func acidGreen() r.Color                           { return hex2rgb("#B0BF1A") }
func aero() r.Color                                { return hex2rgb("#7CB9E8") }
func africanViolet() r.Color                       { return hex2rgb("#B284BE") }
func airSuperiorityBlue() r.Color                  { return hex2rgb("#72A0C1") }
func aliceBlue() r.Color                           { return hex2rgb("#F0F8FF") }
func alizarin() r.Color                            { return hex2rgb("#DB2D43") }
func alloyOrange() r.Color                         { return hex2rgb("#C46210") }
func almond() r.Color                              { return hex2rgb("#EED9C4") }
func amaranthDeepPurple() r.Color                  { return hex2rgb("#9F2B68") }
func amaranthPink() r.Color                        { return hex2rgb("#F19CBB") }
func amaranthPurple() r.Color                      { return hex2rgb("#AB274F") }
func amazon() r.Color                              { return hex2rgb("#3B7A57") }
func amber() r.Color                               { return hex2rgb("#FFBF00") }
func amethyst() r.Color                            { return hex2rgb("#9966CC") }
func androidGreen() r.Color                        { return hex2rgb("#3DDC84") }
func antiqueBrass() r.Color                        { return hex2rgb("#C88A65") }
func antiqueBronze() r.Color                       { return hex2rgb("#665D1E") }
func antiqueFuchsia() r.Color                      { return hex2rgb("#915C83") }
func antiqueRuby() r.Color                         { return hex2rgb("#841B2D") }
func antiqueWhite() r.Color                        { return hex2rgb("#FAEBD7") }
func apricot() r.Color                             { return hex2rgb("#FBCEB1") }
func aqua() r.Color                                { return hex2rgb("#00FFFF") }
func aquamarine() r.Color                          { return hex2rgb("#7FFFD4") }
func arcticLime() r.Color                          { return hex2rgb("#D0FF14") }
func artichokeGreen() r.Color                      { return hex2rgb("#4B6F44") }
func arylideYellow() r.Color                       { return hex2rgb("#E9D66B") }
func ashGray() r.Color                             { return hex2rgb("#B2BEB5") }
func atomicTangerine() r.Color                     { return hex2rgb("#FF9966") }
func aureolin() r.Color                            { return hex2rgb("#FDEE00") }
func azure() r.Color                               { return hex2rgb("#007FFF") }
func azureX11Webr() r.Color                        { return hex2rgb("#F0FFFF") }
func babyBlue() r.Color                            { return hex2rgb("#89CFF0") }
func babyBlueEyes() r.Color                        { return hex2rgb("#A1CAF1") }
func babyPink() r.Color                            { return hex2rgb("#F4C2C2") }
func babyPowder() r.Color                          { return hex2rgb("#FEFEFA") }
func bakerMillerPink() r.Color                     { return hex2rgb("#FF91AF") }
func bananaMania() r.Color                         { return hex2rgb("#FAE7B5") }
func barbiePink() r.Color                          { return hex2rgb("#DA1884") }
func barnRed() r.Color                             { return hex2rgb("#7C0A02") }
func battleshipGrey() r.Color                      { return hex2rgb("#848482") }
func beauBlue() r.Color                            { return hex2rgb("#BCD4E6") }
func beaver() r.Color                              { return hex2rgb("#9F8170") }
func beige() r.Color                               { return hex2rgb("#F5F5DC") }
func bDazzledBlue() r.Color                        { return hex2rgb("#2E5894") }
func bigDipORuby() r.Color                         { return hex2rgb("#9C2542") }
func bisque() r.Color                              { return hex2rgb("#FFE4C4") }
func bistre() r.Color                              { return hex2rgb("#3D2B1F") }
func bistreBrown() r.Color                         { return hex2rgb("#967117") }
func bitterLemon() r.Color                         { return hex2rgb("#CAE00D") }
func bittersweet() r.Color                         { return hex2rgb("#FE6F5E") }
func black() r.Color                               { return hex2rgb("#000000") }
func blackBean() r.Color                           { return hex2rgb("#3D0C02") }
func blackCoral() r.Color                          { return hex2rgb("#54626F") }
func blackOlive() r.Color                          { return hex2rgb("#3B3C36") }
func blackShadows() r.Color                        { return hex2rgb("#BFAFB2") }
func blanchedAlmond() r.Color                      { return hex2rgb("#FFEBCD") }
func blastOffBronze() r.Color                      { return hex2rgb("#A57164") }
func bleuDeFrance() r.Color                        { return hex2rgb("#318CE7") }
func blizzardBlue() r.Color                        { return hex2rgb("#ACE5EE") }
func bloodRed() r.Color                            { return hex2rgb("#660000") }
func blue() r.Color                                { return hex2rgb("#0000FF") }
func blueCrayola() r.Color                         { return hex2rgb("#1F75FE") }
func blueMunsell() r.Color                         { return hex2rgb("#0093AF") }
func blueNcs() r.Color                             { return hex2rgb("#0087BD") }
func bluePantone() r.Color                         { return hex2rgb("#0018A8") }
func bluePigment() r.Color                         { return hex2rgb("#333399") }
func blueBell() r.Color                            { return hex2rgb("#A2A2D0") }
func blueGrayCrayola() r.Color                     { return hex2rgb("#6699CC") }
func blueJeans() r.Color                           { return hex2rgb("#5DADEC") }
func blueSapphire() r.Color                        { return hex2rgb("#126180") }
func blueViolet() r.Color                          { return hex2rgb("#8A2BE2") }
func blueYonder() r.Color                          { return hex2rgb("#5072A7") }
func bluetiful() r.Color                           { return hex2rgb("#3C69E7") }
func blush() r.Color                               { return hex2rgb("#DE5D83") }
func bole() r.Color                                { return hex2rgb("#79443B") }
func bone() r.Color                                { return hex2rgb("#E3DAC9") }
func brickRed() r.Color                            { return hex2rgb("#CB4154") }
func brightLilac() r.Color                         { return hex2rgb("#D891EF") }
func brightYellowCrayola() r.Color                 { return hex2rgb("#FFAA1D") }
func britishRacingGreen() r.Color                  { return hex2rgb("#004225") }
func bronze() r.Color                              { return hex2rgb("#CD7F32") }
func brown() r.Color                               { return hex2rgb("#964B00") }
func brownSugar() r.Color                          { return hex2rgb("#AF6E4D") }
func budGreen() r.Color                            { return hex2rgb("#7BB661") }
func buff() r.Color                                { return hex2rgb("#FFC680") }
func burgundy() r.Color                            { return hex2rgb("#800020") }
func burlywood() r.Color                           { return hex2rgb("#DEB887") }
func burnishedBrown() r.Color                      { return hex2rgb("#A17A74") }
func burntOrange() r.Color                         { return hex2rgb("#CC5500") }
func burntSienna() r.Color                         { return hex2rgb("#E97451") }
func burntUmber() r.Color                          { return hex2rgb("#8A3324") }
func byzantine() r.Color                           { return hex2rgb("#BD33A4") }
func byzantium() r.Color                           { return hex2rgb("#702963") }
func cadetBlue() r.Color                           { return hex2rgb("#5F9EA0") }
func cadetGrey() r.Color                           { return hex2rgb("#91A3B0") }
func cadmiumGreen() r.Color                        { return hex2rgb("#006B3C") }
func cadmiumOrange() r.Color                       { return hex2rgb("#ED872D") }
func cafeAuLait() r.Color                          { return hex2rgb("#A67B5B") }
func cafeNoir() r.Color                            { return hex2rgb("#4B3621") }
func cambridgeBlue() r.Color                       { return hex2rgb("#A3C1AD") }
func camel() r.Color                               { return hex2rgb("#C19A6B") }
func cameoPink() r.Color                           { return hex2rgb("#EFBBCC") }
func canary() r.Color                              { return hex2rgb("#FFFF99") }
func canaryYellow() r.Color                        { return hex2rgb("#FFEF00") }
func candyPink() r.Color                           { return hex2rgb("#E4717A") }
func cardinal() r.Color                            { return hex2rgb("#C41E3A") }
func caribbeanGreen() r.Color                      { return hex2rgb("#00CC99") }
func carmine() r.Color                             { return hex2rgb("#960018") }
func carmineMp() r.Color                           { return hex2rgb("#D70040") }
func carnationPink() r.Color                       { return hex2rgb("#FFA6C9") }
func carnelian() r.Color                           { return hex2rgb("#B31B1B") }
func carolinaBlue() r.Color                        { return hex2rgb("#56A0D3") }
func carrotOrange() r.Color                        { return hex2rgb("#ED9121") }
func catawba() r.Color                             { return hex2rgb("#703642") }
func cedarChest() r.Color                          { return hex2rgb("#C95A49") }
func celadon() r.Color                             { return hex2rgb("#ACE1AF") }
func celeste() r.Color                             { return hex2rgb("#B2FFFF") }
func cerise() r.Color                              { return hex2rgb("#DE3163") }
func cerulean() r.Color                            { return hex2rgb("#007BA7") }
func ceruleanBlue() r.Color                        { return hex2rgb("#2A52BE") }
func ceruleanFrost() r.Color                       { return hex2rgb("#6D9BC3") }
func ceruleanCrayola() r.Color                     { return hex2rgb("#1DACD6") }
func ceruleanRgb() r.Color                         { return hex2rgb("#0040FF") }
func champagne() r.Color                           { return hex2rgb("#F7E7CE") }
func champagnePink() r.Color                       { return hex2rgb("#F1DDCF") }
func charcoal() r.Color                            { return hex2rgb("#36454F") }
func charmPink() r.Color                           { return hex2rgb("#E68FAC") }
func chartreuseWeb() r.Color                       { return hex2rgb("#80FF00") }
func cherryBlossomPink() r.Color                   { return hex2rgb("#FFB7C5") }
func chestnut() r.Color                            { return hex2rgb("#954535") }
func chiliRed() r.Color                            { return hex2rgb("#E23D28") }
func chinaPink() r.Color                           { return hex2rgb("#DE6FA1") }
func chineseRed() r.Color                          { return hex2rgb("#AA381E") }
func chineseViolet() r.Color                       { return hex2rgb("#856088") }
func chineseYellow() r.Color                       { return hex2rgb("#FFB200") }
func choteTraditional() r.Color                    { return hex2rgb("#7B3F00") }
func choteWeb() r.Color                            { return hex2rgb("#D2691E") }
func cinereous() r.Color                           { return hex2rgb("#98817B") }
func cinnabar() r.Color                            { return hex2rgb("#E34234") }
func cinnamonSatin() r.Color                       { return hex2rgb("#CD607E") }
func citrine() r.Color                             { return hex2rgb("#E4D00A") }
func citron() r.Color                              { return hex2rgb("#9FA91F") }
func claret() r.Color                              { return hex2rgb("#7F1734") }
func coffee() r.Color                              { return hex2rgb("#6F4E37") }
func mbiaBlue() r.Color                            { return hex2rgb("#B9D9EB") }
func congoPink() r.Color                           { return hex2rgb("#F88379") }
func coolGrey() r.Color                            { return hex2rgb("#8C92AC") }
func copper() r.Color                              { return hex2rgb("#B87333") }
func copperCrayola() r.Color                       { return hex2rgb("#DA8A67") }
func copperPenny() r.Color                         { return hex2rgb("#AD6F69") }
func copperRed() r.Color                           { return hex2rgb("#CB6D51") }
func copperRose() r.Color                          { return hex2rgb("#996666") }
func coquelicot() r.Color                          { return hex2rgb("#FF3800") }
func coral() r.Color                               { return hex2rgb("#FF7F50") }
func coralPink() r.Color                           { return hex2rgb("#F88379") }
func cordovan() r.Color                            { return hex2rgb("#893F45") }
func corn() r.Color                                { return hex2rgb("#FBEC5D") }
func cornflowerBlue() r.Color                      { return hex2rgb("#6495ED") }
func cornsilk() r.Color                            { return hex2rgb("#FFF8DC") }
func cosmicCobalt() r.Color                        { return hex2rgb("#2E2D88") }
func cosmicLatte() r.Color                         { return hex2rgb("#FFF8E7") }
func coyoteBrown() r.Color                         { return hex2rgb("#81613C") }
func cottonCandy() r.Color                         { return hex2rgb("#FFBCD9") }
func cream() r.Color                               { return hex2rgb("#FFFDD0") }
func crimson() r.Color                             { return hex2rgb("#DC143C") }
func crimsonUa() r.Color                           { return hex2rgb("#9E1B32") }
func culturedPearl() r.Color                       { return hex2rgb("#F5F5F5") }
func cyan() r.Color                                { return hex2rgb("#00FFFF") }
func cyanProcess() r.Color                         { return hex2rgb("#00B7EB") }
func cyberGrape() r.Color                          { return hex2rgb("#58427C") }
func cyberYellow() r.Color                         { return hex2rgb("#FFD300") }
func cyclamen() r.Color                            { return hex2rgb("#F56FA1") }
func dandelion() r.Color                           { return hex2rgb("#FED85D") }
func darkBrown() r.Color                           { return hex2rgb("#654321") }
func darkByzantium() r.Color                       { return hex2rgb("#5D3954") }
func darkCyan() r.Color                            { return hex2rgb("#008B8B") }
func darkElectricBlue() r.Color                    { return hex2rgb("#536878") }
func darkGoldenrod() r.Color                       { return hex2rgb("#B8860B") }
func darkGreenX11() r.Color                        { return hex2rgb("#006400") }
func darkJungleGreen() r.Color                     { return hex2rgb("#1A2421") }
func darkKhaki() r.Color                           { return hex2rgb("#BDB76B") }
func darkLava() r.Color                            { return hex2rgb("#483C32") }
func darkLiverHorses() r.Color                     { return hex2rgb("#543D37") }
func darkMagenta() r.Color                         { return hex2rgb("#8B008B") }
func darkOliveGreen() r.Color                      { return hex2rgb("#556B2F") }
func darkOrange() r.Color                          { return hex2rgb("#FF8C00") }
func darkOrchid() r.Color                          { return hex2rgb("#9932CC") }
func darkPurple() r.Color                          { return hex2rgb("#301934") }
func darkRed() r.Color                             { return hex2rgb("#8B0000") }
func darkSalmon() r.Color                          { return hex2rgb("#E9967A") }
func darkSeaGreen() r.Color                        { return hex2rgb("#8FBC8F") }
func darkSienna() r.Color                          { return hex2rgb("#3C1414") }
func darkSkyBlue() r.Color                         { return hex2rgb("#8CBED6") }
func darkSlateBlue() r.Color                       { return hex2rgb("#483D8B") }
func darkSlateGray() r.Color                       { return hex2rgb("#2F4F4F") }
func darkSpringGreen() r.Color                     { return hex2rgb("#177245") }
func darkTurquoise() r.Color                       { return hex2rgb("#00CED1") }
func darkViolet() r.Color                          { return hex2rgb("#9400D3") }
func davysGrey() r.Color                           { return hex2rgb("#555555") }
func deepCerise() r.Color                          { return hex2rgb("#DA3287") }
func deepChampagne() r.Color                       { return hex2rgb("#FAD6A5") }
func deepChestnut() r.Color                        { return hex2rgb("#B94E48") }
func deepJungleGreen() r.Color                     { return hex2rgb("#004B49") }
func deepPink() r.Color                            { return hex2rgb("#FF1493") }
func deepSaffron() r.Color                         { return hex2rgb("#FF9933") }
func deepSkyBlue() r.Color                         { return hex2rgb("#00BFFF") }
func deepSpaceSparkle() r.Color                    { return hex2rgb("#4A646C") }
func deepTaupe() r.Color                           { return hex2rgb("#7E5E60") }
func denim() r.Color                               { return hex2rgb("#1560BD") }
func denimBlue() r.Color                           { return hex2rgb("#2243B6") }
func desert() r.Color                              { return hex2rgb("#C19A6B") }
func desertSand() r.Color                          { return hex2rgb("#EDC9AF") }
func dimGray() r.Color                             { return hex2rgb("#696969") }
func dodgerBlue() r.Color                          { return hex2rgb("#1E90FF") }
func drabDarkBrown() r.Color                       { return hex2rgb("#4A412A") }
func dukeBlue() r.Color                            { return hex2rgb("#00009C") }
func dutchWhite() r.Color                          { return hex2rgb("#EFDFBB") }
func ebony() r.Color                               { return hex2rgb("#555D50") }
func ecru() r.Color                                { return hex2rgb("#C2B280") }
func eerieBlack() r.Color                          { return hex2rgb("#1B1B1B") }
func eggplant() r.Color                            { return hex2rgb("#614051") }
func eggshell() r.Color                            { return hex2rgb("#F0EAD6") }
func electricLime() r.Color                        { return hex2rgb("#CCFF00") }
func electricPurple() r.Color                      { return hex2rgb("#BF00FF") }
func electricViolet() r.Color                      { return hex2rgb("#8F00FF") }
func emerald() r.Color                             { return hex2rgb("#50C878") }
func eminence() r.Color                            { return hex2rgb("#6C3082") }
func englishLavender() r.Color                     { return hex2rgb("#B48395") }
func englishRed() r.Color                          { return hex2rgb("#AB4B52") }
func englishVermillion() r.Color                   { return hex2rgb("#CC474B") }
func englishViolet() r.Color                       { return hex2rgb("#563C5C") }
func erin() r.Color                                { return hex2rgb("#00FF40") }
func etonBlue() r.Color                            { return hex2rgb("#96C8A2") }
func fallow() r.Color                              { return hex2rgb("#C19A6B") }
func faluRed() r.Color                             { return hex2rgb("#801818") }
func fandango() r.Color                            { return hex2rgb("#B53389") }
func fandangoPink() r.Color                        { return hex2rgb("#DE5285") }
func fawn() r.Color                                { return hex2rgb("#E5AA70") }
func fernGreen() r.Color                           { return hex2rgb("#4F7942") }
func fieldDrab() r.Color                           { return hex2rgb("#6C541E") }
func fieryRose() r.Color                           { return hex2rgb("#FF5470") }
func finn() r.Color                                { return hex2rgb("#683068") }
func firebrick() r.Color                           { return hex2rgb("#B22222") }
func fireEngineRed() r.Color                       { return hex2rgb("#CE2029") }
func flame() r.Color                               { return hex2rgb("#E25822") }
func flax() r.Color                                { return hex2rgb("#EEDC82") }
func flirt() r.Color                               { return hex2rgb("#A2006D") }
func floralWhite() r.Color                         { return hex2rgb("#FFFAF0") }
func forestGreenWeb() r.Color                      { return hex2rgb("#228B22") }
func frenchBeige() r.Color                         { return hex2rgb("#A67B5B") }
func frenchBistre() r.Color                        { return hex2rgb("#856D4D") }
func frenchBlue() r.Color                          { return hex2rgb("#0072BB") }
func frenchFuchsia() r.Color                       { return hex2rgb("#FD3F92") }
func frenchLilac() r.Color                         { return hex2rgb("#86608E") }
func frenchLime() r.Color                          { return hex2rgb("#9EFD38") }
func frenchMauve() r.Color                         { return hex2rgb("#D473D4") }
func frenchPink() r.Color                          { return hex2rgb("#FD6C9E") }
func frenchRaspberry() r.Color                     { return hex2rgb("#C72C48") }
func frenchSkyBlue() r.Color                       { return hex2rgb("#77B5FE") }
func frenchViolet() r.Color                        { return hex2rgb("#8806CE") }
func frostbite() r.Color                           { return hex2rgb("#E936A7") }
func fuchsia() r.Color                             { return hex2rgb("#FF00FF") }
func fuchsiaCrayola() r.Color                      { return hex2rgb("#C154C1") }
func fulvous() r.Color                             { return hex2rgb("#E48400") }
func fuzzyWuzzy() r.Color                          { return hex2rgb("#87421F") }
func gainsboro() r.Color                           { return hex2rgb("#DCDCDC") }
func gamboge() r.Color                             { return hex2rgb("#E49B0F") }
func genericViridian() r.Color                     { return hex2rgb("#007F66") }
func ghostWhite() r.Color                          { return hex2rgb("#F8F8FF") }
func glaucous() r.Color                            { return hex2rgb("#6082B6") }
func glossyGrape() r.Color                         { return hex2rgb("#AB92B3") }
func goGreen() r.Color                             { return hex2rgb("#00AB66") }
func goldMetallic() r.Color                        { return hex2rgb("#D4AF37") }
func goldWebGolden() r.Color                       { return hex2rgb("#FFD700") }
func goldCrayola() r.Color                         { return hex2rgb("#E6BE8A") }
func goldFusion() r.Color                          { return hex2rgb("#85754E") }
func goldenBrown() r.Color                         { return hex2rgb("#996515") }
func goldenPoppy() r.Color                         { return hex2rgb("#FCC200") }
func goldenYellow() r.Color                        { return hex2rgb("#FFDF00") }
func goldenrod() r.Color                           { return hex2rgb("#DAA520") }
func gothamGreen() r.Color                         { return hex2rgb("#00573F") }
func graniteGray() r.Color                         { return hex2rgb("#676767") }
func grannySmithApple() r.Color                    { return hex2rgb("#A8E4A0") }
func grayWeb() r.Color                             { return hex2rgb("#808080") }
func grayX11Gray() r.Color                         { return hex2rgb("#BEBEBE") }
func green() r.Color                               { return hex2rgb("#00FF00") }
func greenCrayola() r.Color                        { return hex2rgb("#1CAC78") }
func greenWeb() r.Color                            { return hex2rgb("#008000") }
func greenMunsell() r.Color                        { return hex2rgb("#00A877") }
func greenNcs() r.Color                            { return hex2rgb("#009F6B") }
func greenPantone() r.Color                        { return hex2rgb("#00AD43") }
func greenPigment() r.Color                        { return hex2rgb("#00A550") }
func greenBlue() r.Color                           { return hex2rgb("#1164B4") }
func greenLizard() r.Color                         { return hex2rgb("#A7F432") }
func greenSheen() r.Color                          { return hex2rgb("#6EAEA1") }
func gunmetal() r.Color                            { return hex2rgb("#2a3439") }
func hansaYellow() r.Color                         { return hex2rgb("#E9D66B") }
func harlequin() r.Color                           { return hex2rgb("#3FFF00") }
func harvestGold() r.Color                         { return hex2rgb("#DA9100") }
func heatWave() r.Color                            { return hex2rgb("#FF7A00") }
func heliotrope() r.Color                          { return hex2rgb("#DF73FF") }
func heliotropeGray() r.Color                      { return hex2rgb("#AA98A9") }
func hollywoodCerise() r.Color                     { return hex2rgb("#F400A1") }
func honoluluBlue() r.Color                        { return hex2rgb("#006DB0") }
func hookersGreen() r.Color                        { return hex2rgb("#49796B") }
func hotMagenta() r.Color                          { return hex2rgb("#FF1DCE") }
func hotPink() r.Color                             { return hex2rgb("#FF69B4") }
func hunterGreen() r.Color                         { return hex2rgb("#355E3B") }
func iceberg() r.Color                             { return hex2rgb("#71A6D2") }
func illuminatingEmerald() r.Color                 { return hex2rgb("#319177") }
func imperialRed() r.Color                         { return hex2rgb("#ED2939") }
func inchworm() r.Color                            { return hex2rgb("#B2EC5D") }
func independence() r.Color                        { return hex2rgb("#4C516D") }
func indiaGreen() r.Color                          { return hex2rgb("#138808") }
func indianRed() r.Color                           { return hex2rgb("#CD5C5C") }
func indianYellow() r.Color                        { return hex2rgb("#E3A857") }
func indigo() r.Color                              { return hex2rgb("#6A5DFF") }
func indigoDye() r.Color                           { return hex2rgb("#00416A") }
func internationalKleinBlue() r.Color              { return hex2rgb("#130a8f") }
func internationalOrangeEngineering() r.Color      { return hex2rgb("#BA160C") }
func internationalOrangeGoldenGateBridge() r.Color { return hex2rgb("#C0362C") }
func irresistible() r.Color                        { return hex2rgb("#B3446C") }
func isabelline() r.Color                          { return hex2rgb("#F4F0EC") }
func italianSkyBlue() r.Color                      { return hex2rgb("#B2FFFF") }
func ivory() r.Color                               { return hex2rgb("#FFFFF0") }
func japaneseCarmine() r.Color                     { return hex2rgb("#9D2933") }
func japaneseViolet() r.Color                      { return hex2rgb("#5B3256") }
func jasmine() r.Color                             { return hex2rgb("#F8DE7E") }
func jazzberryJam() r.Color                        { return hex2rgb("#A50B5E") }
func jet() r.Color                                 { return hex2rgb("#343434") }
func jonquil() r.Color                             { return hex2rgb("#F4CA16") }
func juneBud() r.Color                             { return hex2rgb("#BDDA57") }
func jungleGreen() r.Color                         { return hex2rgb("#29AB87") }
func kellyGreen() r.Color                          { return hex2rgb("#4CBB17") }
func keppel() r.Color                              { return hex2rgb("#3AB09E") }
func keyLime() r.Color                             { return hex2rgb("#E8F48C") }
func khakiWeb() r.Color                            { return hex2rgb("#C3B091") }
func khakiX11LightKhaki() r.Color                  { return hex2rgb("#F0E68C") }
func kobe() r.Color                                { return hex2rgb("#882D17") }
func kobi() r.Color                                { return hex2rgb("#E79FC4") }
func kobicha() r.Color                             { return hex2rgb("#6B4423") }
func ksuPurple() r.Color                           { return hex2rgb("#512888") }
func languidLavender() r.Color                     { return hex2rgb("#D6CADD") }
func lapisLazuli() r.Color                         { return hex2rgb("#26619C") }
func laserLemon() r.Color                          { return hex2rgb("#FFFF66") }
func laurelGreen() r.Color                         { return hex2rgb("#A9BA9D") }
func lava() r.Color                                { return hex2rgb("#CF1020") }
func lavenderFloral() r.Color                      { return hex2rgb("#B57EDC") }
func lavenderWeb() r.Color                         { return hex2rgb("#E6E6FA") }
func lavenderBlue() r.Color                        { return hex2rgb("#CCCCFF") }
func lavenderBlush() r.Color                       { return hex2rgb("#FFF0F5") }
func lavenderGray() r.Color                        { return hex2rgb("#C4C3D0") }
func lawnGreen() r.Color                           { return hex2rgb("#7CFC00") }
func lemon() r.Color                               { return hex2rgb("#FFF700") }
func lemonChiffon() r.Color                        { return hex2rgb("#FFFACD") }
func lemonCurry() r.Color                          { return hex2rgb("#CCA01D") }
func lemonGlacier() r.Color                        { return hex2rgb("#FDFF00") }
func lemonMeringue() r.Color                       { return hex2rgb("#F6EABE") }
func lemonYellow() r.Color                         { return hex2rgb("#FFF44F") }
func lemonYellowCrayola() r.Color                  { return hex2rgb("#FFFF9F") }
func liberty() r.Color                             { return hex2rgb("#545AA7") }
func lightBlue() r.Color                           { return hex2rgb("#ADD8E6") }
func lightCoral() r.Color                          { return hex2rgb("#F08080") }
func lightCornflowerBlue() r.Color                 { return hex2rgb("#93CCEA") }
func lightCyan() r.Color                           { return hex2rgb("#E0FFFF") }
func lightFrenchBeige() r.Color                    { return hex2rgb("#C8AD7F") }
func lightGoldenrodYellow() r.Color                { return hex2rgb("#FAFAD2") }
func lightGray() r.Color                           { return hex2rgb("#D3D3D3") }
func lightGreen() r.Color                          { return hex2rgb("#90EE90") }
func lightOrange() r.Color                         { return hex2rgb("#FED8B1") }
func lightPeriwinkle() r.Color                     { return hex2rgb("#C5CBE1") }
func lightPink() r.Color                           { return hex2rgb("#FFB6C1") }
func lightPurple() r.Color                         { return hex2rgb("#D8BFD8") }
func lightSalmon() r.Color                         { return hex2rgb("#FFA07A") }
func lightSeaGreen() r.Color                       { return hex2rgb("#20B2AA") }
func lightSkyBlue() r.Color                        { return hex2rgb("#87CEFA") }
func lightSlateGray() r.Color                      { return hex2rgb("#778899") }
func lightSteelBlue() r.Color                      { return hex2rgb("#B0C4DE") }
func lightYellow() r.Color                         { return hex2rgb("#FFFFE0") }
func lilac() r.Color                               { return hex2rgb("#C8A2C8") }
func lilacLuster() r.Color                         { return hex2rgb("#AE98AA") }
func limerWheel() r.Color                          { return hex2rgb("#BFFF00") }
func limeWebX11Green() r.Color                     { return hex2rgb("#00FF00") }
func limeGreen() r.Color                           { return hex2rgb("#32CD32") }
func linGreen() r.Color                            { return hex2rgb("#195905") }
func linen() r.Color                               { return hex2rgb("#FAF0E6") }
func lion() r.Color                                { return hex2rgb("#DECC9C") }
func liseranPurple() r.Color                       { return hex2rgb("#DE6FA1") }
func littleBoyBlue() r.Color                       { return hex2rgb("#6CA0DC") }
func liver() r.Color                               { return hex2rgb("#674C47") }
func liverDogs() r.Color                           { return hex2rgb("#B86D29") }
func liverOrgan() r.Color                          { return hex2rgb("#6C2E1F") }
func liverChestnut() r.Color                       { return hex2rgb("#987456") }
func livid() r.Color                               { return hex2rgb("#6699CC") }
func macaroniAndCheese() r.Color                   { return hex2rgb("#FFBD88") }
func madderLake() r.Color                          { return hex2rgb("#CC3336") }
func magenta() r.Color                             { return hex2rgb("#FF00FF") }
func magentaCrayola() r.Color                      { return hex2rgb("#F653A6") }
func magentaDye() r.Color                          { return hex2rgb("#CA1F7B") }
func magentaPantone() r.Color                      { return hex2rgb("#D0417E") }
func magentaProcess() r.Color                      { return hex2rgb("#FF0090") }
func magentaHaze() r.Color                         { return hex2rgb("#9F4576") }
func magicMint() r.Color                           { return hex2rgb("#AAF0D1") }
func magnolia() r.Color                            { return hex2rgb("#F2E8D7") }
func mahogany() r.Color                            { return hex2rgb("#C04000") }
func maize() r.Color                               { return hex2rgb("#FBEC5D") }
func maizeCrayola() r.Color                        { return hex2rgb("#F2C649") }
func majorelleBlue() r.Color                       { return hex2rgb("#6050DC") }
func malachite() r.Color                           { return hex2rgb("#0BDA51") }
func manatee() r.Color                             { return hex2rgb("#979AAA") }
func mandarin() r.Color                            { return hex2rgb("#F37A48") }
func mango() r.Color                               { return hex2rgb("#FDBE02") }
func mangoTango() r.Color                          { return hex2rgb("#FF8243") }
func mantis() r.Color                              { return hex2rgb("#74C365") }
func mardiGras() r.Color                           { return hex2rgb("#880085") }
func marigold() r.Color                            { return hex2rgb("#EAA221") }
func marianBlue() r.Color                          { return hex2rgb("#00488B") }
func maroonCrayola() r.Color                       { return hex2rgb("#C32148") }
func maroonWeb() r.Color                           { return hex2rgb("#800000") }
func maroonX11() r.Color                           { return hex2rgb("#B03060") }
func mauve() r.Color                               { return hex2rgb("#E0B0FF") }
func mauveTaupe() r.Color                          { return hex2rgb("#915F6D") }
func mauvelous() r.Color                           { return hex2rgb("#EF98AA") }
func maximumBlue() r.Color                         { return hex2rgb("#47ABCC") }
func maximumBlueGreen() r.Color                    { return hex2rgb("#30BFBF") }
func maximumBluePurple() r.Color                   { return hex2rgb("#ACACE6") }
func maximumGreen() r.Color                        { return hex2rgb("#5E8C31") }
func maximumGreenYellow() r.Color                  { return hex2rgb("#D9E650") }
func maximumPurple() r.Color                       { return hex2rgb("#733380") }
func maximumRed() r.Color                          { return hex2rgb("#D92121") }
func maximumRedPurple() r.Color                    { return hex2rgb("#A63A79") }
func maximumYellow() r.Color                       { return hex2rgb("#FAFA37") }
func maximumYellowRed() r.Color                    { return hex2rgb("#F2BA49") }
func mayGreen() r.Color                            { return hex2rgb("#4C9141") }
func mayaBlue() r.Color                            { return hex2rgb("#73C2FB") }
func mediumAquamarine() r.Color                    { return hex2rgb("#66DDAA") }
func mediumBlue() r.Color                          { return hex2rgb("#0000CD") }
func mediumCandyAppleRed() r.Color                 { return hex2rgb("#E2062C") }
func mediumCarmine() r.Color                       { return hex2rgb("#AF4035") }
func mediumChampagne() r.Color                     { return hex2rgb("#F3E5AB") }
func mediumOrchid() r.Color                        { return hex2rgb("#BA55D3") }
func mediumPurple() r.Color                        { return hex2rgb("#9370DB") }
func mediumSeaGreen() r.Color                      { return hex2rgb("#3CB371") }
func mediumSlateBlue() r.Color                     { return hex2rgb("#7B68EE") }
func mediumSpringGreen() r.Color                   { return hex2rgb("#00FA9A") }
func mediumTurquoise() r.Color                     { return hex2rgb("#48D1CC") }
func mediumVioletRed() r.Color                     { return hex2rgb("#C71585") }
func mellowApricot() r.Color                       { return hex2rgb("#F8B878") }
func mellowYellow() r.Color                        { return hex2rgb("#F8DE7E") }
func melon() r.Color                               { return hex2rgb("#FEBAAD") }
func metallicGold() r.Color                        { return hex2rgb("#D3AF37") }
func metallicSeaweed() r.Color                     { return hex2rgb("#0A7E8C") }
func metallicSunburst() r.Color                    { return hex2rgb("#9C7C38") }
func mexicanPink() r.Color                         { return hex2rgb("#E4007C") }
func middleBlue() r.Color                          { return hex2rgb("#7ED4E6") }
func middleBlueGreen() r.Color                     { return hex2rgb("#8DD9CC") }
func middleBluePurple() r.Color                    { return hex2rgb("#8B72BE") }
func middleGrey() r.Color                          { return hex2rgb("#8B8680") }
func middleGreen() r.Color                         { return hex2rgb("#4D8C57") }
func middleGreenYellow() r.Color                   { return hex2rgb("#ACBF60") }
func middlePurple() r.Color                        { return hex2rgb("#D982B5") }
func middleRed() r.Color                           { return hex2rgb("#E58E73") }
func middleRedPurple() r.Color                     { return hex2rgb("#A55353") }
func middleYellow() r.Color                        { return hex2rgb("#FFEB00") }
func middleYellowRed() r.Color                     { return hex2rgb("#ECB176") }
func midnight() r.Color                            { return hex2rgb("#702670") }
func midnightBlue() r.Color                        { return hex2rgb("#191970") }
func midnightGreenEagleGreen() r.Color             { return hex2rgb("#004953") }
func mikadoYellow() r.Color                        { return hex2rgb("#FFC40C") }
func mimiPink() r.Color                            { return hex2rgb("#FFDAE9") }
func mindaro() r.Color                             { return hex2rgb("#E3F988") }
func ming() r.Color                                { return hex2rgb("#36747D") }
func minionYellow() r.Color                        { return hex2rgb("#F5E050") }
func mint() r.Color                                { return hex2rgb("#3EB489") }
func mintCream() r.Color                           { return hex2rgb("#F5FFFA") }
func mintGreen() r.Color                           { return hex2rgb("#98FF98") }
func mistyMoss() r.Color                           { return hex2rgb("#BBB477") }
func mistyRose() r.Color                           { return hex2rgb("#FFE4E1") }
func moccasin() r.Color                            { return hex2rgb("#FFE4B5") }
func modeBeige() r.Color                           { return hex2rgb("#967117") }
func monaLisa() r.Color                            { return hex2rgb("#FF948E") }
func morningBlue() r.Color                         { return hex2rgb("#8DA399") }
func mossGreen() r.Color                           { return hex2rgb("#8A9A5B") }
func mountainMeadow() r.Color                      { return hex2rgb("#30BA8F") }
func mountbattenPink() r.Color                     { return hex2rgb("#997A8D") }
func msuGreen() r.Color                            { return hex2rgb("#18453B") }
func mulberry() r.Color                            { return hex2rgb("#C54B8C") }
func mulberryCrayola() r.Color                     { return hex2rgb("#C8509B") }
func mustard() r.Color                             { return hex2rgb("#FFDB58") }
func myrtleGreen() r.Color                         { return hex2rgb("#317873") }
func mystic() r.Color                              { return hex2rgb("#D65282") }
func mysticMaroon() r.Color                        { return hex2rgb("#AD4379") }
func nadeshikoPink() r.Color                       { return hex2rgb("#F6ADC6") }
func naplesYellow() r.Color                        { return hex2rgb("#FADA5E") }
func navajoWhite() r.Color                         { return hex2rgb("#FFDEAD") }
func navyBlue() r.Color                            { return hex2rgb("#000080") }
func navyBlueCrayola() r.Color                     { return hex2rgb("#1974D2") }
func neonBlue() r.Color                            { return hex2rgb("#4666FF") }
func neonGreen() r.Color                           { return hex2rgb("#39FF14") }
func neonFuchsia() r.Color                         { return hex2rgb("#FE4164") }
func newCar() r.Color                              { return hex2rgb("#214FC6") }
func newYorkPink() r.Color                         { return hex2rgb("#D7837F") }
func nickel() r.Color                              { return hex2rgb("#727472") }
func nonPhotoBlue() r.Color                        { return hex2rgb("#A4DDED") }
func nyanza() r.Color                              { return hex2rgb("#E9FFDB") }
func ocherOchre() r.Color                          { return hex2rgb("#CC7722") }
func oldBurgundy() r.Color                         { return hex2rgb("#43302E") }
func oldGold() r.Color                             { return hex2rgb("#CFB53B") }
func oldLace() r.Color                             { return hex2rgb("#FDF5E6") }
func oldLavender() r.Color                         { return hex2rgb("#796878") }
func oldMauve() r.Color                            { return hex2rgb("#673147") }
func oldRose() r.Color                             { return hex2rgb("#C08081") }
func oldSilver() r.Color                           { return hex2rgb("#848482") }
func olive() r.Color                               { return hex2rgb("#808000") }
func oliveDrab3() r.Color                          { return hex2rgb("#6B8E23") }
func oliveDrab7() r.Color                          { return hex2rgb("#3C341F") }
func oliveGreen() r.Color                          { return hex2rgb("#B5B35C") }
func olivine() r.Color                             { return hex2rgb("#9AB973") }
func onyx() r.Color                                { return hex2rgb("#353839") }
func opal() r.Color                                { return hex2rgb("#A8C3BC") }
func operaMauve() r.Color                          { return hex2rgb("#B784A7") }
func orange() r.Color                              { return hex2rgb("#FF8000") }
func orangeCrayola() r.Color                       { return hex2rgb("#FF7538") }
func orangePantone() r.Color                       { return hex2rgb("#FF5800") }
func orangeWeb() r.Color                           { return hex2rgb("#FFA500") }
func orangePeel() r.Color                          { return hex2rgb("#FF9F00") }
func orangeRed() r.Color                           { return hex2rgb("#FF681F") }
func orangeRedCrayola() r.Color                    { return hex2rgb("#FF5349") }
func orangeSoda() r.Color                          { return hex2rgb("#FA5B3D") }
func orangeYellow() r.Color                        { return hex2rgb("#F5BD1F") }
func orangeYellowCrayola() r.Color                 { return hex2rgb("#F8D568") }
func orchid() r.Color                              { return hex2rgb("#DA70D6") }
func orchidPink() r.Color                          { return hex2rgb("#F2BDCD") }
func orchidCrayola() r.Color                       { return hex2rgb("#E29CD2") }
func outerSpaceCrayola() r.Color                   { return hex2rgb("#2D383A") }
func outrageousOrange() r.Color                    { return hex2rgb("#FF6E4A") }
func oxblood() r.Color                             { return hex2rgb("#4A0000") }
func oxfordBlue() r.Color                          { return hex2rgb("#002147") }
func ouCrimsonRed() r.Color                        { return hex2rgb("#841617") }
func pacificBlue() r.Color                         { return hex2rgb("#1CA9C9") }
func pakistanGreen() r.Color                       { return hex2rgb("#006600") }
func palatinatePurple() r.Color                    { return hex2rgb("#682860") }
func paleAqua() r.Color                            { return hex2rgb("#BED3E5") }
func paleCerulean() r.Color                        { return hex2rgb("#9BC4E2") }
func paleDogwood() r.Color                         { return hex2rgb("#ED7A9B") }
func palePink() r.Color                            { return hex2rgb("#FADADD") }
func palePurplePantone() r.Color                   { return hex2rgb("#FAE6FA") }
func paleSpringBud() r.Color                       { return hex2rgb("#ECEBBD") }
func pansyPurple() r.Color                         { return hex2rgb("#78184A") }
func paoloVeroneseGreen() r.Color                  { return hex2rgb("#009B7D") }
func papayaWhip() r.Color                          { return hex2rgb("#FFEFD5") }
func paradisePink() r.Color                        { return hex2rgb("#E63E62") }
func parchment() r.Color                           { return hex2rgb("#F1E9D2") }
func parisGreen() r.Color                          { return hex2rgb("#50C878") }
func pastelPink() r.Color                          { return hex2rgb("#DEA5A4") }
func patriarch() r.Color                           { return hex2rgb("#800080") }
func paua() r.Color                                { return hex2rgb("#1F005E") }
func paynesGrey() r.Color                          { return hex2rgb("#536878") }
func peach() r.Color                               { return hex2rgb("#FFE5B4") }
func peachCrayola() r.Color                        { return hex2rgb("#FFCBA4") }
func peachPuff() r.Color                           { return hex2rgb("#FFDAB9") }
func pear() r.Color                                { return hex2rgb("#D1E231") }
func pearlyPurple() r.Color                        { return hex2rgb("#B768A2") }
func periwinkle() r.Color                          { return hex2rgb("#CCCCFF") }
func periwinkleCrayola() r.Color                   { return hex2rgb("#C3CDE6") }
func permanentGeraniumLake() r.Color               { return hex2rgb("#E12C2C") }
func persianBlue() r.Color                         { return hex2rgb("#1C39BB") }
func persianGreen() r.Color                        { return hex2rgb("#00A693") }
func persianIndigo() r.Color                       { return hex2rgb("#32127A") }
func persianOrange() r.Color                       { return hex2rgb("#D99058") }
func persianPink() r.Color                         { return hex2rgb("#F77FBE") }
func persianPlum() r.Color                         { return hex2rgb("#701C1C") }
func persianRed() r.Color                          { return hex2rgb("#CC3333") }
func persianRose() r.Color                         { return hex2rgb("#FE28A2") }
func persimmon() r.Color                           { return hex2rgb("#EC5800") }
func petunia() r.Color                             { return hex2rgb("#470659") }
func pewterBlue() r.Color                          { return hex2rgb("#8BA8B7") }
func phlox() r.Color                               { return hex2rgb("#DF00FF") }
func phthaloBlue() r.Color                         { return hex2rgb("#000F89") }
func phthaloGreen() r.Color                        { return hex2rgb("#123524") }
func picoteeBlue() r.Color                         { return hex2rgb("#2E2787") }
func pictorialCarmine() r.Color                    { return hex2rgb("#C30B4E") }
func piggyPink() r.Color                           { return hex2rgb("#FDDDE6") }
func pineGreen() r.Color                           { return hex2rgb("#01796F") }
func pineGreen2() r.Color                          { return hex2rgb("#2A2F23") }
func pink() r.Color                                { return hex2rgb("#FFC0CB") }
func pinkPantone() r.Color                         { return hex2rgb("#D74894") }
func pinkLace() r.Color                            { return hex2rgb("#FFDDF4") }
func pinkLavender() r.Color                        { return hex2rgb("#D8B2D1") }
func pinkSherbet() r.Color                         { return hex2rgb("#F78FA7") }
func pistachio() r.Color                           { return hex2rgb("#93C572") }
func platinum() r.Color                            { return hex2rgb("#E5E4E2") }
func plum() r.Color                                { return hex2rgb("#8E4585") }
func plumWeb() r.Color                             { return hex2rgb("#DDA0DD") }
func plumpPurple() r.Color                         { return hex2rgb("#5946B2") }
func polishedPine() r.Color                        { return hex2rgb("#5DA493") }
func pompAndPower() r.Color                        { return hex2rgb("#86608E") }
func popstar() r.Color                             { return hex2rgb("#BE4F62") }
func portlandOrange() r.Color                      { return hex2rgb("#FF5A36") }
func powderBlue() r.Color                          { return hex2rgb("#B0E0E6") }
func prairieGold() r.Color                         { return hex2rgb("#E1CA7A") }
func princetonOrange() r.Color                     { return hex2rgb("#F58025") }
func processCyan() r.Color                         { return hex2rgb("#00B9F2") }
func prune() r.Color                               { return hex2rgb("#701C1C") }
func prussianBlue() r.Color                        { return hex2rgb("#003153") }
func psychedelicPurple() r.Color                   { return hex2rgb("#DF00FF") }
func puce() r.Color                                { return hex2rgb("#CC8899") }
func pullmanBrownUpsBrown() r.Color                { return hex2rgb("#644117") }
func pumpkin() r.Color                             { return hex2rgb("#FF7518") }
func purple() r.Color                              { return hex2rgb("#6A0DAD") }
func purpleWeb() r.Color                           { return hex2rgb("#800080") }
func purpleMunsell() r.Color                       { return hex2rgb("#9F00C5") }
func purpleX11() r.Color                           { return hex2rgb("#A020F0") }
func purpleMountainMajesty() r.Color               { return hex2rgb("#9678B6") }
func purpleNavy() r.Color                          { return hex2rgb("#4E5180") }
func purplePizzazz() r.Color                       { return hex2rgb("#FE4EDA") }
func purplePlum() r.Color                          { return hex2rgb("#9C51B6") }
func queenBlue() r.Color                           { return hex2rgb("#436B95") }
func queenPink() r.Color                           { return hex2rgb("#E8CCD7") }
func quickSilver() r.Color                         { return hex2rgb("#A6A6A6") }
func quinacridoneMagenta() r.Color                 { return hex2rgb("#8E3A59") }
func radicalRed() r.Color                          { return hex2rgb("#FF355E") }
func raisinBlack() r.Color                         { return hex2rgb("#242124") }
func rajah() r.Color                               { return hex2rgb("#FBAB60") }
func raspberry() r.Color                           { return hex2rgb("#E30B5D") }
func raspberryGlace() r.Color                      { return hex2rgb("#915F6D") }
func raspberryRose() r.Color                       { return hex2rgb("#B3446C") }
func rawSienna() r.Color                           { return hex2rgb("#D68A59") }
func rawUmber() r.Color                            { return hex2rgb("#826644") }
func razzleDazzleRose() r.Color                    { return hex2rgb("#FF33CC") }
func razzmatazz() r.Color                          { return hex2rgb("#E3256B") }
func razzmicBerry() r.Color                        { return hex2rgb("#8D4E85") }
func rebeccaPurple() r.Color                       { return hex2rgb("#663399") }
func red() r.Color                                 { return hex2rgb("#FF0000") }
func redCrayola() r.Color                          { return hex2rgb("#EE204D") }
func redMunsell() r.Color                          { return hex2rgb("#F2003C") }
func redNcs() r.Color                              { return hex2rgb("#C40233") }
func redPantone() r.Color                          { return hex2rgb("#ED2939") }
func redPigment() r.Color                          { return hex2rgb("#ED1C24") }
func redRyb() r.Color                              { return hex2rgb("#FE2712") }
func redOrange() r.Color                           { return hex2rgb("#FF5349") }
func redOcherRedOchre() r.Color                    { return hex2rgb("#913831") }
func redOrangeCrayola() r.Color                    { return hex2rgb("#FF681F") }
func redOrangerWheel() r.Color                     { return hex2rgb("#FF4500") }
func redPurple() r.Color                           { return hex2rgb("#E40078") }
func redSalsa() r.Color                            { return hex2rgb("#FD3A4A") }
func redViolet() r.Color                           { return hex2rgb("#C71585") }
func redVioletCrayola() r.Color                    { return hex2rgb("#C0448F") }
func redVioletrWheel() r.Color                     { return hex2rgb("#922B3E") }
func redwood() r.Color                             { return hex2rgb("#A45A52") }
func resolutionBlue() r.Color                      { return hex2rgb("#002387") }
func rhythm() r.Color                              { return hex2rgb("#777696") }
func richBlack() r.Color                           { return hex2rgb("#004040") }
func richBlackFogra29() r.Color                    { return hex2rgb("#010B13") }
func richBlackFogra39() r.Color                    { return hex2rgb("#010203") }
func rifleGreen() r.Color                          { return hex2rgb("#444C38") }
func robinEggBlue() r.Color                        { return hex2rgb("#00CCCC") }
func rocketMetallic() r.Color                      { return hex2rgb("#8A7F80") }
func rojoSpanishRed() r.Color                      { return hex2rgb("#A91101") }
func romanSilver() r.Color                         { return hex2rgb("#838996") }
func rose() r.Color                                { return hex2rgb("#FF0080") }
func roseBonbon() r.Color                          { return hex2rgb("#F9429E") }
func roseDust() r.Color                            { return hex2rgb("#9E5E6F") }
func roseEbony() r.Color                           { return hex2rgb("#674846") }
func roseMadder() r.Color                          { return hex2rgb("#E32636") }
func rosePink() r.Color                            { return hex2rgb("#FF66CC") }
func rosePompadour() r.Color                       { return hex2rgb("#ED7A9B") }
func roseRed() r.Color                             { return hex2rgb("#C21E56") }
func roseTaupe() r.Color                           { return hex2rgb("#905D5D") }
func roseVale() r.Color                            { return hex2rgb("#AB4E52") }
func rosewood() r.Color                            { return hex2rgb("#65000B") }
func rossoCorsa() r.Color                          { return hex2rgb("#D40000") }
func rosyBrown() r.Color                           { return hex2rgb("#BC8F8F") }
func royalBlueDark() r.Color                       { return hex2rgb("#002366") }
func royalBlueLight() r.Color                      { return hex2rgb("#4169E1") }
func royalPurple() r.Color                         { return hex2rgb("#7851A9") }
func royalYellow() r.Color                         { return hex2rgb("#FADA5E") }
func ruber() r.Color                               { return hex2rgb("#CE4676") }
func rubineRed() r.Color                           { return hex2rgb("#D10056") }
func ruby() r.Color                                { return hex2rgb("#E0115F") }
func rubyRed() r.Color                             { return hex2rgb("#9B111E") }
func rufous() r.Color                              { return hex2rgb("#A81C07") }
func russet() r.Color                              { return hex2rgb("#80461B") }
func russianGreen() r.Color                        { return hex2rgb("#679267") }
func russianViolet() r.Color                       { return hex2rgb("#32174D") }
func rust() r.Color                                { return hex2rgb("#B7410E") }
func rustyRed() r.Color                            { return hex2rgb("#DA2C43") }
func sacramentoStateGreen() r.Color                { return hex2rgb("#043927") }
func saddleBrown() r.Color                         { return hex2rgb("#8B4513") }
func safetyOrange() r.Color                        { return hex2rgb("#FF7800") }
func safetyOrangeBlazeOrange() r.Color             { return hex2rgb("#FF6700") }
func safetyYellow() r.Color                        { return hex2rgb("#EED202") }
func saffron() r.Color                             { return hex2rgb("#F4C430") }
func sage() r.Color                                { return hex2rgb("#BCB88A") }
func stPatricksBlue() r.Color                      { return hex2rgb("#23297A") }
func salmon() r.Color                              { return hex2rgb("#FA8072") }
func salmonPink() r.Color                          { return hex2rgb("#FF91A4") }
func sand() r.Color                                { return hex2rgb("#C2B280") }
func sandDune() r.Color                            { return hex2rgb("#967117") }
func sandyBrown() r.Color                          { return hex2rgb("#F4A460") }
func sapGreen() r.Color                            { return hex2rgb("#507D2A") }
func sapphire() r.Color                            { return hex2rgb("#0F52BA") }
func sapphireBlue() r.Color                        { return hex2rgb("#0067A5") }
func sapphireCrayola() r.Color                     { return hex2rgb("#2D5DA1") }
func satinSheenGold() r.Color                      { return hex2rgb("#CBA135") }
func scarlet() r.Color                             { return hex2rgb("#FF2400") }
func schaussPink() r.Color                         { return hex2rgb("#FF91AF") }
func schoolBusYellow() r.Color                     { return hex2rgb("#FFD800") }
func screaminGreen() r.Color                       { return hex2rgb("#66FF66") }
func seaGreen() r.Color                            { return hex2rgb("#2E8B57") }
func seaGreenCrayola() r.Color                     { return hex2rgb("#00FFCD") }
func seance() r.Color                              { return hex2rgb("#612086") }
func sealBrown() r.Color                           { return hex2rgb("#59260B") }
func seashell() r.Color                            { return hex2rgb("#FFF5EE") }
func secret() r.Color                              { return hex2rgb("#764374") }
func selectiveYellow() r.Color                     { return hex2rgb("#FFBA00") }
func sepia() r.Color                               { return hex2rgb("#704214") }
func shadow() r.Color                              { return hex2rgb("#8A795D") }
func shadowBlue() r.Color                          { return hex2rgb("#778BA5") }
func shamrockGreen() r.Color                       { return hex2rgb("#009E60") }
func sheenGreen() r.Color                          { return hex2rgb("#8FD400") }
func shimmeringBlush() r.Color                     { return hex2rgb("#D98695") }
func shinyShamrock() r.Color                       { return hex2rgb("#5FA778") }
func shockingPink() r.Color                        { return hex2rgb("#FC0FC0") }
func shockingPinkCrayola() r.Color                 { return hex2rgb("#FF6FFF") }
func sienna() r.Color                              { return hex2rgb("#882D17") }
func silver() r.Color                              { return hex2rgb("#C0C0C0") }
func silverCrayola() r.Color                       { return hex2rgb("#C9C0BB") }
func silverMetallic() r.Color                      { return hex2rgb("#AAA9AD") }
func silverChalice() r.Color                       { return hex2rgb("#ACACAC") }
func silverPink() r.Color                          { return hex2rgb("#C4AEAD") }
func silverSand() r.Color                          { return hex2rgb("#BFC1C2") }
func sinopia() r.Color                             { return hex2rgb("#CB410B") }
func sizzlingRed() r.Color                         { return hex2rgb("#FF3855") }
func sizzlingSunrise() r.Color                     { return hex2rgb("#FFDB00") }
func skobeloff() r.Color                           { return hex2rgb("#007474") }
func skinr() r.Color                               { return hex2rgb("#FFDEAD") }
func skyBlue() r.Color                             { return hex2rgb("#87CEEB") }
func skyBlueCrayola() r.Color                      { return hex2rgb("#76D7EA") }
func skyMagenta() r.Color                          { return hex2rgb("#CF71AF") }
func slateBlue() r.Color                           { return hex2rgb("#6A5ACD") }
func slateGray() r.Color                           { return hex2rgb("#708090") }
func slimyGreen() r.Color                          { return hex2rgb("#299617") }
func smitten() r.Color                             { return hex2rgb("#C84186") }
func smokyBlack() r.Color                          { return hex2rgb("#100C08") }
func snow() r.Color                                { return hex2rgb("#FFFAFA") }
func solidPink() r.Color                           { return hex2rgb("#893843") }
func sonicSilver() r.Color                         { return hex2rgb("#757575") }
func spaceCadet() r.Color                          { return hex2rgb("#1D2951") }
func spanishBistre() r.Color                       { return hex2rgb("#807532") }
func spanishBlue() r.Color                         { return hex2rgb("#0070B8") }
func spanishCarmine() r.Color                      { return hex2rgb("#D10047") }
func spanishGray() r.Color                         { return hex2rgb("#989898") }
func spanishGreen() r.Color                        { return hex2rgb("#009150") }
func spanishOrange() r.Color                       { return hex2rgb("#E86100") }
func spanishPink() r.Color                         { return hex2rgb("#F7BFBE") }
func spanishRed() r.Color                          { return hex2rgb("#E60026") }
func spanishSkyBlue() r.Color                      { return hex2rgb("#00FFFE") }
func spanishViolet() r.Color                       { return hex2rgb("#4C2882") }
func spanishViridian() r.Color                     { return hex2rgb("#007F5C") }
func springBud() r.Color                           { return hex2rgb("#A7FC00") }
func springFrost() r.Color                         { return hex2rgb("#87FF2A") }
func springGreen() r.Color                         { return hex2rgb("#00FF80") }
func springGreenCrayola() r.Color                  { return hex2rgb("#ECEBBD") }
func starCommandBlue() r.Color                     { return hex2rgb("#007BB8") }
func steelBlue() r.Color                           { return hex2rgb("#4682B4") }
func steelPink() r.Color                           { return hex2rgb("#CC33CC") }
func stilDeGrainYellow() r.Color                   { return hex2rgb("#FADA5E") }
func straw() r.Color                               { return hex2rgb("#E4D96F") }
func strawberry() r.Color                          { return hex2rgb("#FA5053") }
func strawberryBlonde() r.Color                    { return hex2rgb("#FF9361") }
func strongLimeGreen() r.Color                     { return hex2rgb("#33CC33") }
func sugarPlum() r.Color                           { return hex2rgb("#914E75") }
func sunglow() r.Color                             { return hex2rgb("#FFCC33") }
func sunray() r.Color                              { return hex2rgb("#E3AB57") }
func sunset() r.Color                              { return hex2rgb("#FAD6A5") }
func superPink() r.Color                           { return hex2rgb("#CF6BA9") }
func sweetBrown() r.Color                          { return hex2rgb("#A83731") }
func syracuseOrange() r.Color                      { return hex2rgb("#D44500") }
func tan() r.Color                                 { return hex2rgb("#D2B48C") }
func tanCrayola() r.Color                          { return hex2rgb("#D99A6C") }
func tangerine() r.Color                           { return hex2rgb("#F28500") }
func tangoPink() r.Color                           { return hex2rgb("#E4717A") }
func tartOrange() r.Color                          { return hex2rgb("#FB4D46") }
func taupe() r.Color                               { return hex2rgb("#483C32") }
func taupeGray() r.Color                           { return hex2rgb("#8B8589") }
func teaGreen() r.Color                            { return hex2rgb("#D0F0C0") }
func teaRose() r.Color                             { return hex2rgb("#F4C2C2") }
func teal() r.Color                                { return hex2rgb("#008080") }
func tealBlue() r.Color                            { return hex2rgb("#367588") }
func technobotanica() r.Color                      { return hex2rgb("#00FFBF") }
func telemagenta() r.Color                         { return hex2rgb("#CF3476") }
func tenneTawny() r.Color                          { return hex2rgb("#CD5700") }
func terraCotta() r.Color                          { return hex2rgb("#E2725B") }
func thistle() r.Color                             { return hex2rgb("#D8BFD8") }
func thulianPink() r.Color                         { return hex2rgb("#DE6FA1") }
func tickleMePink() r.Color                        { return hex2rgb("#FC89AC") }
func tiffanyBlue() r.Color                         { return hex2rgb("#0ABAB5") }
func timberwolf() r.Color                          { return hex2rgb("#DBD7D2") }
func titaniumYellow() r.Color                      { return hex2rgb("#EEE600") }
func tomato() r.Color                              { return hex2rgb("#FF6347") }
func tourmaline() r.Color                          { return hex2rgb("#86A1A9") }
func tropicalRainforest() r.Color                  { return hex2rgb("#00755E") }
func trueBlue() r.Color                            { return hex2rgb("#2D68C4") }
func trypanBlue() r.Color                          { return hex2rgb("#1C05B3") }
func tuftsBlue() r.Color                           { return hex2rgb("#3E8EDE") }
func tumbleweed() r.Color                          { return hex2rgb("#DEAA88") }
func turquoise() r.Color                           { return hex2rgb("#40E0D0") }
func turquoiseBlue() r.Color                       { return hex2rgb("#00FFEF") }
func turquoiseGreen() r.Color                      { return hex2rgb("#A0D6B4") }
func turtleGreen() r.Color                         { return hex2rgb("#8A9A5B") }
func tuscan() r.Color                              { return hex2rgb("#FAD6A5") }
func tuscanBrown() r.Color                         { return hex2rgb("#6F4E37") }
func tuscanRed() r.Color                           { return hex2rgb("#7C4848") }
func tuscanTan() r.Color                           { return hex2rgb("#A67B5B") }
func tuscany() r.Color                             { return hex2rgb("#C09999") }
func twilightLavender() r.Color                    { return hex2rgb("#8A496B") }
func tyrianPurple() r.Color                        { return hex2rgb("#66023C") }
func uaBlue() r.Color                              { return hex2rgb("#0033AA") }
func uaRed() r.Color                               { return hex2rgb("#D9004C") }
func ultramarine() r.Color                         { return hex2rgb("#3F00FF") }
func ultramarineBlue() r.Color                     { return hex2rgb("#4166F5") }
func ultraPink() r.Color                           { return hex2rgb("#FF6FFF") }
func ultraRed() r.Color                            { return hex2rgb("#FC6C85") }
func umber() r.Color                               { return hex2rgb("#635147") }
func unbleachedSilk() r.Color                      { return hex2rgb("#FFDDCA") }
func unitedNationsBlue() r.Color                   { return hex2rgb("#009EDB") }
func universityOfPennsylvaniaRed() r.Color         { return hex2rgb("#A50021") }
func unmellowYellow() r.Color                      { return hex2rgb("#FFFF66") }
func upForestGreen() r.Color                       { return hex2rgb("#014421") }
func upMaroon() r.Color                            { return hex2rgb("#7B1113") }
func upsdellRed() r.Color                          { return hex2rgb("#AE2029") }
func uranianBlue() r.Color                         { return hex2rgb("#AFDBF5") }
func usafaBlue() r.Color                           { return hex2rgb("#004F98") }
func vanDykeBrown() r.Color                        { return hex2rgb("#664228") }
func vanilla() r.Color                             { return hex2rgb("#F3E5AB") }
func vanillaIce() r.Color                          { return hex2rgb("#F38FA9") }
func vantgBlue() r.Color                           { return hex2rgb("#5271FF") }
func vegasGold() r.Color                           { return hex2rgb("#C5B358") }
func venetianRed() r.Color                         { return hex2rgb("#C80815") }
func verdigris() r.Color                           { return hex2rgb("#43B3AE") }
func vermilion() r.Color                           { return hex2rgb("#E34234") }
func vermilion2() r.Color                          { return hex2rgb("#D9381E") }
func veronica() r.Color                            { return hex2rgb("#A020F0") }
func violet() r.Color                              { return hex2rgb("#8000FF") }
func electricVioletRgb() r.Color                   { return hex2rgb("#8F00FF") }
func violetCrayola() r.Color                       { return hex2rgb("#963D7F") }
func violetRyb() r.Color                           { return hex2rgb("#8601AF") }
func violetWeb() r.Color                           { return hex2rgb("#EE82EE") }
func violetBlue() r.Color                          { return hex2rgb("#324AB2") }
func violetBlueCrayola() r.Color                   { return hex2rgb("#766EC8") }
func violetRed() r.Color                           { return hex2rgb("#F75394") }
func violetRedperbang() r.Color                    { return hex2rgb("#F0599C") }
func viridian() r.Color                            { return hex2rgb("#40826D") }
func viridianGreen() r.Color                       { return hex2rgb("#009698") }
func vividBurgundy() r.Color                       { return hex2rgb("#9F1D35") }
func vividSkyBlue() r.Color                        { return hex2rgb("#00CCFF") }
func vividTangerine() r.Color                      { return hex2rgb("#FFA089") }
func vividViolet() r.Color                         { return hex2rgb("#9F00FF") }
func volt() r.Color                                { return hex2rgb("#CEFF00") }
func warmBlack() r.Color                           { return hex2rgb("#004242") }
func weezyBlue() r.Color                           { return hex2rgb("#189BCC") }
func wheat() r.Color                               { return hex2rgb("#F5DEB3") }
func white() r.Color                               { return hex2rgb("#FFFFFF") }
func wildBlueYonder() r.Color                      { return hex2rgb("#A2ADD0") }
func wildOrchid() r.Color                          { return hex2rgb("#D470A2") }
func wildStrawberry() r.Color                      { return hex2rgb("#FF43A4") }
func wildWatermelon() r.Color                      { return hex2rgb("#FC6C85") }
func willpowerOrange() r.Color                     { return hex2rgb("#FD5800") }
func windsorTan() r.Color                          { return hex2rgb("#A75502") }
func wine() r.Color                                { return hex2rgb("#722F37") }
func wineRed() r.Color                             { return hex2rgb("#B11226") }
func wineDregs() r.Color                           { return hex2rgb("#673147") }
func winterSky() r.Color                           { return hex2rgb("#FF007C") }
func wintergreenDream() r.Color                    { return hex2rgb("#56887D") }
func wisteria() r.Color                            { return hex2rgb("#C9A0DC") }
func woodBrown() r.Color                           { return hex2rgb("#C19A6B") }
func xanadu() r.Color                              { return hex2rgb("#738678") }
func xander() r.Color                              { return hex2rgb("#44500C") }
func xanthic() r.Color                             { return hex2rgb("#EEED09") }
func xanthous() r.Color                            { return hex2rgb("#F1B42F") }
func xboxGreen() r.Color                           { return hex2rgb("#0E7A0D") }
func xiaomiOrange() r.Color                        { return hex2rgb("#FD4900") }
func xumo() r.Color                                { return hex2rgb("#413639") }
func yaleBlue() r.Color                            { return hex2rgb("#00356B") }
func yellow() r.Color                              { return hex2rgb("#FFFF00") }
func yellowCrayola() r.Color                       { return hex2rgb("#FCE883") }
func yellowMunsell() r.Color                       { return hex2rgb("#EFCC00") }
func yellowNcs() r.Color                           { return hex2rgb("#FFD300") }
func yellowPantone() r.Color                       { return hex2rgb("#FEDF00") }
func yellowProcess() r.Color                       { return hex2rgb("#FFEF00") }
func yellowRyb() r.Color                           { return hex2rgb("#FEFE33") }
func yellowGreen() r.Color                         { return hex2rgb("#9ACD32") }
func yellowGreenCrayola() r.Color                  { return hex2rgb("#C5E384") }
func yellowGreenrWheel() r.Color                   { return hex2rgb("#30B21A") }
func yellowOrange() r.Color                        { return hex2rgb("#FFAE42") }
func yellowOrangerWheel() r.Color                  { return hex2rgb("#FF9505") }
func yellowRose() r.Color                          { return hex2rgb("#FFF000") }
func yellowSunshine() r.Color                      { return hex2rgb("#FFF700") }
func yinmnBlue() r.Color                           { return hex2rgb("#2E5090") }
func zafferZaffre() r.Color                        { return hex2rgb("#0014A8") }
func zarqa() r.Color                               { return hex2rgb("#FF4500") }
func zeal() r.Color                                { return hex2rgb("#91E0B7") }
func zebraWhite() r.Color                          { return hex2rgb("#F5F5F5") }
func zincGray() r.Color                            { return hex2rgb("#655B55") }
func zincWhite() r.Color                           { return hex2rgb("#FDF8FF") }
func zinnwalditeBrown() r.Color                    { return hex2rgb("#2C1608") }
func zinzolin() r.Color                            { return hex2rgb("#6C0277") }
func zirconGray() r.Color                          { return hex2rgb("#807473") }
func zomp() r.Color                                { return hex2rgb("#39A78E") }
func zydeco() r.Color                              { return hex2rgb("#20483F") }
