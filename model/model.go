package model

type Request interface {
	ToString() string
	PrettyString() string
}

type RequestImpl struct{}

type Route struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type Trip struct {
	RequestImpl
	RuleId      int    `json:"ruleId"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	// Route     Route   `json:"route"`
	Agency       string `json:"agency"`
	Supplier     string `json:"supplier"`
	Airline      string `json:"airline"`
	BasePrice    int    `json:"basePrice"`
	MarkUp       int    `json:"markup"`
	PayablePrice int    `json:"payablePrice"`
}

type amountType string

const (
	FIXED      amountType = "FIXED"
	PERCENTAGE            = "PERCENTAGE"
)

type Rule struct {
	RequestImpl
	Routes      []Route    `json:"routes"`
	Airlines    []string   `json:"airlines"`
	Agencies    []string   `json:"agencies"`
	Suppliers   []string   `json:"suppliers"`
	AmountType  amountType `json:"amountType"`
	AmountValue int        `json:"amountValue"`
}

type status string

const (
	SUCCESS status = "SUCCESS"
	FAILED         = "FAILED"
)

type Response struct {
	Message string `json:"message"`
	Status  status `json:"status"`
}
