package gdson

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type GdType interface{}

type Gdson struct {
	Condition GdType `json:"condition"`
	Dialogue  GdType `json:"dialogue"`
	Movement  GdType `json:"movement"`
}

func (g *Gdson) Add(args []string) error {
	var errArgs []string
	var re *regexp.Regexp
	keyValue := make(map[string]string)
	for _, keyValuePair := range args[0:] {
		re = regexp.MustCompile(`^(\w+)\=(\w+)$`)
		if found := re.FindStringSubmatch(keyValuePair); found != nil {
			keyValue[found[1]] = found[2]
		} else {
			errArgs = append(errArgs, keyValuePair)
		}
	}
	if len(errArgs) == len(args) {
		return errors.New("no parsable args")
	}

	var gdStruct GdType
	switch args[0] {
	case "Condition":
		gdStruct = &g.Condition
	case "Dialogue":
		gdStruct = &g.Dialogue
	case "Movement":
		gdStruct = &g.Movement
	}
	fields := getStructFields(gdStruct)
	var successCount int
	var fieldFound bool
	for _, f := range fields {
		fieldFound = false
		for k, v := range keyValue {
			if strings.EqualFold(f.Name, k) {
				fmt.Println(f.Type.String())
				fieldFound = true
				err := setStructFieldByName(gdStruct, f.Name, v)
				if err != nil {
					errArgs = append(errArgs, k+"="+v)
				} else {
					successCount += 1
				}
			}
			if !fieldFound {
				errArgs = append(errArgs, k+"="+v)
			}
		}
	}

	if len(errArgs) > 0 {
		fmt.Printf(`Unable to parse the following args:\n\t%s\n`, strings.Join(errArgs, ","))
	}
	if successCount == 0 {
		return errors.New("nothing was added")
	}
	return nil
}

func (g Gdson) String() string {
	ret, err := json.Marshal(g)
	if err != nil {
		return err.Error()
	}
	return string(ret)
}
