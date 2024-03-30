package main

import (
	"testing"
)

var inputs = map[string]interface{}{
	"firstname": "John",
	"lastname":  "Doe",
	"age":       30,
	"email":     "john@example.com",
	"address":   "123 Main St",
	"tel":       "1234567890",
}

// optimizeがtrueの場合のベンチマークテスト
func BenchmarkNewProgramOptimizeTrue(b *testing.B) {
	prog, err := NewProgram(true)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		prog.Eval(inputs)
	}
}

// optimizeがfalseの場合のベンチマークテスト
func BenchmarkNewProgramOptimizeFalse(b *testing.B) {
	prog, err := NewProgram(false)
	if err != nil {
		b.Error(err)
	}
	inputs := map[string]interface{}{
		"num": 10,
	}
	for i := 0; i < b.N; i++ {
		prog.Eval(inputs)
	}
}

func BenchmarkRawCode(b *testing.B) {
	u := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john@example.com",
		Address:   "123 Main St",
		Tel:       "1234567890",
	}
	for i := 0; i < b.N; i++ {
		RawCode(u)
	}
}
