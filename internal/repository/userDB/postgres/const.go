package postgres

const (
	getUserByEmail = `
	SELECT * FROM users WHERE email = $1 and is_active = true
	`
	getUserByID = `
	SELECT * FROM users WHERE id = $1 and is_active = true
	`
	insertUser = `
	INSERT INTO users(email, address, password)
	VALUES ($1,$2,$3)
	`
	updateUser = `
	UPDATE users SET 
		email = $2,
		address = $3
	WHERE id = $1
	`
	deleteUser = `
	UPDATE users SET 
		is_active = false
	WHERE id = $1
	`
)
