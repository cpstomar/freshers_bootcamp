package main
import (
	"fmt"
	"math/rand"
)
type Employee interface {
	salary(timeperiod int) int
}
type Fulltime struct {
	basicpay int
}
type Contactor struct {
	basicpay int
}
type Freelancer struct {
	basicpay int
}
func (e Fulltime) salary(days int) int {
	return days * e.basicpay
}
func (e Contactor) salary(days int) int {
	return days * e.basicpay
}
func (e Freelancer) salary(hours int) int {
	return hours * e.basicpay
}
func main(){
	f := Fulltime{basicpay:500}
	c := Contactor{basicpay: 100}
	fr := Freelancer{basicpay: 10}
	data := [3]Employee{f, c, fr}
	for idx, elem := range data {
		if idx == 2{
			fmt.Println(elem.salary(rand.Intn(200)))
		}else {
			fmt.Println(elem.salary(rand.Intn(28)))
		}
	}
}
