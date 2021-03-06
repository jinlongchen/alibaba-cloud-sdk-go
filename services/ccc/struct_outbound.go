package ccc

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

// Outbound is a nested struct in ccc response
type Outbound struct {
	AnswerRate                   float64 `json:"AnswerRate" xml:"AnswerRate"`
	AverageDialingTime           float64 `json:"AverageDialingTime" xml:"AverageDialingTime"`
	AverageTalkTime              float64 `json:"AverageTalkTime" xml:"AverageTalkTime"`
	AverageWorkTime              int64   `json:"AverageWorkTime" xml:"AverageWorkTime"`
	CallsAnswered                int64   `json:"CallsAnswered" xml:"CallsAnswered"`
	CallsDialed                  int64   `json:"CallsDialed" xml:"CallsDialed"`
	MaxDialingTime               int64   `json:"MaxDialingTime" xml:"MaxDialingTime"`
	MaxTalkTime                  int64   `json:"MaxTalkTime" xml:"MaxTalkTime"`
	MaxWorkTime                  int64   `json:"MaxWorkTime" xml:"MaxWorkTime"`
	SatisfactionIndex            float64 `json:"SatisfactionIndex" xml:"SatisfactionIndex"`
	SatisfactionSurveysOffered   int64   `json:"SatisfactionSurveysOffered" xml:"SatisfactionSurveysOffered"`
	SatisfactionSurveysResponded int64   `json:"SatisfactionSurveysResponded" xml:"SatisfactionSurveysResponded"`
	TotalDialingTime             int64   `json:"TotalDialingTime" xml:"TotalDialingTime"`
	TotalHoldTime                int64   `json:"TotalHoldTime" xml:"TotalHoldTime"`
	TotalTalkTime                int64   `json:"TotalTalkTime" xml:"TotalTalkTime"`
	TotalWorkTime                int64   `json:"TotalWorkTime" xml:"TotalWorkTime"`
}
