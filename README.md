# Ctl9 (T9 keyboard for Controllers)

<!--TOC-->

- [Ctl9 (T9 keyboard for Controllers)](#ctl9-t9-keyboard-for-controllers)
  - [How-to](#how-to)
  - [Features](#features)
  - [Roadmap](#roadmap)
    - [Done](#done)
    - [TODOs](#todos)
    - [Investigate](#investigate)
  - [FAQ](#faq)
    - [Linux input issues](#linux-input-issues)
    - [Show debug information](#show-debug-information)

<!--TOC-->

__This currently just a proof of concept, but some things are already working fairly well, layer support is lacking.__

This is inspired by the [PSP Homebrew AFKIM messanger client](https://www.gamebrew.org/wiki/AFKIM_Away_From_Keyboard_Instant_Messenger_PSP) keyboard.

End goal is to enable Steam Deck style devices with a better way of writing.

## How-to

1. Download binary from github releases.
2. Start application.
    - If you have issues getting inputs on linux there might be a uinput premissin issue, see [FAQ](#linux-input-issues) at bottom
3. Enter any text input field
    - Start menus / task bars / other apps that reacts to controller input by them selvs are difficult to work with
4. Show/hide interface with key combo depending on your controller type:
    - select+start (psx)
    - share+options (ps3)
    - back+start (xbox)
    - similar on other standardized controllers
5. Hold the left joystick towards the direction of the button you wish to press, and use the action keys to press a button / letter


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

## FAQ

### Linux input issues

The underlying library that emulates a keyboard is [keybd_event](https://github.com/micmonay/keybd_event) which has the following information on how to give access to enable inputs:


On Linux this library uses **uinput**, which on the major distributions requires root permissions.

The easy solution is executing through root user (by using `sudo`). A worse way is by changing the executable's permissions by using `chmod`.

__Secure Linux Example__
```bash
sudo groupadd uinput
sudo usermod -a -G uinput my_username
sudo udevadm control --reload-rules
echo "SUBSYSTEM==\"misc\", KERNEL==\"uinput\", GROUP=\"uinput\", MODE=\"0660\"" | sudo tee /etc/udev/rules.d/uinput.rules
echo uinput | sudo tee /etc/modules-load.d/uinput.conf
```

Another subtlety on Linux: it is important after creating the `keybd_event` to **wait 2 seconds before running first keyboard actions**.

### Show debug information

If you activate the application window (click it in the taskbar) and press the Right Control key on your keyboard you will get debug information printed on the application. Click Right Control again to toggle it off.

_Currently the screen does not properly resize when needed causing overflow text to get cut off._
