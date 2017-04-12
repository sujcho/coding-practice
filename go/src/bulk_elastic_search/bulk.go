package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/olivere/elastic.v5"
)

//type Input map[string]interface{}
type Input struct {
	Index  string                 `json:"_index"`
	Type   string                 `json:"_type"`
	Id     string                 `json:"_id"`
	Score  int64                  `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}
type Inputs []Input

func main() {

	//reads input set
	inputs := GetInputs()

	//elastic
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200/"))
	if err != nil {
		panic(err)
	}

	bulkRequest := client.Bulk()
	bulkRequest = BuildIndexRequest(inputs, bulkRequest)

	// Do sends the bulk requests to Elasticsearch
	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		// ...
	}
	indexed := bulkResponse.Indexed()
	for _, index := range indexed {
		fmt.Println(index.Id)
	}
}

//build multiple bulk requests
func BuildIndexRequest(entries Inputs, s *elastic.BulkService) *elastic.BulkService {
	for _, entry := range entries {
		req := WrapperBulkIndexRequest(entry.Index, entry.Type, entry.Id, entry.Source)
		s = s.Add(req)
	}
	return s
}

//!!!! handle score later!!
//wrapper index NewBulkIndexRequest
func WrapperBulkIndexRequest(theIndex string, theType string, theId string, theSource map[string]interface{}) *elastic.BulkIndexRequest {
	source, _ := json.Marshal(theSource)
	return elastic.NewBulkIndexRequest().Index(theIndex).Type(theType).Id(theId).Doc(string(source))
}

func PrintInputs(inputs Inputs) {
	for _, input := range inputs {
		data, _ := json.Marshal(input.Source)
		fmt.Println(string(data))
	}
}
func GetInputs() Inputs {

	raw, err := ioutil.ReadFile("MyInputs.json")
	if err != nil {
		fmt.Print("not found")
	}

	var all Inputs
	json.Unmarshal(raw, &all)
	return all
}
