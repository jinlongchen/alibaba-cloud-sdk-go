package ecs

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

func (client *Client) DescribeInvocationResults(request *DescribeInvocationResultsRequest) (response *DescribeInvocationResultsResponse, err error) {
	response = CreateDescribeInvocationResultsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInvocationResultsWithChan(request *DescribeInvocationResultsRequest) (<-chan *DescribeInvocationResultsResponse, <-chan error) {
	responseChan := make(chan *DescribeInvocationResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInvocationResults(request)
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

func (client *Client) DescribeInvocationResultsWithCallback(request *DescribeInvocationResultsRequest, callback func(response *DescribeInvocationResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInvocationResultsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInvocationResults(request)
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

type DescribeInvocationResultsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CommandId            string           `position:"Query" name:"CommandId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	InvokeId             string           `position:"Query" name:"InvokeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	InvokeRecordStatus   string           `position:"Query" name:"InvokeRecordStatus"`
}

type DescribeInvocationResultsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Invocation struct {
		PageSize          int `json:"PageSize" xml:"PageSize"`
		PageNumber        int `json:"PageNumber" xml:"PageNumber"`
		TotalCount        int `json:"TotalCount" xml:"TotalCount"`
		InvocationResults struct {
			InvocationResult []struct {
				CommandId          string `json:"CommandId" xml:"CommandId"`
				InvokeId           string `json:"InvokeId" xml:"InvokeId"`
				InstanceId         string `json:"InstanceId" xml:"InstanceId"`
				FinishedTime       string `json:"FinishedTime" xml:"FinishedTime"`
				Output             string `json:"Output" xml:"Output"`
				InvokeRecordStatus string `json:"InvokeRecordStatus" xml:"InvokeRecordStatus"`
				ExitCode           int    `json:"ExitCode" xml:"ExitCode"`
			} `json:"InvocationResult" xml:"InvocationResult"`
		} `json:"InvocationResults" xml:"InvocationResults"`
	} `json:"Invocation" xml:"Invocation"`
}

func CreateDescribeInvocationResultsRequest() (request *DescribeInvocationResultsRequest) {
	request = &DescribeInvocationResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInvocationResults", "ecs", "openAPI")
	return
}

func CreateDescribeInvocationResultsResponse() (response *DescribeInvocationResultsResponse) {
	response = &DescribeInvocationResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
