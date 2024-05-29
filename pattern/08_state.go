package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type driverState interface {
	state(*someDriver)
}

type someDriver struct {
	state driverState
}

func (dr *someDriver) GetState() driverState {
	return dr.state
}

func (dr *someDriver) setState(state driverState) {
	dr.state = state
}

type AlcoholState struct {
	Description string
}

func (state *AlcoholState) state(dr *someDriver) {
	dr.setState(state)
}

type NormalState struct {
	Description string
}

func (state *NormalState) state(dr *someDriver) {
	dr.setState(state)
}
