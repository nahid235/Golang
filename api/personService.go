package api

func NewList() PersonDetails {
	return PersonDetails{}
}

func GetPersonList() []PersonDetails {
	if PersonList == nil || len(PersonList) == 0 {
		PersonList = []PersonDetails{}
		return PersonList
	}
	return PersonList
}

// add person details service
func AddPersonDetailsService(data PersonDetails) {
	SavePersonDetailsToDB(data)
	GetAllPersonDetailsFromDB()
}

// Update person details Service
func UpdatePersonDetailsService(data PersonDetails) {
	UpdatePersonDetailsFromDB(data)
	GetAllPersonDetailsFromDB()

}

// Delete person details Service
func DeletePersonDetailsService(data PersonDetails) {
	DeletePersonDetailsFromDB(data.Guid)
	GetAllPersonDetailsFromDB()
}
