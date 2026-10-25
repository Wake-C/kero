package main

import (
	"os"
	"path/filepath"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

// kero carries JetBrains Mono, with Nerd Fonts' icons (resources/fonts,
// SIL Open Font License, see OFL.txt there): the terminals' font unless
// the user picks another, and the icons of prompts and of agents' TUIs.

// bundledFont is the family the fonts register as.
const bundledFont = "JetBrainsMono Nerd Font"

// defaultTermFamily is the empty font-family: kero's Nerd Font for Latin
// and icons, then CJK monospaced families people install for terminals,
// then the system's. A name that is not there is skipped, so Latin stays
// JetBrains.
const defaultTermFamily = bundledFont + ", Sarasa Mono SC, Maple Mono NF CN, Noto Sans Mono CJK SC, Noto Sans Mono CJK, monospace"

// fontsLoaded tells that the fonts registered.
var fontsLoaded bool

// loadFonts registers the fonts kero carries: from the app's resources,
// where `mygo build` and `mygo dev` put them, else from the source tree,
// as under `go run`.
func loadFonts() {
	var dirs []string
	if res, err := mygo.App.Path(mygo.PathResources); err == nil {
		dirs = append(dirs, filepath.Join(res, "fonts"))
	}
	dirs = append(dirs, filepath.Join("resources", "fonts"))
	for _, dir := range dirs {
		files, _ := filepath.Glob(filepath.Join(dir, "*.ttf"))
		for _, f := range files {
			data, err := os.ReadFile(f)
			if err != nil {
				continue
			}
			if ui.RegisterFont(data, bundledFont) == nil {
				fontsLoaded = true
			}
		}
		if fontsLoaded {
			return
		}
	}
}

// termFamily returns the family terminals draw with: the user's, else
// kero's Nerd Font and CJK monospaced fallbacks. A name that is not
// installed is skipped.
func (a *App) termFamily() string {
	if f := a.set.FontFamily; f != "" {
		return f
	}
	return defaultTermFamily
}
