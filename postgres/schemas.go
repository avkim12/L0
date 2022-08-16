package postgres

const createSchema = `
CREATE TABLE IF NOT EXISTS posts
(
	id SERIAL PRIMARY KEY,
	model JSON
)
`

var insertOrderSchema = `
INSERT INTO posts(id, model) VALUES($1,$2) RETURNING id
`

var selectOrderSchema = `
SELECT * FROM orders WHERE id = $1
`