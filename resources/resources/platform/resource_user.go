package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"permissions": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	var (
		endpoint = "https://cloud.tenable.com/users"
		username = d.Get("username").(string)
		password = d.Get("passowrd").(string)
		permissions = d.Get("permissions").(int)
		name = d.Get("name").(string)
		email = d.Get("email").(string)
		type = d.Get("type").(string)
	)
	d.SetId(username)
	return resourceUserRead(d, m)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceUserRead(d, m)
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
