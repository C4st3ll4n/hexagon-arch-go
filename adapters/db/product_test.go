package db_test

import (
	"database/sql"
	"github.com/c4st3ll4n/go-hexagon/adapters/db"
	"github.com/c4st3ll4n/go-hexagon/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := "CREATE TABLE products(id varchar(255), name varchar(255), price float, status varchar(50))"

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product", 0, "disabled")`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product 2"
	product.Price = 10

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, productResult.GetName(), product.GetName())
	require.Equal(t, productResult.GetPrice(), product.GetPrice())
	require.Equal(t, productResult.GetStatus(), product.GetStatus())

	product.Status = "enabled"
	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, productResult.GetStatus(), product.GetStatus())

}
