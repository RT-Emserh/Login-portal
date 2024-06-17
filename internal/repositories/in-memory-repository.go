package repositories

import (
	"gorm.io/gorm"
	"microservicos.com/cmd/config"
	entites "microservicos.com/internal/entities"
)

type CategoryRepository struct {
	db *gorm.DB
}

// bom aqui nessa parte do repositorio vou colocar a conexão do banco de dados, assim o codigo em si não precisa
// se importa com que banco estamos nós conectando

// NewCategoryRepository cria uma nova instância do repositório usando a conexão GORM
func NewInMemoryCategoryRepository() *CategoryRepository {
	db := config.GetDatabase()
	db.AutoMigrate(&entites.Category{})
	return &CategoryRepository{
		db: db,
	}
}

// Save salva uma categoria no banco de dados
func (r *CategoryRepository) Save(category *entites.Category) error {
	return r.db.Create(category).Error
}

// FindAll recupera todas as categorias do banco de dados
func (r *CategoryRepository) FindAll() ([]*entites.Category, error) {
	var categories []*entites.Category
	err := r.db.Find(&categories).Error
	return categories, err
}
