package postgres

var insertOrderSchema = `
INSERT INTO posts(uid, model) VALUES($1,$2) RETURNING uid
`

var selectOrderSchema = `
SELECT * FROM orders WHERE uid = $1
`