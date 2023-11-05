package Repositories

import (
	"sepet/Database"
	"sepet/Models"
)

func NewCartRepository(db Database.Database) *CartRepository {
	return &CartRepository{database: db}
}

type CartRepository struct {
	database Database.Database
}

func (repo *CartRepository) GetItem(key string, keyName string) (*Models.Cart, error) {
	var cart *Models.Cart
	if err := repo.database.GetItem(key, keyName, cart); err != nil {
		return nil, err
	}
	return cart, nil
}

func (repo *CartRepository) PutItem(item Models.Cart) error {
	return repo.database.PutItem(item)
}
