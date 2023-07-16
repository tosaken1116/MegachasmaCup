package directive

import (
	"context"
	"errors"
	"fmt"
	"megachasma/internal"
	"megachasma/middleware/auth"

	"github.com/99designs/gqlgen/graphql"
)

var Directive internal.DirectiveRoot = internal.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fmt.Println("authentication")
	if _, ok := auth.GetUserID(ctx); !ok {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
