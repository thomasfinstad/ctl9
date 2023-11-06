package app

import (
	"fmt"
	"image/color"
	"log"

	"github.com/gookit/goutil/arrutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/thomasfinstad/ctl9/internal/pkg/t9"
)

// NewApplication returns the main app to use with ebiten
func NewApplication() *Application {
	return &Application{
		debug:  false,
		hidden: false,

		width:  1920,
		height: 1080,
	}
}

// Application is an ebiten compliant t9 app
type Application struct {
	// Meta
	debug  bool
	hidden bool

	// Display
	width  int
	height int

	// Inputs
	controllerIDActive *ebiten.GamepadID
	controllerIDList   []ebiten.GamepadID

	// T9 / Main application logic
	t9 *t9.T9
}

// Layout return the app sizes
func (app *Application) Layout(outsideWidth, outsideHeight int) (int, int) {
	return app.width, app.height
}

// Update processes the application logic
func (app *Application) Update() error {
	log.Printf("--- TPS [%0.3f] ---", ebiten.ActualTPS())
	controllerConnectionStatusHandler(app)

	if app.t9 == nil {
		return nil
	}

	if app.t9.Display() != nil {
		app.width = app.t9.Display().Bounds().Dx()
		app.height = app.t9.Display().Bounds().Dy()
	}
	ebiten.SetWindowSize(app.width, app.height)

	app.t9.Process()

	app.hidden = app.t9.Hidden()

	// check if debug should be enabled
	debugTrigger := ebiten.KeyControlRight
	keys := inpututil.AppendJustPressedKeys([]ebiten.Key{})
	if arrutil.In(debugTrigger, keys) {
		log.Printf("%s press detected, setting debug mode to: %t", debugTrigger, !app.debug)
		app.debug = !app.debug
	}

	return nil
}

// Draw draws the t9 application if not hidden
func (app *Application) Draw(screen *ebiten.Image) {
	if app == nil {
		log.Printf("No app has been configured")
		return
	}

	if app.hidden {
		log.Printf("App visible: %t", app.hidden)
		return
	}

	/*
		T9
	*/
	if app.t9 != nil {
		drawImgOpt := &ebiten.DrawImageOptions{}
		if app.t9.Display() != nil {
			screen.DrawImage(app.t9.Display(), drawImgOpt)
		}
	}

	/*
		Debug Text
	*/
	if app.debug {
		debugMsg := fmt.Sprintf("--- TPS [%7.3f] --- FPS [%7.3f] ---\n", ebiten.ActualTPS(), ebiten.ActualFPS())

		// Controllers
		if len(app.controllerIDList) > 0 {
			for _, controllerID := range app.controllerIDList {
				if ebiten.IsStandardGamepadLayoutAvailable(controllerID) {
					debugMsg += normalControllerInfoStr(controllerID)
				} else {
					debugMsg += "WARNING: non-standard layout detected, issues in showing info may occur!\n"
					debugMsg += normalControllerInfoStr(controllerID)
				}
				debugMsg += "\n"
			}
		} else {
			debugMsg = "Please connect your gamepad.\n"
		}

		// Keyboard + TouchID
		debugMsg += keyboardInfoStr()

		// Add semitransparant background for easier text readability
		vector.DrawFilledRect(
			screen,
			0.0,
			0.0,
			float32(screen.Bounds().Dx()),
			float32(screen.Bounds().Dy()),
			color.RGBA{R: 50, G: 50, B: 50, A: 150},
			false,
		)

		// Print text to screen, needs to be the last Draw action to be on top
		ebitenutil.DebugPrint(screen, debugMsg)
	}
}
