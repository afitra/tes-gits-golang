package publisher

type PublisherFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatPublisher(publisher Publisher) PublisherFormatter {

	formatter := PublisherFormatter{
		ID:   publisher.ID,
		Name: publisher.Name,
	}
	return formatter
}
