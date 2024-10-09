package core

import "errors"

//primary port (ขาเข้า)
type OrderService interface {
	//method
	CreateOrder(order Order) error
}

//รับข้อมูลของ business logic core สามารถเรียกใช้งาน secondary port ได้
type OrderServiceImpl struct {
	repo OrderRepository
}

func newOrderService(repo OrderRepository) OrderService {
	return &OrderServiceImpl{repo: repo}
}

//business logic func
func (s *OrderServiceImpl) CreateOrder(order Order) error {

	if order.Total <= 0 {
		return errors.New("total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	//return when no err found
	return nil
}
