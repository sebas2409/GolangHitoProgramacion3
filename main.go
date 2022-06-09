package main

import (
	"GolangHitoProgramacion3/db"
	"GolangHitoProgramacion3/models/entities"
	"GolangHitoProgramacion3/models/structure"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func main() {

	viper.SetConfigFile("./envs/.env")
	viper.ReadInConfig()
	var dbUrl = viper.Get("DB_URL").(string)
	var database = db.Init(dbUrl)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var votaciones []entities.Votation
		var votacion []structure.Campos

		result, _ := database.Select("estetica", "vulnerabilidad", "funcionalidad", "control_errores", "rendimiento").Find(&votaciones).Rows()
		for result.Next() {
			database.ScanRows(result, &votacion)
		}
		c.JSON(http.StatusOK, votacion)
	})

	r.POST("/adduser", func(c *gin.Context) {
		body := entities.Votation{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var vote entities.Votation
		vote.User = body.User
		vote.Estetica = body.Estetica
		vote.ControlErrores = body.ControlErrores
		vote.Funcionalidad = body.Funcionalidad
		vote.Rendimiento = body.Rendimiento
		vote.Vulnerabilidad = body.Vulnerabilidad

		fmt.Println(vote)
		if result := database.Create(&vote); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		c.JSON(http.StatusOK, &vote)
	})

	if err := r.Run(); err != nil {
		return
	}
}
