package story

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

func DecodeToJson(file io.Reader) (Story, error) {
	var story Story
	r := json.NewDecoder(file)
	err := r.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, err
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
