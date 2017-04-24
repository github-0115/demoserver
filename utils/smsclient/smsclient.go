package smsclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/inconshreveable/log15"
)

const rawUrl = "http://api.synhey.com/json?key=dqvtnz&secret=4wwXBZ6b&"

var (
	checksuccessTemplate = "尊敬的 %s，您的账户已经通过审核，您可以登录开发者平台使用密钥进行开发接入了。"
	updateBillTemplate   = "尊敬的 %s 用户，您 %s 的账单金额有发生变动，请尽快到开发者平台处理，以免影响正常使用。"
	generateBillTemplate = "尊敬的 %s 用户，您 %s 的API使用情况及费用已产生，请尽快到开发者平台关注并处理，以免影响正常使用。"
	billnotpassTemplate  = "尊敬的 %s，我们无法核实您的结算信息，请确保填写的信息正确。"
	packageapplyTemplate = "尊敬的 %s，您的账户已经成功变更成 %s，有任何问题，请与我们联系：bd@deepir.com"
)

func SendVerifySMS(phone string, username string) string {
	var msg = fmt.Sprintf(checksuccessTemplate, username)
	log.Info("msg = " + msg)

	smsUrl := rawUrl + "to=" + phone + "&text=" + url.QueryEscape(msg)
	log.Info("smsUrl = " + smsUrl)
	log.Info("after escape, smsUrl = " + smsUrl)

	resp, err := http.Get(smsUrl)
	if err != nil {
		log.Error("http get failed." + err.Error())
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp read error." + err.Error())
		return err.Error()
	}

	log.Info(string(body))
	return ""
}

func SendGenerateBillSMS(phone string, username string, timeStr string) string {
	var msg = fmt.Sprintf(generateBillTemplate, username, timeStr)
	log.Info("msg = " + msg)

	smsUrl := rawUrl + "to=" + phone + "&text=" + url.QueryEscape(msg)
	log.Info("smsUrl = " + smsUrl)
	log.Info("after escape, smsUrl = " + smsUrl)

	resp, err := http.Get(smsUrl)
	if err != nil {
		log.Error("http get failed." + err.Error())
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp read error." + err.Error())
		return err.Error()
	}

	log.Info(string(body))
	return ""
}

func SendUpdateBillSMS(phone string, username string, timeStr string) string {
	var msg = fmt.Sprintf(updateBillTemplate, username, timeStr)
	log.Info("msg = " + msg)

	smsUrl := rawUrl + "to=" + phone + "&text=" + url.QueryEscape(msg)
	log.Info("smsUrl = " + smsUrl)
	log.Info("after escape, smsUrl = " + smsUrl)

	resp, err := http.Get(smsUrl)
	if err != nil {
		log.Error("http get failed." + err.Error())
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp read error." + err.Error())
		return err.Error()
	}

	log.Info(string(body))
	return ""
}

func SendNotPassBillSMS(phone string, username string) string {
	var msg = fmt.Sprintf(billnotpassTemplate, username)
	log.Info("msg = " + msg)

	smsUrl := rawUrl + "to=" + phone + "&text=" + url.QueryEscape(msg)
	log.Info("smsUrl = " + smsUrl)
	log.Info("after escape, smsUrl = " + smsUrl)

	resp, err := http.Get(smsUrl)
	if err != nil {
		log.Error("http get failed." + err.Error())
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp read error." + err.Error())
		return err.Error()
	}

	log.Info(string(body))
	return ""
}

func SendPackageSMS(phone string, username string, utype string) string {
	var msg = fmt.Sprintf(packageapplyTemplate, username, utype)
	log.Info("msg = " + msg)

	smsUrl := rawUrl + "to=" + phone + "&text=" + url.QueryEscape(msg)
	log.Info("smsUrl = " + smsUrl)
	log.Info("after escape, smsUrl = " + smsUrl)

	resp, err := http.Get(smsUrl)
	if err != nil {
		log.Error("http get failed." + err.Error())
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp read error." + err.Error())
		return err.Error()
	}

	log.Info(string(body))
	return ""
}
