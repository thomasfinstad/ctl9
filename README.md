# Ctl9 (T9 keyboard for Controllers)

__This currently just a proof of concept, but some things are already working fairly well, layer support is lacking.__

This is inspired by the [PSP Homebrew AFKIM messanger client](https://www.gamebrew.org/wiki/AFKIM_Away_From_Keyboard_Instant_Messenger_PSP) keyboard.

End goal is to enable Steam Deck style devices with a better way of writing.

## Features

- Linux support (amd64)
- Windows support (amd64) (unverifed: i386, arm, arm64)
- Hide/Unhide overlay

## Roadmap

### Done

- [x] Controller button press detection
- [x] Controller button hold detection
- [x] Emulate keyboard keystrokes
- [x] Transparancy support
- [x] Click through support (overlay style)
- [x] Hide support (don't draw window)
- [x] Always on top (does not work with taskbar/start menu)
- [x] Configurable layout/layer/section support

### TODOs

- [ ] Better debug view
- [ ] Support multi key actions (eg: SHIFT+1 = !)
- [ ] Support button press indicatior in the same vain as section indicator
- [ ] Configuration file support
- [ ] Custom layouts
- [ ] Integrated window movement / placement
- [ ] Update available notification
- [ ] Packaging
    - [ ] Flatpak

### Investigate

- [ ] https://torvaney.github.io/projects/t9-optimised
- [ ] Mouse movement and click support
- [ ] Unicode / Emoji support
- [ ] Possibility of hijacking all controller inputs (windows bar/start search reacts to inputs)
- [ ] Fix Always on top with taskbar/start menu
