package layout

import (
	"encoding/json"
	"io"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/micmonay/keybd_event"
	"github.com/thomasfinstad/ctl9/internal/pkg/constraint"
)

/*
NewLayout returns a new Layout based on the main.json file in the provided layoyutDir

At least one layer is required per layout, multiple are supported,
and recommended to enable more writing options.

Main layout config (main.json) structure:

	{"name":"<layout name>", "layers":[{"name":"<layer name>","config":"<layer config json file name>"}, ...]}

Layer dedicated config structure:...
*/
func NewLayout(kb *keybd_event.KeyBonding, controllerID ebiten.GamepadID, layoutDir fs.FS) *Layout {
	mainConfigFile, err := layoutDir.Open("main.json")
	if err != nil {
		log.Fatal("unable to open file: ", err)
	}
	defer mainConfigFile.Close()

	mainConfigJSON, err := io.ReadAll(mainConfigFile)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	layoutConfigData := make(map[string]any)
	err = json.Unmarshal(mainConfigJSON, &layoutConfigData)
	if err != nil {
		log.Fatal("Error when loading main.json", err)
	}

	layoutName := layoutConfigData["name"].(string)
	log.Printf("Configuring layout: %s", layoutName)
	layout := &Layout{
		Name: layoutName,
	}

	// ToggleHidden
	if layerActivatorData, ok := layoutConfigData["toggle-hidden"].(map[string]any); ok {
		layout.toggleHidden = &activator{
			RequireConstraint: true,
			Constraint:        constraint.NewConstraint(controllerID, layerActivatorData),
		}
		log.Printf("Visibility toggle: %#v", layout.toggleHidden)
	}

	// Layers
	for _, layerData := range layoutConfigData["layers"].([]any) {
		layer := NewLayer(kb, controllerID, layoutDir, layerData.(map[string]any))
		layout.layers = append(layout.layers, layer)
	}

	return layout
}

// Layout hold the t9 parts
type Layout struct {
	Name         string
	view         *ebiten.Image
	toggleHidden *activator
	Hidden       bool
	layers       []*Layer
}

// Process traverses all sub components and processes the first one that satisfies its constraints
func (layout *Layout) Process() (tookAction bool) {
	log.Printf("Processing Layout: %s", layout.Name)

	if layout == nil {
		return false
	}

	if layout.toggleHidden != nil {
		toggleHidden := layout.toggleHidden.Active()
		if toggleHidden {
			log.Printf("Setting visibility: %t", !layout.Hidden)
			layout.Hidden = !layout.Hidden
		}
	}

	for _, layer := range layout.layers {
		if layer.Active() {
			layout.view = layer.view
			return layer.Process()
		}
	}
	return false
}

// View returns the layout virtual screen
func (layout *Layout) View() *ebiten.Image {
	return layout.view
}
