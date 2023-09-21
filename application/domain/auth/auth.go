package auth

type Token struct {
	Context string `json:"context"`
	Token   string `json:"token"`
	Secret  string `json:"secret"`
}
