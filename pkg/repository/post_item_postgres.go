package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	ginger_api "github.com/petrrnnt/ginger-api"
	"strings"
)

type PostItemPostgres struct {
	db *sqlx.DB
}

func NewPostItemPostgres(db *sqlx.DB) *PostItemPostgres {
	return &PostItemPostgres{db: db}
}

func (r *PostItemPostgres) Create(listId int, item ginger_api.PostItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", postItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *PostItemPostgres) GetAll(userId int, listId int) ([]ginger_api.PostItem, error) {
	var items []ginger_api.PostItem
	query := fmt.Sprintf(`SELECT pi.id, pi.title, pi.description, pi.done FROM %s pi INNER JOIN %s li on li.item_id = pi.id 
		INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		postItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *PostItemPostgres) GetById(userId int, itemId int) (ginger_api.PostItem, error) {
	var item ginger_api.PostItem
	query := fmt.Sprintf(`SELECT pi.id, pi.title, pi.description, pi.done FROM %s pi INNER JOIN %s li on li.item_id = pi.id 
		INNER JOIN %s ul on ul.list_id = li.list_id WHERE pi.id = $1 AND ul.user_id = $2`,
		postItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *PostItemPostgres) Delete(userId int, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s pi USING %s li, %s ul
									WHERE pi.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND pi.id = $2`,
		postItemsTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *PostItemPostgres) Update(userId int, itemId int, input ginger_api.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s pi SET %s FROM %s li, %s ul 
                    WHERE pi.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND pi.id = $%d`,
		postItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}
