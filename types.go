package vld

type Validatable interface {
	validate() error
}
