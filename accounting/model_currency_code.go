/*
Xero Accounting API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.40.0
Contact: api@xero.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CurrencyCode 3 letter alpha code for the currency – see list of currency codes
type CurrencyCode string

// List of CurrencyCode
const (
	CURRENCYCODE_AED CurrencyCode = "AED"
	CURRENCYCODE_AFN CurrencyCode = "AFN"
	CURRENCYCODE_ALL CurrencyCode = "ALL"
	CURRENCYCODE_AMD CurrencyCode = "AMD"
	CURRENCYCODE_ANG CurrencyCode = "ANG"
	CURRENCYCODE_AOA CurrencyCode = "AOA"
	CURRENCYCODE_ARS CurrencyCode = "ARS"
	CURRENCYCODE_AUD CurrencyCode = "AUD"
	CURRENCYCODE_AWG CurrencyCode = "AWG"
	CURRENCYCODE_AZN CurrencyCode = "AZN"
	CURRENCYCODE_BAM CurrencyCode = "BAM"
	CURRENCYCODE_BBD CurrencyCode = "BBD"
	CURRENCYCODE_BDT CurrencyCode = "BDT"
	CURRENCYCODE_BGN CurrencyCode = "BGN"
	CURRENCYCODE_BHD CurrencyCode = "BHD"
	CURRENCYCODE_BIF CurrencyCode = "BIF"
	CURRENCYCODE_BMD CurrencyCode = "BMD"
	CURRENCYCODE_BND CurrencyCode = "BND"
	CURRENCYCODE_BOB CurrencyCode = "BOB"
	CURRENCYCODE_BRL CurrencyCode = "BRL"
	CURRENCYCODE_BSD CurrencyCode = "BSD"
	CURRENCYCODE_BTN CurrencyCode = "BTN"
	CURRENCYCODE_BWP CurrencyCode = "BWP"
	CURRENCYCODE_BYN CurrencyCode = "BYN"
	CURRENCYCODE_BYR CurrencyCode = "BYR"
	CURRENCYCODE_BZD CurrencyCode = "BZD"
	CURRENCYCODE_CAD CurrencyCode = "CAD"
	CURRENCYCODE_CDF CurrencyCode = "CDF"
	CURRENCYCODE_CHF CurrencyCode = "CHF"
	CURRENCYCODE_CLF CurrencyCode = "CLF"
	CURRENCYCODE_CLP CurrencyCode = "CLP"
	CURRENCYCODE_CNY CurrencyCode = "CNY"
	CURRENCYCODE_COP CurrencyCode = "COP"
	CURRENCYCODE_CRC CurrencyCode = "CRC"
	CURRENCYCODE_CUC CurrencyCode = "CUC"
	CURRENCYCODE_CUP CurrencyCode = "CUP"
	CURRENCYCODE_CVE CurrencyCode = "CVE"
	CURRENCYCODE_CZK CurrencyCode = "CZK"
	CURRENCYCODE_DJF CurrencyCode = "DJF"
	CURRENCYCODE_DKK CurrencyCode = "DKK"
	CURRENCYCODE_DOP CurrencyCode = "DOP"
	CURRENCYCODE_DZD CurrencyCode = "DZD"
	CURRENCYCODE_EEK CurrencyCode = "EEK"
	CURRENCYCODE_EGP CurrencyCode = "EGP"
	CURRENCYCODE_ERN CurrencyCode = "ERN"
	CURRENCYCODE_ETB CurrencyCode = "ETB"
	CURRENCYCODE_EUR CurrencyCode = "EUR"
	CURRENCYCODE_FJD CurrencyCode = "FJD"
	CURRENCYCODE_FKP CurrencyCode = "FKP"
	CURRENCYCODE_GBP CurrencyCode = "GBP"
	CURRENCYCODE_GEL CurrencyCode = "GEL"
	CURRENCYCODE_GHS CurrencyCode = "GHS"
	CURRENCYCODE_GIP CurrencyCode = "GIP"
	CURRENCYCODE_GMD CurrencyCode = "GMD"
	CURRENCYCODE_GNF CurrencyCode = "GNF"
	CURRENCYCODE_GTQ CurrencyCode = "GTQ"
	CURRENCYCODE_GYD CurrencyCode = "GYD"
	CURRENCYCODE_HKD CurrencyCode = "HKD"
	CURRENCYCODE_HNL CurrencyCode = "HNL"
	CURRENCYCODE_HRK CurrencyCode = "HRK"
	CURRENCYCODE_HTG CurrencyCode = "HTG"
	CURRENCYCODE_HUF CurrencyCode = "HUF"
	CURRENCYCODE_IDR CurrencyCode = "IDR"
	CURRENCYCODE_ILS CurrencyCode = "ILS"
	CURRENCYCODE_INR CurrencyCode = "INR"
	CURRENCYCODE_IQD CurrencyCode = "IQD"
	CURRENCYCODE_IRR CurrencyCode = "IRR"
	CURRENCYCODE_ISK CurrencyCode = "ISK"
	CURRENCYCODE_JMD CurrencyCode = "JMD"
	CURRENCYCODE_JOD CurrencyCode = "JOD"
	CURRENCYCODE_JPY CurrencyCode = "JPY"
	CURRENCYCODE_KES CurrencyCode = "KES"
	CURRENCYCODE_KGS CurrencyCode = "KGS"
	CURRENCYCODE_KHR CurrencyCode = "KHR"
	CURRENCYCODE_KMF CurrencyCode = "KMF"
	CURRENCYCODE_KPW CurrencyCode = "KPW"
	CURRENCYCODE_KRW CurrencyCode = "KRW"
	CURRENCYCODE_KWD CurrencyCode = "KWD"
	CURRENCYCODE_KYD CurrencyCode = "KYD"
	CURRENCYCODE_KZT CurrencyCode = "KZT"
	CURRENCYCODE_LAK CurrencyCode = "LAK"
	CURRENCYCODE_LBP CurrencyCode = "LBP"
	CURRENCYCODE_LKR CurrencyCode = "LKR"
	CURRENCYCODE_LRD CurrencyCode = "LRD"
	CURRENCYCODE_LSL CurrencyCode = "LSL"
	CURRENCYCODE_LTL CurrencyCode = "LTL"
	CURRENCYCODE_LVL CurrencyCode = "LVL"
	CURRENCYCODE_LYD CurrencyCode = "LYD"
	CURRENCYCODE_MAD CurrencyCode = "MAD"
	CURRENCYCODE_MDL CurrencyCode = "MDL"
	CURRENCYCODE_MGA CurrencyCode = "MGA"
	CURRENCYCODE_MKD CurrencyCode = "MKD"
	CURRENCYCODE_MMK CurrencyCode = "MMK"
	CURRENCYCODE_MNT CurrencyCode = "MNT"
	CURRENCYCODE_MOP CurrencyCode = "MOP"
	CURRENCYCODE_MRO CurrencyCode = "MRO"
	CURRENCYCODE_MRU CurrencyCode = "MRU"
	CURRENCYCODE_MUR CurrencyCode = "MUR"
	CURRENCYCODE_MVR CurrencyCode = "MVR"
	CURRENCYCODE_MWK CurrencyCode = "MWK"
	CURRENCYCODE_MXN CurrencyCode = "MXN"
	CURRENCYCODE_MXV CurrencyCode = "MXV"
	CURRENCYCODE_MYR CurrencyCode = "MYR"
	CURRENCYCODE_MZN CurrencyCode = "MZN"
	CURRENCYCODE_NAD CurrencyCode = "NAD"
	CURRENCYCODE_NGN CurrencyCode = "NGN"
	CURRENCYCODE_NIO CurrencyCode = "NIO"
	CURRENCYCODE_NOK CurrencyCode = "NOK"
	CURRENCYCODE_NPR CurrencyCode = "NPR"
	CURRENCYCODE_NZD CurrencyCode = "NZD"
	CURRENCYCODE_OMR CurrencyCode = "OMR"
	CURRENCYCODE_PAB CurrencyCode = "PAB"
	CURRENCYCODE_PEN CurrencyCode = "PEN"
	CURRENCYCODE_PGK CurrencyCode = "PGK"
	CURRENCYCODE_PHP CurrencyCode = "PHP"
	CURRENCYCODE_PKR CurrencyCode = "PKR"
	CURRENCYCODE_PLN CurrencyCode = "PLN"
	CURRENCYCODE_PYG CurrencyCode = "PYG"
	CURRENCYCODE_QAR CurrencyCode = "QAR"
	CURRENCYCODE_RON CurrencyCode = "RON"
	CURRENCYCODE_RSD CurrencyCode = "RSD"
	CURRENCYCODE_RUB CurrencyCode = "RUB"
	CURRENCYCODE_RWF CurrencyCode = "RWF"
	CURRENCYCODE_SAR CurrencyCode = "SAR"
	CURRENCYCODE_SBD CurrencyCode = "SBD"
	CURRENCYCODE_SCR CurrencyCode = "SCR"
	CURRENCYCODE_SDG CurrencyCode = "SDG"
	CURRENCYCODE_SEK CurrencyCode = "SEK"
	CURRENCYCODE_SGD CurrencyCode = "SGD"
	CURRENCYCODE_SHP CurrencyCode = "SHP"
	CURRENCYCODE_SKK CurrencyCode = "SKK"
	CURRENCYCODE_SLE CurrencyCode = "SLE"
	CURRENCYCODE_SLL CurrencyCode = "SLL"
	CURRENCYCODE_SOS CurrencyCode = "SOS"
	CURRENCYCODE_SRD CurrencyCode = "SRD"
	CURRENCYCODE_STN CurrencyCode = "STD"
	CURRENCYCODE_STD CurrencyCode = "STN"
	CURRENCYCODE_SVC CurrencyCode = "SVC"
	CURRENCYCODE_SYP CurrencyCode = "SYP"
	CURRENCYCODE_SZL CurrencyCode = "SZL"
	CURRENCYCODE_THB CurrencyCode = "THB"
	CURRENCYCODE_TJS CurrencyCode = "TJS"
	CURRENCYCODE_TMT CurrencyCode = "TMT"
	CURRENCYCODE_TND CurrencyCode = "TND"
	CURRENCYCODE_TOP CurrencyCode = "TOP"
	CURRENCYCODE_TRY_LIRA CurrencyCode = "TRY"
	CURRENCYCODE_TTD CurrencyCode = "TTD"
	CURRENCYCODE_TWD CurrencyCode = "TWD"
	CURRENCYCODE_TZS CurrencyCode = "TZS"
	CURRENCYCODE_UAH CurrencyCode = "UAH"
	CURRENCYCODE_UGX CurrencyCode = "UGX"
	CURRENCYCODE_USD CurrencyCode = "USD"
	CURRENCYCODE_UYU CurrencyCode = "UYU"
	CURRENCYCODE_UZS CurrencyCode = "UZS"
	CURRENCYCODE_VEF CurrencyCode = "VEF"
	CURRENCYCODE_VES CurrencyCode = "VES"
	CURRENCYCODE_VND CurrencyCode = "VND"
	CURRENCYCODE_VUV CurrencyCode = "VUV"
	CURRENCYCODE_WST CurrencyCode = "WST"
	CURRENCYCODE_XAF CurrencyCode = "XAF"
	CURRENCYCODE_XCD CurrencyCode = "XCD"
	CURRENCYCODE_XOF CurrencyCode = "XOF"
	CURRENCYCODE_XPF CurrencyCode = "XPF"
	CURRENCYCODE_YER CurrencyCode = "YER"
	CURRENCYCODE_ZAR CurrencyCode = "ZAR"
	CURRENCYCODE_ZMW CurrencyCode = "ZMW"
	CURRENCYCODE_ZMK CurrencyCode = "ZMK"
	CURRENCYCODE_ZWD CurrencyCode = "ZWD"
	CURRENCYCODE_EMPTY_CURRENCY CurrencyCode = ""
)

// All allowed values of CurrencyCode enum
var AllowedCurrencyCodeEnumValues = []CurrencyCode{
	"AED",
	"AFN",
	"ALL",
	"AMD",
	"ANG",
	"AOA",
	"ARS",
	"AUD",
	"AWG",
	"AZN",
	"BAM",
	"BBD",
	"BDT",
	"BGN",
	"BHD",
	"BIF",
	"BMD",
	"BND",
	"BOB",
	"BRL",
	"BSD",
	"BTN",
	"BWP",
	"BYN",
	"BYR",
	"BZD",
	"CAD",
	"CDF",
	"CHF",
	"CLF",
	"CLP",
	"CNY",
	"COP",
	"CRC",
	"CUC",
	"CUP",
	"CVE",
	"CZK",
	"DJF",
	"DKK",
	"DOP",
	"DZD",
	"EEK",
	"EGP",
	"ERN",
	"ETB",
	"EUR",
	"FJD",
	"FKP",
	"GBP",
	"GEL",
	"GHS",
	"GIP",
	"GMD",
	"GNF",
	"GTQ",
	"GYD",
	"HKD",
	"HNL",
	"HRK",
	"HTG",
	"HUF",
	"IDR",
	"ILS",
	"INR",
	"IQD",
	"IRR",
	"ISK",
	"JMD",
	"JOD",
	"JPY",
	"KES",
	"KGS",
	"KHR",
	"KMF",
	"KPW",
	"KRW",
	"KWD",
	"KYD",
	"KZT",
	"LAK",
	"LBP",
	"LKR",
	"LRD",
	"LSL",
	"LTL",
	"LVL",
	"LYD",
	"MAD",
	"MDL",
	"MGA",
	"MKD",
	"MMK",
	"MNT",
	"MOP",
	"MRO",
	"MRU",
	"MUR",
	"MVR",
	"MWK",
	"MXN",
	"MXV",
	"MYR",
	"MZN",
	"NAD",
	"NGN",
	"NIO",
	"NOK",
	"NPR",
	"NZD",
	"OMR",
	"PAB",
	"PEN",
	"PGK",
	"PHP",
	"PKR",
	"PLN",
	"PYG",
	"QAR",
	"RON",
	"RSD",
	"RUB",
	"RWF",
	"SAR",
	"SBD",
	"SCR",
	"SDG",
	"SEK",
	"SGD",
	"SHP",
	"SKK",
	"SLE",
	"SLL",
	"SOS",
	"SRD",
	"STD",
	"STN",
	"SVC",
	"SYP",
	"SZL",
	"THB",
	"TJS",
	"TMT",
	"TND",
	"TOP",
	"TRY",
	"TTD",
	"TWD",
	"TZS",
	"UAH",
	"UGX",
	"USD",
	"UYU",
	"UZS",
	"VEF",
	"VES",
	"VND",
	"VUV",
	"WST",
	"XAF",
	"XCD",
	"XOF",
	"XPF",
	"YER",
	"ZAR",
	"ZMW",
	"ZMK",
	"ZWD",
	"",
}

func (v *CurrencyCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CurrencyCode(value)
	for _, existing := range AllowedCurrencyCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CurrencyCode", value)
}

// NewCurrencyCodeFromValue returns a pointer to a valid CurrencyCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCurrencyCodeFromValue(v string) (*CurrencyCode, error) {
	ev := CurrencyCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CurrencyCode: valid values are %v", v, AllowedCurrencyCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CurrencyCode) IsValid() bool {
	for _, existing := range AllowedCurrencyCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CurrencyCode value
func (v CurrencyCode) Ptr() *CurrencyCode {
	return &v
}

type NullableCurrencyCode struct {
	value *CurrencyCode
	isSet bool
}

func (v NullableCurrencyCode) Get() *CurrencyCode {
	return v.value
}

func (v *NullableCurrencyCode) Set(val *CurrencyCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyCode(val *CurrencyCode) *NullableCurrencyCode {
	return &NullableCurrencyCode{value: val, isSet: true}
}

func (v NullableCurrencyCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

