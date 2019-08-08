package util

import (
	"flag"
)

// SelectEditor sets a flag for selecting an editor
func SelectEditor(editor *string) {
	flag.StringVar(editor, "editor", "mvim", "which editor would you like to open the file with")
}

// RenameEditor changes it to the executable file in $PATH
func RenameEditor(editor string) string {
	switch editor {
	case "atom":
		return "atom"
	case "macvim", "mvim":
		return "mvim"
	case "code", "code-insiders":
		return "code-insiders"
	case "vimr":
		return "vimr"
	case "mate", "textmate":
		return "mate"
	default:
		panic("Unknown editor:" + editor)
	}
}
