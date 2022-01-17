package views

import "embed"

//go:embed *
var files embed.FS

func GetFiles() embed.FS {
	return files
}