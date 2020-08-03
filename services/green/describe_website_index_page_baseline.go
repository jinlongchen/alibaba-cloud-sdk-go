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

// DescribeWebsiteIndexPageBaseline invokes the green.DescribeWebsiteIndexPageBaseline API synchronously
// api document: https://help.aliyun.com/api/green/describewebsiteindexpagebaseline.html
func (client *Client) DescribeWebsiteIndexPageBaseline(request *DescribeWebsiteIndexPageBaselineRequest) (response *DescribeWebsiteIndexPageBaselineResponse, err error) {
	response = CreateDescribeWebsiteIndexPageBaselineResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebsiteIndexPageBaselineWithChan invokes the green.DescribeWebsiteIndexPageBaseline API asynchronously
// api document: https://help.aliyun.com/api/green/describewebsiteindexpagebaseline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebsiteIndexPageBaselineWithChan(request *DescribeWebsiteIndexPageBaselineRequest) (<-chan *DescribeWebsiteIndexPageBaselineResponse, <-chan error) {
	responseChan := make(chan *DescribeWebsiteIndexPageBaselineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebsiteIndexPageBaseline(request)
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

// DescribeWebsiteIndexPageBaselineWithCallback invokes the green.DescribeWebsiteIndexPageBaseline API asynchronously
// api document: https://help.aliyun.com/api/green/describewebsiteindexpagebaseline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebsiteIndexPageBaselineWithCallback(request *DescribeWebsiteIndexPageBaselineRequest, callback func(response *DescribeWebsiteIndexPageBaselineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebsiteIndexPageBaselineResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebsiteIndexPageBaseline(request)
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

// DescribeWebsiteIndexPageBaselineRequest is the request struct for api DescribeWebsiteIndexPageBaseline
type DescribeWebsiteIndexPageBaselineRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// DescribeWebsiteIndexPageBaselineResponse is the response struct for api DescribeWebsiteIndexPageBaseline
type DescribeWebsiteIndexPageBaselineResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	CreateTime     string `json:"CreateTime" xml:"CreateTime"`
	Snapshot       string `json:"Snapshot" xml:"Snapshot"`
	BaseLineStatus string `json:"BaseLineStatus" xml:"BaseLineStatus"`
}

// CreateDescribeWebsiteIndexPageBaselineRequest creates a request to invoke DescribeWebsiteIndexPageBaseline API
func CreateDescribeWebsiteIndexPageBaselineRequest() (request *DescribeWebsiteIndexPageBaselineRequest) {
	request = &DescribeWebsiteIndexPageBaselineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeWebsiteIndexPageBaseline", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeWebsiteIndexPageBaselineResponse creates a response to parse from DescribeWebsiteIndexPageBaseline response
func CreateDescribeWebsiteIndexPageBaselineResponse() (response *DescribeWebsiteIndexPageBaselineResponse) {
	response = &DescribeWebsiteIndexPageBaselineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
