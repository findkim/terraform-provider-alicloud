package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Alarm is a nested struct in emr response
type Alarm struct {
	Name               string `json:"Name" xml:"Name"`
	MetricName         string `json:"MetricName" xml:"MetricName"`
	Period             int    `json:"Period" xml:"Period"`
	Threshold          string `json:"Threshold" xml:"Threshold"`
	EvaluationCount    int    `json:"EvaluationCount" xml:"EvaluationCount"`
	StartTime          int    `json:"StartTime" xml:"StartTime"`
	EndTime            int    `json:"EndTime" xml:"EndTime"`
	SilenceTime        int    `json:"SilenceTime" xml:"SilenceTime"`
	NotifyType         int    `json:"NotifyType" xml:"NotifyType"`
	Enable             bool   `json:"Enable" xml:"Enable"`
	State              string `json:"State" xml:"State"`
	ComparisonOperator string `json:"ComparisonOperator" xml:"ComparisonOperator"`
	ContactGroups      string `json:"ContactGroups" xml:"ContactGroups"`
}
