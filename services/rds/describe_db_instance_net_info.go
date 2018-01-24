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

func (client *Client) DescribeDBInstanceNetInfo(request *DescribeDBInstanceNetInfoRequest) (response *DescribeDBInstanceNetInfoResponse, err error) {
	response = CreateDescribeDBInstanceNetInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDBInstanceNetInfoWithChan(request *DescribeDBInstanceNetInfoRequest) (<-chan *DescribeDBInstanceNetInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceNetInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceNetInfo(request)
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

func (client *Client) DescribeDBInstanceNetInfoWithCallback(request *DescribeDBInstanceNetInfoRequest, callback func(response *DescribeDBInstanceNetInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceNetInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceNetInfo(request)
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

type DescribeDBInstanceNetInfoRequest struct {
	*requests.RpcRequest
	DBInstanceId             string           `position:"Query" name:"DBInstanceId"`
	Flag                     string           `position:"Query" name:"Flag"`
	ClientToken              string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceNetRWSplitType string           `position:"Query" name:"DBInstanceNetRWSplitType"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeDBInstanceNetInfoResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	InstanceNetworkType string `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	DBInstanceNetInfos  struct {
		DBInstanceNetInfo []struct {
			Upgradeable          string `json:"Upgradeable" xml:"Upgradeable"`
			ExpiredTime          string `json:"ExpiredTime" xml:"ExpiredTime"`
			ConnectionString     string `json:"ConnectionString" xml:"ConnectionString"`
			IPAddress            string `json:"IPAddress" xml:"IPAddress"`
			IPType               string `json:"IPType" xml:"IPType"`
			Port                 string `json:"Port" xml:"Port"`
			VPCId                string `json:"VPCId" xml:"VPCId"`
			VSwitchId            string `json:"VSwitchId" xml:"VSwitchId"`
			ConnectionStringType string `json:"ConnectionStringType" xml:"ConnectionStringType"`
			MaxDelayTime         string `json:"MaxDelayTime" xml:"MaxDelayTime"`
			DistributionType     string `json:"DistributionType" xml:"DistributionType"`
			SecurityIPGroups     struct {
				SecurityIPGroup []struct {
					SecurityIPGroupName string `json:"SecurityIPGroupName" xml:"SecurityIPGroupName"`
					SecurityIPs         string `json:"SecurityIPs" xml:"SecurityIPs"`
				} `json:"securityIPGroup" xml:"securityIPGroup"`
			} `json:"SecurityIPGroups" xml:"SecurityIPGroups"`
			DBInstanceWeights struct {
				DBInstanceWeight []struct {
					DBInstanceId   string `json:"DBInstanceId" xml:"DBInstanceId"`
					DBInstanceType string `json:"DBInstanceType" xml:"DBInstanceType"`
					Availability   string `json:"Availability" xml:"Availability"`
					Weight         string `json:"Weight" xml:"Weight"`
				} `json:"DBInstanceWeight" xml:"DBInstanceWeight"`
			} `json:"DBInstanceWeights" xml:"DBInstanceWeights"`
		} `json:"DBInstanceNetInfo" xml:"DBInstanceNetInfo"`
	} `json:"DBInstanceNetInfos" xml:"DBInstanceNetInfos"`
}

func CreateDescribeDBInstanceNetInfoRequest() (request *DescribeDBInstanceNetInfoRequest) {
	request = &DescribeDBInstanceNetInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceNetInfo", "rds", "openAPI")
	return
}

func CreateDescribeDBInstanceNetInfoResponse() (response *DescribeDBInstanceNetInfoResponse) {
	response = &DescribeDBInstanceNetInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
