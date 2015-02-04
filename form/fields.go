package form

type Field struct {
	Required string
}

type FieldHandler interface {
	Validate()
}
