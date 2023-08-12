package entity

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  *int   `json:"user_id"`
}

type PostFilter struct {
	Limit   int
	ID      int
	Title   string
	Content string
	UserId  int
}
