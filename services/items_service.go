package services

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

//itemsServiceInterface includes all methods that can be called only on items
type itemsServiceInterface interface {
	GetItem()
	SaveItem()
}

type itemsService struct{}

func (s *itemsService) GetItem()  {}
func (s *itemsService) SaveItem() {}
