package book

type BookDataInput struct {
	Name        string `json:"name" binding:"required"`
	AuthorID    int    `json:"author_id" binding:"required"`
	PublisherID int    `json:"publisher_id" binding:"required"`
}
type BookParamInput struct {
	ID int `uri:"id" binding:"required"`
}
