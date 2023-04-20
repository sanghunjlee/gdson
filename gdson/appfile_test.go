package gdson

import (
	"os"
	"testing"
)

func TestInitSaveLoad(t *testing.T) {
	var isFailed = false
	app := &App{AppFile: NewAppFile(), Gdson: &Gdson{}}
	app.Initialize("test")

	test_condition := &Condition{
		Weekday: []Weekday{Monday},
	}

	test_dialogue := &Dialogue{
		Id:   0,
		Text: []string{"Hello", "World"},
		Next: 0,
	}

	app.Gdson = &Gdson{
		Condition: *test_condition,
		Dialogue:  *test_dialogue,
	}

	err := app.AppFile.Save(app.Gdson)
	if err != nil {
		isFailed = true
		t.Errorf(`Saving the app resulted in an error: %s`, err)
	}

	ex_gdson := app.Gdson
	app.Gdson = &Gdson{}
	loaded_gdson, err := app.AppFile.Load()
	if err != nil {
		isFailed = true
		t.Errorf(`Loading the app resulted in an error: %s`, err)
	}

	if loaded_gdson.String() != ex_gdson.String() {
		isFailed = true
		t.Errorf(`%s is not equal to %s`, loaded_gdson, ex_gdson)
	}

	if !isFailed {
		_ = os.RemoveAll("Output")
	}
}

func TestSetPath(t *testing.T) {
	app := &App{AppFile: NewAppFile(), Gdson: &Gdson{}}
	if err := app.AppFile.SetPath("test"); err == nil {
		t.Error(`Expected error during SetPath("test") but yielded nil error`)
	}
	if err := app.AppFile.Initialize("test"); err != nil {
		t.Errorf(`Initialize yielded an error: %s`, err)
	}
	if err := app.AppFile.SetPath("test"); err != nil {
		t.Errorf(`SetPath yielded an error: %s`, err)
	} else {
		t.Logf(`Test is successful: app.AppFile.Path = %s`, app.AppFile.Path)
		_ = os.RemoveAll("Output")
	}
}
