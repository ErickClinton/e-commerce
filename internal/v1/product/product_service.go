package product

import (
	"eccomerce/internal/v1/entity"
	"eccomerce/internal/v1/product/dto"
	"eccomerce/pkg/utils"
	"encoding/json"
)

type ProductService interface {
	utils.Service[dto.CreateProductRequest, entity.Product]
	UpdateById(dto *dto.UpdateProductRequest, id int) error
}

type product struct {
	repository ProductRepository
}

func NewProductService(repository ProductRepository) ProductService {
	return &product{repository: repository}

}
func (s *product) Create(product *dto.CreateProductRequest) error {
	userJSON, _ := json.MarshalIndent(product, "", "    ")
	utils.Logger.Info().Msgf("Start method create %v", string(userJSON))
	entityProduct := &entity.Product{
		Description: product.Description,
		Title:       product.Title,
		Price:       product.Price,
		UserId:      product.UserId,
	}
	return s.repository.Create(entityProduct)
}

func (s *product) GetByID(id uint) (*entity.Product, error) {
	idJSON, _ := json.MarshalIndent(id, "", "    ")
	utils.Logger.Info().Msgf("Start method GetByID %v", string(idJSON))
	return s.repository.GetByID(id)
}

func (s *product) Update(product *dto.CreateProductRequest) error {
	userJSON, _ := json.MarshalIndent(product, "", "    ")
	utils.Logger.Info().Msgf("Start method Update %v", string(userJSON))
	entityProduct := &entity.Product{
		Description: product.Description,
		Title:       product.Title,
		Price:       product.Price,
	}
	return s.repository.Update(entityProduct)
}

func (s *product) UpdateById(product *dto.UpdateProductRequest, id int) error {
	userJSON, error := json.MarshalIndent(product, "", "    ")
	if error != nil {
		return error
	}
	utils.Logger.Info().Msgf("Start method Update %v", string(userJSON))

	return s.repository.UpdateById(product, id)
}

func (s *product) Delete(id uint) error {
	idJSON, _ := json.MarshalIndent(id, "", "    ")
	utils.Logger.Info().Msgf("Start method Delete %v", string(idJSON))
	return s.repository.Delete(id)
}
