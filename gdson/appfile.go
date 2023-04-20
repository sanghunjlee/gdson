package gdson

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type AppFile struct {
	Found bool
	Path  string
}

const OutputDir string = "Output"
const DefaultName string = "Dialogue"

func NewAppFile() *AppFile {
	return &AppFile{
		Found: false,
		Path:  "",
	}
}

func (f *AppFile) Initialize(input string) error {
	outputDirExist, err := IsExist(OutputDir)
	if err != nil {
		return err
	}
	if err == nil && !outputDirExist {
		os.Mkdir(OutputDir, 0777)
	}

	var fileName string
	if fileName = input; fileName == "" {
		entries, err := os.ReadDir(OutputDir)
		if err != nil {
			return err
		}

		var fileId int
		var maxId int = 0
		var ids []int
		for _, e := range entries {
			if strings.HasPrefix(e.Name(), DefaultName) {
				var extIndex = strings.LastIndex(e.Name(), ".")
				var id, err = strconv.Atoi(e.Name()[len(DefaultName):extIndex])
				if err != nil {
					return err
				}
				if maxId < id {
					maxId = id
				}
				ids = append(ids, id)
			}
		}
		var found bool
		fileId = maxId + 1
		for i := 0; i < maxId; i++ {
			found = false
			for _, id := range ids {
				if id == i {
					found = true
					break
				}
			}
			if !found {
				fileId = i
				break
			}
		}

		fileName = DefaultName + strconv.Itoa(fileId)
	}
	filePath := filepath.Join(OutputDir, fileName+".json")
	newFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	newFile.Write([]byte("[]"))
	if err := newFile.Close(); err != nil {
		return err
	}

	f.Path = filePath
	return nil
}

func (f *AppFile) Load() (*Gdson, error) {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	var gdson *Gdson
	jsonError := json.Unmarshal(data, &gdson)
	if jsonError != nil {
		return nil, jsonError
	}
	f.Found = true
	return gdson, nil
}

func (f *AppFile) Save(gdson *Gdson) error {
	data, _ := json.MarshalIndent(gdson, "", "  ")
	if err := os.WriteFile(f.Path, []byte(data), 0644); err != nil {
		return err
	}
	return nil
}

func (f *AppFile) SetPath(input string) error {
	filePath := filepath.Join(OutputDir, input+".json")
	if isExist, err := IsExist(filePath); err != nil {
		return err
	} else if !isExist {
		return os.ErrNotExist
	}
	f.Path = filePath
	return nil
}
