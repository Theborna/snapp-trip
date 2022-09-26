package model

func (city City) validate() (err error) {
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

func (trip Trip) validate() (err error) {
	errors := []error{trip.Agency.validate(), trip.Airline.validate(), trip.Destination.validate(),
		trip.Origin.validate(), trip.Supplier.validate()}
	for _, _err := range errors {
		if _err != nil {
			return _err
		}
	}
	return nil
}

func (rule Rule) validate() (err error) {
	return nil
}
