package requests

type ItemPartner struct {
	D struct {
		PartnerFunction string `json:"PartnerFunction"`
		BusinessPartner int    `json:"BusinessPartner"`
	} `json:"d"`
}
