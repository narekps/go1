/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3, y1, y2, y3 int

	fmt.Print("Введите X1:")
	fmt.Scanf("%d\n", &x1)
	fmt.Print("Введите Y1:")
	fmt.Scanf("%d\n", &y1)

	fmt.Print("Введите X2:")
	fmt.Scanf("%d\n", &x2)
	fmt.Print("Введите Y2:")
	fmt.Scanf("%d\n", &y2)

	fmt.Print("Введите X3:")
	fmt.Scanf("%d\n", &x3)
	fmt.Print("Введите Y3:")
	fmt.Scanf("%d\n", &y3)

	line1, line2, line3 := calcLine(x1, y1, x2, y2), calcLine(x2, y2, x3, y3), calcLine(x1, y1, x3, y3)

	if !isTriangle(x1, x2, x3, y1, y2, y3) {
		fmt.Println("Нельзя построить треугольник.")
		return
	}

	fmt.Printf("Lines: %.2f, %.2f, %.2f\n", line1, line2, line3)

	fmt.Printf("Площадь: %.2f\n", calcSquare(x1, x2, x3, y1, y2, y3))

	if checkTriangle(line1, line2, line3) {
		fmt.Println("Треугольник является прямоугольным")
	} else {
		fmt.Println("Треугольник не является прямоугольным")
	}
}

func calcLine(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2.0) + math.Pow(float64(y1-y2), 2.0))
}

func isTriangle(x1, x2, x3, y1, y2, y3 int) bool {
	return (y3-y1)*(x2-x1) != (x3-x1)*(y2-y1)
}

func calcSquare(x1, x2, x3, y1, y2, y3 int) float64 {
	var s float64
	s = math.Abs(float64((x2-x1)*(y3-y1)-(x3-x1)*(y2-y1))) / 2

	return s
}

func checkTriangle(line1, line2, line3 float64) bool {
	if line1 > line2 && line1 > line3 {
		return math.Floor(line1*line1) == math.Floor(line2*line2)+math.Floor(line3*line3) || math.Ceil(line1*line1) == math.Ceil(line2*line2)+math.Ceil(line3*line3)
	}
	if line2 > line1 && line2 > line3 {
		return math.Floor(line2*line2) == math.Floor(line1*line1)+math.Floor(line3*line3) || math.Ceil(line2*line2) == math.Ceil(line1*line1)+math.Ceil(line3*line3)
	}
	if line3 > line1 && line3 > line2 {
		return math.Floor(line3*line3) == math.Floor(line1*line1)+math.Floor(line2*line2) || math.Ceil(line3*line3) == math.Ceil(line1*line1)+math.Ceil(line2*line2)
	}

	return false
}
