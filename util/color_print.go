package util

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type rcp struct {
}

func randomColor() *color.Color {
	fgColors := []color.Attribute{
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
		color.FgWhite,
	}

	_ = []color.Attribute{
		color.BgRed,
		color.BgGreen,
		color.BgYellow,
		color.BgBlue,
		color.BgMagenta,
		color.BgCyan,
		color.BgWhite,
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	fgIndex := rand.Intn(len(fgColors))

	return color.New(fgColors[fgIndex])
}

func (*rcp) RandomColorPrint(content string) {
	random := randomColor()
	random.Print(content)
}

func (*rcp) RandomColorPrintln(content string) {
	random := randomColor()
	random.Println(content)
}
