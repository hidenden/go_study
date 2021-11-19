package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func read_decision_log1(data []byte) {
	fmt.Println("READ DECISOIN LOG Part.1")
	var f interface{}
	if jerr := json.Unmarshal(data, &f); jerr != nil {
		fmt.Println("JSON Unmarshal error:" + jerr.Error())
		return
	}
	json_array, ok := f.([]interface{})
	if !ok {
		fmt.Println("Invalid decision log data")
		return
	}
	for i, v := range json_array {
		fmt.Printf("Index:%d\n", i)
		json_object, ok := v.(map[string]interface{})
		if !ok {
			fmt.Printf("Invalid Log Data\n")
			continue
		}
		did, ok := json_object["decision_id"]
		if !ok {
			fmt.Printf("Invalid Log Data\n")
			continue
		}
		did_s, ok := did.(string)
		if !ok {
			fmt.Printf("Invalid Log Data\n")
			continue
		}
		fmt.Printf("Decision id:%s", did_s)
		// fmt.Printf("Data:%v\n", json_object)
		log_data, m_err := json.Marshal(&json_object)
		if m_err != nil {
			fmt.Printf("JSON Marshal error:" + m_err.Error() + "\n")
			continue
		}
		fmt.Printf("JSON DATA:%s\n", string(log_data))
	}
}

type DecisionLog struct {
	DecisionId string `json:"decision_id"`
}

func read_decision_log2(data []byte) {
	var err error
	var log_array []DecisionLog

	fmt.Println("READ DECISOIN LOG Part.2")

	if err = json.Unmarshal(data, &log_array); err != nil {
		fmt.Println("JSON Unmarshal error:" + err.Error())
		return
	}

	for i, v := range log_array {
		fmt.Printf("Index:%d\n", i)
		fmt.Printf("Data:%v\n", v)
		did := v.DecisionId
		fmt.Printf("Decision id:%s\n", did)

		log_data, m_err := json.Marshal(&v)
		if m_err != nil {
			fmt.Printf("JSON Marshal error:" + m_err.Error() + "\n")
			continue
		}
		fmt.Printf("JSON DATA:%s\n", string(log_data))
	}
}

func read_decision_log3(data []byte) {
	fmt.Println("READ DECISOIN LOG Part.3")
	var json_array []map[string]interface{}
	if jerr := json.Unmarshal(data, &json_array); jerr != nil {
		fmt.Println("JSON Unmarshal error:" + jerr.Error())
		return
	}
	for i, json_object := range json_array {
		fmt.Printf("Index:%d\n", i)
		did, ok := json_object["decision_id"]
		if !ok {
			fmt.Printf("Invalid Log Data\n")
			continue
		}
		did_s, ok := did.(string)
		if !ok {
			fmt.Printf("Invalid Log Data\n")
			continue
		}
		fmt.Printf("Decision id:%s", did_s)
		// fmt.Printf("Data:%v\n", json_object)
		log_data, m_err := json.Marshal(&json_object)
		if m_err != nil {
			fmt.Printf("JSON Marshal error:" + m_err.Error() + "\n")
			continue
		}
		fmt.Printf("JSON DATA:%s\n", string(log_data))
	}
}

func main() {
	fmt.Println("Hello JSON World")
	data, err := ioutil.ReadFile("log.json")
	if err != nil {
		fmt.Println("Unable to open JSON File:" + err.Error())
		return
	}
	//	read_decision_log1(data)

	//	read_decision_log2(data)

	read_decision_log3(data)
}
