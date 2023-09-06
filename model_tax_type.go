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

// TaxType See Tax Types – can only be used on update calls
type TaxType string

// List of TaxType
const (
	OUTPUT TaxType = "OUTPUT"
	INPUT TaxType = "INPUT"
	CAPEXINPUT TaxType = "CAPEXINPUT"
	EXEMPTEXPORT TaxType = "EXEMPTEXPORT"
	EXEMPTEXPENSES TaxType = "EXEMPTEXPENSES"
	EXEMPTCAPITAL TaxType = "EXEMPTCAPITAL"
	EXEMPTOUTPUT TaxType = "EXEMPTOUTPUT"
	INPUTTAXED TaxType = "INPUTTAXED"
	BASEXCLUDED TaxType = "BASEXCLUDED"
	GSTONCAPIMPORTS TaxType = "GSTONCAPIMPORTS"
	GSTONIMPORTS TaxType = "GSTONIMPORTS"
	NONE TaxType = "NONE"
	INPUT2 TaxType = "INPUT2"
	ZERORATED TaxType = "ZERORATED"
	OUTPUT2 TaxType = "OUTPUT2"
	CAPEXINPUT2 TaxType = "CAPEXINPUT2"
	CAPEXOUTPUT TaxType = "CAPEXOUTPUT"
	CAPEXOUTPUT2 TaxType = "CAPEXOUTPUT2"
	CAPEXSRINPUT TaxType = "CAPEXSRINPUT"
	CAPEXSROUTPUT TaxType = "CAPEXSROUTPUT"
	ECACQUISITIONS TaxType = "ECACQUISITIONS"
	ECZRINPUT TaxType = "ECZRINPUT"
	ECZROUTPUT TaxType = "ECZROUTPUT"
	ECZROUTPUTSERVICES TaxType = "ECZROUTPUTSERVICES"
	EXEMPTINPUT TaxType = "EXEMPTINPUT"
	REVERSECHARGES TaxType = "REVERSECHARGES"
	RRINPUT TaxType = "RRINPUT"
	RROUTPUT TaxType = "RROUTPUT"
	SRINPUT TaxType = "SRINPUT"
	SROUTPUT TaxType = "SROUTPUT"
	ZERORATEDINPUT TaxType = "ZERORATEDINPUT"
	ZERORATEDOUTPUT TaxType = "ZERORATEDOUTPUT"
	BLINPUT TaxType = "BLINPUT"
	DSOUTPUT TaxType = "DSOUTPUT"
	EPINPUT TaxType = "EPINPUT"
	ES33_OUTPUT TaxType = "ES33OUTPUT"
	ESN33_OUTPUT TaxType = "ESN33OUTPUT"
	IGDSINPUT2 TaxType = "IGDSINPUT2"
	IMINPUT2 TaxType = "IMINPUT2"
	MEINPUT TaxType = "MEINPUT"
	NRINPUT TaxType = "NRINPUT"
	OPINPUT TaxType = "OPINPUT"
	OSOUTPUT TaxType = "OSOUTPUT"
	TXESSINPUT TaxType = "TXESSINPUT"
	TXN33_INPUT TaxType = "TXN33INPUT"
	TXPETINPUT TaxType = "TXPETINPUT"
	TXREINPUT TaxType = "TXREINPUT"
	INPUT3 TaxType = "INPUT3"
	INPUT4 TaxType = "INPUT4"
	OUTPUT3 TaxType = "OUTPUT3"
	OUTPUT4 TaxType = "OUTPUT4"
	SROUTPUT2 TaxType = "SROUTPUT2"
	TXCA TaxType = "TXCA"
	SRCAS TaxType = "SRCAS"
	BLINPUT2 TaxType = "BLINPUT2"
	DRCHARGESUPPLY20 TaxType = "DRCHARGESUPPLY20"
	DRCHARGE20 TaxType = "DRCHARGE20"
	DRCHARGESUPPLY5 TaxType = "DRCHARGESUPPLY5"
	DRCHARGE5 TaxType = "DRCHARGE5"
	BADDEBTRELIEF TaxType = "BADDEBTRELIEF"
	IGDSINPUT3 TaxType = "IGDSINPUT3"
	SROVR TaxType = "SROVR"
	TOURISTREFUND TaxType = "TOURISTREFUND"
	TXRCN33 TaxType = "TXRCN33"
	TXRCRE TaxType = "TXRCRE"
	TXRCESS TaxType = "TXRCESS"
	TXRCTS TaxType = "TXRCTS"
	OUTPUTY23 TaxType = "OUTPUTY23"
	DSOUTPUTY23 TaxType = "DSOUTPUTY23"
	INPUTY23 TaxType = "INPUTY23"
	IMINPUT2_Y23 TaxType = "IMINPUT2Y23"
	IGDSINPUT2_Y23 TaxType = "IGDSINPUT2Y23"
	TXPETINPUTY23 TaxType = "TXPETINPUTY23"
	TXESSINPUTY23 TaxType = "TXESSINPUTY23"
	TXN33_INPUTY23 TaxType = "TXN33INPUTY23"
	TXREINPUTY23 TaxType = "TXREINPUTY23"
	TXCAY23 TaxType = "TXCAY23"
	BADDEBTRELIEFY23 TaxType = "BADDEBTRELIEFY23"
	IGDSINPUT3_Y23 TaxType = "IGDSINPUT3Y23"
	SROVRRSY23 TaxType = "SROVRRSY23"
	SROVRLVGY23 TaxType = "SROVRLVGY23"
	SRLVGY23 TaxType = "SRLVGY23"
	TXRCN33_Y23 TaxType = "TXRCN33Y23"
	TXRCREY23 TaxType = "TXRCREY23"
	TXRCESSY23 TaxType = "TXRCESSY23"
	TXRCTSY23 TaxType = "TXRCTSY23"
	IM TaxType = "IM"
	IMY23 TaxType = "IMY23"
	IMESS TaxType = "IMESS"
	IMESSY23 TaxType = "IMESSY23"
	IMN33 TaxType = "IMN33"
	IMN33_Y23 TaxType = "IMN33Y23"
	IMRE TaxType = "IMRE"
	IMREY23 TaxType = "IMREY23"
	BADDEBTRECOVERY TaxType = "BADDEBTRECOVERY"
	BADDEBTRECOVERYY23 TaxType = "BADDEBTRECOVERYY23"
)

// All allowed values of TaxType enum
var AllowedTaxTypeEnumValues = []TaxType{
	"OUTPUT",
	"INPUT",
	"CAPEXINPUT",
	"EXEMPTEXPORT",
	"EXEMPTEXPENSES",
	"EXEMPTCAPITAL",
	"EXEMPTOUTPUT",
	"INPUTTAXED",
	"BASEXCLUDED",
	"GSTONCAPIMPORTS",
	"GSTONIMPORTS",
	"NONE",
	"INPUT2",
	"ZERORATED",
	"OUTPUT2",
	"CAPEXINPUT2",
	"CAPEXOUTPUT",
	"CAPEXOUTPUT2",
	"CAPEXSRINPUT",
	"CAPEXSROUTPUT",
	"ECACQUISITIONS",
	"ECZRINPUT",
	"ECZROUTPUT",
	"ECZROUTPUTSERVICES",
	"EXEMPTINPUT",
	"REVERSECHARGES",
	"RRINPUT",
	"RROUTPUT",
	"SRINPUT",
	"SROUTPUT",
	"ZERORATEDINPUT",
	"ZERORATEDOUTPUT",
	"BLINPUT",
	"DSOUTPUT",
	"EPINPUT",
	"ES33OUTPUT",
	"ESN33OUTPUT",
	"IGDSINPUT2",
	"IMINPUT2",
	"MEINPUT",
	"NRINPUT",
	"OPINPUT",
	"OSOUTPUT",
	"TXESSINPUT",
	"TXN33INPUT",
	"TXPETINPUT",
	"TXREINPUT",
	"INPUT3",
	"INPUT4",
	"OUTPUT3",
	"OUTPUT4",
	"SROUTPUT2",
	"TXCA",
	"SRCAS",
	"BLINPUT2",
	"DRCHARGESUPPLY20",
	"DRCHARGE20",
	"DRCHARGESUPPLY5",
	"DRCHARGE5",
	"BADDEBTRELIEF",
	"IGDSINPUT3",
	"SROVR",
	"TOURISTREFUND",
	"TXRCN33",
	"TXRCRE",
	"TXRCESS",
	"TXRCTS",
	"OUTPUTY23",
	"DSOUTPUTY23",
	"INPUTY23",
	"IMINPUT2Y23",
	"IGDSINPUT2Y23",
	"TXPETINPUTY23",
	"TXESSINPUTY23",
	"TXN33INPUTY23",
	"TXREINPUTY23",
	"TXCAY23",
	"BADDEBTRELIEFY23",
	"IGDSINPUT3Y23",
	"SROVRRSY23",
	"SROVRLVGY23",
	"SRLVGY23",
	"TXRCN33Y23",
	"TXRCREY23",
	"TXRCESSY23",
	"TXRCTSY23",
	"IM",
	"IMY23",
	"IMESS",
	"IMESSY23",
	"IMN33",
	"IMN33Y23",
	"IMRE",
	"IMREY23",
	"BADDEBTRECOVERY",
	"BADDEBTRECOVERYY23",
}

func (v *TaxType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaxType(value)
	for _, existing := range AllowedTaxTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaxType", value)
}

// NewTaxTypeFromValue returns a pointer to a valid TaxType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaxTypeFromValue(v string) (*TaxType, error) {
	ev := TaxType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaxType: valid values are %v", v, AllowedTaxTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaxType) IsValid() bool {
	for _, existing := range AllowedTaxTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaxType value
func (v TaxType) Ptr() *TaxType {
	return &v
}

type NullableTaxType struct {
	value *TaxType
	isSet bool
}

func (v NullableTaxType) Get() *TaxType {
	return v.value
}

func (v *NullableTaxType) Set(val *TaxType) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxType) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxType(val *TaxType) *NullableTaxType {
	return &NullableTaxType{value: val, isSet: true}
}

func (v NullableTaxType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

