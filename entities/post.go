package entities

type Posts struct {
	Post_id     uint   `json:"post_id"`
	Title       string `json:"post_title"`
	Description string `json:"post_description"`
	Created_at  string `json:"created_at"`
	Update_at   string `json:"updated_at"`
}
