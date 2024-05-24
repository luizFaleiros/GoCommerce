package database

import (
	"github.com/luizFaleiros/GoCommerce/internal/entity"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func (p ProductDB) Update(product *entity.Product) (*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductDB) DeleteById(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewProductDb(db *gorm.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (p ProductDB) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p ProductDB) FindById(id string) (*entity.Product, error) {
	var product *entity.Product
	err := p.DB.Where("id = ?", id).First(&product).Error
	return product, err
}

func (p ProductDB) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p ProductDB) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	var products []*entity.Product
	var err error
	if sort == "" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
		if err != nil {
			return nil, err
		}
		return products, nil
	}
	err = p.DB.Order("created_at " + sort).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
