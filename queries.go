package main

//FORMS
const (
	InsertWorkshopForm = `INSERT INTO users (name, surname, mail, birth) 
  											SELECT ?, ?, ?, ? FROM DUAL
												WHERE NOT EXISTS 
  											(SELECT user_id FROM reservations WHERE user_id=? AND workshop_id=?)`
)

const (
	InsertItemsForm = ``
)

const (
	InsertEventForm = ``
)

//CONFIRM
const (
	ConfirmWorkshop = ``
)

const (
	ConfirmItem = ``
)
const (
	ConfirmEvent = ``
)

//CANCEL
const (
	CancelWorkshop = ``
)

const (
	CancelItem = ``
)
const (
	CancelEvent = ``
)

//DISPLAY
const (
	DisplayItemsByLocation = ``
)

const (
	DisplayItemsByType = ``
)

const (
	DisplayAllItems = ``
)

const (
	DisplayItem = ``
)

const (
	DisplayAllEvents = ``
)

const (
	DisplayEvent = ``
)

const (
	DisplayWorkshops = ``
)
