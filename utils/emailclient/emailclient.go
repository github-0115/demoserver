package emailclient

import (
	cfg "ManageCenter/config"
	"fmt"

	log "github.com/inconshreveable/log15"
	"gopkg.in/gomail.v2"
)

var (
	mailUsername = cfg.Cfg.Mail.Username
	mailPassword = cfg.Cfg.Mail.Password
	mailHost     = cfg.Cfg.Mail.Host
	mailPort     = cfg.Cfg.Mail.Port
	S            gomail.SendCloser
)

func init() {

	var d = gomail.NewDialer(mailHost, mailPort, mailUsername, mailPassword)
	var err error
	S, err = d.Dial()
	if err != nil {
		log.Error(fmt.Sprintf("mail err%v", err))
	}
}

func SendPlainMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]消息推广!")
	m.SetBody("text/plain", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send plain mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendHtmlMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]消息推广!")
	m.SetBody("text/html", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send html mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendBillMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]账单生成通知!")
	m.SetBody("text/plain", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send generate the bill mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendGenerateBillHtmlMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]账单生成通知!")
	m.SetBody("text/html", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send generate the bill mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendUpdateBillHtmlMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]账单金额变动通知!")
	m.SetBody("text/html", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send update the bill mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func BusinessMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]欢迎您申请使用深图的服务")
	m.SetHeader("Content-Type", "text/html; charset=UTF-8")
	m.SetBody("text/html", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send business  email to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendPackageApplyMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]用户类型变更通知!")
	m.SetBody("text/plain", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send package apply mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendUseTotalErrMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]用户使用量错误!")
	m.SetBody("text/html", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send api use total  mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendMail(subject string, body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]"+subject)
	m.SetBody("text/html", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send api use total  mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}

func SendUseTotalMail(body string, email string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", mailUsername)
	m.SetHeader("To", email)
	m.SetHeader("subject", "[深图智服]接口调用量即将超量，请及时续费!")
	m.SetBody("text/plain", body)

	if err := gomail.Send(S, m); err != nil {
		log.Error("Couldn't send use level remind mail to" + email + ",err=" + err.Error())
		return err
	}

	return nil
}
