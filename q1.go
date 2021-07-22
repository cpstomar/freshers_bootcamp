package main
import (
	"fmt"
	"math/rand"
)
type Matrix struct {
	row int
	col int
	matrix [][]int
}
func (m Matrix) getrow() int {
	return m.row
}
func (m Matrix) getcolumn() int {
	return m.col
}
func (m Matrix) setelem(i int , j int, val int) bool {
	if i >= m.row && j >= m.col {
		return false
	}
	m.matrix[i][j] = val
	return true
}
func (m Matrix) add(M Matrix) bool {
	if M.row != m.row && M.col != m.col {
		return false
	}
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			m.matrix[i][j] += M.matrix[i][j]
		}
	}
	return true
}
func (m Matrix) print() bool {
	fmt.Println("{" )
	for i := 0; i < m.row; i++ {
		fmt.Printf("\t%s", "{ ")
		for j := 0; j < m.col; j++ {
			fmt.Printf("%d,",m.matrix[i][j])
		}
		fmt.Println("}")
	}
	fmt.Println("}")
	return true
}
func generateMatrix(r int, c int) Matrix{
	m := Matrix{
		row: r,
		col: c,
	}
	m.matrix = make([][]int,  r)
	for i := 0; i < m.row; i++ {
		m.matrix[i] = make([]int, c)
		for j := 0; j < m.col; j++ {
			m.matrix[i][j] = rand.Intn(10)
		}
	}
	return m
}

func main(){
	m1 := generateMatrix(3, 3)
	m2 := generateMatrix(3, 3)
	fmt.Println(m1.getrow())
	fmt.Println(m1.getcolumn())
	m1.print()
	m2.print()
	m1.add(m2)
	m1.print()
}
