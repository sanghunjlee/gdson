package gdson

import "encoding/json"

type Gdson struct {
	Condition Condition `json:"condition"`
	Dialogue  Dialogue  `json:"dialogue"`
	Movement  Movement  `json:"movement"`
}

func (g Gdson) String() string {
	ret, err := json.Marshal(g)
	if err != nil {
		return err.Error()
	}
	return string(ret)
}
