
package sts

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

// AssumeRole invokes the sts.AssumeRole API synchronously
// api document: https://help.aliyun.com/api/sts/assumerole.html
func (client *Client) AssumeRole(request *AssumeRoleRequest) (response *AssumeRoleResponse, err error) {
response = CreateAssumeRoleResponse()
err = client.DoAction(request, response)
return
}

// AssumeRoleWithChan invokes the sts.AssumeRole API asynchronously
// api document: https://help.aliyun.com/api/sts/assumerole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssumeRoleWithChan(request *AssumeRoleRequest) (<-chan *AssumeRoleResponse, <-chan error) {
responseChan := make(chan *AssumeRoleResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.AssumeRole(request)
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

// AssumeRoleWithCallback invokes the sts.AssumeRole API asynchronously
// api document: https://help.aliyun.com/api/sts/assumerole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssumeRoleWithCallback(request *AssumeRoleRequest, callback func(response *AssumeRoleResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *AssumeRoleResponse
var err error
defer close(result)
response, err = client.AssumeRole(request)
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

// AssumeRoleRequest is the request struct for api AssumeRole
type AssumeRoleRequest struct {
*requests.RpcRequest
                    RoleSessionName     string `position:"Query" name:"RoleSessionName"`
                    Policy     string `position:"Query" name:"Policy"`
                    RoleArn     string `position:"Query" name:"RoleArn"`
                    DurationSeconds     requests.Integer `position:"Query" name:"DurationSeconds"`
}


// AssumeRoleResponse is the response struct for api AssumeRole
type AssumeRoleResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            Credentials Credentials  `json:"Credentials" xml:"Credentials"`
            AssumedRoleUser AssumedRoleUser  `json:"AssumedRoleUser" xml:"AssumedRoleUser"`
}

// CreateAssumeRoleRequest creates a request to invoke AssumeRole API
func CreateAssumeRoleRequest() (request *AssumeRoleRequest) {
request = &AssumeRoleRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Sts", "2015-04-01", "AssumeRole", "sts", "openAPI")
return
}

// CreateAssumeRoleResponse creates a response to parse from AssumeRole response
func CreateAssumeRoleResponse() (response *AssumeRoleResponse) {
response = &AssumeRoleResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}


