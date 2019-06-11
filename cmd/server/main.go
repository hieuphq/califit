package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/hieuphq/califit/src/interfaces/config"
	"github.com/hieuphq/califit/src/interfaces/http"
	"github.com/hieuphq/califit/src/interfaces/repository"
	"github.com/hieuphq/califit/src/registry"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	// r.Use(middleware.Logger())

	r.Use(gin.Recovery())

	cls := config.DefaultConfigLoaders()
	cfg := config.LoadConfig(cls)

	store, closeFn := repository.NewPostgresStore(&cfg)
	defer closeFn()

	container := registry.NewDefaultContainer(store)
	http.Apply(r, container)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":" + cfg.Port)
}
