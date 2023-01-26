# Terminal Emulators
*updated 01/26/2023*

The following is a (far from exhaustive) list of some popular terminal emulators you might want to consider when doing work on the command line.

## Windows
* [Windows Terminal](https://apps.microsoft.com/store/detail/windows-terminal/9N0DX20HK701?hl=en-us&gl=us&SilentAuth=1&rtc=1) The official MS Terminal app is actually pretty good! Definitely the best place to start if you're unfamiliar.
* [Cmder](https://cmder.app/) - A well-designed app by a terminal enthusiast. Fully featured and portable with Linux commands integrated into the bundled "Git For Windows", which also comes with Git, so there's that!
* [MobaXterm](https://mobaxterm.mobatek.net/) - If you're looking for batteries-included, this is it, but it might be a little much for most users.

You can also enable [WSL](https://learn.microsoft.com/en-us/windows/wsl/install) if you're feeling a bit more adventurous. This will enable to to run a Linux environment alongside Windows.
## macOS
* [iTerm2](https://iterm2.com/) - The classic is still the best. You can't go wrong with iTerm2. It's what I use every day.
* [Alacrittu](https://github.com/alacritty/alacritty) - Very fast, if that matters to you, thanks to its OpenGL rendering engine. Configuration is done with text files though, and you'll have to locate them on disk, so it's not nearly as easy to get set up as iTerm2. Works on Linux, too.
* [Warp](https://www.warp.dev/) - The new kid on the block, but this is definitely some next-level stuff with advanced views, graphical integration, multi cursors, and the lot. Just make sure you read the [privacy statements](https://www.warp.dev/privacy) and [opt-out of telemetry](https://docs.warp.dev/getting-started/privacy#how-to-disable-telemetry-and-crash-reporting) if you're concerned about it.

## Linux
Let's be honest, you probably know what you're doing when it comes to terminal emulators, but here are some options anyways.
* [Kitty](https://github.com/kovidgoyal/kitty/) - Like Alacritty, it's a hardware-accelerated terminal. Also like Alacritty, you will need to configure it with text, but at least the config file is well-documented and easy to access.
* Alacritty also works on Linux
* [Foot](https://codeberg.org/dnkl/foot) - Fast and featureful with commandline graphics, thanks to libsixel. Also application integration for things like loading links in the browser for you, if you want.
