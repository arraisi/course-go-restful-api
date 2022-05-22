package category

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	Save(ctx context.Context, request Category) (tx *sql.Tx, category Category, err error)
	Update(ctx context.Context, request Category) (tx *sql.Tx, category Category, err error)
	Delete(ctx context.Context, id int) (tx *sql.Tx, err error)
	FindById(ctx context.Context, id int) (category Category, err error)
	FindAll(ctx context.Context) (categories []Category, err error)
}

type repository struct {
	DB *sql.DB
}

func NewRepository(DB *sql.DB) *repository {
	return &repository{DB: DB}
}

func (r *repository) Save(ctx context.Context, request Category) (tx *sql.Tx, category Category, err error) {
	SQL := "INSERT INTO category (name) VALUES($1) RETURNING id"
	tx, err = r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return tx, request, err
	}

	statement, err := tx.PrepareContext(ctx, SQL)
	if err != nil {
		fmt.Printf("Query insert gwp_campaigns err:%v", err)
		return tx, request, err
	}

	err = statement.QueryRowContext(ctx, request.Name).Scan(&request.Id)
	if err != nil {
		fmt.Printf("Execute query insert category err:%v", err)
		return tx, request, err
	}

	return tx, request, nil
}

func (r *repository) Update(ctx context.Context, request Category) (tx *sql.Tx, category Category, err error) {
	SQL := "UPDATE category SET name=$1 WHERE id=$2"
	tx, err = r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return tx, request, err
	}

	statement, err := tx.PrepareContext(ctx, SQL)
	if err != nil {
		fmt.Printf("Query update category err:%v", err)
		return tx, request, err
	}

	_, err = statement.ExecContext(ctx, request.Name, request.Id)
	if err != nil {
		fmt.Printf("Execute update category err:%v", err)
		return tx, request, err
	}

	return tx, request, nil
}

func (r *repository) Delete(ctx context.Context, id int) (tx *sql.Tx, err error) {
	SQL := "DELETE FROM category WHERE id=$1"
	tx, err = r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return tx, err
	}

	statement, err := tx.PrepareContext(ctx, SQL)
	if err != nil {
		fmt.Printf("Query delete category err:%v", err)
		return tx, err
	}

	_, err = statement.ExecContext(ctx, id)
	if err != nil {
		fmt.Printf("Execute delete category err:%v", err)
		return tx, err
	}

	return tx, err
}

func (r *repository) FindById(ctx context.Context, id int) (category Category, err error) {
	SQL := "SELECT id, name FROM category WHERE id=$1"
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		fmt.Printf("Query find by id category err:%v", err)
		return category, err
	}

	err = tx.QueryRowContext(ctx, SQL, id).Scan(&category.Id, &category.Name)
	if err != nil {
		fmt.Printf("Execute find by id category err:%v", err)
		return category, err
	}

	return category, err
}

func (r *repository) FindAll(ctx context.Context) (categories []Category, err error) {
	SQL := "SELECT id, name FROM category ORDER BY id"
	statement, err := r.DB.PrepareContext(ctx, SQL)
	if err != nil {
		fmt.Printf("Query find all category err:%v", err)
		return categories, err
	}

	rows, err := statement.QueryContext(ctx)
	if err != nil {
		fmt.Printf("Execute find all category err:%v", err)
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var data Category
		err := rows.Scan(&data.Id, &data.Name)
		if err != nil {
			return categories, err
		}
		categories = append(categories, data)
	}

	if err = rows.Err(); err != nil {
		return categories, err
	}

	return categories, err
}
