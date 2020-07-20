package edas

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

// RollbackChangeOrder invokes the edas.RollbackChangeOrder API synchronously
// api document: https://help.aliyun.com/api/edas/rollbackchangeorder.html
func (client *Client) RollbackChangeOrder(request *RollbackChangeOrderRequest) (response *RollbackChangeOrderResponse, err error) {
	response = CreateRollbackChangeOrderResponse()
	err = client.DoAction(request, response)
	return
}

// RollbackChangeOrderWithChan invokes the edas.RollbackChangeOrder API asynchronously
// api document: https://help.aliyun.com/api/edas/rollbackchangeorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RollbackChangeOrderWithChan(request *RollbackChangeOrderRequest) (<-chan *RollbackChangeOrderResponse, <-chan error) {
	responseChan := make(chan *RollbackChangeOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RollbackChangeOrder(request)
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

// RollbackChangeOrderWithCallback invokes the edas.RollbackChangeOrder API asynchronously
// api document: https://help.aliyun.com/api/edas/rollbackchangeorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RollbackChangeOrderWithCallback(request *RollbackChangeOrderRequest, callback func(response *RollbackChangeOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RollbackChangeOrderResponse
		var err error
		defer close(result)
		response, err = client.RollbackChangeOrder(request)
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

// RollbackChangeOrderRequest is the request struct for api RollbackChangeOrder
type RollbackChangeOrderRequest struct {
	*requests.RoaRequest
	ChangeOrderId string `position:"Query" name:"ChangeOrderId"`
}

// RollbackChangeOrderResponse is the response struct for api RollbackChangeOrder
type RollbackChangeOrderResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRollbackChangeOrderRequest creates a request to invoke RollbackChangeOrder API
func CreateRollbackChangeOrderRequest() (request *RollbackChangeOrderRequest) {
	request = &RollbackChangeOrderRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "RollbackChangeOrder", "/pop/v5/oam/changeorder/rollback", "edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateRollbackChangeOrderResponse creates a response to parse from RollbackChangeOrder response
func CreateRollbackChangeOrderResponse() (response *RollbackChangeOrderResponse) {
	response = &RollbackChangeOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
