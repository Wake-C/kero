package main

import (
	"strings"
	"testing"
)

// An empty font-family draws with kero's Nerd Font, then CJK monospaced
// fallbacks, so Chinese in a TUI does not drop to a proportional face
// when the user has one of those families.
func TestTermFamilyDefault(t *testing.T) {
	a := testApp(t)
	if got := a.termFamily(); got != defaultTermFamily {
		t.Errorf("empty font-family %q, want %q", got, defaultTermFamily)
	}
	if !strings.HasPrefix(defaultTermFamily, bundledFont+",") {
		t.Errorf("default stack %q should start with %q", defaultTermFamily, bundledFont)
	}
	for _, name := range []string{"Sarasa Mono SC", "Maple Mono NF CN", "Noto Sans Mono CJK SC", "Noto Sans Mono CJK", "monospace"} {
		if !strings.Contains(defaultTermFamily, name) {
			t.Errorf("default stack %q lacks %q", defaultTermFamily, name)
		}
	}
	a.set.FontFamily = "Iosevka Term"
	if got := a.termFamily(); got != "Iosevka Term" {
		t.Errorf("a chosen family %q", got)
	}
	a.set.FontFamily = "monospace"
	if got := a.termFamily(); got != "monospace" {
		t.Errorf("System picked, terminal font %q", got)
	}
}
