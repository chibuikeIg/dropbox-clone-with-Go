package ports

type FileMetaDataDBRepository interface {
	Create(data any) (any, error)
	Find(condition []string, dataModel any) any
	Table(tname string) FileMetaDataDBRepository
	Where(data [][]any) FileMetaDataDBRepository
	Update(data map[string]string) (any, error)
	First() (any, error)
	Get() (any, error)
	Delete() error
}
