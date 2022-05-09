package vo

type City struct {
	Name string `json:"name"`
	Country string `json:"country"`
	Condition string `json:"condition,omitempty"`
}