package smsgate

// 7.02.2015 23:02
// (c) Dmitriy Blokhin (sv.dblokhin@gmail.com), www.webjinn.ru

import (
	"net/url"
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"errors"
)

var (
	ErrNoPhones = errors.New("Phones doesnt exists")
)

type TRANSLIT_MODE int
type CHARSET_MODE string

const (
	TRANSLIT_NO TRANSLIT_MODE = iota
	TRANSLIT_F1 // -> translit
	TRANSLIT_F2 // -> mpaHc/Ium

	CHARSET_UTF8 CHARSET_MODE    = "utf-8"
	CHARSET_WINDOWS CHARSET_MODE = "windows-1251"
	CHARSET_KOI8 CHARSET_MODE    = "koi8-r"
)

type TSMSCenter struct {
	login    string
	hash     string
	phones   []string
	sender   string
	translit TRANSLIT_MODE
	charset  CHARSET_MODE
	valid    int // Hours of live time of sms

}

func New(login, pass string) *TSMSCenter {
	return &TSMSCenter {
		login: url.QueryEscape(login),
		hash: url.QueryEscape(pass),
		translit: TRANSLIT_NO,
		charset: CHARSET_UTF8,
		valid: 1,
	}
}

func (smsc *TSMSCenter) Send(msg string) ([]byte, error) {
	if len(smsc.phones) == 0 {
		return nil, ErrNoPhones
	}
	recvs := url.QueryEscape(strings.Join(smsc.phones, ","))

	command := fmt.Sprintf("http://smsc.ru/sys/send.php?login=%s&psw=%s&phones=%s&mes=%s&sender=%s&charset=%s&translit=%d&valid=%d",
		smsc.login, smsc.hash, recvs, url.QueryEscape(msg), smsc.sender, smsc.charset, smsc.translit, smsc.valid)

	resp, err := http.Get(command)
	if err != nil {
		return nil, err
	}
	
	return ioutil.ReadAll(resp.Body)
}

func (smsc *TSMSCenter) AddPhones(phones []string) {
	smsc.phones = append(smsc.phones, phones...)
}

func (smsc *TSMSCenter) AddPhone(phone string) {
	smsc.phones = append(smsc.phones, phone)
}

func (smsc *TSMSCenter) SetSender(s string) {
	smsc.sender = url.QueryEscape(s)
}

func (smsc *TSMSCenter) Translit(mode TRANSLIT_MODE) {
	smsc.translit = mode
}

func (smsc *TSMSCenter) Charset(mode CHARSET_MODE) {
	smsc.charset = mode
}

func (smsc *TSMSCenter) Valid(hours int) {
	smsc.valid = hours
}

func Send(login, hash, sender, msg string, phones[]string) ([]byte, error) {
	login = url.QueryEscape(login)
	msg = url.QueryEscape(msg)
	sender = url.QueryEscape(sender)
	receivers := url.QueryEscape(strings.Join(phones, ","))


	command := fmt.Sprintf("http://smsc.ru/sys/send.php?login=%s&psw=%s&phones=%s&mes=%s&sender=%s&charset=utf-8",
		login, hash, receivers, msg, sender)

	resp, err := http.Get(command)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
