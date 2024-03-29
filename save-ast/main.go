package main

import (
	"log"
	"os"

	"github.com/google/cel-go/cel"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/proto"
)

func main() {
	env, err := cel.NewEnv(
		cel.Variable("name", cel.StringType),
	)
	if err != nil {
		log.Fatalf("Failed to create CEL environment: %v", err)
	}

	// Read the serialized AST from the file
	astBytes, err := os.ReadFile("./gen/ast.pb")
	if err != nil {
		log.Fatalf("Failed to read AST from file: %v", err)
	}

	// Deserialize the AST from Protocol Buffers binary format
	var astPb exprpb.CheckedExpr
	if err := proto.Unmarshal(astBytes, &astPb); err != nil {
		log.Fatalf("Failed to deserialize AST: %v", err)
	}

	// Recover the AST structure
	ast := cel.CheckedExprToAst(&astPb)

	// Create a Program from the AST
	prg, err := env.Program(ast, cel.EvalOptions(cel.OptTrackState, cel.OptExhaustiveEval))
	if err != nil {
		log.Fatalf("Failed to create program: %v", err)
	}

	// Evaluate the Program with a given variable
	out, _, err := prg.Eval(map[string]interface{}{
		"name": "World",
	})
	if err != nil {
		log.Fatalf("Evaluation failed: %v", err)
	}

	log.Printf("Result: %v\n", out)
}
