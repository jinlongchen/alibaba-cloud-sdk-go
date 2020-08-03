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

// CreateKeyword invokes the green.CreateKeyword API synchronously
// api document: https://help.aliyun.com/api/green/createkeyword.html
func (client *Client) CreateKeyword(request *CreateKeywordRequest) (response *CreateKeywordResponse, err error) {
	response = CreateCreateKeywordResponse()
	err = client.DoAction(request, response)
	return
}

// CreateKeywordWithChan invokes the green.CreateKeyword API asynchronously
// api document: https://help.aliyun.com/api/green/createkeyword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateKeywordWithChan(request *CreateKeywordRequest) (<-chan *CreateKeywordResponse, <-chan error) {
	responseChan := make(chan *CreateKeywordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateKeyword(request)
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

// CreateKeywordWithCallback invokes the green.CreateKeyword API asynchronously
// api document: https://help.aliyun.com/api/green/createkeyword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateKeywordWithCallback(request *CreateKeywordRequest, callback func(response *CreateKeywordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateKeywordResponse
		var err error
		defer close(result)
		response, err = client.CreateKeyword(request)
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

// CreateKeywordRequest is the request struct for api CreateKeyword
type CreateKeywordRequest struct {
	*requests.RpcRequest
	Keywords     string           `position:"Query" name:"Keywords"`
	KeywordLibId requests.Integer `position:"Query" name:"KeywordLibId"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	Lang         string           `position:"Query" name:"Lang"`
}

// CreateKeywordResponse is the response struct for api CreateKeyword
type CreateKeywordResponse struct {
	*responses.BaseResponse
	RequestId          string         `json:"RequestId" xml:"RequestId"`
	SuccessCount       int            `json:"SuccessCount" xml:"SuccessCount"`
	InvalidKeywordList []string       `json:"InvalidKeywordList" xml:"InvalidKeywordList"`
	ValidKeywordList   []ValidKeyword `json:"validKeywordList" xml:"validKeywordList"`
}

// CreateCreateKeywordRequest creates a request to invoke CreateKeyword API
func CreateCreateKeywordRequest() (request *CreateKeywordRequest) {
	request = &CreateKeywordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "CreateKeyword", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateKeywordResponse creates a response to parse from CreateKeyword response
func CreateCreateKeywordResponse() (response *CreateKeywordResponse) {
	response = &CreateKeywordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
