package ports

type ApiGatewayRepository interface {
	Create(data any) (any, error)
	Find(condition []string, dataModel any) any
	Table(tname string) ApiGatewayRepository
}
