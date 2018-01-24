package domain_intl

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

func (client *Client) SaveBatchTaskForTransferProhibitionLock(request *SaveBatchTaskForTransferProhibitionLockRequest) (response *SaveBatchTaskForTransferProhibitionLockResponse, err error) {
	response = CreateSaveBatchTaskForTransferProhibitionLockResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SaveBatchTaskForTransferProhibitionLockWithChan(request *SaveBatchTaskForTransferProhibitionLockRequest) (<-chan *SaveBatchTaskForTransferProhibitionLockResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForTransferProhibitionLockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForTransferProhibitionLock(request)
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

func (client *Client) SaveBatchTaskForTransferProhibitionLockWithCallback(request *SaveBatchTaskForTransferProhibitionLockRequest, callback func(response *SaveBatchTaskForTransferProhibitionLockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForTransferProhibitionLockResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForTransferProhibitionLock(request)
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

type SaveBatchTaskForTransferProhibitionLockRequest struct {
	*requests.RpcRequest
	DomainName   *[]string        `position:"Query" name:"DomainName"  type:"Repeated"`
	Status       requests.Boolean `position:"Query" name:"Status"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
}

type SaveBatchTaskForTransferProhibitionLockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

func CreateSaveBatchTaskForTransferProhibitionLockRequest() (request *SaveBatchTaskForTransferProhibitionLockRequest) {
	request = &SaveBatchTaskForTransferProhibitionLockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForTransferProhibitionLock", "", "")
	return
}

func CreateSaveBatchTaskForTransferProhibitionLockResponse() (response *SaveBatchTaskForTransferProhibitionLockResponse) {
	response = &SaveBatchTaskForTransferProhibitionLockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
