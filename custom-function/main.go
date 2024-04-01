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
			// グローバルな関数なのでOverloadを使う
			// 命名規則はfunc_argType_argType
			cel.Overload("concatStr_string_string",
				// 引数の型を指定
				[]*cel.Type{cel.StringType, cel.StringType},
				// 戻り値の型を指定
				cel.StringType,
				// 引数が２つの関数を用意するのでBinaryBinding
				cel.BinaryBinding(concatFunc),
			),
		),
	)
	if err != nil {
		log.Fatalf("failed to create env: %s\n", err)
	}

	// 評価式でカスタム関数を呼び出す。今回は変数なし
	expr := `concatStr('Hello', 'World')`
	ast, iss := env.Compile(expr)
	if iss != nil && iss.Err() != nil {
		log.Fatalf("failed to compile expression: %v", iss.Err())
	}

	p, err := env.Program(ast)
	if err != nil {
		log.Fatalf("failed to create program: %s\n", err)
	}

	// 変数なしなので空のmapを渡す
	out, _, err := p.Eval(map[string]interface{}{})
	if err != nil {
		log.Fatalf("evaluation error: %s\n", err)
	}

	fmt.Println("Result:", out)
}

func globalWithVariableArg(a, b string) {
	env, err := cel.NewEnv(
		// 引数を変数として定義
		cel.Variable("arg1", cel.StringType),
		cel.Variable("arg2", cel.StringType),
		cel.Function("concatStr",
			// グローバルな関数なのでOverloadを使う
			cel.Overload("concatStr_string_string",
				// 引数の型を指定
				[]*cel.Type{cel.StringType, cel.StringType},
				// 戻り値の型を指定
				cel.StringType,
				// 引数が２つの関数を用意するのでBinaryBinding
				cel.BinaryBinding(concatFunc),
			),
		),
	)
	if err != nil {
		log.Fatalf("failed to create env: %s\n", err)
	}

	// 評価式でカスタム関数を呼び出す。今回は変数あり
	expr := `concatStr(arg1, arg2)`
	ast, iss := env.Compile(expr)
	if iss != nil && iss.Err() != nil {
		log.Fatalf("failed to compile expression: %v", iss.Err())
	}

	p, err := env.Program(ast)
	if err != nil {
		log.Fatalf("failed to create program: %s\n", err)
	}

	// 入力値として変数を渡す
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
		// メンバ、引数を変数として定義
		cel.Variable("target", cel.StringType),
		cel.Variable("arg1", cel.StringType),
		cel.Function("concatStr",
			// メンバ関数の場合はMemberOverloadを使う
			// 命名規則はtargetType_func_argType_argType
			cel.MemberOverload("string_concatStr_string",
				// 引数の型を指定。第一引数がメンバの型になる
				[]*cel.Type{cel.StringType, cel.StringType},
				// 戻り値の型を指定
				cel.StringType,
				// 引数が２つの関数を用意するのでBinaryBinding
				cel.BinaryBinding(concatFunc),
			),
		),
	)
	if err != nil {
		log.Fatalf("failed to create env: %s\n", err)
	}

	// メンバ関数の場合は、メンバに関数を生やして記述する
	expr := `target.concatStr(arg1)`
	ast, iss := env.Compile(expr)
	if iss != nil && iss.Err() != nil {
		log.Fatalf("failed to compile expression: %v", iss.Err())
	}

	p, err := env.Program(ast)
	if err != nil {
		log.Fatalf("failed to create program: %s\n", err)
	}

	out, _, err := p.Eval(map[string]interface{}{
		"target": a,
		"arg1":   b,
	})
	if err != nil {
		log.Fatalf("evaluation error: %s\n", err)
	}

	fmt.Println("Result:", out)
}
