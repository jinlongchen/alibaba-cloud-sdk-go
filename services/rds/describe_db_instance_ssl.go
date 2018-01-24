package rds

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

func (client *Client) DescribeDBInstanceSSL(request *DescribeDBInstanceSSLRequest) (response *DescribeDBInstanceSSLResponse, err error) {
	response = CreateDescribeDBInstanceSSLResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDBInstanceSSLWithChan(request *DescribeDBInstanceSSLRequest) (<-chan *DescribeDBInstanceSSLResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceSSLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceSSL(request)
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

func (client *Client) DescribeDBInstanceSSLWithCallback(request *DescribeDBInstanceSSLRequest, callback func(response *DescribeDBInstanceSSLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceSSLResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceSSL(request)
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

type DescribeDBInstanceSSLRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeDBInstanceSSLResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	ConnectionString    string `json:"ConnectionString" xml:"ConnectionString"`
	SSLExpireTime       string `json:"SSLExpireTime" xml:"SSLExpireTime"`
	RequireUpdate       string `json:"RequireUpdate" xml:"RequireUpdate"`
	RequireUpdateReason string `json:"RequireUpdateReason" xml:"RequireUpdateReason"`
}

func CreateDescribeDBInstanceSSLRequest() (request *DescribeDBInstanceSSLRequest) {
	request = &DescribeDBInstanceSSLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceSSL", "rds", "openAPI")
	return
}

func CreateDescribeDBInstanceSSLResponse() (response *DescribeDBInstanceSSLResponse) {
	response = &DescribeDBInstanceSSLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
