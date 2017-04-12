package main

import (
	"context"
	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

func main() {

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://localhost:9200").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Index a tweet (using JSON serialization)
	//	tweet := Tweet{User: "olivere", Message: "Take Five"}

	input1 := `{"srcInst": "fg29806:srchostid:IP:starttime:pid",
		"reason":     132,
		"srcIP":      "172.17.0.1:53740",
		"httpPath":   "/genres",
		"nanoMode":   1,
		"destInst":   "fg32156",
		"srcAppInfo": "_",
		"srcPort":    0, "dir": 1,
		"linkProps":   "com.nano.sftestb.moviesui com.nano.sftestc.movie 3 62358ab04 1 132",
		"policyOrder": 0, "protocol": 3, "destPort": 8080, "tokenPresent": true,
		"startTime": "2017-03-29T22:48:15.9045126Z",
		"destComp":  "com.nano.sftestc.movie",
		"srcComp":   "com.nano.sftestb.moviesui",
		"id":        "477f4ee1-308f-4deb-9478-9b01e28e9748",
		"serviceId": "62358ab04", "policyEvalCount": 0,
		"destIP":         "127.0.0.1",
		"policyHitCount": 0, "destppInfo": "_",
		"policyGroupId": "IN_PolicyGrp", "srcHostInfo": "62358ab06",
		"count": 1, "hostId": "fg32156", "sessionId": "123ghfgrfe",
		"envId": "preprod", "nanoAction": 3, "destHostInfo": "62358ab04",
		"policyId": "", "tenantId": "63800913", "endTime": "2017-03-29T22:48:20.336869408Z",
		"destUrl": "172.17.0.1:9005"}`

	put1, err := client.Index().
		Index("dev_nanodemoxxx1_traffic_logs").
		Type("kafka").
		Id("AVscQNTJsyPShSb8zXLe").
		BodyJson(input1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed input %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
