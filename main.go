package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	println("Requesting file:", requestedFilename)
	fileData, err := os.ReadFile(requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}

	res.Write(fileData)
}

func openFile(_ *menu.CallbackData) {

}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	AppMenu.AddText("Index", keys.CmdOrCtrl("i"), openFile)
	AppMenu.AddText("About", keys.CmdOrCtrl("b"), openFile)

	FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "go-vue-issuelist",
		Width:  960,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		Menu:      AppMenu,
		Frameless: true,
		// CSSDragProperty: "widows",
		// CSSDragValue:    "1",
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.SetContext(ctx)
			// otherStruct.SetContext(ctx)
		},
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
			// otherStruct
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
