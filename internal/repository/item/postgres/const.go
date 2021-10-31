package postgres

const (
	getItemsByIDs = `
	SELECT * FROM items WHERE id = any($1)
	`
	reduceItemStockByID = `
	UPDATE items SET 
		quantity = quantity - $2
	WHERE id = $1
	`
)
