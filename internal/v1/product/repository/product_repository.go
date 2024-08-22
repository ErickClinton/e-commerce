package repository

import (
	"eccomerce/internal/v1/entity"
	"eccomerce/internal/v1/product/dto"
	"eccomerce/pkg/utils"
	"gorm.io/gorm"
)

type ProductRepository interface {
	utils.Repository[entity.Product]
	UpdateById(updateDto *dto.UpdateProductRequest, id int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetByID(id uint) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *entity.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) UpdateById(updateProductDto *dto.UpdateProductRequest, id int) error {
	return r.db.Model(&entity.Product{}).Where("id = ?", id).Updates(updateProductDto).First(&updateProductDto, id).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}
