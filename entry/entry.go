package entry

type Entry struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	CodeBlock  string `json:"codeblock"`
	Notes      string `json:"notes"`
	Categories string `json:"categories"`
}
