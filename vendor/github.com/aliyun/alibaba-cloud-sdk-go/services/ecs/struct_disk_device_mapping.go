package ecs

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

// DiskDeviceMapping is a nested struct in ecs response
type DiskDeviceMapping struct {
	OSSBucket       string `json:"OSSBucket" xml:"OSSBucket"`
	OSSObject       string `json:"OSSObject" xml:"OSSObject"`
	Format          string `json:"Format" xml:"Format"`
	Device          string `json:"Device" xml:"Device"`
	ImportOSSObject string `json:"ImportOSSObject" xml:"ImportOSSObject"`
	Encrypted       bool   `json:"Encrypted" xml:"Encrypted"`
	Progress        string `json:"Progress" xml:"Progress"`
	DiskImageSize   int    `json:"DiskImageSize" xml:"DiskImageSize"`
	Size            string `json:"Size" xml:"Size"`
	RemainTime      int    `json:"RemainTime" xml:"RemainTime"`
	SnapshotId      string `json:"SnapshotId" xml:"SnapshotId"`
	ImportOSSBucket string `json:"ImportOSSBucket" xml:"ImportOSSBucket"`
	Type            string `json:"Type" xml:"Type"`
}
