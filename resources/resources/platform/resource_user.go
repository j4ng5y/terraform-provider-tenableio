package resources

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/j4ng5y/terraform-provider-tenableio/resources/resources"

	"github.com/hashicorp/terraform/helper/schema"
)

type UserCreateSchema struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Permissions int    `json:"permissions"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Type        string `json:"type"`
}

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
		endpoint    = "https://cloud.tenable.com/users"
		username    = d.Get("username").(string)
		password    = d.Get("passowrd").(string)
		permissions = d.Get("permissions").(int)
		name        = d.Get("name").(string)
		email       = d.Get("email").(string)
		t           = d.Get("type").(string)
	)
	d.SetId(username)

	j, err := json.Marshal(&UserCreateSchema{
		Username:    username,
		Password:    password,
		Permissions: permissions,
		Name:        name,
		Email:       email,
		Type:        t})
	if err != nil {
		log.Fatal(err)
	}
	T := &resources.TenableIORequest{
		Endpoint: endpoint,
		Method:   http.MethodGet,
		Credentials: &resources.TenableIORequestCredentials{
			AccessKey: d.Get("accessKey").(string),
			SecretKey: d.Get("secretKey").(string),
		},
		Body: j,
	}
	resp, err := T.Do()
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

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
