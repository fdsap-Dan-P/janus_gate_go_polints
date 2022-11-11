package controller

import (
	"JanusGate/middleware/go-utils/database"
	"JanusGate/struct/response"

	"github.com/gofiber/fiber/v2"
)

// Developer			Roldan
// UponOpen	godoc
// @summary 	  		kPLUS UPON OPEN
// @Description	  		Provide the data that will be used by kPLUS upon opening the application
// @Tags		  		kPLUS
// @Produce		  		json
// @Success		  		200 {object} response.KplusUponOpenResponse
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/kplus/kplus_upon_open [get]
func KplusUponOpen(c *fiber.Ctx) error {

	UponOpenResponse := []response.KplusUponOpenResponse{}
	// TryResponse := []response.TryResponse{}

	if fetchErr := database.DBConn.Debug().Table("c_splash_screen").Where("show = 'true'").Find(&UponOpenResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch data",
			Data:    fetchErr,
		})
	}
	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    database.DBConn.Debug().Table("c_splash_screen").Find(&UponOpenResponse),
	})

	// // database.DBConn.Debug().Table("try").Create(&UponOpenResponse)
	// // database.DBConn.Debug().Table("try").Find(&TryResponse)

	// return c.JSON(response.ResponseModel{
	// 	RetCode: "200",
	// 	Message: "success",
	// 	Data:   TryResponse,
	// })
}
