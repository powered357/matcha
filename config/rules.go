package config

import ()

const (
	DB_HOST   = "localhost"
	DB_NAME   = "matcha_db"
	DB_USER   = "bsabre"
	DB_PASSWD = "23"
	DB_TYPE   = "postgres"

	MAIL_FROM   = "bsabre.cat@gmail.com"
	MAIL_HOST   = "smtp.gmail.com"
	MAIL_PASSWD = "den619392"

	PASSWD_MIN_LEN    = 6
	MAIL_MAX_LEN      = 30
	MAIL_MIN_LEN      = 6
	NAME_MAX_LEN      = 30
	BIOGRAPHY_MAX_LEN = 300

	NOTIF_MAX_LEN   = 250
	MESSAGE_MAX_LEN = 300
	PHOTO_MAX_LEN   = 300
	DEVICE_MAX_LEN  = 150

	RED       = "\033[31m"
	GREEN     = "\033[32m"
	YELLOW    = "\033[33m"
	BLUE      = "\033[34m"
	RED_BG    = "\033[41;30m"
	GREEN_BG  = "\033[42;30m"
	YELLOW_BG = "\033[43;30m"
	BLUE_BG   = "\033[44;30m"
	NO_COLOR  = "\033[m"
)

type User struct {
	Uid         int    `json:"uid"`
	Mail        string `json:"mail,,omitempty"`
	Passwd      string `json:"-"`
	Fname       string `json:"fname"`
	Lname       string `json:"lname"`
	Age         int    `json:"age,,omitempty"`
	Gender      string `json:"gender,,omitempty"`
	Orientation string `json:"orientation,,omitempty"`
	Biography   string `json:"biography,,omitempty"`
	AvaPhotoID  int    `json:"avaPhotoID,,omitempty"`
	AccType     string `json:"-"`
	Rating      int    `json:"rating"`
}

type Notif struct {
	Nid         int    `json:"nid"`
	UidSender   int    `json:"uidSender"`
	UidReceiver int    `json:"uidReceiver"`
	Body        string `json:"body"`
}

type Message struct {
	Mid         int    `json:"nid"`
	UidSender   int    `json:"uidSender"`
	UidReceiver int    `json:"uidReceiver"`
	Body        string `json:"body"`
}

type Device struct {
	Id          int    `json:"id"`
	Uid  		int    `json:"uid"`
	Device      string `json:"device"`
}
