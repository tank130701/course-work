package todo_item

const (
	createItemQuery      = `INSERT INTO todo_items (title, description) values ($1, $2) RETURNING id`
	createListItemsQuery = `INSERT INTO todo_items (list_id, item_id) values ($1, $2)`
	getAllQuery          = `SELECT ti.id, ti.title, ti.description, ti.done 
								FROM todo_items ti INNER JOIN lists_items li on li.item_id = ti.id
								INNER JOIN users ul on ul.list_id = li.list_id 
							WHERE li.list_id = $1 AND ul.user_id = $2`
	getByIdQuery = `SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`
	deleteQuery = `DELETE FROM items ti USING itemsList li, users ul 
					WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`
)
