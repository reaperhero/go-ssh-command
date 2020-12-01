package main

import (
	"git.mudu.tv/reaperhero/go-ssh-command/model"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	client := model.NewPasswordClientPassword("root","password","127.0.0.1")
	session,err := client.NewSession()
	if err !=nil {
		log.Info(err)
	}
	session.Stdout = os.Stdin
	if err := session.Run("ls /");err != nil{
		log.Warn(err)
	}
	return
}


