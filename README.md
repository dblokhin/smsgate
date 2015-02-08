# SMS-gateway для СМС-Центра на Golang

Простая реализация API для рассылки SMS-сообщений (http://smsc.ru/api/http/)

## Пример использования

```
sms := smsgate.New("Login", "PASSWD_or_HASH")
sms.SetSender("SenderName")
sms.Send("Message")
```
