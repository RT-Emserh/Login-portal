package middlewares

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func TelefoneVerefication(telefone string) error {
	d := telefone //"98984369949"
	num, err := phonenumbers.Parse(d, "BR")
	if err != nil {
		fmt.Println("Erro ao analisar o número de telefone:", err)
		return err
	}

	// Verifique se o número é válido
	if !phonenumbers.IsValidNumber(num) {
		fmt.Println("O número de telefone não é válido.")
		return fmt.Errorf("não é valido")
	}

	fmt.Println("O número de telefone é válido.")
	return nil
}
