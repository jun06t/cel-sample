package main

import (
	"fmt"
	"log"

	"github.com/google/cel-go/cel"
)

func main() {
	// CEL環境の設定
	env, err := cel.NewEnv(
		// 'num' という名前の整数型の変数を宣言
		cel.Variable("num", cel.IntType),
	)
	if err != nil {
		log.Fatalf("Failed to create CEL environment: %v", err)
	}

	// 入力値が偶数かどうかをチェックするCEL式
	expr := `num % 2 == 0`

	// 式のコンパイル
	ast, issues := env.Compile(expr)
	if issues != nil && issues.Err() != nil {
		log.Fatalf("Compile error: %v", issues.Err())
	}

	// プログラムの生成
	prg, err := env.Program(ast)
	if err != nil {
		log.Fatalf("Program creation error: %v", err)
	}

	// 評価する入力値
	inputs := map[string]interface{}{
		"num": 9, // この値を変更して異なる入力で試すことができます
	}

	// プログラムの評価
	result, _, err := prg.Eval(inputs)
	if err != nil {
		log.Fatalf("Evaluation error: %v", err)
	}

	// 結果の出力
	fmt.Printf("Is %v an even number? %v\n", inputs["num"], result.Value().(bool))
}