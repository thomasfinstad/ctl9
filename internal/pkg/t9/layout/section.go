package layout

import (
	"bytes"
	"image"
	"io"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/micmonay/keybd_event"
	"github.com/thomasfinstad/ctl9/internal/pkg/constraint"
)

/*
NewSection returns a new section based on the layer json config file in the provided layoyutDir

section config structure:

	 {
		"activator": <constraint config>,
		"image": {
			"file": <string with image file name to show when section is active>,
			"placement": [<int x value>,<int y value>]
		},
		"inputs": [
			<input config>,
			...
		]
	}
*/
func NewSection(kb *keybd_event.KeyBonding, controllerID ebiten.GamepadID, layoutDirPath fs.FS, sectionName string, sectionConfig map[string]any) *Section {
	log.Printf("Configuring section: %s", sectionName)
	section := &Section{
		Name: sectionName,
	}

	// Activator
	if sectionActivatorConfig, ok := sectionConfig["activator"].(map[string]any); ok {
		section.activator = &activator{Constraint: constraint.NewConstraint(controllerID, sectionActivatorConfig)}
	} else {
		section.activator = &activator{Constraint: nil}
	}

	// Image when active
	if imageConfig, ok := sectionConfig["image"].(map[string]any); ok {

		imageFile, err := layoutDirPath.Open(imageConfig["file"].(string))
		if err != nil {
			log.Fatal("unable to open section image file", imageConfig, err)
		}
		defer imageFile.Close()

		imageData, err := io.ReadAll(imageFile)
		if err != nil {
			log.Fatal(err)
		}
		Img, _, err := image.Decode(bytes.NewReader(imageData))
		if err != nil {
			log.Fatal(err)
		}
		section.image = &sectionImage{
			Image: ebiten.NewImageFromImage(Img),
			X:     int(imageConfig["placement"].([]any)[0].(float64)),
			Y:     int(imageConfig["placement"].([]any)[1].(float64)),
		}
	} else {
		section.image = nil
	}

	// Inputs
	for _, inputConfig := range sectionConfig["inputs"].([]any) {
		input := NewInput(kb, controllerID, inputConfig.(map[string]any))
		section.inputs = append(section.inputs, input)
	}

	return section
}

// Section holds information about a layer section containing inputs
type Section struct {
	Name      string
	activator *activator
	image     *sectionImage
	inputs    []*Input
}

type sectionImage struct {
	Image *ebiten.Image
	X     int
	Y     int
}

// Active checks if all constraints are satisfied and this section should be processed
func (section *Section) Active() (shouldBeActive bool) {
	active := section.activator.Active()
	if active {
		log.Printf("Active Section: %s", section.Name)
	}
	return active
}

// Process traverses all sub components and processes the first one that satisfies its constraints
func (section *Section) Process() (tookAction bool) {
	log.Printf("Processing Section: %s Inputs: %d", section.Name, len(section.inputs))
	for _, input := range section.inputs {
		if input.Active() {
			return input.Process()
		}
	}
	return false
}
