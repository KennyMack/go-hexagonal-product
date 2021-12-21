package db

import (
	"database/sql"
	"github.com/kennymack/go-hexagonal-product/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("SELECT ID, NAME, PRICE, STATUS FROM products WHERE ID = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`INSERT INTO products (ID, NAME, PRICE, STATUS) 
                                                   VALUES (?, ? , ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`UPDATE products 
                                        SET NAME = ?, 
                                            PRICE = ?, 
                                            STATUS = ?
                                      WHERE ID = ?`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetId())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	result, err := p.Get(product.GetId())

	if result == nil {
		_, err = p.create(product)

		if err != nil {
			return nil, err
		}
	} else {
		_, err = p.update(product)

		if err != nil {
			return nil, err
		}
	}

	return product, nil
}
