package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
)

func ipServer() *schema.Resource {
	return &schema.Resource{
		Create: demo_001Create,
		Read:   demo_001Read,
		Update: demo_001Update,
		Delete: demo_001Delete,

		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func demo_001Create(d *schema.ResourceData, m interface{}) error {
	ip := d.Get("ip").(string)

	d.SetId(ip)

	return demo_001Read(d, m)
}

func demo_001Read(d *schema.ResourceData, m interface{}) error {

	myip := d.Get("ip").(string)
	url := ""

	if myip == "ipv4" {
		url = "https://api.ipify.org/?format=json"
	} else {
		url = "https://api64.ipify.org/?format=json"
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}

	defer resp.Body.Close()

	type Summary struct {
		Ip string `json:"ip"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var summary = new(Summary)
	err2 := json.Unmarshal(body, &summary)
	if err2 != nil {
		fmt.Println("whoops:", err2)
	}

	d.SetId(summary.Ip)

	return nil
}

func demo_001Update(d *schema.ResourceData, m interface{}) error {
	return demo_001Read(d, m)
}

func demo_001Delete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
