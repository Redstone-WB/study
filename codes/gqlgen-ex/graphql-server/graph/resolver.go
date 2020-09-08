package graph

import "github.com/redstone-wb/gqlgen-todos/graph/model"

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	videos []*model.Video  // 여기에 * 추가해줘야
}
