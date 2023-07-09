package shipment

type ShipmentModel struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	salt     []byte
	hash     []byte
}
