package forms

type Keyword struct {
	Keyword string `json:"keyword" valid:"required"`
}
