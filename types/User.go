package types

type User struct {
	Id        int64  `json:"id"`
	Bio       string `json:"bio"`
	Username  string `json:"username"`
	Points    int64  `json:"reputation_count"`
	Cover     string `json:"cover_url"`
	Avatar    string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
}

type Submission struct {
	Id        string
	Title     string
	Link      string
	Cover     Media
	Points    int64
	Upvotes   int64
	Downvotes int64
	Comments  int64
	Views     int64
	IsAlbum   bool
}
