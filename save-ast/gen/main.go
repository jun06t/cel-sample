package main

import (
	"log"
	"os"

	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/proto"
)

func main() {
	env, err := cel.NewEnv(
		cel.Variable("name", cel.StringType),
	)
	if err != nil {
		log.Fatalf("Failed to create CEL environment: %v", err)
	}

	ast, iss := env.Compile(`"Hello, " + name + "!"`)
	if iss.Err() != nil {
		log.Fatalf("Failed to compile expression: %v", iss.Err())
	}
	expr, err := cel.AstToCheckedExpr(ast)
	if err != nil {
		log.Fatalf("Failed to convert an Ast to an protobuf: %v", err)
	}

	// Serialize the AST to Protocol Buffers binary format
	astBytes, err := proto.Marshal(expr)
	if err != nil {
		log.Fatalf("Failed to serialize AST: %v", err)
	}

	// Save the serialized AST to a file
	if err := os.WriteFile("ast.pb", astBytes, 0644); err != nil {
		log.Fatalf("Failed to write AST to file: %v", err)
	}
}
