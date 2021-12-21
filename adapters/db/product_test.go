package db_test

import (
	"database/sql"
	"github.com/kennymack/go-hexagonal-product/adapters/db"
	"github.com/kennymack/go-hexagonal-product/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
                ID string, 
                NAME string, 
                STATUS string, 
                PRICE float
              );`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, _ = stmt.Exec()

}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products (ID, NAME, STATUS, PRICE)
                             VALUES ('abc1', 'Produto 1', 'enabled', 10);`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, _ = stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()

	defer Db.Close()

	productDB := db.NewProductDb(Db)

	product, err := productDB.Get("abc1")

	require.Nil(t, err)

	require.Equal(t, "Produto 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()

	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Produto 2"
	product.Status = application.ENABLED
	product.Price = 25

	productResult, err := productDb.Save(product)

	require.Nil(t, err)

	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

	product.Status = application.DISABLED

	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())
}



