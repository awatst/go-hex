package core

// secondary port (port ขาออก ลงเบสน่ะ)
type OrderRepository interface {
	Save(order Order) error
}
