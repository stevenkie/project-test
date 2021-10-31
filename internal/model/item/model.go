package item

type Item struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Quantity int32  `db:"quantity"`
	Price    int32  `db:"price"`
}
