<h1 align="center">multi-roblox-macos</h1>

<div align="center">

<p><i>Multi-instance Roblox for modern macOS, fully compatible with macOS Sonoma, Sequoia and 27.2+ on Apple Silicon (M1–M4) & Intel.</i></p>

[![macOS](https://img.shields.io/badge/macOS-27.2%2B%20%7C%20Sequoia%20%7C%20Sonoma-000000?style=for-the-badge&logo=apple&logoColor=white)](https://apple.com)
[![Apple Silicon](https://img.shields.io/badge/Apple%20Silicon-M1--M4%20Native-grey?style=for-the-badge&logo=apple&logoColor=white)](https://apple.com)
[![Roblox](https://img.shields.io/badge/Roblox-Multi--Instance-E2231A?style=for-the-badge&logo=roblox&logoColor=white)](https://roblox.com)
[![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![Release](https://img.shields.io/badge/Release-v1.0.0-success?style=for-the-badge)](https://github.com/rhythmcreative/multi-roblox-macos/releases)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

</div>

___

[![Typing SVG](https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=23&pause=1000&color=F7F7F7&vCenter=true&width=435&height=30&lines=ABOUT)](https://git.io/typing-svg)

This is an updated, modernized fork of [`Insadem/multi-roblox-macos`](https://github.com/Insadem/multi-roblox-macos) rewritten specifically to support modern macOS releases, including **macOS 27.2**, macOS Sequoia, and macOS Sonoma.

It allows you to run multiple independent Roblox accounts and client windows on your Mac at the same time without them closing each other.

___

[![Typing SVG](https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=23&pause=1000&color=F7F7F7&vCenter=true&width=435&height=30&lines=FEATURES)](https://git.io/typing-svg)

- **macOS 27.2+ AMFI / Gatekeeper Fix**: Resolves the `Permission Denied (Code 257)` launch failure by automatically applying ad-hoc re-signing (`codesign --force --deep --sign -`) and stripping quarantine attributes.
- **Ultra-Fast APFS Cloning**: Instantaneous Copy-on-Write (`cp -cR`) launches instances in <0.3s without consuming duplicate disk space.
- **Universal Binary**: 100% native execution on Apple Silicon (`arm64`, M1/M2/M3/M4) as well as Intel (`x86_64`).
- **Dynamic Semaphore Watcher**: Eliminates launch race conditions with a background pulse monitor for `/RobloxPlayerUniq`.
- **Proper URL Scheme Registration**: Declares `roblox-player` and `roblox` schemes in `CFBundleURLTypes` and registers directly with LaunchServices.
- **Modern Process Management**: Powered by Darwin's native `libproc` API (`proc_listpids`, `proc_pidinfo`) instead of fragile static struct offsets.
- **Auto Clean-up**: Restores default Roblox handlers upon exit or cancellation (`SIGINT`/`SIGTERM`) and cleans up temporary clones.

___

[![Typing SVG](https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=23&pause=1000&color=F7F7F7&vCenter=true&width=435&height=30&lines=INSTALLATION)](https://git.io/typing-svg)

1. Download **`multiroblox-1.0.0-macos.zip`** from [Releases](https://github.com/rhythmcreative/multi-roblox-macos/releases/tag/v1.0.0).
2. Unzip and drag **`multiroblox.app`** to your `/Applications` folder (or run it from Downloads).
3. If macOS displays a security warning, open Terminal and run:
   ```bash
   xattr -cr /Applications/multiroblox.app
   ```

___

[![Typing SVG](https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=23&pause=1000&color=F7F7F7&vCenter=true&width=435&height=30&lines=USAGE)](https://git.io/typing-svg)

1. Launch **`multiroblox.app`** (a small control window will appear).
2. Go to [roblox.com](https://www.roblox.com) in your browser, log in to your first account, and press **Play** on any game.
3. Switch accounts or open a private/different browser profile and press **Play** on another game.
4. Both instances run side-by-side without interference!
5. When finished, close the app or click **Close all Roblox instances** to cleanly terminate all processes and restore standard handlers.

___

[![Typing SVG](https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=23&pause=1000&color=F7F7F7&vCenter=true&width=435&height=30&lines=BUILD+FROM+SOURCE)](https://git.io/typing-svg)

```bash
# Clone repository
git clone https://github.com/rhythmcreative/multi-roblox-macos.git
cd multi-roblox-macos

# Build Universal Binary & sign app bundle
./build.sh
```

___

<div align="center">

<p>Made with ❤️ from rhythmcreative.</p>

</div>
