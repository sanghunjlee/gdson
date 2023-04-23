package gdson

import (
	"errors"
	"fmt"
	"os"
	"regexp"
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
		return
	}
	fmt.Printf("Dialogue file loaded: %s\n", a.AppFile.Path)
}

func (a *App) Add(args []string) {
	var typeFound = false
	var re *regexp.Regexp
	re, _ = regexp.Compile(`(?i)^condi(?:tion)s?$`)
	if re.MatchString(args[0]) {
		typeFound = true
		args[0] = "Condition"
	}

	re, _ = regexp.Compile(`(?i)^dialog(?:ue)s?$`)
	if re.MatchString(args[0]) {
		typeFound = true
		args[0] = "Dialogue"
	}

	re, _ = regexp.Compile(`(?i)move(?:ment)s?$`)
	if re.MatchString((args[0])) {
		typeFound = true
		args[0] = "Movement"
	}
	if !typeFound {
		fmt.Printf("The type entered ('%s') is not a valid type.\n", args[0])
		fmt.Println("To check what are valid types, check with 'gdson help add'.")
		return
	}

	err := a.Gdson.Add(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s entry added successfully!", args[0])
}

func (a *App) load() error {
	gd, err := a.AppFile.Load()
	if err != nil {
		return err
	}
	a.Gdson = gd
	return nil
}

func (a *App) save() error {
	err := a.AppFile.Save(a.Gdson)
	return err
}
