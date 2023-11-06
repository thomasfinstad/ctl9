package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thomasfinstad/ctl9/internal/app"
)

func main() {
	kb := app.NewApplication()

	ebiten.SetWindowTitle("Contoller T9 Keyboard)")
	ebiten.SetTPS(30)
	ebiten.SetVsyncEnabled(true)

	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetRunnableOnUnfocused(true)

	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true
	op.SkipTaskbar = false
	op.InitUnfocused = true

	if err := ebiten.RunGameWithOptions(kb, op); err != nil {
		log.Fatal(err)
	}
}
