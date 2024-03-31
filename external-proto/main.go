package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/google/cel-go/cel"
	pb "github.com/jun06t/cel-sample/external-proto/proto"
	rpcpb "google.golang.org/genproto/googleapis/rpc/context/attribute_context"
	"google.golang.org/protobuf/types/known/structpb"
	tpb "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	runExternalProto()
	runInternalProto()
}

func runExternalProto() {
	env, _ := cel.NewEnv(
		cel.Types(&rpcpb.AttributeContext_Request{}),
		cel.Variable("request",
			cel.ObjectType("google.rpc.context.AttributeContext.Request"),
		),
	)

	ast, iss := env.Compile(
		`request.auth.claims.group == 'admin'
            || request.auth.principal == 'user:me@acme.co'`,
	)
	if iss != nil && iss.Err() != nil {
		log.Fatalf("Compile error: %v", iss.Err())
	}
	if !reflect.DeepEqual(ast.OutputType(), cel.BoolType) {
		log.Fatalf("Got unexpected output type: %v", ast.OutputType())
	}

	prog, err := env.Program(ast)
	if err != nil {
		log.Fatalf("Program construction error: %v", err)
	}

	claims := map[string]string{"group": "admin"}
	input := request(auth("user:me@acme.co", claims), time.Now())
	out, _, err := prog.Eval(input)
	if err != nil {
		log.Fatalf("Evaluation error: %v", err)
	}

	fmt.Println("Is authorized user?", out)
}

func runInternalProto() {
	env, _ := cel.NewEnv(
		cel.Types(&pb.HelloRequest{}),
		cel.Variable("request",
			cel.ObjectType("helloworld.HelloRequest"),
		),
	)

	ast, iss := env.Compile(
		`request.name == 'Alice' && request.age > 20 && request.man == false`,
	)
	if iss != nil && iss.Err() != nil {
		log.Fatalf("Compile error: %v", iss.Err())
	}
	if !reflect.DeepEqual(ast.OutputType(), cel.BoolType) {
		log.Fatalf("Got unexpected output type: %v", ast.OutputType())
	}

	prog, err := env.Program(ast)
	if err != nil {
		log.Fatalf("Program construction error: %v", err)
	}

	input := map[string]any{
		"request": &pb.HelloRequest{
			Name: "Alice",
			Age:  21,
			Man:  false,
		},
	}
	out, _, err := prog.Eval(input)
	if err != nil {
		log.Fatalf("Evaluation error: %v", err)
	}

	fmt.Println("Is permitted user?", out)
}

// auth constructs a `google.rpc.context.AttributeContext.Auth` message.
func auth(user string, claims map[string]string) *rpcpb.AttributeContext_Auth {
	claimFields := make(map[string]*structpb.Value)
	for k, v := range claims {
		claimFields[k] = structpb.NewStringValue(v)
	}
	return &rpcpb.AttributeContext_Auth{
		Principal: user,
		Claims:    &structpb.Struct{Fields: claimFields},
	}
}

// request constructs a `google.rpc.context.AttributeContext.Request` message.
func request(auth *rpcpb.AttributeContext_Auth, t time.Time) map[string]any {
	req := &rpcpb.AttributeContext_Request{
		Auth: auth,
		Time: &tpb.Timestamp{Seconds: t.Unix()},
	}
	return map[string]any{"request": req}
}
