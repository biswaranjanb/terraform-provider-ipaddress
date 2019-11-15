package main

import "github.com/hashicorp/terraform/helper/schema"

func resourceServer() *schema.Resource {
	return &schema.Resource{ // This defines the data schema and CRUD operations of the resource. This is the only thing we require to create a resource.
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{ // TForm's schema automatically enforces validation and type casting.
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
//resourceCreate models the method for resource creation
func resourceCreate(d *schema.ResourceData, m interface{}) error {
	   address := d.Get("address").(string)
       d.SetId(address)
       return resourceRead(d, m)
} 

//resourceRead models the method for reading the resource
func resourceRead(d *schema.ResourceData, m interface{}) error{
	return nil
}

//resourceUpdate models the method for updating the resource
func resourceUpdate(d *schema.ResourceData, m interface{}) error{	
	d.Partial(true)
	if d.HasChange("address"){
		d.SetPartial("address")
	}
	d.Partial(false)
	return resourceRead(d,m)
} 
//resourceDelete models the method for deleting the resource.
func resourceDelete(d *schema.ResourceData, m interface{}) error{
	d.SetId("")
	return nil
}