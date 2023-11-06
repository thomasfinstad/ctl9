package t9

import (
	"io/fs"
	"log"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/micmonay/keybd_event"
	"github.com/thomasfinstad/ctl9/internal/pkg/t9/layout"
)

// NewT9 creates a new t9 keyboard and binds it to the controller and assets provided
func NewT9(controllerID ebiten.GamepadID, assetsDir fs.FS) *T9 {
	t := &T9{}

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Fatal("unable to create keybonding", err)
	}
	t.kb = &kb

	// For linux, it is very important to wait 2 seconds
	// ^ that was from the library README, does not seem to be true for me, but better safe than sorry
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	layoutsPath, err := fs.Sub(assetsDir, "layouts")
	if err != nil {
		log.Fatal("unable to open layouts dir", err)
	}

	alphaLayoutPath, err := fs.Sub(layoutsPath, "alpha-no")
	if err != nil {
		log.Fatal("unable to open alpha-no dir", err)
	}
	t.layout = layout.NewLayout(t.kb, controllerID, alphaLayoutPath)

	return t
}

// T9 is a virtual keyboard application
type T9 struct {
	kb     *keybd_event.KeyBonding
	layout *layout.Layout
}

// Process traverses all sub components and processes their logic
func (t *T9) Process() {
	if t.layout == nil {
		return
	}

	if !t.layout.Process() {
		log.Printf("T9 no action taken")
	}
}

// Hidden returns if the t9 screen show be hidden (or visible if not)
func (t *T9) Hidden() bool {
	return t.layout.Hidden
}

// Display returns the T9 virtual screen
func (t *T9) Display() *ebiten.Image {
	if t.layout == nil {
		return nil
	}
	return t.layout.View()
}
