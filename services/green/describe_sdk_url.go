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

// DescribeSdkUrl invokes the green.DescribeSdkUrl API synchronously
// api document: https://help.aliyun.com/api/green/describesdkurl.html
func (client *Client) DescribeSdkUrl(request *DescribeSdkUrlRequest) (response *DescribeSdkUrlResponse, err error) {
	response = CreateDescribeSdkUrlResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSdkUrlWithChan invokes the green.DescribeSdkUrl API asynchronously
// api document: https://help.aliyun.com/api/green/describesdkurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSdkUrlWithChan(request *DescribeSdkUrlRequest) (<-chan *DescribeSdkUrlResponse, <-chan error) {
	responseChan := make(chan *DescribeSdkUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSdkUrl(request)
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

// DescribeSdkUrlWithCallback invokes the green.DescribeSdkUrl API asynchronously
// api document: https://help.aliyun.com/api/green/describesdkurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSdkUrlWithCallback(request *DescribeSdkUrlRequest, callback func(response *DescribeSdkUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSdkUrlResponse
		var err error
		defer close(result)
		response, err = client.DescribeSdkUrl(request)
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

// DescribeSdkUrlRequest is the request struct for api DescribeSdkUrl
type DescribeSdkUrlRequest struct {
	*requests.RpcRequest
	Debug    requests.Boolean `position:"Query" name:"Debug"`
	SourceIp string           `position:"Query" name:"SourceIp"`
	Id       requests.Integer `position:"Query" name:"Id"`
	Lang     string           `position:"Query" name:"Lang"`
}

// DescribeSdkUrlResponse is the response struct for api DescribeSdkUrl
type DescribeSdkUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SdkUrl    string `json:"SdkUrl" xml:"SdkUrl"`
}

// CreateDescribeSdkUrlRequest creates a request to invoke DescribeSdkUrl API
func CreateDescribeSdkUrlRequest() (request *DescribeSdkUrlRequest) {
	request = &DescribeSdkUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeSdkUrl", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSdkUrlResponse creates a response to parse from DescribeSdkUrl response
func CreateDescribeSdkUrlResponse() (response *DescribeSdkUrlResponse) {
	response = &DescribeSdkUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
