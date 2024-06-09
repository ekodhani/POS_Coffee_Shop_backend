package config

import (
	controllers "backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}

	router.Use(cors.New(config))

	//Portal Dhani
	pd := router.Group("/pos")
	{
		// Version 1
		pd_v1 := pd.Group("/v1")
		{
			pd_v1.POST("/menu", controllers.Menu)
			pd_v1.POST("/getMenuOrder", controllers.MenuOrder)
			pd_v1.POST("/kategori", controllers.Kategori)
			pd_v1.POST("/addMenu", controllers.AddMenu)
			pd_v1.POST("/editMenu", controllers.EditMenu)
			pd_v1.POST("/filterMenu", controllers.FilterMenu)
			pd_v1.POST("/deleteMenu", controllers.HapusMenu)
			pd_v1.POST("/inStockMenu", controllers.InStockMenu)
			pd_v1.POST("/addTransaksi", controllers.AddTransaksi)
			pd_v1.POST("/getHistoryTransaksi", controllers.GetDataTransaksi)
			pd_v1.POST("/getDashboard", controllers.GetDashboard)
			pd_v1.POST("/getDataChart", controllers.GetChart)
			pd_v1.POST("/cancelTransaksi", controllers.CancelTransaksi)
		}
	}
	return router
}
