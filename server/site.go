package server

import (
	"encoding/json"
	"fmt"
	"snapp-trip/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Airline struct {
	// data
	Name    string `json:"name" gorm:"name"`
	NameEng string `json:"name_eng" gorm:"name_eng"`
	NameFar string `json:"name_far" gorm:"name_far"`
}

func indexHandler(c *fiber.Ctx, db *gorm.DB) error {
	var todos []string
	var airlines []Airline
	db.Model(&airlines)
	db.Table("airline").Find(&airlines)
	for _, airline := range airlines {
		res, _ := json.Marshal(airline)
		todos = append(todos, string(res))
	}

	path2 := "data/samples/create_rule_request.json"
	rule_json_data := GetFromPath(path2)
	rules := model.RulesFromJson(rule_json_data)
	fmt.Println(rules[0].Airlines)
	fmt.Println(rules[0].Routes)
	println(rules[0].ToString())
	fmt.Println(rules[0].PrettyString())
	rules[0].Info()

	SendToDB(&rules[0])

	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}
func SendToDB(rule *model.Rule) error {
	con := GetConnection()
	err := con.Table("rule").Create(&rule).Error
	// statement := `
	// INSERT INTO 'rule (amount_value, amount_type, suppliers, agencies, airlines, routes, rule_id)
	// VALUES ($1, $2, $3, $4, $5, $6, $7)`
	// err := con.Raw(statement, rule.AmountValue, rule.AmountType, rule.Suppliers, rule.Agencies, rule.Airlines, rule.Routes, rule.RuleId).Error
	return err
}
