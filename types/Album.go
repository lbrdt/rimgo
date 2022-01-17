package types

type Album struct {
	Id        string
	Title     string
	Views     int64
	Upvotes   int64
	Downvotes int64
	CreatedAt string
	UpdatedAt string
	Media     []string
}
