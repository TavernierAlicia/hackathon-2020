package main

import (
	"github.com/gin-gonic/gin"
)

//HANDLE FORMS
func bookItemForm(c *gin.Context) {
	c.HTML(200, "form-book_item.html", nil)
}

func bookWorkForm(c *gin.Context) {
	c.HTML(200, "form-book_work.html", nil)
}

func bookEventForm(c *gin.Context) {
	c.HTML(200, "form-book_event.html", nil)
}

//RECEPT FORMS
/*
func receptBookEvent(c *gin.Context) {

	c.Request.ParseForm()

	//make some conversion
	objS0, err := strconv.Atoi(strings.Join(c.Request.PostForm["objectsize0"], " "))
	objW0, err := strconv.Atoi(strings.Join(c.Request.PostForm["objectweight0"], " "))
	objS1, err := strconv.Atoi(strings.Join(c.Request.PostForm["objectsize1"], " "))
	objW1, err := strconv.Atoi(strings.Join(c.Request.PostForm["objectweight1"], " "))
	objS2, err := strconv.Atoi(strings.Join(c.Request.PostForm["objectsize2"], " "))
	objW2, err := strconv.Atoi(strings.Join(c.Request.PostForm["objectweight2"], " "))

	if err != nil {
		fmt.Println("Conversion Error")
	}

	data := EventForm{
		Mail:    strings.Join(c.Request.PostForm["mail"], " "),
		Name:    strings.Join(c.Request.PostForm["name"], " "),
		Surname: strings.Join(c.Request.PostForm["surname"], " "),
		Birth:   strings.Join(c.Request.PostForm["birth"], " "),
		EventId: 1, //not still line this

		ObjectName0:   strings.Join(c.Request.PostForm["objectname0"], " "),
		ObjectType0:   strings.Join(c.Request.PostForm["type0"], " "),
		ObjectState0:  strings.Join(c.Request.PostForm["objectstate0"], " "),
		ObjectSize0:   objS0,
		ObjectWeight0: objW0,

		ObjectName1:   strings.Join(c.Request.PostForm["objectname1"], " "),
		ObjectType1:   strings.Join(c.Request.PostForm["subject1"], " "),
		ObjectState1:  strings.Join(c.Request.PostForm["objectstate1"], " "),
		ObjectSize1:   objS1,
		ObjectWeight1: objW1,

		ObjectName2:   strings.Join(c.Request.PostForm["objectname2"], " "),
		ObjectType2:   strings.Join(c.Request.PostForm["subject2"], " "),
		ObjectState2:  strings.Join(c.Request.PostForm["objectstate2"], " "),
		ObjectSize2:   objS2,
		ObjectWeight2: objW2,
	}
	path := c.FullPath()

	database(data, req)
}

func receptBookItem(c *gin.Context) {

	c.Request.ParseForm()

	database(ItemsForm{
		Mail:       strings.Join(c.Request.PostForm["mail"], " "),
		Name:       strings.Join(c.Request.PostForm["name"], " "),
		Surname:    strings.Join(c.Request.PostForm["surname"], " "),
		Birth:      strings.Join(c.Request.PostForm["birth"], " "),
		PickupDate: strings.Join(c.Request.PostForm["pickupdate"], " "),
		ItemId:     1, //not still like this
	}, id, req)
	database(data, req)
}

func receptBookWork(c *gin.Context) {

	c.Request.ParseForm()

	data := WorkshopForm{
		Mail:       strings.Join(c.Request.PostForm["mail"], " "),
		Name:       strings.Join(c.Request.PostForm["name"], " "),
		Surname:    strings.Join(c.Request.PostForm["surname"], " "),
		Birth:      strings.Join(c.Request.PostForm["birth"], " "),
		WorkshopId: 1, //not still like this
	}
	database(data, req)
}

	//check send mail
	if response == true {
		c.Redirect(http.StatusMovedPermanently, "/success")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/error")
	}
*/
