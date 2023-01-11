package author

type AuthorFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatAuthor(author Author) AuthorFormatter {

	formatter := AuthorFormatter{
		ID:   author.ID,
		Name: author.Name,
	}
	return formatter
}
