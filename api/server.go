package api

import (
	"encoding/json"
	"log"
	"net/http"

	"snapp-trip/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	trouble_decoding = "trouble decoding request"
	trouble_encoding = "trouble encoding response"
)

type Server struct {
	*mux.Router

	rules []model.Rule
}

func NewServer() *Server {
	router := mux.NewRouter()
	// this part was to show an html too but i cant be bothered to do so
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("./views/")))
	s := &Server{
		Router: router,
		rules:  []model.Rule{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/create-rule", s.CreateRule()).Methods("POST")
	s.HandleFunc("/change-price", s.ChangePrice()).Methods("POST")
}

func (s *Server) CreateRule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("got query %s", r.URL)
		var rules []model.Rule
		if err := json.NewDecoder(r.Body).Decode(&rules); err != nil {
			errorResponse := model.Response{
				Message: &trouble_decoding,
				Status:  model.FAILED,
			}
			http.Error(w, errorResponse.PrettyString(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		for _, rule := range rules {
			if err := rule.Validate(); err != nil {
				message := err.Error()
				errorResponse := model.Response{
					Message: &message,
					Status:  model.FAILED,
				}
				http.Error(w, errorResponse.PrettyString(), http.StatusBadRequest)
				return
			}
			log.Printf("got the rule %s", rule.PrettyString())
			rule.RuleId = uuid.New()
			s.addRule(rule)
		}
		errorResponse := model.Response{
			Status: model.SUCCESS,
		}
		w.Write([]byte(errorResponse.PrettyString()))
		log.Println("no errors")
	}
}

func (s *Server) ChangePrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("got query %s", r.URL)
		var trips []model.Trip
		if err := json.NewDecoder(r.Body).Decode(&trips); err != nil {
			errorResponse := model.Response{
				Message: &trouble_decoding,
				Status:  model.FAILED,
			}
			http.Error(w, errorResponse.PrettyString(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		for idx, _ := range trips {
			trips[idx].SetRule(s.rulesForTrip(trips[idx]))
			// log.Printf("got the trip %s", trips[idx].PrettyString())
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(trips); err != nil {
			errorResponse := model.Response{
				Message: &trouble_encoding,
				Status:  model.FAILED,
			}
			http.Error(w, errorResponse.PrettyString(), http.StatusBadRequest)
			return
		}
		log.Println("no errors")
	}
}

func (s *Server) rulesForTrip(trip model.Trip) []model.Rule { // should be modified to use database
	var rules []model.Rule
	for _, rule := range s.rules {
		if rule.ValidForTrip(trip) {
			rules = append(rules, rule)
		}
	}
	return rules
}

func (s *Server) addRule(rule model.Rule) { // should be modified to use database
	s.rules = append(s.rules, rule)
}
