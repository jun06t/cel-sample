package main

import (
	"fmt"
	"os"

	"github.com/google/cel-go/cel"
)

func genAst() (*cel.Env, *cel.Ast, error) {
	env, err := cel.NewEnv(
		cel.Variable("firstname", cel.StringType),
		cel.Variable("lastname", cel.StringType),
		cel.Variable("age", cel.IntType),
		cel.Variable("email", cel.StringType),
		cel.Variable("address", cel.StringType),
		cel.Variable("tel", cel.StringType),
	)
	if err != nil {
		return nil, nil, err
	}
	b, err := os.ReadFile("./expr.txt")
	if err != nil {
		return nil, nil, err
	}

	expr := string(b)

	ast, issues := env.Compile(expr)
	if issues != nil && issues.Err() != nil {
		return nil, nil, issues.Err()
	}
	return env, ast, nil
}

func NewProgram(optimize bool) (cel.Program, error) {
	env, ast, err := genAst()
	if err != nil {
		return nil, err
	}

	if optimize {
		prog, err := env.Program(ast, cel.EvalOptions(cel.OptOptimize))
		if err != nil {
			return nil, err
		}
		return prog, nil
	}
	prog, err := env.Program(ast)
	if err != nil {
		return nil, err
	}
	return prog, nil
}

func main() {
	prog, err := NewProgram(true)
	if err != nil {
		panic(err)
	}
	inputs := map[string]interface{}{
		"firstname": "John",
		"lastname":  "Doe",
		"age":       30,
		"email":     "john@example.com",
		"address":   "123 Main St",
		"tel":       "1234567890",
	}
	out, _, err := prog.Eval(inputs)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
