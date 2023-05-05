package lib

import (
	"net/http"
	"sync"

	"bitbucket.org/andreychernih/tweemote/models"
	"github.com/RangelReale/osin"
)

var osinServer *osin.Server
var osinServerOnce sync.Once

func GetOsinServer() *osin.Server {
	osinServerOnce.Do(func() {
		gormDb := models.Connect()
		db := gormDb.DB()
		if db == nil {
			panic("db is not supposed to be nil")
		}

		storage, err := CreateOsinStorage(db)
		if err != nil {
			panic(err)
		}

		sconfig := osin.NewServerConfig()
		sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
		sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE,
			osin.REFRESH_TOKEN, osin.PASSWORD, osin.CLIENT_CREDENTIALS, osin.ASSERTION}
		sconfig.AllowGetAccessRequest = true
		sconfig.AllowClientSecretInParams = true
		sconfig.AccessExpiration = 3600

		osinServer = osin.NewServer(sconfig, storage)
	})

	return osinServer
}

func VerifyAuth(s *osin.Server, r *http.Request, w http.ResponseWriter) (*models.User, error) {
	resp := s.NewResponse()
	defer resp.Close()

	if ir := s.HandleInfoRequest(resp, r); ir == nil {
		w.WriteHeader(http.StatusUnauthorized)
		osin.OutputJSON(resp, w, r)

		return nil, nil
	} else {
		ad := ir.AccessData
		userId := ad.UserData.(string)
		db := models.Connect()
		var user models.User
		if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
			return nil, err
		}

		return &user, nil
	}
}
