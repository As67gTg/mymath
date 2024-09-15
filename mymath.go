package mymath

import "math"
//корень из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
//округление до наим целого большего или равноего х
func Ceil(x float64) float64 {
	return math.Ceil(x)
}
//Округление до наиб целого меньшего или равного х
func Floor(x float64) float64 {
	return math.Floor(x)
}
//возведение х в степень у
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}
//максимальное число из пары
func Max(x, y float64) float64 {
	return math.Max(x, y)
}
//минимальнео число из пары
func Min(x, y float64) float64 {
	return math.Min(x, y)
}
//модуль
func Abs(x float64) float64 {
	return math.Abs(x)
}
func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}

