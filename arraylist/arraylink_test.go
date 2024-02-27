package arraylist

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	ArrayLink()
}

func TestArrayList2(t *testing.T) {
	a := Make(0, 3)
	fmt.Println("cap:", a.Cap(), "len:", a.Len(), "array:", PrintArray(a))

	a.Append(10)
	fmt.Println("cap:", a.Cap(), "len:", a.Len(), "array:", PrintArray(a))

	a.Append(9)
	fmt.Println("cap:", a.Cap(), "len:", a.Len(), "array:", PrintArray(a))

	a.AppendMany(8, 7, 6)
	fmt.Println("cap:", a.Cap(), "len:", a.Len(), "array:", PrintArray(a))

}
