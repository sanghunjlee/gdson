package gdson

import (
	"errors"
	"fmt"
	"os"
)

const (
	VERSION string = "0.0.1"
	ISO8601 string = "2006-01-02T15:04:05Z07:00"
)

type App struct {
	AppFile *AppFile
	Gdson   *Gdson
}

func NewApp() *App {
	app := &App{
		AppFile: NewAppFile(),
		Gdson:   &Gdson{},
	}
	return app
}

func (a *App) Initialize(input string) {
	err := a.AppFile.Initialize(input)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println("The file already exists!")
			fmt.Println("Load the existing file with 'gdson load' instead.")
			return
		}
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("New dialogue file initialized as: %s\n", a.AppFile.Path)
}

func (a *App) LoadFile(input string) {
	err := a.AppFile.SetPath(input)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The file does not exist!")
			fmt.Println("Initialize the file with 'gdson init' instead.")
			return
		}
		fmt.Println(err.Error())
	}
}
