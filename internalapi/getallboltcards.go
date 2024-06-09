package internalapi

import (
	"net/http"
	"strconv"

	"github.com/boltcard/boltcard/db"
	"github.com/boltcard/boltcard/resp_err"
	log "github.com/sirupsen/logrus"
)

func Getallboltcards(w http.ResponseWriter, r *http.Request) {
	if db.Get_setting("FUNCTION_INTERNAL_API") != "ENABLE" {
		msg := "getboltcard: internal API function is not enabled"
		log.Debug(msg)
		resp_err.Write_message(w, msg)
		return
	}

	// log the request
	log.Info("getallboltcards API request")

	// get the card record

	s, err := db.Get_all_cards()
	if err != nil {
		msg := "getallboltcard: No cards exist in the database"
		log.Warn(msg)
		resp_err.Write_message(w, msg)
		return
	}

	count := strconv.Itoa(len(s))
	arr := "["
	for i, c := range s {
		if i != 0 {
			arr += ","
		}
		arr += `{"card_id": "` + strconv.Itoa(c.Card_id) + `",` + `"card_name": "` + c.Card_name + `",` +
			`"uid": "` + c.Db_uid + `",` + `"last_counter_value": "` + strconv.Itoa(int(c.Last_counter_value)) + `",` +
			`"lnurlw_enable": "` + c.Lnurlw_enable + `",` +
			`"lnurlw_request_timeout_sec": "` + strconv.Itoa(c.Lnurlw_request_timeout_sec) + `",` +
			`"tx_limit_sats": "` + strconv.Itoa(c.Tx_limit_sats) + `",` +
			`"day_limit_sats": "` + strconv.Itoa(c.Day_limit_sats) + `", ` +
			`"lnurlp_enable": "` + c.Lnurlp_enable + `",` +
			`"Allow_negative_balance": "` + c.Allow_negative_balance + `",` +
			`"Email_enable": "` + c.Email_enable + `",` + `"Email_address": "` + c.Email_address + `",` +
			`"pin_enable": "` + c.Pin_enable + `", ` + `"pin_limit_sats": "` + strconv.Itoa(c.Pin_limit_sats) + `",` +
			`"wiped": "` + c.Wiped + `"}`
	}
	arr += "]"

	jsonData := []byte(`{"status":"OK",` + `"cards_count": "` + count + `","all_cards":` + arr + `}`)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
