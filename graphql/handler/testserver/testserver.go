package testserver

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
)

// New provides a server for use in tests that isn't relying on generated code. It isnt a perfect reproduction of
// a generated server, but it aims to be good enough to test the handler package without relying on codegen.
func New() *TestServer {
	next := make(chan struct{})
	completeSubscription := make(chan struct{})

	schema := gqlparser.MustLoadSchema(&ast.Source{Input: `
		type Query {
			name: String!
			find(id: Int!): String!
		}
		type Mutation {
			name: String!
		}
		type Subscription {
			name: String!
		}
	`})

	srv := &TestServer{
		next:                 next,
		completeSubscription: completeSubscription,
	}

	srv.Server = handler.New(&graphql.ExecutableSchemaMock{
		ExecFunc: func(ctx context.Context) graphql.ResponseHandler {
			opCtx := graphql.GetOperationContext(ctx)
			switch opCtx.Operation.Operation {
			case ast.Query:
				ran := false
				// If the query contains @defer, we will mimic a deferred response.
				if strings.Contains(opCtx.RawQuery, "@defer") {
					initialResponse := true
					return func(context context.Context) *graphql.Response {
						select {
						case <-ctx.Done():
							return nil
						case <-next:
							if initialResponse {
								initialResponse = false
								hasNext := true
								return &graphql.Response{
									Data:    []byte(`{"name":null}`),
									HasNext: &hasNext,
								}
							}
							hasNext := false
							return &graphql.Response{
								Data:    []byte(`{"name":"test"}`),
								HasNext: &hasNext,
							}
						case <-completeSubscription:
							return nil
						}
					}
				}
				return func(ctx context.Context) *graphql.Response {
					if ran {
						return nil
					}
					ran = true
					// Field execution happens inside the generated code, lets simulate some of it.
					ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
						Object: "Query",
						Field: graphql.CollectedField{
							Field: &ast.Field{
								Name:       "name",
								Alias:      "name",
								Definition: schema.Types["Query"].Fields.ForName("name"),
							},
						},
					})
					res, err := graphql.GetOperationContext(ctx).
						ResolverMiddleware(ctx, func(ctx context.Context) (any, error) {
							return &graphql.Response{Data: []byte(`{"name":"test"}`)}, nil
						})
					if err != nil {
						panic(err)
					}
					return res.(*graphql.Response)
				}
			case ast.Mutation:
				return graphql.OneShot(graphql.ErrorResponse(ctx, "mutations are not supported"))
			case ast.Subscription:
				return func(context context.Context) *graphql.Response {
					select {
					case <-ctx.Done():
						return nil
					case <-next:
						return &graphql.Response{
							Data: []byte(`{"name":"test"}`),
						}
					case <-completeSubscription:
						return nil
					}
				}
			default:
				return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
			}
		},
		SchemaFunc: func() *ast.Schema {
			return schema
		},
		ComplexityFunc: func(ctx context.Context, typeName string, fieldName string, childComplexity int, args map[string]any) (i int, b bool) {
			return srv.complexity, true
		},
	})
	return srv
}

// NewError provides a server for use in resolver error tests that isn't relying on generated code. It isnt a perfect reproduction of
// a generated server, but it aims to be good enough to test the handler package without relying on codegen.
func NewError() *TestServer {
	next := make(chan struct{})

	schema := gqlparser.MustLoadSchema(&ast.Source{Input: `
		type Query {
			name: String!
		}
	`})

	srv := &TestServer{
		next: next,
	}

	srv.Server = handler.New(&graphql.ExecutableSchemaMock{
		ExecFunc: func(ctx context.Context) graphql.ResponseHandler {
			opCtx := graphql.GetOperationContext(ctx)
			switch opCtx.Operation.Operation {
			case ast.Query:
				ran := false
				return func(ctx context.Context) *graphql.Response {
					if ran {
						return nil
					}
					ran = true

					graphql.AddError(ctx, errors.New("resolver error"))

					return &graphql.Response{
						Data: []byte(`null`),
					}
				}
			case ast.Mutation:
				return graphql.OneShot(graphql.ErrorResponse(ctx, "mutations are not supported"))
			case ast.Subscription:
				return graphql.OneShot(graphql.ErrorResponse(ctx, "subscription are not supported"))
			default:
				return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
			}
		},
		SchemaFunc: func() *ast.Schema {
			return schema
		},
		ComplexityFunc: func(ctx context.Context, typeName string, fieldName string, childComplexity int, args map[string]any) (i int, b bool) {
			return srv.complexity, true
		},
	})
	return srv
}

type TestServer struct {
	*handler.Server
	next                 chan struct{}
	completeSubscription chan struct{}
	complexity           int
}

func (s *TestServer) SendNextSubscriptionMessage() {
	select {
	case s.next <- struct{}{}:
	case <-time.After(1 * time.Second):
		fmt.Println("WARNING: no active subscription")
	}
}

func (s *TestServer) SendCompleteSubscriptionMessage() {
	select {
	case s.completeSubscription <- struct{}{}:
	case <-time.After(1 * time.Second):
		fmt.Println("WARNING: no active subscription")
	}
}

func (s *TestServer) SetCalculatedComplexity(complexity int) {
	s.complexity = complexity
}
