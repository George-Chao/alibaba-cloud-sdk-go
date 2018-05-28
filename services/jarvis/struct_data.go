package jarvis

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

// Data is a nested struct in jarvis response
type Data struct {
	GmtExpire     string `json:"GmtExpire" xml:"GmtExpire"`
	PunishCount   int    `json:"PunishCount" xml:"PunishCount"`
	Protocol      string `json:"Protocol" xml:"Protocol"`
	SrcIP         string `json:"SrcIP" xml:"SrcIP"`
	SrcPort       int    `json:"SrcPort" xml:"SrcPort"`
	DstPort       int    `json:"DstPort" xml:"DstPort"`
	FeedBack      int    `json:"FeedBack" xml:"FeedBack"`
	AutoConfig    int    `json:"AutoConfig" xml:"AutoConfig"`
	Reason        string `json:"Reason" xml:"Reason"`
	GmtCreate     string `json:"GmtCreate" xml:"GmtCreate"`
	GroupId       int    `json:"GroupId" xml:"GroupId"`
	PunishType    string `json:"PunishType" xml:"PunishType"`
	RegionId      string `json:"RegionId" xml:"RegionId"`
	PunishResult  string `json:"PunishResult" xml:"PunishResult"`
	Status        string `json:"Status" xml:"Status"`
	GmtRealExpire string `json:"GmtRealExpire" xml:"GmtRealExpire"`
	SrcUid        string `json:"SrcUid" xml:"SrcUid"`
	DstIP         string `json:"DstIP" xml:"DstIP"`
	Items         []Item `json:"Items" xml:"Items"`
}