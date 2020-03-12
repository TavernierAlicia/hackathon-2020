package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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
	c.HTML(200, "give.html", nil)
}
func workshopPage(c *gin.Context) {
	c.HTML(200, "workshop.html", nil)
}

func shopPage(c *gin.Context) {
	//data := Preview{}
	display := DbConnect()
	for _, items := range display.Items {
		fmt.Println(items.Id)
	}
	//fmt.Println(display)
	//t := template.Must(template.New("shop").Parse(display))
	c.HTML(200, "shop.html", nil)
}

func legalPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "legal.html", data)
}

func showCategory(c *gin.Context) {
	c.HTML(200, "shop.html", nil)
}

func showLocation(c *gin.Context) {
	c.HTML(200, "shop.html", nil)
}

func showItem(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "item.html", data)
}

func reinsertionPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "reinsertion.html", data)
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
