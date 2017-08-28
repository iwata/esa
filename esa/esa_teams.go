package esa

// Team represents a esa team.
// ref. https://docs.esa.io/posts/102#4-2-0
type Team struct {
	Name        string `json:"name"`
	Privacy     string `json:"privacy"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	URL         string `json:"url"`
}

func (t Team) String() string {
	return Stringify(t)
}
