package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

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

type Source struct {
	SrcInst string `json:"srcInst"`
	Reason  int    `json:"reason"`
}

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
	_ = indexed
	/*
		for _, index := range indexed {
			fmt.Println(index.Id)
		}
	*/
	// Search with a term query
	termQuery := elastic.NewTermQuery("srcInst", "62358ab04")

	searchResult, err := client.Search().
		Index("dev_nanodemoxxx1_traffic_logs"). // search in index "twitter"
		Type("kafka").
		Query(termQuery).        // specify the query
		Sort("srcInst", true).   // sort by "user" field, ascending
		From(0).Size(10).        // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.
	var src Source
	for _, item := range searchResult.Each(reflect.TypeOf(src)) {
		if t, ok := item.(Source); ok {
			fmt.Printf("input by %s: %d\n", t.SrcInst, t.Reason)
		}
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
