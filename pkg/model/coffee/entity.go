package coffee

type Coffee struct {
	ID          int64   `json: "id"`
	Variety     string  `json: "variety"`
	Bitterness  float32 `json: "bitterness"`
	Description string  `json: description`
}
