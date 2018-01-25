/*
* Golang reflect update struct to tag
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package main

import (
	"fmt"
	"github.com/jeffotoni/gsreflect"
)

// struct config
type Config struct {
	Domain  string `default:"s3go.gov"`
	Schema  string `default:"https"`
	Port    string `default:"9002"`
	Db      string `default:"postgresql"`
	Cluster string `default:"10"`
	Passwd  string `default:"x37c$%2"`
	Login   string `default:"postgres"`
}

func main() {

	c := Config{}

	err := gsr.Default(&c)

	if err != nil {

		fmt.Println("error: ", err)
		return
	}

	fmt.Println("Domain: ", c.Domain)
	fmt.Println("Schema: ", c.Schema)
	fmt.Println("Port: ", c.Port)
	fmt.Println("Db: ", c.Db)
	fmt.Println("Cluster: ", c.Cluster)
	fmt.Println("Passwd: ", c.Passwd)
	fmt.Println("Login: ", c.Login)

}
