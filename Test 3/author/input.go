package author

type AuthorDataInput struct {
	Name string `json:"name" binding:"required"`
}
type AuthorParamInput struct {
	ID int `uri:"id" binding:"required"`
}
