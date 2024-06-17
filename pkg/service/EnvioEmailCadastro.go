package service

/*
import (
	"github.com/MakMoinee/go-mith/pkg/email"
	"github.com/gocraft/work"
	"microservicos.com/pkg/processors"
)

// bom coloquei essa função no service uma vez que os dados do email e o textos só podem
// ser utilizadas para uma unica funcionalidade
func (c *processors.Context) EnvEmailJob(job *work.Job) error {
	email := job.ArgString("email_address")
	name := job.ArgString("name")
	cpf := job.ArgString("cpf")

	if err := job.ArgError(); err != nil {
		return err
	}

	return service.EnvEmail(email, name, cpf)
}

// bom coloquei essa função no service uma vez que os dados do email e o textos só podem
// ser utilizadas para uma unica funcionalidade
func EnvEmail(Email string) error {
	emailService := email.NewEmailService(587, "smtp.gmail.com", "coordenacaofopag.emserh@gmail.com", "hewa gkbm mdlx mnmp")

	isSend, err := emailService.SendEmail(Email, "Test", "Email cadastrado com sucesso sr."+name+"\n"+"seu login é constituido pelo seu cpf e sua senha\nCpf: "+cpf+"\nSenha: ")
	if err != nil {
		return err
	}
	if isSend {
		return nil
	}
	return nil
}
*/
