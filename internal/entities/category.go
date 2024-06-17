package entites

import (
	"fmt"
	"time"

	"microservicos.com/pkg/middlewares"
	"microservicos.com/pkg/util"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Category struct {
	ID        uint
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Telefone  string `json:"telefone"`
	Cpf       string `json:"Cpf"gorm:"unique"`
	Password  string `json:"Password"`
	CreatedAt time.Time
	UpdatesAt time.Time
}

// é sempre bom fazer a regra de negocios dentro da minha entidade e não decidir o banco de dados que mande na regra de negocios
// Bom sempre que desejarem criar uma category nova eles vão utilizar nossa função não nosso type
func NewCategory(name, email, telefone, cpf, password string) (*Category, error) {
	passwordcode := util.Sha256Encoder(password)
	category := &Category{
		Name:      name,
		Email:     email,
		Telefone:  telefone,
		Cpf:       cpf,
		Password:  passwordcode,
		CreatedAt: time.Now(),
		UpdatesAt: time.Now(),
	}

	// aqui validamos se a regra de negocios está certa ou seja validamos se o nome tem mais de 5 letras
	err := category.Isvalid()
	if err != nil {
		return nil, err
	}
	// category.Password = util.Sha256Encoder(category.Password)

	// fmt.Printf("senha no banco %s", category.Password)
	// busines rules
	return category, nil
}

// vamos criar um metodo, para devolver se tem algum erro na nossa regra de negocios
func (c *Category) Isvalid() error {
	if len(c.Name) < 5 {
		return fmt.Errorf("o nome é muito pequeno não pode ter menos de 5 letras: %d", len(c.Name))
	}

	if err := middlewares.EmailVerifier(c.Name, c.Email); err != nil {
		return fmt.Errorf("o Email %s não é valido", c.Email)
	}

	if err := middlewares.TelefoneVerefication(c.Telefone); err != nil {
		return fmt.Errorf("O numero de telefone %s não é valido", c.Telefone)
	}

	if err := middlewares.Cpfverification(c.Cpf); err != nil {
		return fmt.Errorf("O cpf %s não é valido", c.Cpf)
	}

	return nil
}
