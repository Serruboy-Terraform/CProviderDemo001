package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
)

func public_ip() *schema.Resource {
	return &schema.Resource{
		Create: public_ipCreate,
		Read:   public_ipRead,
		Update: public_ipUpdate,
		Delete: public_ipDelete,

		Schema: map[string]*schema.Schema{
			"iptype": &schema.Schema{
				Type:     schema.TypeString,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func public_ipCreate(d *schema.ResourceData, m interface{}) error {
	uuidWithHyphen := (uuid.New()).String()

	myip := d.Get("iptype").(string)
	url := ""

	if myip == "v4" {
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

	d.Set("result", summary.Ip)
	d.SetId(uuidWithHyphen)

	return public_ipRead(d, m)
}

func public_ipRead(d *schema.ResourceData, m interface{}) error {
	myip := d.Get("result").(string)

	d.Set("result", myip)
	return nil

}

func public_ipUpdate(d *schema.ResourceData, m interface{}) error {
	return public_ipRead(d, m)
}

func public_ipDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
