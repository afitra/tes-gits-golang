package publisher

type PublisherDataInput struct {
	Name string `json:"name" binding:"required"`
}
type PublisherParamInput struct {
	ID int `uri:"id" binding:"required"`
}
