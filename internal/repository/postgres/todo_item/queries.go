package todo_item

const (
	createTaskQuery = `INSERT INTO tasks (title, description, status, user_id, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	getAllQuery     = `SELECT tasks.id, tasks.title, tasks.description, tasks.status
								FROM tasks
							WHERE user_id = $1 AND tasks.category_id = $2`
	getByIdQuery = `SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`
	deleteQuery = `DELETE FROM tasks WHERE tasks.id = $1`
)
