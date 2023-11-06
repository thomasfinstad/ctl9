package layout

import (
	"bytes"
	"encoding/json"
	"image"
	"io"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/micmonay/keybd_event"
	"github.com/thomasfinstad/ctl9/internal/pkg/constraint"
)

/*
NewLayer returns a new layer based on the layer json config file in the provided layoyutDir

layer config structure:

	{
	    "activator": <constraint config / null for default layer>,
	    "background": <string with image file name>,
	    "sections" :{
	        <string section name>: <section config>,
			...
		}
	}
*/
func NewLayer(kb *keybd_event.KeyBonding, controllerID ebiten.GamepadID, layoutDir fs.FS, layerData map[string]any) *Layer {
	layerName := layerData["name"].(string)
	log.Printf("Configuring layer: %s", layerName)
	layer := &Layer{
		Name: layerName,
	}

	/* Main layer config */
	layerConfigFile, err := layoutDir.Open(layerData["config"].(string))
	if err != nil {
		log.Fatal("unable to open config file", layerData, err)
	}
	defer layerConfigFile.Close()

	layerConfigJSON, err := io.ReadAll(layerConfigFile)
	if err != nil {
		log.Fatal("unable to read config file", err)
	}

	layerConfig := make(map[string]any)
	err = json.Unmarshal(layerConfigJSON, &layerConfig)
	if err != nil {
		log.Fatal("Error when loading layer config", err)
	}

	/* Layer dedicated config */

	// Background
	if bgName, ok := layerConfig["background"]; ok {
		bfFile, err := layoutDir.Open(bgName.(string))
		if err != nil {
			log.Fatal("unable to open background file", bgName, err)
		}
		defer bfFile.Close()

		bgData, err := io.ReadAll(bfFile)
		if err != nil {
			log.Fatal("unable to read background file", err)
		}
		bgImg, _, err := image.Decode(bytes.NewReader(bgData))
		if err != nil {
			log.Fatal(err)
		}
		layer.background = ebiten.NewImageFromImage(bgImg)
	} else {
		layer.background = nil
	}

	// Activator
	if layerActivatorData, ok := layerConfig["activator"].(map[string]any); ok {
		layer.activator = &activator{Constraint: constraint.NewConstraint(controllerID, layerActivatorData)}
	} else {
		layer.activator = &activator{Constraint: nil}
	}

	// Layer Sections (t9)
	for sectionName, sectionConfig := range layerConfig["sections"].(map[string]any) {
		section := NewSection(kb, controllerID, layoutDir, sectionName, sectionConfig.(map[string]any))
		layer.sections = append(layer.sections, section)
	}

	return layer
}

// Layer holds the sections
type Layer struct {
	Name       string
	activator  *activator
	background *ebiten.Image
	view       *ebiten.Image
	sections   []*Section
}

// Active checks if all constraints are satisfied and this layer should be processed
func (layer *Layer) Active() (shouldBeActive bool) {
	active := layer.activator.Active()
	if active {
		log.Printf("Active Layer: %s", layer.Name)
	}
	return active
}

// Process traverses all sub components and processes the first one that satisfies its constraints
func (layer *Layer) Process() (tookAction bool) {
	log.Printf("Processing Layer: %s", layer.Name)

	if layer.background != nil {
		layer.view = ebiten.NewImageFromImage(layer.background)
	}

	// Sections
	for _, section := range layer.sections {
		if section.Active() {

			// Activate section image
			if section.image != nil {
				imgOpts := &ebiten.DrawImageOptions{}
				imgOpts.GeoM.Translate(float64(section.image.X), float64(section.image.Y))

				layer.view.DrawImage(section.image.Image, imgOpts)
			}

			// Buttons
			return section.Process()
		}
	}
	return false
}
