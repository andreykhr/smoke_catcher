package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"smocker_catcher/models/v1"
)

func main() {
	defer fmt.Println("Завершено!")

	scenarioPtr := flag.String("scenario", "foo", "Сценарий тестов на yaml")

	flag.Parse()

	fmt.Println("scenario:", *scenarioPtr)

	b, err := ioutil.ReadFile(*scenarioPtr)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	scenarioYaml := string(b)
	log.Println("Scenario Yaml")
	log.Println(scenarioYaml)

	scenario := v1.Scenario{}

	err2 := yaml.Unmarshal([]byte(scenarioYaml), &scenario)
	if err2 != nil {
		log.Fatalf("error: %v", err2)
	}

	d, err := yaml.Marshal(&scenario)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("--- scenario dump:\n%s\n\n", string(d))

	for _, element := range scenario.Tests {
		resp, err := http.Get(element.Test.Request.Server)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		//Проверка кода ответ
		if resp.StatusCode != element.Test.Response.HTTPCode {
			log.Fatalf("Http status code wrong. Expected [%d] but was [%d]",
				element.Test.Response.HTTPCode, resp.StatusCode)
			continue
		}

		//Проверка заголовков ответов
		for _, header := range element.Test.Response.Headers {
			log.Println(header)
		}

		for true {
			bs := make([]byte, 1014)
			n, err := resp.Body.Read(bs)
			fmt.Println(string(bs[:n]))

			if n == 0 || err != nil {
				break
			}
		}
	}
}
