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

// Partition is a nested struct in emr response
type Partition struct {
	ConsumerOffset        int64  `json:"ConsumerOffset" xml:"ConsumerOffset"`
	ConsumerInstanceOwner string `json:"ConsumerInstanceOwner" xml:"ConsumerInstanceOwner"`
	ConsumerGroupId       string `json:"ConsumerGroupId" xml:"ConsumerGroupId"`
	LogEndOffset          int64  `json:"LogEndOffset" xml:"LogEndOffset"`
	DataSourceId          string `json:"DataSourceId" xml:"DataSourceId"`
	PartitionComment      string `json:"PartitionComment" xml:"PartitionComment"`
	TableId               string `json:"TableId" xml:"TableId"`
	PartitionPath         string `json:"PartitionPath" xml:"PartitionPath"`
	Location              string `json:"Location" xml:"Location"`
	PartitionName         string `json:"PartitionName" xml:"PartitionName"`
	Lag                   int64  `json:"Lag" xml:"Lag"`
	GmtCreate             int64  `json:"GmtCreate" xml:"GmtCreate"`
	PartitionType         string `json:"PartitionType" xml:"PartitionType"`
	DatabaseId            string `json:"DatabaseId" xml:"DatabaseId"`
	PartitionId           int    `json:"PartitionId" xml:"PartitionId"`
	BucketNum             int    `json:"BucketNum" xml:"BucketNum"`
	GmtModified           int64  `json:"GmtModified" xml:"GmtModified"`
}
