package main

import (
	"log"
	"net/http"

	db "github.com/ScriptMang/conch/internal/bikeshop"
	"github.com/gin-gonic/gin"
)

// configs gin router and renders index-page
func setRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Conch Index-Page",
			"crud1": "Insert Table-Row Data",
			"crud2": "Perform a Read Operation",
		})
	})
	return r
}

// Takes Post request data of the types: url-encoded or json
// and binds it, to the struct 'invs'.
// When passed to insert-op its used as a bridge
// to add a new invoice.
func show(r *gin.Engine) *gin.Engine {
	r.POST("/crud1", func(c *gin.Context) {

		var invs db.Invoice
		if err := c.ShouldBind(&invs); err != nil {
			log.Fatalf("Error Binding: %v\n", err)
		}

		db.InsertOp(invs)
		c.JSON(http.StatusOK, invs)
	})
	return r
}

// renders the form page that's needed to create an invoice
func addInvoice(r *gin.Engine) *gin.Engine {
	r.GET("/crud1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "crud1.tmpl", gin.H{
			"title":   "Crud1",
			"details": "Add Invoice",
		})
	})
	return r
}

// reads the tablerows from the database
func readData(r *gin.Engine) *gin.Engine {
	r.GET("/crud2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "crud2.tmpl", gin.H{
			"title": "Crud2",
			"details": "Request the First Name, Last Name, Product," +
				"Price and Quantity of all the customers who's" +
				"invoice has a unit price over $13 ",
			"rslt": db.ReadOp(),
		})
	})
	return r
}

func main() {
	r := setRouter()
	r = readData(r)
	r = addInvoice(r)
	r = show(r)
	r.Run()
}
