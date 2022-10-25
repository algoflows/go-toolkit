package toolkit

import (
	"fmt"
	"testing"
)

func TestTools_RandomString(t *testing.T) {
	var tools Tools // Instantiate the module

	s := tools.RandomString(10) // Call the RandomString method
	fmt.Println("Random string:", s)
	
	if len(s) != 10 {
		t.Error("Random string length is not 10")
	}
}