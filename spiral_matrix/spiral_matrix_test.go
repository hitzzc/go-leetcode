package spiral_matrix

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	fmt.Printf("%v\n", spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
}
