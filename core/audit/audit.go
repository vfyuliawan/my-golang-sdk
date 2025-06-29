package audit

import (
	"time"
)

type Payload struct {
	UserId           int64     `json:"userId"`
	AccountNumber    string    `json:"accountNumber"`
	ChannelCode      string    `json:"channelCode"`
	ActivityDate     time.Time `json:"activityDate"`
	ActivityCategory string    `json:"activityCategory"`
	ActivityName     string    `json:"activityName"`
	ActivityRefCode  string    `json:"activityRefCode"`
	ActivityStatus   string    `json:"activityStatus"`
	DeviceId         string    `json:"deviceId"`
	PhoneOs          string    `json:"phoneOs"`
	PhoneType        string    `json:"phoneType"`
	PhoneBrand       string    `json:"phoneBrand"`
	UserAgent        string    `json:"userAgent"`
	IpAddress        string    `json:"ipAddress"`
	AppVersion       string    `json:"appVersion"`
	CoreRefNo        string    `json:"coreRefNo"`
	TransactionId    string    `json:"transactionId"`
	SourceSystem     string    `json:"sourceSystem"`
	ErrorCode        string    `json:"errorCode"`
	ErrorMessage     string    `json:"errorMessage"`
	AdditionalData   string    `json:"additionalData"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	CreatedBy        string    `json:"createdBy"`
}
