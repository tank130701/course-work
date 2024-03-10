package auth

const (
	createUserQuery = `insert into users (
                   username,
                   password_hash
                   ) values ($1, $2, $3)
                    RETURNING id`
	getUserQuery = `SELECT id FROM users WHERE username = $1 AND password_hash = $2`
)
