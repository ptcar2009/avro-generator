package demo

type Testing struct {
	Other  *Other `json:"other"`
	Inline struct {
		Single int `json:"other"`
	}
	K   int
	Opa []int
}

type Other struct {
	Field int `json:"field"`
}
