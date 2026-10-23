# multi-roblox-macos (Updated for Modern macOS & 27.2+)

A modern, high-performance Multi-Instance manager for Roblox on macOS.

This is an updated fork of [Insadem/multi-roblox-macos](https://github.com/Insadem/multi-roblox-macos) rewritten and hardened for modern macOS versions (macOS 14 Sonoma, macOS 15 Sequoia, and macOS 27.2+).

---

## 🚀 Key Improvements & Fixes

1. **macOS 27.2+ & AMFI / Gatekeeper Compatibility**:
   - Modern macOS enforces strict code signing and sealed resources via AMFI (Apple Mobile File Integrity). In older versions, modifying `Info.plist` broke the signature, resulting in `Launch failed (NSCocoaErrorDomain Code 257 / Permission denied)`.
   - This fork automatically strips quarantine extended attributes (`xattr -cr`) and re-signs cloned instances on-the-fly (`codesign --force --deep --sign -`), guaranteeing 100% launch success.

2. **Instant APFS Clone-on-Write**:
   - Uses APFS copy-on-write (`cp -cR`) instead of slow recursive disk copy. Launching new instances is virtually instantaneous (<0.3s) and takes zero extra disk space until modified.

3. **Universal Binary (Native Apple Silicon + Intel)**:
   - Built natively for Apple Silicon (M1/M2/M3/M4, `arm64`) and Intel (`x86_64`) as a universal binary.

4. **Dynamic Semaphore Watcher**:
   - Solves the startup race condition by actively pulsing and monitoring the `/RobloxPlayerUniq` POSIX semaphore during app launch and in the background.

5. **Proper LaunchServices & Deep Link Registration**:
   - `multiroblox.app` correctly declares `roblox-player` and `roblox` schemes in `CFBundleURLTypes` and registers with `lsregister`.
   - Launches new instances with `open -n` and restores the original Roblox player handler upon exit.

6. **Robust Process Management via `libproc`**:
   - Replaced fragile hardcoded `kinfo_proc` struct offsets with Darwin's native `libproc` API (`proc_listpids`, `proc_pidinfo`), ensuring seamless process tracking across any macOS kernel version.

---

## 🛠️ Usage

### Quick Start:
1. Download `multiroblox.app` from [Releases](https://github.com/rhythmcreative/multi-roblox-macos/releases).
2. If downloaded from the internet, remove quarantine attribute:
   ```bash
   xattr -cr /path/to/multiroblox.app
   ```
3. Open `multiroblox.app`.
4. Navigate to the Roblox website in your browser, log in to your first account, and click **Play**.
5. Switch accounts or open a different browser/profile and click **Play** on another game.
6. Both instances will run simultaneously without conflict.

---

## 🔨 Building from Source

### Requirements:
- macOS with Xcode Command Line Tools installed (`xcode-select --install`)
- Go 1.23 or newer

### Build:
Run the included build script:
```bash
./build.sh
```
This produces a Universal Binary (`arm64` + `x86_64`), installs it into `multiroblox.app`, and applies ad-hoc codesigning.

---

## 📄 License
MIT License.
