package layout

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/micmonay/keybd_event"
	"github.com/thomasfinstad/ctl9/internal/pkg/constraint"
)

/*
NewInput returns a new input based on the layer json config file in the provided layoyutDir

input config structure:

	{"activator":<constraint config>,"action": <keyboard key code>}
*/
func NewInput(kb *keybd_event.KeyBonding, controllerID ebiten.GamepadID, inputConfig map[string]any) *Input {
	log.Printf("Configuring input: %#v", inputConfig)
	input := &Input{
		kb: kb,
	}

	// Activator
	if inputActivatorConfig, ok := inputConfig["activator"].(map[string]any); ok {
		input.activator = &activator{
			Constraint:        constraint.NewConstraint(controllerID, inputActivatorConfig),
			RequireConstraint: true,
		}
	} else {
		input.activator = &activator{
			Constraint:        nil,
			RequireConstraint: true,
		}
	}

	// Action
	input.action = inputConfig["action"].(string)
	if _, ok := micmonayKeybdEventMapping[input.action]; !ok {
		log.Fatalf("invalid key bound as input action: %s", input.action)
	}

	return input
}

// Input holds configuration for what keys to press and when
type Input struct {
	activator *activator
	action    string
	kb        *keybd_event.KeyBonding
}

// Active checks if all constraints are satisfied and this input should be processed
func (input *Input) Active() (shouldBeActive bool) {
	active := input.activator.Active()
	if active {
		log.Printf("Active Input: %#v", input)
	}
	return active
}

// Process executes the configured action
func (input *Input) Process() (tookAction bool) {
	log.Printf("Action taken: %s", input.action)

	// Select keys to be pressed
	input.kb.SetKeys(micmonayKeybdEventMapping[input.action])

	// Press the selected keys
	err := input.kb.Launching()
	if err != nil {
		log.Fatal("unable to click keyboard keys", err)
	}

	return true
}

/*
"autogenerated" mapping from string to constants.
It will only allow in key bindings that are valid for both of linux and windows.
Mac is not supportable because VK_BACKSPACE is not a thing for mac...

******** Commands to generate list, in created "mapping.tmp" file ********

curl -s "https://raw.githubusercontent.com/micmonay/keybd_event/master/keybd_linux.go" \
	| sed -n '/const (/,/)/s/^[[:space:]]*\(VK_[^[:space:]]*\).\+/"\1": keybd_event.\1,/p' \
	| sort > lin.tmp

curl -s "https://raw.githubusercontent.com/micmonay/keybd_event/master/keybd_windows.go" \
	| sed -n '/const (/,/)/s/^[[:space:]]*\(VK_[^[:space:]]*\).\+/"\1": keybd_event.\1,/p' \
	| sort > win.tmp

sort <(cat lin.tmp win.tmp | sort | uniq -d) \
	> mappings.tmp \
	&& rm lin.tmp win.tmp

******** If Mac was also supported these two commands would replace the last one above ********

curl -s "https://raw.githubusercontent.com/micmonay/keybd_event/master/keybd_darwin.go" \
	| sed -n '/const (/,/)/s/^[[:space:]]*\(VK_[^[:space:]]*\).\+/"\1": keybd_event.\1,/p' \
	| sort > mac.tmp

sort <(cat lin.tmp win.tmp | sort | uniq -d) \
	mac.tmp | uniq -d > mappings.tmp \
	&& rm lin.tmp win.tmp mac.tmp
*/

var micmonayKeybdEventMapping = map[string]int{
	"VK_0":          keybd_event.VK_0,
	"VK_1":          keybd_event.VK_1,
	"VK_2":          keybd_event.VK_2,
	"VK_3":          keybd_event.VK_3,
	"VK_4":          keybd_event.VK_4,
	"VK_5":          keybd_event.VK_5,
	"VK_6":          keybd_event.VK_6,
	"VK_7":          keybd_event.VK_7,
	"VK_8":          keybd_event.VK_8,
	"VK_9":          keybd_event.VK_9,
	"VK_A":          keybd_event.VK_A,
	"VK_APOSTROPHE": keybd_event.VK_APOSTROPHE,
	"VK_BACKSLASH":  keybd_event.VK_BACKSLASH,
	"VK_BACKSPACE":  keybd_event.VK_BACKSPACE,
	"VK_B":          keybd_event.VK_B,
	"VK_CAPSLOCK":   keybd_event.VK_CAPSLOCK,
	"VK_C":          keybd_event.VK_C,
	"VK_COMMA":      keybd_event.VK_COMMA,
	"VK_DELETE":     keybd_event.VK_DELETE,
	"VK_D":          keybd_event.VK_D,
	"VK_DOT":        keybd_event.VK_DOT,
	"VK_DOWN":       keybd_event.VK_DOWN,
	"VK_E":          keybd_event.VK_E,
	"VK_END":        keybd_event.VK_END,
	"VK_ENTER":      keybd_event.VK_ENTER,
	"VK_EQUAL":      keybd_event.VK_EQUAL,
	"VK_ESC":        keybd_event.VK_ESC,
	"VK_F10":        keybd_event.VK_F10,
	"VK_F11":        keybd_event.VK_F11,
	"VK_F12":        keybd_event.VK_F12,
	"VK_F1":         keybd_event.VK_F1,
	"VK_F2":         keybd_event.VK_F2,
	"VK_F3":         keybd_event.VK_F3,
	"VK_F4":         keybd_event.VK_F4,
	"VK_F5":         keybd_event.VK_F5,
	"VK_F6":         keybd_event.VK_F6,
	"VK_F7":         keybd_event.VK_F7,
	"VK_F8":         keybd_event.VK_F8,
	"VK_F9":         keybd_event.VK_F9,
	"VK_F":          keybd_event.VK_F,
	"VK_G":          keybd_event.VK_G,
	"VK_GRAVE":      keybd_event.VK_GRAVE,
	"VK_H":          keybd_event.VK_H,
	"VK_HOME":       keybd_event.VK_HOME,
	"VK_I":          keybd_event.VK_I,
	"VK_INSERT":     keybd_event.VK_INSERT,
	"VK_J":          keybd_event.VK_J,
	"VK_K":          keybd_event.VK_K,
	"VK_KP0":        keybd_event.VK_KP0,
	"VK_KP1":        keybd_event.VK_KP1,
	"VK_KP2":        keybd_event.VK_KP2,
	"VK_KP3":        keybd_event.VK_KP3,
	"VK_KP4":        keybd_event.VK_KP4,
	"VK_KP5":        keybd_event.VK_KP5,
	"VK_KP6":        keybd_event.VK_KP6,
	"VK_KP7":        keybd_event.VK_KP7,
	"VK_KP8":        keybd_event.VK_KP8,
	"VK_KP9":        keybd_event.VK_KP9,
	"VK_KPASTERISK": keybd_event.VK_KPASTERISK,
	"VK_KPDOT":      keybd_event.VK_KPDOT,
	"VK_KPMINUS":    keybd_event.VK_KPMINUS,
	"VK_KPPLUS":     keybd_event.VK_KPPLUS,
	"VK_LEFTBRACE":  keybd_event.VK_LEFTBRACE,
	"VK_LEFT":       keybd_event.VK_LEFT,
	"VK_L":          keybd_event.VK_L,
	"VK_MINUS":      keybd_event.VK_MINUS,
	"VK_M":          keybd_event.VK_M,
	"VK_N":          keybd_event.VK_N,
	"VK_NUMLOCK":    keybd_event.VK_NUMLOCK,
	"VK_O":          keybd_event.VK_O,
	"VK_PAGEDOWN":   keybd_event.VK_PAGEDOWN,
	"VK_PAGEUP":     keybd_event.VK_PAGEUP,
	"VK_PAUSE":      keybd_event.VK_PAUSE,
	"VK_P":          keybd_event.VK_P,
	"VK_Q":          keybd_event.VK_Q,
	"VK_RESERVED":   keybd_event.VK_RESERVED,
	"VK_RIGHTBRACE": keybd_event.VK_RIGHTBRACE,
	"VK_RIGHT":      keybd_event.VK_RIGHT,
	"VK_R":          keybd_event.VK_R,
	"VK_SCROLLLOCK": keybd_event.VK_SCROLLLOCK,
	"VK_SEMICOLON":  keybd_event.VK_SEMICOLON,
	"VK_S":          keybd_event.VK_S,
	"VK_SLASH":      keybd_event.VK_SLASH,
	"VK_SP10":       keybd_event.VK_SP10,
	"VK_SP11":       keybd_event.VK_SP11,
	"VK_SP12":       keybd_event.VK_SP12,
	"VK_SP1":        keybd_event.VK_SP1,
	"VK_SP2":        keybd_event.VK_SP2,
	"VK_SP3":        keybd_event.VK_SP3,
	"VK_SP4":        keybd_event.VK_SP4,
	"VK_SP5":        keybd_event.VK_SP5,
	"VK_SP6":        keybd_event.VK_SP6,
	"VK_SP7":        keybd_event.VK_SP7,
	"VK_SP8":        keybd_event.VK_SP8,
	"VK_SP9":        keybd_event.VK_SP9,
	"VK_SPACE":      keybd_event.VK_SPACE,
	"VK_TAB":        keybd_event.VK_TAB,
	"VK_T":          keybd_event.VK_T,
	"VK_U":          keybd_event.VK_U,
	"VK_UP":         keybd_event.VK_UP,
	"VK_V":          keybd_event.VK_V,
	"VK_W":          keybd_event.VK_W,
	"VK_X":          keybd_event.VK_X,
	"VK_Y":          keybd_event.VK_Y,
	"VK_Z":          keybd_event.VK_Z,
}
