// Паттерн Мост (Bridge)
//

package bridge

// Тип Carer, описывает интерфейс автомобиля
type Carer interface {
	Rase() string // автомобиль умеет ездить (для этого и нужен двигатель)
}

// Тип Enginer, описывает интерфейс двигателя
// Каждый двигатель должен его реализовать
type Enginer interface {
	GetSound() string // Возвращает звук двигателя
}

// Тип Car, реализует автомобиль
type Car struct {
	engine Enginer
}

// Специфический Request.
// Будем его адаптировать
func (self *Car) Rase() string {
	return self.engine.GetSound()
}

// Тип EngineSuzuki, реализует двигатель Suzuki
type EngineSuzuki struct {
}

// Метод отвечает за завук двигателя
func (self *EngineSuzuki) GetSound() string {
	return "SssuuuuZzzuuuuKkiiiii"
}

// Тип EngineHonda, реализует двигатель Honda
type EngineHonda struct {
}

// Метод отвечает за завук двигателя
func (self *EngineHonda) GetSound() string {
	return "HhoooNnnnnnnnnDddaaaaaaa"
}

// Тип EngineLada, реализует двигатель АвтоВаза
type EngineLada struct {
}

// Метод отвечает за завук двигателя
func (self *EngineLada) GetSound() string {
	return "PhhhhPhhhhPhPhPhPhPh"
}
