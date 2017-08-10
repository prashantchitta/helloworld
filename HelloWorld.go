// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// A simple example exposing fictional RPC latencies with different types of
// random distributions (uniform, normal, and exponential) as Prometheus
// metrics.
package main

//https://raw.githubusercontent.com/prometheus/client_golang/master/examples/random/main.go

import (
	"time"
	"os"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v5"
	"github.com/Pallinder/go-randomdata"

	//"k8s.io/GoPlay/vendor/github.com/Pallinder/go-randomdata"
	//elastic "k8s.io/GoPlay/vendor/gopkg.in/olivere/elastic.v5"
	"golang.org/x/net/context"
)

const ELASTIC_DEFAULT_URL = "http://localhost:9200"

type Tweet struct {
	User    string `json:"user"`
	Email string `json:"message"`
	Time    time.Time `json:"time"`
}

func main() {
	ctx := context.Background()
	elasticUrl := os.Getenv("ELASTIC_URL")

	if len(elasticUrl) == 0 {
		fmt.Println("Using the default elastic url:", ELASTIC_DEFAULT_URL)
		elasticUrl = ELASTIC_DEFAULT_URL
	} else {
		fmt.Println("Environment variable for Elastic found. Using it :", elasticUrl)
	}

	client, err := elastic.NewClient(
		elastic.SetURL(elasticUrl),
		elastic.SetBasicAuth("elastic", "changeme"))

	if err != nil {
		// Handle error
	}

	for  {
		tweet := Tweet{User: randomdata.SillyName(), Email: randomdata.Email(), Time:    time.Now()}
		_, err = client.Index().Index("twitter").Type("tweet").BodyJson(tweet).Do(ctx)
		client.Stop();
		if err != nil {
			// Handle error
			panic(err)
		}
		fmt.Println("Tweet :", tweet)
		time.Sleep(1000 * time.Millisecond)
	}
}
