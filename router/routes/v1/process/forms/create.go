package forms

// CreateForm - form for creating process
type CreateForm struct {
	Duration int         `json:"duration"`
	Data     interface{} `json:"data"`
}
