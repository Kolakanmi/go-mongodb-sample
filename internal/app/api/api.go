package api

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/handler/auth"
	user2 "github.com/Kolakanmi/go-mongodb-sample/internal/app/handler/user"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/repository/auth/mongo"
	mongo2 "github.com/Kolakanmi/go-mongodb-sample/internal/app/repository/user/mongo"
	auth2 "github.com/Kolakanmi/go-mongodb-sample/internal/app/service/auth"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/user"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/db/mongodb"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/health"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/middleware"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/router"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/jwt"
	log2 "github.com/Kolakanmi/go-mongodb-sample/internal/pkg/log"
	"log"
	"net/http"
)

func NewRouter() (http.Handler, error) {
	mConf := mongodb.LoadConfigFromFile()
	db, err := mongodb.ConnectDB(mConf)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	userRepo := mongo2.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user2.NewHandler(userService)

	jwtConf := jwt.Config{JWTSecretKey: "secret-key"}
	jwtGen := jwt.New(jwtConf)
	authRepo := mongo.NewAuthRepository(db)


	authServ := auth2.NewService(jwtGen, authRepo)
	authServ.RegisterService("user", userService)
	authHandler := auth.NewHandler(authServ)

	routes := []router.Route{
		{
			Path: "/readiness",
			Method: http.MethodGet,
			Handler: health.Readiness(),
		},
	}

	routes = append(routes, authHandler.Routes()...)
	routes = append(routes, userHandler.Routes()...)

	rConf := router.GetEmptyConfig()
	rConf.Routes = routes
	rConf.GlobalMiddlewares = []router.Middleware{
		middleware.Recover,
		log2.NewHTTPContextHandler(log2.Root()),
		auth2.UserAuthMiddleware(jwtGen, authRepo),
		middleware.StatusResponseWriter,
		middleware.HTTPRequestResponseInfo(nil),
	}

	r, err := router.New(rConf)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	health.Ready()
	return middleware.Cors(r), nil
}
