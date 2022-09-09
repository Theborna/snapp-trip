package model

type Request interface {
	ToString() string
	PrettyString() string
}

type RequestImpl struct{}

type (
	amountType string
	status     string
)

const ( // enumms
	FIXED      amountType = "FIXED"
	PERCENTAGE amountType = "PERCENTAGE"
)
const (
	SUCCESS status = "SUCCESS"
	FAILED  status = "FAILED"
)

type (
	City     string
	Agency   string
	Supplier string
	Airline  string
)

type Route struct {
	Origin      City `json:"origin"`
	Destination City `json:"destination"`
}

type Trip struct {
	RequestImpl
	RuleId       int      `json:"ruleId"`
	Origin       City     `json:"origin"`
	Destination  City     `json:"destination"`
	Agency       Agency   `json:"agency"`
	Supplier     Supplier `json:"supplier"`
	Airline      Airline  `json:"airline"`
	BasePrice    int      `json:"basePrice"`
	MarkUp       int      `json:"markup"`
	PayablePrice int      `json:"payablePrice"`
}

type Rule struct {
	RequestImpl
	Routes      []Route    `json:"routes"`
	Airlines    []Airline  `json:"airlines"`
	Agencies    []Agency   `json:"agencies"`
	Suppliers   []Supplier `json:"suppliers"`
	AmountType  amountType `json:"amountType"`
	AmountValue int        `json:"amountValue"`
}

type Response struct {
	Message string `json:"message"`
	Status  status `json:"status"`
}
