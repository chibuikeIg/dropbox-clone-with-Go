package ports

type ObjectStorageRepository interface {
	CreateMultipartUpload(data map[string]*string) (string, error)
	UploadPart(data map[string]any) (string, error)
	CompleteMultipartUpload(data map[string]any) (string, error)
	AbortMultipartUpload(data map[string]any) (bool, error)
}

type FileUploadDBRepository interface {
	Table(table string) FileUploadDBRepository
	Get() (any, error)
	Create(data any) (any, error)
	Delete() error
	Find(condition []string, dataModel any) any
	Where(data [][]any) FileUploadDBRepository
}
