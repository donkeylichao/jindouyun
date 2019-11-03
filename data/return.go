package data

type OptionData struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}

type OptionsData struct {
	Options []OptionData `json:"options"`
}
