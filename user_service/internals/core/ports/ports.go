package ports

type UserDBRepository interface {
	Create(data any) (any, error)
	Find(condition []string, dataModel any) any
	Table(tname string) UserDBRepository
}
