package main

//FORMS

/* AVOID DOUBLES
const (
	InsertWorkshopForm1 = `INSERT INTO users (name, surname, mail, birth)
  											SELECT ?, ?, ? ? FROM DUAL?
												WHERE NOT EXISTS
												(SELECT user_id FROM users WHERE user_id=?)`
)
*/

const (
	InsertWorkshopForm = `INSERT INTO user (name, surname, mail, birth_date) 
												VALUES ("Clarence", "Chapelier", "clacla@yahoo.com", "1999/03/14")
												INSERT INTO reservations (user_id, workshop_id, type, status) 
												VALUES (1, 2, "workshop", "pending");
												;`
)

const (
	InsertItemsForm = `INSERT INTO user (name, surname, mail, birth_date) 
										VALUES ("Clarence", "Chapelier", "clacla@yahoo.com", "1999/03/14");
										INSERT INTO appointements (user_id, item_id, date_begin, date_end, status) 
										VALUES (2, 3, '2020-03-12', '2020-03-12' + interval 10 day, 'pending');`
)

const (
	InsertEventForm = `INSERT INTO user (name, surname, mail, birth_date) 
										VALUES ("Clarence", "Chapelier", "clacla@yahoo.com", "1999/03/14"); 
										INSERT INTO reservations (user_id, event_id, type, status) 
										VALUES (1, 2, "event", "pending");`
)

const (
	InsertObj = `INSERT INTO objects (user_id, name, type, height, weight, rate)
								VALUES (2, "truelle", "Quincaillerie/Outillage", NULL, 450, "use");`
)

//CONFIRM
const (
	ConfirmWorkshopEvent = `UPDATE reservations 
													SET reservations.status = "confirmed" 
													WHERE reservations.user_id = 1`
)

const (
	ConfirmItem = `UPDATE appointements, items
								SET appointements.status = "confirmed", items.availlable = 0
								WHERE appointements.item_id = items.id
								AND appointements.item_id = 1;
	
	`
)

//CANCEL
const (
	CancelWorkshopEvent = `UPDATE reservations 
												SET reservations.status = "cancelled" 
												WHERE reservations.user_id = 1`
)

const (
	CancelItem = `UPDATE appointements, items
								SET appointements.status = "cancelled", items.availlable = 1
								WHERE appointements.item_id = items.id
								AND appointements.item_id = 1;`
)

//DISPLAY
const (
	DisplayItemsByLocation = `SELECT * items WHERE location=?`
)

const (
	DisplayItemsByType = `SELECT * FROM items WHERE type=?`
)

const (
	DisplayAllItems = `SELECT id, name, type, location, price  FROM items WHERE availlable=1`
)

const (
	DisplayItem = `SELECT * FROM items WHERE id=?`
)

const (
	DisplayAllEvents = `SELECT * FROM events`
)

const (
	DisplayEvent = `SELECT * FROM events WHERE event_id=?`
)

const (
	DisplayWorkshops = `SELECT * FROM workshops`
)
