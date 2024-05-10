package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/cel-go/cel"
)

func main() {
	// CEL環境の設定
	env, err := cel.NewEnv(
		cel.Variable("num", cel.IntType),
		cel.Variable("users", cel.ListType(cel.MapType(cel.StringType, cel.AnyType))),
	)
	if err != nil {
		log.Fatalf("Failed to create CEL environment: %v", err)
	}
	b, err := os.ReadFile("./expr.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	expr := string(b)

	ast, issues := env.Compile(expr)
	if issues != nil && issues.Err() != nil {
		log.Fatalf("Compile error: %v", issues.Err())
	}

	prg, err := env.Program(ast)
	if err != nil {
		log.Fatalf("Program creation error: %v", err)
	}

	// Usersは含めない
	req := Request{
		Num: 10,
	}
	data, _ := json.Marshal(req)
	fmt.Println(string(data))
	var inputs map[string]interface{}
	if err := json.Unmarshal([]byte(data), &inputs); err != nil {
		log.Fatalf("JSON Unmarshal error: %v\n", err)
	}

	result, _, err := prg.Eval(inputs)
	if err != nil {
		log.Fatalf("Evaluation error: %v", err)
	}

	fmt.Printf("found?", result.Value())
}

type Request struct {
	Num   int    `json:"num"`
	Users []User `json:"users"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
