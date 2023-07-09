package shipment

type ShipmentPortRepositories interface {
	GetShipment(id string) ([]ShipmentModel, error)
	CreateShipment(user ShipmentModel) ([]ShipmentModel, error)
	UpdateShipment(params ShipmentModel) ([]ShipmentModel, error)
}
