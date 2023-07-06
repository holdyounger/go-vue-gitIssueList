package service

import (
	"context"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx  context.Context
	menu menu.Menu
	// 其他字段...
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func openFile(_ *menu.CallbackData) {

}

func (a *App) NewMenu() *menu.Menu {
	a.menu = *menu.NewMenu()
	FileMenu := a.menu.AddSubmenu("File")
	a.menu.AddText("Index", keys.CmdOrCtrl("i"), openFile)
	a.menu.AddText("About", keys.CmdOrCtrl("b"), openFile)

	FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(a.ctx)
	})

	return &a.menu
}

func GetContext(a *App) *context.Context {
	return &a.ctx
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	a.NewMenu()
}

func (a *App) Shutdown(ctx context.Context) {
	fmt.Println("Good Bye,Ming")
}

func (b *App) BeforeClose(ctx context.Context) (prevent bool) {
	dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Quit?",
		Message: "Are you sure you want to quit?",
	})

	if err != nil {
		return false
	}
	return dialog != "Yes"
}

func (a *App) SetContext(ctx context.Context) {
	// WindowSetTitle(ctx, "GitIssue")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Greet returns a greeting for the given name
func (a *App) PrintGreet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Close() {
	// runtime.Quit(a.ctx)
}

// SelectFile 选择需要处理的文件
func (a *App) SelectFile(title string, filetype string) string {
	// runtime.LogPrint(a.ctx, fmt.Sprintf("FileType: %s", filetype))
	log.Info(fmt.Sprintf("FileType: %s", filetype))

	if title == "" {
		title = "选择文件"
	}
	if filetype == "" {
		filetype = "*.log;*.txt;*.rdb;*.aof"
	}
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: title,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "文本数据",
				Pattern:     filetype,
			},
		},
	})
	if err != nil {
		return fmt.Sprintf("err %s!", err)
	}
	return selection
}

func LogPrint(context context.Context, s string) {
	panic("unimplemented")
}
