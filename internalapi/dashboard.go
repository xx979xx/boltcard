package internalapi

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/boltcard/boltcard/db"
	log "github.com/sirupsen/logrus"
	qrcode "github.com/skip2/go-qrcode"
)

type Dashboard_data struct {
	Stats_data        map[string]int
	Cards_data        []map[string]string
	Transactions_data []map[string]string
	Withdraws_data    []map[string]string
	Deposits_data     []map[string]string
	Template_data     map[string]string
}

func make_qr_from_text(url string) {
	err := qrcode.WriteFile(url, qrcode.Medium, 256, "static/img/qr.png")
	if err != nil {
		msg := "dashboard: Can not make qr: "
		log.Warn(msg, err)
	}
}

func (dat *Dashboard_data) getdashboarddata() {
	cards, err := db.Get_all_cards()
	if err != nil {
		msg := "dashboard: Can not get cards data"
		log.Warn(msg, err)
		return
	}
	dat.Stats_data = map[string]int{
		"num_cards":              len(cards),
		"num_active_cards":       len(cards),
		"num_succeeded_payments": 0,
		"payments_total_sats":    0,
		"num_settled_deposit":    0,
		"receipts_total_sats":    0,
	}

	for _, card := range cards {
		c := map[string]string{
			"card_id":                    strconv.Itoa(card.Card_id),
			"card_name":                  card.Card_name,
			"uid":                        card.Db_uid,
			"last_counter_value":         strconv.Itoa(int(card.Last_counter_value)),
			"lnurlw_enable":              card.Lnurlw_enable,
			"lnurlw_request_timeout_sec": strconv.Itoa(card.Lnurlw_request_timeout_sec),
			"tx_limit_sats":              strconv.Itoa(card.Tx_limit_sats),
			"day_limit_sats":             strconv.Itoa(card.Day_limit_sats),
			"pin_enable":                 card.Pin_enable,
			"pin_limit_sats":             strconv.Itoa(card.Pin_limit_sats),
			"lnurlp_enable":              card.Lnurlp_enable,
			"email_enable":               card.Email_enable,
			"email_address":              card.Email_address,
			"allow_negative_balance":     card.Allow_negative_balance,
			"wiped":                      card.Wiped,
		}
		dat.Cards_data = append(dat.Cards_data, c)
		if card.Wiped == "Y" {
			dat.Stats_data["num_active_cards"]--
		}
	}

	payments, err := db.Get_all_payment()
	if err != nil {
		msg := "dashboard: Can not get withdraw data: "
		log.Warn(msg, err)
		return
	}
	dat.Stats_data["num_payments"] = len(payments)
	for _, payment := range payments {
		p := map[string]string{
			"card_payment_id":     strconv.Itoa(payment.Card_payment_id),
			"card_id":             strconv.Itoa(payment.Card_id),
			"lnurlw_k1":           payment.Lnurlw_k1,
			"lnurlw_request_time": payment.Lnurlw_request_time.Format("02/01/2006 15:04:05"),
			"ln_invoice":          payment.Ln_invoice,
			"amount_sats":         strconv.Itoa(int(payment.Amount_msats / 1000)),
			"paid_flag":           payment.Paid_flag,
			"payment_time":        payment.Payment_time.Format("02/01/2006 15:04:05"),
			"payment_status":      payment.Payment_status,
			"failure_reason":      payment.Failure_reason,
			"payment_status_time": payment.Payment_status_time.Format("02/01/2006 15:04:05"),
		}
		for _, card := range dat.Cards_data {
			if card["card_id"] == p["card_id"] {
				p["card_name"] = card["card_name"]
			}
		}
		dat.Withdraws_data = append(dat.Withdraws_data, p)
		if payment.Payment_status != "FAILED" && payment.Payment_status != "" {
			tr := map[string]string{
				"card_id":        strconv.Itoa(payment.Card_id),
				"card_name":      p["card_name"],
				"tx_id":          strconv.Itoa(payment.Card_payment_id),
				"tx_type":        "payment",
				"tx_amount_sats": strconv.Itoa(int(payment.Amount_msats / 1000)),
				"tx_time":        payment.Payment_status_time.Format("02/01/2006 15:04:05"),
			}
			dat.Transactions_data = append(dat.Transactions_data, tr)
			dat.Stats_data["num_succeeded_payments"]++
			dat.Stats_data["payments_total_sats"] += int(payment.Amount_msats / 1000)
		}
	}

	deposits, err := db.Get_all_deposit()
	if err != nil {
		msg := "dashboard: Can not get deposit data: "
		log.Warn(msg, err)
		return
	}
	dat.Stats_data["num_deposit"] = len(deposits)
	for _, deposit := range deposits {
		d := map[string]string{
			"card_receipt_id":     strconv.Itoa(deposit.Card_receipt_id),
			"card_id":             strconv.Itoa(deposit.Card_id),
			"ln_invoice":          deposit.Ln_invoice,
			"r_hash_hex":          deposit.R_hash_hex,
			"amount_sats":         strconv.Itoa(int(deposit.Amount_msat / 1000)),
			"receipt_status":      deposit.Receipt_status,
			"receipt_status_time": deposit.Receipt_status_time.Format("02/01/2006 15:04:05"),
		}
		for _, card := range dat.Cards_data {
			if card["card_id"] == d["card_id"] {
				d["card_name"] = card["card_name"]
			}
		}
		dat.Deposits_data = append(dat.Deposits_data, d)
		if deposit.Receipt_status == "SETTLED" {
			tr := map[string]string{
				"card_id":        strconv.Itoa(deposit.Card_id),
				"card_name":      d["card_name"],
				"tx_id":          strconv.Itoa(deposit.Card_receipt_id),
				"tx_type":        "payment",
				"tx_amount_sats": strconv.Itoa(int(deposit.Amount_msat / 1000)),
				"tx_time":        deposit.Receipt_status_time.Format("02/01/2006 15:04:05"),
			}
			dat.Transactions_data = append(dat.Transactions_data, tr)
			dat.Stats_data["num_settled_deposit"]++
			dat.Stats_data["receipts_total_sats"] += int(deposit.Amount_msat) / 1000
		}
	}

	dat.Template_data = map[string]string{
		"Android_app_link": "https://somelink-a",
		"Ios_app_link":     "https://somelink-b",
		"Qr_connect_link":  "https://somelink-c",
		"Version_number":   "0.0.1",
		"Settings_link":    "/settings",
	}
	make_qr_from_text(dat.Template_data["Qr_connect_link"])
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

	case "GET":
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			log.Fatal("Parse: ", err)
			return
		}
		dashboard_data := Dashboard_data{}
		dashboard_data.getdashboarddata()
		tmpl.Execute(w, dashboard_data)
		return
	default:
		return
	}

}
