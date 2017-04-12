package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/Jeffail/gabs"
)

func RequestAggWithField(client *http.Client) {

	Field := [4]string{
		"srcInst",
		"destInst",
		"linkProps",
		"destComp",
	}
	//aggregation over all field
	for _, field := range Field {
		//	fmt.Printf("field[%s]\n", field)
		result := SendAggRequest(client, "agg_bucket", field)
		PutToDB(result, field)
	}
}
func RequestAgg() {
	client := &http.Client{}
	RequestAggWithField(client)
}
func SendAggRequest(client *http.Client, Name string, Field string) []byte {

	var URL string = "http://localhost:9200/dev_nanodemoxxx1_traffic_logs/kafka/_search"
	reqBody := gabs.New()
	// or gabs.Consume(jsonObject) to work on an existing map[string]interface{}
	reqBody.Set(Field, "aggs", Name, "terms", "field")
	//fmt.Println(reqBody.String())
	var jsonStr = []byte(reqBody.String())

	req, err := http.NewRequest("GET", URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	agg := ParseAgg(body)

	return agg
}
