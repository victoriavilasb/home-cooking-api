package domain

type Grocery struct {
	Name         string
	Type         string
	PurchaseDate string
	// esse dia de vencimento pode ser gerado atraves do dia compra no momento de inserção no banco
	// vindo da api. Exemplo, digamos que é tipo é tomate e dura tanto tempo, aí preenchemos o due date
	DueDate      string
	IsPerishable bool
	Quantity     Quantity
}

type Quantity struct {
	Value float64
	Type  string
}
