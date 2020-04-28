package workorder

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

// ReplyTicket invokes the workorder.ReplyTicket API synchronously
// api document: https://help.aliyun.com/api/workorder/replyticket.html
func (client *Client) ReplyTicket(request *ReplyTicketRequest) (response *ReplyTicketResponse, err error) {
	response = CreateReplyTicketResponse()
	err = client.DoAction(request, response)
	return
}

// ReplyTicketWithChan invokes the workorder.ReplyTicket API asynchronously
// api document: https://help.aliyun.com/api/workorder/replyticket.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReplyTicketWithChan(request *ReplyTicketRequest) (<-chan *ReplyTicketResponse, <-chan error) {
	responseChan := make(chan *ReplyTicketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReplyTicket(request)
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

// ReplyTicketWithCallback invokes the workorder.ReplyTicket API asynchronously
// api document: https://help.aliyun.com/api/workorder/replyticket.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReplyTicketWithCallback(request *ReplyTicketRequest, callback func(response *ReplyTicketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReplyTicketResponse
		var err error
		defer close(result)
		response, err = client.ReplyTicket(request)
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

// ReplyTicketRequest is the request struct for api ReplyTicket
type ReplyTicketRequest struct {
	*requests.RpcRequest
	Language      string `position:"Query" name:"Language"`
	TicketId      string `position:"Query" name:"TicketId"`
	Content       string `position:"Query" name:"Content"`
	SecretContent string `position:"Query" name:"SecretContent"`
}

// ReplyTicketResponse is the response struct for api ReplyTicket
type ReplyTicketResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReplyTicketRequest creates a request to invoke ReplyTicket API
func CreateReplyTicketRequest() (request *ReplyTicketRequest) {
	request = &ReplyTicketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Workorder", "2020-03-26", "ReplyTicket", "workorder", "openAPI")
	return
}

// CreateReplyTicketResponse creates a response to parse from ReplyTicket response
func CreateReplyTicketResponse() (response *ReplyTicketResponse) {
	response = &ReplyTicketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}