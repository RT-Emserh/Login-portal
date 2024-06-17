package processors

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/MakMoinee/go-mith/pkg/email"
	"github.com/gocraft/work"
	"microservicos.com/pkg/util"
)

type context struct {
	email string
	name  string
	cpf   string
}

func Processors() {
	pool := work.NewWorkerPool(context{}, 30, "application_namespace", util.RedisPool)
	fmt.Println("Iniciando pool de workers")

	// Adiciona middlewares que serão executados para cada trabalho, se necessário

	// Mapeia o nome dos trabalhos para funções de manipulador
	pool.Job("EnvEmail", (*context).EnvEmailJob)

	// Inicia o processamento de trabalhos
	pool.Start()

	// Mantém o pool em execução até receber um sinal para parar
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Para o pool
	pool.Stop()
	fmt.Println("Pool de workers parado")
}

func EnqueueEmailJob(em, nam, cp string) error {
	enqueuer := work.NewEnqueuer("application_namespace", util.RedisPool)
	_, err := enqueuer.Enqueue("EnvEmail", work.Q{"email_address": em, "user_name": nam, "user_matricula": cp})
	return err
}

func (c *context) EnvEmailJob(job *work.Job) error {
	fmt.Println("inicio do envio de email")
	Email := job.ArgString("email_address")
	Name := job.ArgString("user_name")
	CPF := job.ArgString("user_matricula")

	log.Println(Email)

	emailService := email.NewEmailService(587, "smtp.gmail.com", "coordenacaofopag.emserh@gmail.com", "hewa gkbm mdlx mnmp")

	isSend, err := emailService.SendEmail(Email, "Test", "Email cadastrado com sucesso sr(a)."+Name+"\n"+"seu login é constituído pela sua matricula e sua senha\nCpf: "+CPF+"\nSenha: ")

	log.Println("fim processo")
	if err != nil {
		log.Println("error")
		return err
	}
	if isSend {
		return nil
	}
	return nil
}
