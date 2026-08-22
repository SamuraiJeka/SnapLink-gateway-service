package dto

type Link struct {
	Id uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	ShortUrl string `json:"short_url"`
	BaseUrl string `json:"base_url"`
}

type CreateLink struct {
	Url string `json:"url"`
}
