package models

type Entry struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	CodeBlock string `json:"codeblock"`
	Notes     string `json:"notes"`
	Tags      []Tag  `json:"tags"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
