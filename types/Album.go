package types

type Album struct {
	Id                  string
	Title               string
	Views               int64
	Upvotes             int64
	Downvotes           int64
	SharedWithCommunity bool
	CreatedAt           string
	UpdatedAt           string
	Comments            int64
	Media               []Media
}
