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

// UpdateWebsiteInstanceKeyUrl invokes the green.UpdateWebsiteInstanceKeyUrl API synchronously
// api document: https://help.aliyun.com/api/green/updatewebsiteinstancekeyurl.html
func (client *Client) UpdateWebsiteInstanceKeyUrl(request *UpdateWebsiteInstanceKeyUrlRequest) (response *UpdateWebsiteInstanceKeyUrlResponse, err error) {
	response = CreateUpdateWebsiteInstanceKeyUrlResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWebsiteInstanceKeyUrlWithChan invokes the green.UpdateWebsiteInstanceKeyUrl API asynchronously
// api document: https://help.aliyun.com/api/green/updatewebsiteinstancekeyurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateWebsiteInstanceKeyUrlWithChan(request *UpdateWebsiteInstanceKeyUrlRequest) (<-chan *UpdateWebsiteInstanceKeyUrlResponse, <-chan error) {
	responseChan := make(chan *UpdateWebsiteInstanceKeyUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWebsiteInstanceKeyUrl(request)
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

// UpdateWebsiteInstanceKeyUrlWithCallback invokes the green.UpdateWebsiteInstanceKeyUrl API asynchronously
// api document: https://help.aliyun.com/api/green/updatewebsiteinstancekeyurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateWebsiteInstanceKeyUrlWithCallback(request *UpdateWebsiteInstanceKeyUrlRequest, callback func(response *UpdateWebsiteInstanceKeyUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWebsiteInstanceKeyUrlResponse
		var err error
		defer close(result)
		response, err = client.UpdateWebsiteInstanceKeyUrl(request)
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

// UpdateWebsiteInstanceKeyUrlRequest is the request struct for api UpdateWebsiteInstanceKeyUrl
type UpdateWebsiteInstanceKeyUrlRequest struct {
	*requests.RpcRequest
	Urls       string `position:"Query" name:"Urls"`
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// UpdateWebsiteInstanceKeyUrlResponse is the response struct for api UpdateWebsiteInstanceKeyUrl
type UpdateWebsiteInstanceKeyUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateWebsiteInstanceKeyUrlRequest creates a request to invoke UpdateWebsiteInstanceKeyUrl API
func CreateUpdateWebsiteInstanceKeyUrlRequest() (request *UpdateWebsiteInstanceKeyUrlRequest) {
	request = &UpdateWebsiteInstanceKeyUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "UpdateWebsiteInstanceKeyUrl", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateWebsiteInstanceKeyUrlResponse creates a response to parse from UpdateWebsiteInstanceKeyUrl response
func CreateUpdateWebsiteInstanceKeyUrlResponse() (response *UpdateWebsiteInstanceKeyUrlResponse) {
	response = &UpdateWebsiteInstanceKeyUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
