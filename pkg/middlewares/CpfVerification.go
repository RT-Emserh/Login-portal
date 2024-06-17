package middlewares

import "fmt"

var cpfList = []string{"62020735350", "62020735351"}

func Cpfverification(cpf string) error {
	for _, testcpf := range cpfList {
		if cpf == testcpf {
			return nil
		}
	}
	return fmt.Errorf("o cpf não é valido")
}
