package chapter08

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

// Car を満たすように実装
type Basic struct{}

// Car を満たすように実装
type Middle struct{}

// Car を満たすように実装
type Premium struct{}

func Calc(car Car, minutes int) int {
	// TODO: 実装
	return 0
}
