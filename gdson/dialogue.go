package gdson

type Option struct {
	Id   int    `json:"option_id"`
	Text string `json:"option_text"`
	Next int    `json:"option_next"`
}

type Dialogue struct {
	Id      int      `json:"id"`
	Text    []string `json:"text"`
	Options []Option `json:"options"`
	Next    int      `json:"next"`
}
