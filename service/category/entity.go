package category

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryCreateRequest struct {
	Name string `json:"name"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
