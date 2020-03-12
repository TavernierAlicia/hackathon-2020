package main

import (
	"github.com/gin-gonic/gin"
	//"go.uber.org/zap"
)

//define variables
//var (
//	log *zap.Logger
//)

func main() {

	htmlSource := "front/templates/"
	//zap stuff
	//log, _ = zap.NewProduction()
	//defer log.Sync()

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
	router.Static("./front/js", "./front/js")
	router.Static("./front/css", "./front/css")
	router.Static("./front/pictures/misc", "./front/pictures/misc")
	router.Static("./front/pictures/items", "./front/pictures/items")
	router.Static("./front/pictures/events", "./front/pictures/events")
	router.Static("./front/pictures/workshops", "./front/pictures/workshops")

	//ROUTES

	//index routes
	router.NoRoute(errorPage)
	router.GET("/error", errorPage)
	router.GET("/", indexPage)
	router.GET("/index", indexPage)
	//router.GET("/success", successForm)

	//navbar routes
	router.GET("/recycleetvous", indexPage)
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
	router.GET("/shop/item/:item_id", showItem)
	//router.GET("/shop/:item_id/book", bookItemForm)
	//temporary
	router.GET("/shop/book", bookItemForm)
	//router.POST("/shop/:item_id/book", receptBookItem)

	//show event - book event
	//router.GET("/events/:event_id", showEvent)
	//router.GET("/events/:event_id/book", bookEventForm)
	//temporary
	router.GET("/events/book", bookEventForm)
	//router.POST("/events/:event_id/book", receptBookEvent)

	//book workshop
	//router.GET("/works/:work_id/book", bookWorkForm)
	//temporary
	router.GET("/works/book", bookWorkForm)
	//router.POST("/works/:work_id/book", receptBookWork)

	//confirm booking
	router.GET("/confirm/user/:user_id/item/:item_id", confirmItem)
	router.GET("/confirm/user/:user_id/event/:event_id", confirmEvent)
	router.GET("/confirm/user/:user_id/work/:work_id", confirmWorkshop)

	//cancel booking
	router.GET("/cancel/user/:user_id/item/:item_id", cancelItem)
	router.GET("/cancel/user/:user_id/event/:event_id", cancelEvent)
	router.GET("/cancel/user/:user_id/work/:work_id", cancelWorkshop)

	//launch service
	router.Run(":3000")
}
