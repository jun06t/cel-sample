package main

import (
	"fmt"
	"log"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

func main() {
	global()
	globalWithVariableArg("I'm ", "Alice")
	memberWithVariableArg("I'm ", "Bob")
}

func concatFunc(arg1, arg2 ref.Val) ref.Val {
	v1 := arg1.(types.String)
	v2 := arg2.(types.String)
	return types.String(v1 + v2)
}

func global() {
	env, err := cel.NewEnv(
		cel.Function("concatStr",
			cel.Overload("concatStr_string_string",
				[]*cel.Type{cel.StringType, cel.StringType},
				cel.StringType,
				cel.BinaryBinding(concatFunc),
			),
		),
	)
	if err != nil {
		log.Fatalf("failed to create env: %s\n", err)
	}

	expr := `concatStr('Hello', 'World')`
	ast, iss := env.Compile(expr)
	if iss != nil && iss.Err() != nil {
		log.Fatalf("failed to compile expression: %v", iss.Err())
	}

	p, err := env.Program(ast)
	if err != nil {
		log.Fatalf("failed to create program: %s\n", err)
	}

	out, _, err := p.Eval(map[string]interface{}{})
	if err != nil {
		log.Fatalf("evaluation error: %s\n", err)
	}

	fmt.Println("Result:", out)
}

func globalWithVariableArg(a, b string) {
	env, err := cel.NewEnv(
		cel.Variable("arg1", cel.StringType),
		cel.Variable("arg2", cel.StringType),
		cel.Function("concatStr",
			cel.Overload("concatStr_string_string",
				[]*cel.Type{cel.StringType, cel.StringType},
				cel.StringType,
				cel.BinaryBinding(concatFunc),
			),
		),
	)
	if err != nil {
		log.Fatalf("failed to create env: %s\n", err)
	}

	expr := `concatStr(arg1, arg2)`
	ast, iss := env.Compile(expr)
	if iss != nil && iss.Err() != nil {
		log.Fatalf("failed to compile expression: %v", iss.Err())
	}

	p, err := env.Program(ast)
	if err != nil {
		log.Fatalf("failed to create program: %s\n", err)
	}

	out, _, err := p.Eval(map[string]interface{}{
		"arg1": a,
		"arg2": b,
	})
	if err != nil {
		log.Fatalf("evaluation error: %s\n", err)
	}

	fmt.Println("Result:", out)
}

func memberWithVariableArg(a, b string) {
	env, err := cel.NewEnv(
		cel.Variable("arg1", cel.StringType),
		cel.Variable("arg2", cel.StringType),
		cel.Function("concatStr",
			cel.MemberOverload("string_concatStr_string",
				[]*cel.Type{cel.StringType, cel.StringType},
				cel.StringType,
				cel.BinaryBinding(concatFunc),
			),
		),
	)
	if err != nil {
		log.Fatalf("failed to create env: %s\n", err)
	}

	expr := `arg1.concatStr(arg2)`
	ast, iss := env.Compile(expr)
	if iss != nil && iss.Err() != nil {
		log.Fatalf("failed to compile expression: %v", iss.Err())
	}

	p, err := env.Program(ast)
	if err != nil {
		log.Fatalf("failed to create program: %s\n", err)
	}

	out, _, err := p.Eval(map[string]interface{}{
		"arg1": a,
		"arg2": b,
	})
	if err != nil {
		log.Fatalf("evaluation error: %s\n", err)
	}

	fmt.Println("Result:", out)
}
