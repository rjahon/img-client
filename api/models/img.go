package models

type Img struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Body      []byte `json:"body"`
}

type GetImagesModel struct {
	Offset string `json:"offset"`
	Limit  string `json:"limit"`
}
