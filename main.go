package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//define variables
var (
	log *zap.Logger
)

//Preview error struct
type Preview struct {
	Text       string
	ErrorCause string
}

//HANDLE FORMS
func bookItemForm(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "form-book_item.html", data)
}

func bookWorkForm(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "form-book_work.html", data)
}

func bookEventForm(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "form-book_event.html", data)
}

//HANDLE PAGES
func aboutPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "about.html", data)
}

func indexPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "index.html", data)
}

func contactPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "contact.html", data)
}

func givePage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "give.html", data)
}
func workshopPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "workshop.html", data)
}

func shopPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "shop.html", data)
}

func legalPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "legal.html", data)
}

func showCategory(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "shop.html", data)
}

func showLocation(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "shop.html", data)
}

func showItem(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "item.html", data)
}

func reinsertionPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "reinsertion.html", data)
}

func showWork(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "workshop.html", data)
}

func showEvent(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "event.html", data)
}

func confirmItem(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "confirm.html", data)
}

func confirmEvent(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "confirm.html", data)
}

func confirmWorkshop(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "confirm.html", data)
}

func cancelItem(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "confirm.html", data)
}

func cancelEvent(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "confirm.html", data)
}

func cancelWorkshop(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "confirm.html", data)
}

//RECEPT FORMS
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

	EventForm{
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

	//database()
}

func receptBookItem(c *gin.Context) {

	c.Request.ParseForm()

	ItemsForm{
		Mail:       strings.Join(c.Request.PostForm["mail"], " "),
		Name:       strings.Join(c.Request.PostForm["name"], " "),
		Surname:    strings.Join(c.Request.PostForm["surname"], " "),
		Birth:      strings.Join(c.Request.PostForm["birth"], " "),
		PickupDate: strings.Join(c.Request.PostForm["pickupdate"], " "),
		ItemId:     1, //not still like this
	}
}

func receptBookWork(c *gin.Context) {

	c.Request.ParseForm()

	WorkshopForm{
		Mail:       strings.Join(c.Request.PostForm["mail"], " "),
		Name:       strings.Join(c.Request.PostForm["name"], " "),
		Surname:    strings.Join(c.Request.PostForm["surname"], " "),
		Birth:      strings.Join(c.Request.PostForm["birth"], " "),
		WorkshopId: 1, //not still like this
	}
}

/*
	if response == true {
		c.Redirect(http.StatusMovedPermanently, "/success")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/error")
	}
*/

//handle success page
func successForm(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "success.html", data)
}

//handle custom error page
func errorPage(c *gin.Context) {
	var cause string
	//get infos about the path
	host := "127.0.0.1"
	fullpath := host + c.FullPath()

	cause = "une erreur s'est produite"

	//display error page
	data := Preview{
		Text:       fullpath,
		ErrorCause: cause,
	}
	c.HTML(404, "404.html", data)
}

func main() {

	htmlSource := "/front/templates/"
	//zap stuff
	log, _ = zap.NewProduction()
	defer log.Sync()

	//Define router
	router := gin.Default()

	//to include html
	router.LoadHTMLFiles(
		htmlSource+"index.html",
		htmlSource+"404.html",
		htmlSource+"about.html",
		htmlSource+"legal.html",
		htmlSource+"reinsertion.html",
		htmlSource+"give.html",
		htmlSource+"workshops_list.html",
		htmlSource+"workshop.html",
		htmlSource+"shop.html",
		htmlSource+"item.html",
		htmlSource+"confirm.html",
		htmlSource+"cancel.html",
		htmlSource+"contact.html",
		htmlSource+"event.html",
		htmlSource+"success.html",
		htmlSource+"form-book_item.html",
		htmlSource+"form-book_work.html",
		htmlSource+"form-book_event.html",
	)

	//to include js / css / imgs
	router.Static("/front/js", "./js")
	router.Static("/front/css", "./css")
	router.Static("/front/pictures", "./pictures")

	//ROUTES

	//index routes
	router.NoRoute(errorPage)
	router.GET("/error", errorPage)
	router.GET("/", indexPage)
	router.GET("/index", indexPage)
	//router.GET("/success", successForm)

	//navbar routes
	router.GET("/recycleme", indexPage)
	router.GET("/reinsertion", reinsertionPage)
	router.GET("/give", givePage)
	router.GET("/workshop", workshopPage)
	router.GET("/shop", shopPage)
	router.GET("/about", aboutPage)
	router.GET("/legal", legalPage)

	//item - category && location
	router.GET("/shop/category/:category", showCategory)
	router.GET("/shop/location/:location", showLocation)

	//show item / book item
	router.GET("/shop/:item_id", showItem)
	router.GET("/shop/:item_id/book", bookItemForm)
	router.POST("/shop/:item_id/book", receptBookItem)

	//show event - book event
	router.GET("/shop/:event_id", showEvent)
	router.GET("/shop/:event_id/book", bookEventForm)
	router.POST("/shop/:event_id/book", receptBookEvent)

	//book workshop
	router.GET("/shop/:work_id", showWork)
	router.GET("/shop/:work_id/book", bookWorkForm)
	router.POST("/shop/:work_id/book", receptBookWork)

	//confirm booking
	router.GET("/confirm/:user_id/item/:item_id", confirmItem)
	router.GET("/confirm/:user_id/event/:event_id", confirmEvent)
	router.GET("/confirm/:user_id/work/:work_id", confirmWorkshop)

	//cancel booking
	router.GET("http://127.0.0.1/cancel/:user_id/item/:item_id", cancelItem)
	router.GET("http://127.0.0.1/cancel/:user_id/event/:event_id", cancelEvent)
	router.GET("http://127.0.0.1/cancel/:user_id/work/:work_id", cancelWorkshop)

	//launch service
	router.Run(":3000")
}
