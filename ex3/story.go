package story

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strings"
)

type Story map[string]Chapter

type StoryHandler struct {
	s    Story
	tmpl *template.Template
}
type HandlerOptions func(h *StoryHandler)

func WithTemplate(t *template.Template) HandlerOptions {
	return func(h *StoryHandler) {
		h.tmpl = t
	}
}
func NewStoryHandler(s Story, opts ...HandlerOptions) *StoryHandler {
	t := template.Must(template.ParseFiles("../templates/story.html"))
	h := StoryHandler{s, t}
	for _, opt := range opts {
		opt(&h)
	}
	return &h
}
func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := h.tmpl
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	if chapter, ok := h.s[path]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			panic(err)
		}
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
