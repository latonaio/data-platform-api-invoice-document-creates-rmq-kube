package requests

type HeaderPartner struct {
	D struct {
		PartnerFunction         string `json:"PartnerFunction"`
		BusinessPartner         int    `json:"BusinessPartner"`
		BusinessPartnerFullName string `json:"BusinessPartnerFullName"`
		BusinessPartnerName     string `json:"BusinessPartnerName"`
		Organization            string `json:"Organization"`
		Country                 string `json:"Country"`
		Language                string `json:"Language"`
		Currency                string `json:"Currency"`
		ExternalDocumentID      string `json:"ExternalDocumentID"`
		AddressID               int    `json:"AddressID"`
	} `json:"d"`
}
