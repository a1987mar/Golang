package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required"`
}

type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required"`
	Hash string `json:"hash"`
}
