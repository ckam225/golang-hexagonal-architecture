package entity

type Post struct {
	ID      int
	Title   string
	Content string
	UserId  *int
}

type PostFilter struct {
	Limit   int
	ID      int
	Title   string
	Content string
	UserId  int
}
