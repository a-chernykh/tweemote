package ctx

import (
	"context"
	"fmt"

	"bitbucket.org/andreychernih/tweemote/models"
)

type key int

const CtxUserKey key = 0

func UserFromContext(ctx context.Context) *models.User {
	if user, ok := ctx.Value(CtxUserKey).(*models.User); !ok {
		panic(fmt.Sprintf("Context value for user has incorrect type: %v", user))
	} else {
		return user
	}
}
