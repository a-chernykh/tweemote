package tests

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/RangelReale/osin"

	"bitbucket.org/andreychernih/tweemote/ctx"
	"bitbucket.org/andreychernih/tweemote/lib"
	"bitbucket.org/andreychernih/tweemote/models"
)

func StubAuth(r *http.Request, u *models.User) {
	tok := AddValidUserToken(lib.GetOsinServer(), lib.GetDefaultOsinClient(), u)
	r.Header.Set("Authorization", "bearer "+tok.AccessToken)
}

func AddValidUserToken(s *osin.Server, c osin.Client, u *models.User) *osin.AccessData {
	return AddUserToken(s, c, u, 3600)
}

func AddExpiredUserToken(s *osin.Server, c osin.Client, u *models.User) *osin.AccessData {
	return AddUserToken(s, c, u, 0)
}

func AddUserToken(s *osin.Server, c osin.Client, u *models.User, expiresIn int32) *osin.AccessData {
	ret := &osin.AccessData{
		Client:        c,
		AuthorizeData: nil,
		AccessData:    nil,
		AccessToken:   "fake_token_from_tests",
		RefreshToken:  "fake_refresh_token_from_tests",
		RedirectUri:   "/redir",
		CreatedAt:     time.Now(),
		ExpiresIn:     expiresIn,
		UserData:      fmt.Sprintf("%d", u.ID),
		Scope:         "write",
	}
	err := s.Storage.SaveAccess(ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func AddUser(r *http.Request, u *models.User) *http.Request {
	nc := context.WithValue(r.Context(), ctx.CtxUserKey, u)
	return r.WithContext(nc)
}
