package main

import (
	_ "embed"
	"go-vue-gitIssueList/backend/config"
	"go-vue-gitIssueList/backend/image"
	"go-vue-gitIssueList/backend/stat"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
	app := NewApp()

	cfg := config.NewConfig()
	st := stat.NewStat()
	fm := image.NewFileManager(cfg, st)

	err := wails.Run(&options.App{
		Width:            1200,
		Height:           742,
		Title:            "Optimus",
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			app,
			cfg,
			st,
			fm,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
