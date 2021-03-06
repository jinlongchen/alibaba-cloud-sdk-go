package imm

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

// FacesItemInGetImage is a nested struct in imm response
type FacesItemInGetImage struct {
	FaceId            string         `json:"FaceId" xml:"FaceId"`
	FaceConfidence    float64        `json:"FaceConfidence" xml:"FaceConfidence"`
	Age               string         `json:"Age" xml:"Age"`
	Gender            string         `json:"Gender" xml:"Gender"`
	Emotion           string         `json:"Emotion" xml:"Emotion"`
	Attractive        float64        `json:"Attractive" xml:"Attractive"`
	GenderConfidence  float64        `json:"GenderConfidence" xml:"GenderConfidence"`
	GroupId           string         `json:"GroupId" xml:"GroupId"`
	FaceQuality       float64        `json:"FaceQuality" xml:"FaceQuality"`
	EmotionConfidence float64        `json:"EmotionConfidence" xml:"EmotionConfidence"`
	FaceAttributes    FaceAttributes `json:"FaceAttributes" xml:"FaceAttributes"`
	EmotionDetails    EmotionDetails `json:"EmotionDetails" xml:"EmotionDetails"`
}
