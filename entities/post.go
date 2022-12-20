package entities

type Posts struct {
	PostId     		uint   `json:"id"`
	PostAuthor		string `json:"author"`
	PostTitle       string `json:"title"`
	PostDescription string `json:"description"`
	PostUrl			string `json:"url"`
	PostSource		string `json:"source"`
	PostImage		string `json:"image"`
	PostCategory	string `json:"category"`
	PostLanguage	string `json:"language"`
	PostCountry		string `json:"country"`
	PostPublishedAt  string `json:"published_at"`
}
