package server

import (
	"fmt"
	"net/http"

	"bitbucket.org/andreychernih/tweemote/lib"
	"bitbucket.org/andreychernih/tweemote/services"
	"github.com/RangelReale/osin"
)

func CreateAccessToken(w http.ResponseWriter, r *http.Request) {
	resp := lib.GetOsinServer().NewResponse()
	resp.StatusCode = 201
	resp.ErrorStatusCode = 403
	defer resp.Close()

	if ar := lib.GetOsinServer().HandleAccessRequest(resp, r); ar != nil {
		if ar.Type == osin.PASSWORD {
			user, err := services.AuthorizeUser(ar.Username, ar.Password)
			if user != nil && err == nil {
				ar.Authorized = true
				ar.UserData = fmt.Sprintf("%d", user.ID)
			}
		} else if ar.Type == osin.REFRESH_TOKEN {
			ar.Authorized = true
		}

		lib.GetOsinServer().FinishAccessRequest(resp, r, ar)
	}

	//log.Println("Internal error:", resp.InternalError)

	osin.OutputJSON(resp, w, r)
}
