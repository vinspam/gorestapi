package aboutme

type Aboutme struct {
	User string `json:"user"`
	JoinedAt float64 `json:"joined-at"`
	Toptracks string `json:"toptracks"`
	Following string `json:"following"`
}