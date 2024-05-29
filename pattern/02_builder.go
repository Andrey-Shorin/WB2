package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Bicycle struct {
	make   string
	model  string
	color  int
	height int
}

type IBicycleBuilder struct {
	color  int
	height int
}

func (a *IBicycleBuilder) GetResult() *Bicycle {
	if a.height == 29 {
		return &Bicycle{"GT", "Avalanche", a.color, a.height}
	}

	return nil
}

type MountainBikeBuildDirector struct {
	IBicycleBuilder
}

func (a *MountainBikeBuildDirector) Construct() {
	a.height = 29
	a.color = 255
}
func (a *MountainBikeBuildDirector) GetResult() {
	a.IBicycleBuilder.GetResult()
}
