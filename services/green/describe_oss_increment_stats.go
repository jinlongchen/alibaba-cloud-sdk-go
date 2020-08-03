package green

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeOssIncrementStats invokes the green.DescribeOssIncrementStats API synchronously
// api document: https://help.aliyun.com/api/green/describeossincrementstats.html
func (client *Client) DescribeOssIncrementStats(request *DescribeOssIncrementStatsRequest) (response *DescribeOssIncrementStatsResponse, err error) {
	response = CreateDescribeOssIncrementStatsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOssIncrementStatsWithChan invokes the green.DescribeOssIncrementStats API asynchronously
// api document: https://help.aliyun.com/api/green/describeossincrementstats.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOssIncrementStatsWithChan(request *DescribeOssIncrementStatsRequest) (<-chan *DescribeOssIncrementStatsResponse, <-chan error) {
	responseChan := make(chan *DescribeOssIncrementStatsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOssIncrementStats(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeOssIncrementStatsWithCallback invokes the green.DescribeOssIncrementStats API asynchronously
// api document: https://help.aliyun.com/api/green/describeossincrementstats.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOssIncrementStatsWithCallback(request *DescribeOssIncrementStatsRequest, callback func(response *DescribeOssIncrementStatsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOssIncrementStatsResponse
		var err error
		defer close(result)
		response, err = client.DescribeOssIncrementStats(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeOssIncrementStatsRequest is the request struct for api DescribeOssIncrementStats
type DescribeOssIncrementStatsRequest struct {
	*requests.RpcRequest
	StartDate    string `position:"Query" name:"StartDate"`
	ResourceType string `position:"Query" name:"ResourceType"`
	Scene        string `position:"Query" name:"Scene"`
	EndDate      string `position:"Query" name:"EndDate"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeOssIncrementStatsResponse is the response struct for api DescribeOssIncrementStats
type DescribeOssIncrementStatsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	StatList   []Stat `json:"StatList" xml:"StatList"`
}

// CreateDescribeOssIncrementStatsRequest creates a request to invoke DescribeOssIncrementStats API
func CreateDescribeOssIncrementStatsRequest() (request *DescribeOssIncrementStatsRequest) {
	request = &DescribeOssIncrementStatsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeOssIncrementStats", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeOssIncrementStatsResponse creates a response to parse from DescribeOssIncrementStats response
func CreateDescribeOssIncrementStatsResponse() (response *DescribeOssIncrementStatsResponse) {
	response = &DescribeOssIncrementStatsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
