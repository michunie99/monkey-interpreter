package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is Johnny"}
	diff2 := &String{Value: "My name is Johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same contents have diffrent hashes")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same contents have diffrent hashes")
	}

	if diff1.HashKey() == hello1.HashKey() {
		t.Errorf("strings with diffrent contents have same hashes")
	}
}
