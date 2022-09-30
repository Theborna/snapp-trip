package model

func (city City) Validate() (err error) {
	// var _city *City
	// return db.GetConnection().Table("City").Find(_city, "key = ?", city).Error
	return nil
}

func (agency Agency) validate() (err error) {
	return nil
}

func (supplier Supplier) validate() (err error) {
	return nil
}

func (airline Airline) validate() (err error) {
	return nil
}

func (trip Trip) Validate() (err error) {
	errors := []error{trip.Agency.validate(), trip.Airline.validate(), trip.Destination.Validate(),
		trip.Origin.Validate(), trip.Supplier.validate()}
	for _, _err := range errors {
		if _err != nil {
			return _err
		}
	}
	return nil
}

func (rule Rule) Validate() (err error) {
	return nil
}
