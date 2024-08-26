package story

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

type Story map[string]Chapter

type StoryHandler struct {
	s Story
}

func NewStoryHandler(s Story) *StoryHandler {
	return &StoryHandler{s}
}
func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "story.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	err := tmpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}

}
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
