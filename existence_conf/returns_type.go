package existence_conf

type Returns struct {
	ConnectionKey                                string                                        `json:"connection_key"`
	Result                                       bool                                          `json:"result"`
	RedisKey                                     string                                        `json:"redis_key"`
	RuntimeSessionID                             string                                        `json:"runtime_session_id"`
	BusinessPartnerID                            *int                                          `json:"business_partner"`
	Filepath                                     string                                        `json:"filepath"`
	ServiceLabel                                 string                                        `json:"service_label"`
	SupplyChainRelationshipGeneralReturn         *SupplyChainRelationshipGeneralReturn         `json:"SupplyChainRelationshipGeneral,omitempty"`
	SupplyChainRelationshipBillingRelationReturn *SupplyChainRelationshipBillingRelationReturn `json:"SupplyChainRelationshipBillingRelation,omitempty"`
	SupplyChainRelationshipPaymentRelationReturn *SupplyChainRelationshipPaymentRelationReturn `json:"SupplyChainRelationshipPaymentRelation,omitempty"`
	IncotermsReturn                              IncotermsReturn                               `json:"IncotermsReturn"`
	PaymentTermsReturn                           PaymentTermsReturn                            `json:"PaymentTerms"`
	PaymentMethodReturn                          PaymentMethodReturn                           `json:"PaymentMethod"`
	APISchema                                    string                                        `json:"api_schema"`
	Accepter                                     []string                                      `json:"accepter"`
	Deleted                                      bool                                          `json:"deleted"`
}

type SupplyChainRelationshipGeneralReturn struct {
	SupplyChainRelationshipID int `json:"SupplyChainRelationshipID"`
	Buyer                     int `json:"Buyer"`
	Seller                    int `json:"Seller"`
}

type SupplyChainRelationshipBillingRelationReturn struct {
	SupplyChainRelationshipID        int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int `json:"Buyer"`
	Seller                           int `json:"Seller"`
	BillToParty                      int `json:"BillToParty"`
	BillFromParty                    int `json:"BillFromParty"`
}

type SupplyChainRelationshipPaymentRelationReturn struct {
	SupplyChainRelationshipID        int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int `json:"Buyer"`
	Seller                           int `json:"Seller"`
	BillToParty                      int `json:"BillToParty"`
	BillFromParty                    int `json:"BillFromParty"`
	Payer                            int `json:"Payer"`
	Payee                            int `json:"Payee"`
}

type IncotermsReturn struct {
	Incoterms string `json:"Incoterms"`
}

type PaymentTermsReturn struct {
	PaymentTerms string `json:"PaymentTerms"`
}

type PaymentMethodReturn struct {
	PaymentMethod string `json:"PaymentMethod"`
}
