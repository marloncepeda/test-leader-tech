package user

import (
	"api/infra/database"
	"api/internal/shipment"
	"api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShipmentHandler struct {
	shipmentService *shipment.ShipmentService
}

func NewShipmentHandler() *ShipmentHandler {
	shipmentRepo := shipment.NewShipmentRepository(database.InstanceDB().GetConnect())
	shipmentService := shipment.NewShipmentService(shipmentRepo)
	//shipmentRepo and shipmentServices manager errors pending, now returns 1 value in future needs return 2 with errors
	return &ShipmentHandler{
		shipmentService: shipmentService,
	}
}

func (sh *ShipmentHandler) GetUser(c *gin.Context) {

	id := c.Param("id")

	user, err := sh.shipmentService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Response(http.StatusOK, user, ""))
}

func (sh *ShipmentHandler) CreateUser(c *gin.Context) {
	var smodel shipment.ShipmentModel

	if err := c.BindJSON(&smodel); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userCreation, _ := sh.shipmentService.CreateUser(smodel)

	c.JSON(http.StatusOK, utils.Response(http.StatusOK, userCreation, ""))

}

func (sh *ShipmentHandler) UpdateUser(c *gin.Context) {
	var smodel shipment.ShipmentModel

	if err := c.BindJSON(&smodel); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userUpdate, _ := sh.shipmentService.UpdateUser(smodel)

	c.JSON(http.StatusOK, utils.Response(http.StatusOK, userUpdate, ""))
}
