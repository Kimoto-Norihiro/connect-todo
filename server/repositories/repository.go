package repositories

type ITODORepository interface {
	CreateTODO() error
	ListTODOs() error
	UpdateTODO() error
	DeleteTODO() error
}
