package todo_categories

const (
	createCategoryQuery = "INSERT INTO categories (user_id, name) VALUES ($1, $2) RETURNING id"
	getAllQuery         = "SELECT c.id, c.name, c.user_id FROM categories c WHERE c.user_id = $1"
	getById             = "SELECT c.id, c.name, c.user_id FROM categories c WHERE c.id = $1"
	deleteQuery         = "DELETE FROM categories WHERE user_id = $1 AND id = $2"
)
