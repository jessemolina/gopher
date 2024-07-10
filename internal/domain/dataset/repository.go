package dataset

type Repsitory interface {
	Create() error
	Update() error
	Delete() error
	Query() error
}
