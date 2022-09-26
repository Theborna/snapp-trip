package model

import (
	"database/sql/driver"
)

// interfaces
type Request interface {
	ToString() string
	PrettyString() string
}

type NamedData interface {
	validate() error
}

type RequestImpl struct{}

// enums
type (
	amountType string
	status     string
)

const (
	FIXED      amountType = "FIXED"
	PERCENTAGE amountType = "PERCENTAGE"
)
const (
	SUCCESS status = "SUCCESS"
	FAILED  status = "FAILED"
)

// named data
type (
	City     string
	Agency   string
	Supplier string
	Airline  string
)

// data types
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
	// gorm.Model
	RequestImpl
	RuleId      uint64 `gorm:"column:rule_id;primaryKey;autoIncrement;unique"`
	RoutesId    uint64
	Routes      []Route    `json:"routes" gorm:"column:routes;ForeignKey:RoutesId"`
	Airlines    []Airline  `json:"airlines" gorm:"column:airlines"`
	Agencies    []Agency   `json:"agencies" gorm:"column:agencies"`
	Suppliers   []Supplier `json:"suppliers" gorm:"column:suppliers"`
	AmountType  amountType `json:"amountType" sql:"type:ENUM('FIXED','PERCENTAGE')"`
	AmountValue int        `json:"amountValue"`
}

// func (ct *RuleModel) TableName() string {
// 	return "rule"
// }

// type RuleModel struct {
// 	gorm.Model
// 	RequestImpl
// 	RuleId      uint64     `gorm:"column:rule_id;primaryKey;autoIncrement;unique"`
// 	Routes      []uint64   `json:"routes" gorm:"column:routes;ForeignKey:RoutesId"`
// 	Airlines    []Airline  `json:"airlines" gorm:"column:airlines"`
// 	Agencies    []Agency   `json:"agencies" gorm:"column:agencies"`
// 	Suppliers   []Supplier `json:"suppliers" gorm:"column:suppliers"`
// 	AmountType  amountType `json:"amountType" sql:"type:ENUM('FIXED','PERCENTAGE')"`
// 	AmountValue int        `json:"amountValue"`
// }

type Response struct {
	Message string `json:"message"`
	Status  status `json:"status"`
}

type Student struct {
	Name     string
	LastName string
	Height   int
}

type Class struct {
	Level    int
	Name     string
	Students []Student
}

func (ct *amountType) Scan(value interface{}) error {
	*ct = amountType(value.([]byte))
	return nil
}

func (ct *amountType) Value() (driver.Value, error) {
	return string(*ct), nil
}

func (ct *Route) Scan(value interface{}) error {
	// *ct = Route(value.([]byte))
	// *ct = value.([]byte)
	return nil
}

func (ct *Route) Value() (driver.Value, error) {
	return ct.ToString(), nil
}
