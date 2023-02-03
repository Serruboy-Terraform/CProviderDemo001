package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	 myip := d.Get("country").(string)
	 url := ""

	if myip == "ipv4"{
		url = "https://api.ipify.org/?format=json"
	} else {
		url = "https://api64.ipify.org/?format=json"
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}

	bodyString := string(resBody)

	jsondata, err := json.Marshal(bodyString)

	if err != nil {
		log.Println(err)
	}

	log.Print(jsondata)

	d.SetId(bodyString)

	return nil
}

func resourceServer2Update(d *schema.ResourceData, m interface{}) error {
	return resourceServer2Read(d, m)
}

func resourceServer2Delete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
