package util

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Encoder(s string) string {
	fmt.Println(s)
	// isso é uma biblioteca do golang a qual quando colocamos a senha envez da senha aparecer no banco como uma string normal ela aparece como uma string256
	str := sha256.Sum256([]byte(s))
	fmt.Print(str)
	return fmt.Sprintf("%x", str)
}
