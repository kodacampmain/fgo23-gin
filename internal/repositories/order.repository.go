package repositories

import (
	"context"
	"fgo23-gin/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

// pembelian oleh student id 10
// product (1:coffee, 2) (3: cake, 1)
func (o *OrderRepository) CreateTransaction(ctx context.Context, studentId int, transaction models.Transaction) error {
	// awali dengan db transaksi
	tx, err := o.db.Begin(ctx)
	if err != nil {
		return err
	}
	// defer func(cause error) {
	// 	if cause:=recover()
	// }
	defer tx.Rollback(ctx)

	// 1. buat transaksi order
	transQuery := "INSERT INTO transactions (student_id) VALUES ($1) RETURNING id"
	var transactionId int
	if err := tx.QueryRow(ctx, transQuery, studentId).Scan(&transactionId); err != nil {
		return err
	}
	// 2. masukkan detail order
	detQuery := "INSERT INTO transactions_products (transaction_id, product_id, qty) VALUES "
	// building query
	values := []any{transactionId}
	for i, product := range transaction.Products {
		// (transaction_id, product_id, qty)
		if i > 0 {
			detQuery += ","
		}
		detQuery += fmt.Sprintf("($1,$%d,$%d)", len(values)+1, len(values)+2)
		values = append(values, product.ProductId, product.Quantity)
		// if i != len(transaction.Products)-1 {
		// 	detQuery += ","
		// }
	}

	// log.Println("[DEBUG] Transaction Details Query: ", detQuery)
	_, err = tx.Exec(ctx, detQuery, values...)
	if err != nil {
		return err
	}

	// Jangan lupa commit
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
