package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type MyOrder struct {
	a int
	b int
}

type OrderFacade struct {
	MyOrder
}

func (pmt *MyOrder) MakePayment(n int) {
	pmt.a += n
}
func (pmt *MyOrder) MakeCost(n int) {
	pmt.b += n
}
func (pmt *OrderFacade) Increment() {
	pmt.MakeCost(1)
	pmt.MakePayment(1)
}
