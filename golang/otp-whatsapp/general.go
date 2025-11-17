package otpWhatsapp

import (
	"errors"
	"log"
)

type OTPProvider string

const (
	MESSAGE_OTP_CODE    = "[OTP_CODE]"
	MESSAGE_ACTION_CODE = "[ACTION_NAME]"
	MESSAGE_CLIENT_CODE = "[CLIENT_NAME]"
	MESSAGE_TEMPLATE    = "Welcome to TEST. Your request " + MESSAGE_ACTION_CODE + " requires OTP.  Here is your OTP CODE: " + MESSAGE_OTP_CODE

	VENDOR_WATZAP OTPProvider = "watzap"
	VENDOR_FONNTE OTPProvider = "fonnte"

	COUNTRY_PHONE_SEPARATOR = "."
	COUNTRY_PHONE_PREFIX    = "+"

	ERR_NOT_FOUND_VENDOR      = "not found vendor"
	ERR_PHONE_EMPTY           = "phone empty"
	ERR_NOT_AVAILABLE_SERVICE = "not available service"

	COUNTRY_CODE_INDONESIA = "62"
)

type OTPVendor struct {
	Title          string `db:"title"               form:"title"               json:"title"`
	APIKey         string `db:"api_key"             form:"api_key"             json:"api_key"`
	Url            string `db:"url"                 form:"url"                 json:"url"`
	NumberKey      string `db:"number_key"          form:"number_key"          json:"number_key"`
	DefaultMessage string `db:"default_message"     form:"default_message"     json:"default_message"`
}

type OTPProviderAPI interface {
	SendOtp(phone, otpCode string) (string, error)
}

func InitVendor(vendorName OTPProvider, apiKey, numberKey, defaultMessage string) (OTPProviderAPI, error) {
	var otp OTPProviderAPI
	defMessage := MESSAGE_TEMPLATE
	if defaultMessage != "" {
		defMessage = defaultMessage
	}
	switch OTPProvider(vendorName) {
	case VENDOR_WATZAP:
		var otpW WatzapOtp
		otpW.Vendor = OTPVendor{
			Title:          string(vendorName),
			APIKey:         apiKey,
			DefaultMessage: defMessage,
		}
		otpW.Request.ApiKey = apiKey
		otpW.Request.NumberKey = numberKey
		otpW.Request.Message = defMessage
		otp = &otpW
	case VENDOR_FONNTE:
		var otpF FonnteOtp
		otpF.Vendor = OTPVendor{
			Title:          string(vendorName),
			APIKey:         apiKey,
			DefaultMessage: defMessage,
		}
		otpF.Request.Target = numberKey
		otpF.Request.Message = defMessage
		otpF.Request.CountryCode = COUNTRY_CODE_INDONESIA
		otp = &otpF
	default:
		return nil, errors.New(ERR_NOT_FOUND_VENDOR)
	}

	return otp, nil
}

func SendOTP(vendor OTPProviderAPI, phone, otpCode string) (string, error) {
	resp, err := vendor.SendOtp(phone, otpCode)
	if err != nil {
		log.Println("error when send OTP . Err : ", err.Error())
		return "", err
	}
	return resp, nil
}
