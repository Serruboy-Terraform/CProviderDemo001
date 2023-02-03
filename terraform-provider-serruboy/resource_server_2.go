package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
	"log"
	"net/http"
)

func resourceServer2() *schema.Resource {
	return &schema.Resource{
		Create: resourceServer2Create,
		Read:   resourceServer2Read,
		Update: resourceServer2Update,
		Delete: resourceServer2Delete,

		Schema: map[string]*schema.Schema{
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServer2Create(d *schema.ResourceData, m interface{}) error {
	country := d.Get("country").(string)

	d.SetId(country)

	return resourceServer2Read(d, m)
}

func resourceServer2Read(d *schema.ResourceData, m interface{}) error {
	resp, err := http.Get("https://fakerapi.it/api/v1/persons?_locale=" + d.Id())

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//	body, err := io.ReadAll(resp.Body)

	log.Print(resp.Body)

	// d.Set("country", body)

	return nil
}

func resourceServer2Update(d *schema.ResourceData, m interface{}) error {
	return resourceServer2Read(d, m)
}

func resourceServer2Delete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
