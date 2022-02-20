package facebook

import "errors"

type FacebookData struct {
	IdFacebookData     int    `json:"id_facebook_data"`
	FacebookPageAdress string `json:"facebook_page_adress"`
	FacebookPageId     string `json:"facebook_page_id"`
	FacebookToken      string `json:"facebook_token"`
	FacebookUserToken  string `json:"unique_client_id"`
	ActiveStatus       bool   `json:"active_status"`
}

func NewFacebookData() *FacebookData {
	return &FacebookData{}
}

func (f *FacebookData) IsValid() error {
	if f.FacebookPageAdress == "" || f.FacebookPageId == "" || f.FacebookToken == "" || f.FacebookUserToken == "" {
		return errors.New("Invalid Facebook data")
	}

	return nil
}
