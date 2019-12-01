package lazada_models

type AuthorizationResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpriesIn    int64  `json:"expires_in"`
	Account      string `json:"account"`
}
