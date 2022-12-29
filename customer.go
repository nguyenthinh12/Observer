package observer

import "fmt"


/*Customer is struct*/
type Customer struct {
	ID string
}

/*Updete is function*/
func (c *Customer) Updete(itemName string) {
	fmt.Println("Sending email to customer %s for item %s\n", c.ID, itemName)
}

/*GetID is function*/
func (c *Customer) GetID()string {
	return c.ID
}