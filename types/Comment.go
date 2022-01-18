package types

type Comment struct {
	Comments  []Comment
	User			User
	Id        string
	Comment   string
	Upvotes   int64
	Downvotes int64
	Platform  string
	CreatedAt string
	RelTime		string
	UpdatedAt string
	DeletedAt string
}
