package models


type TokenResponse struct {
	Token string `json:"access_token"`
}

type UserDetailsResponse struct {
	Login string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	Name string `json:"name"`
	TwitterUsername string `json:"twitter_username"`
	Followers int `json:"followers"`
	Following int `json:"following"`
	Repos int `json:"public_repos"`
	Bio string `json:"bio"`
}


