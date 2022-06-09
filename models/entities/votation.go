package entities

type Votation struct {
	VotationId     int  `json:"votation_id" gorm:"primary_key;auto_increment"`
	User           User `json:"user"`
	UserId         int  `json:"user_id" gorm:"ForeignKey:UserId"`
	Estetica       int  `json:"estetica"`
	Vulnerabilidad int  `json:"vulnerabilidad"`
	Funcionalidad  int  `json:"funcionalidad"`
	ControlErrores int  `json:"controlErrores"`
	Rendimiento    int  `json:"rendimiento"`
}
