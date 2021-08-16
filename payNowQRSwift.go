package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func crc16(data []byte) uint16 {

	var crcTable = []uint16{
		0x0000, 0x1021, 0x2042, 0x3063, 0x4084, 0x50a5,
		0x60c6, 0x70e7, 0x8108, 0x9129, 0xa14a, 0xb16b,
		0xc18c, 0xd1ad, 0xe1ce, 0xf1ef, 0x1231, 0x0210,
		0x3273, 0x2252, 0x52b5, 0x4294, 0x72f7, 0x62d6,
		0x9339, 0x8318, 0xb37b, 0xa35a, 0xd3bd, 0xc39c,
		0xf3ff, 0xe3de, 0x2462, 0x3443, 0x0420, 0x1401,
		0x64e6, 0x74c7, 0x44a4, 0x5485, 0xa56a, 0xb54b,
		0x8528, 0x9509, 0xe5ee, 0xf5cf, 0xc5ac, 0xd58d,
		0x3653, 0x2672, 0x1611, 0x0630, 0x76d7, 0x66f6,
		0x5695, 0x46b4, 0xb75b, 0xa77a, 0x9719, 0x8738,
		0xf7df, 0xe7fe, 0xd79d, 0xc7bc, 0x48c4, 0x58e5,
		0x6886, 0x78a7, 0x0840, 0x1861, 0x2802, 0x3823,
		0xc9cc, 0xd9ed, 0xe98e, 0xf9af, 0x8948, 0x9969,
		0xa90a, 0xb92b, 0x5af5, 0x4ad4, 0x7ab7, 0x6a96,
		0x1a71, 0x0a50, 0x3a33, 0x2a12, 0xdbfd, 0xcbdc,
		0xfbbf, 0xeb9e, 0x9b79, 0x8b58, 0xbb3b, 0xab1a,
		0x6ca6, 0x7c87, 0x4ce4, 0x5cc5, 0x2c22, 0x3c03,
		0x0c60, 0x1c41, 0xedae, 0xfd8f, 0xcdec, 0xddcd,
		0xad2a, 0xbd0b, 0x8d68, 0x9d49, 0x7e97, 0x6eb6,
		0x5ed5, 0x4ef4, 0x3e13, 0x2e32, 0x1e51, 0x0e70,
		0xff9f, 0xefbe, 0xdfdd, 0xcffc, 0xbf1b, 0xaf3a,
		0x9f59, 0x8f78, 0x9188, 0x81a9, 0xb1ca, 0xa1eb,
		0xd10c, 0xc12d, 0xf14e, 0xe16f, 0x1080, 0x00a1,
		0x30c2, 0x20e3, 0x5004, 0x4025, 0x7046, 0x6067,
		0x83b9, 0x9398, 0xa3fb, 0xb3da, 0xc33d, 0xd31c,
		0xe37f, 0xf35e, 0x02b1, 0x1290, 0x22f3, 0x32d2,
		0x4235, 0x5214, 0x6277, 0x7256, 0xb5ea, 0xa5cb,
		0x95a8, 0x8589, 0xf56e, 0xe54f, 0xd52c, 0xc50d,
		0x34e2, 0x24c3, 0x14a0, 0x0481, 0x7466, 0x6447,
		0x5424, 0x4405, 0xa7db, 0xb7fa, 0x8799, 0x97b8,
		0xe75f, 0xf77e, 0xc71d, 0xd73c, 0x26d3, 0x36f2,
		0x0691, 0x16b0, 0x6657, 0x7676, 0x4615, 0x5634,
		0xd94c, 0xc96d, 0xf90e, 0xe92f, 0x99c8, 0x89e9,
		0xb98a, 0xa9ab, 0x5844, 0x4865, 0x7806, 0x6827,
		0x18c0, 0x08e1, 0x3882, 0x28a3, 0xcb7d, 0xdb5c,
		0xeb3f, 0xfb1e, 0x8bf9, 0x9bd8, 0xabbb, 0xbb9a,
		0x4a75, 0x5a54, 0x6a37, 0x7a16, 0x0af1, 0x1ad0,
		0x2ab3, 0x3a92, 0xfd2e, 0xed0f, 0xdd6c, 0xcd4d,
		0xbdaa, 0xad8b, 0x9de8, 0x8dc9, 0x7c26, 0x6c07,
		0x5c64, 0x4c45, 0x3ca2, 0x2c83, 0x1ce0, 0x0cc1,
		0xef1f, 0xff3e, 0xcf5d, 0xdf7c, 0xaf9b, 0xbfba,
		0x8fd9, 0x9ff8, 0x6e17, 0x7e36, 0x4e55, 0x5e74,
		0x2e93, 0x3eb2, 0x0ed1, 0x1ef0,
	}

	var crc uint16
	crc = 0xffff
	for _, v := range data {
		j := (uint16(v) ^ (crc >> 8)) & 0xFF
		crc = crcTable[uint8(j)] ^ (crc << 8)
	}
	return (crc ^ 0) & 0xFFFF
}

func generatePayNowQRString(beneficiaryType, beneficiary, beneficiaryName, transactionAmount, referenceNumber, expiryDate string, isEditable bool) string {

	zeroConst := "0"
	//Payload Format Indicator String
	const payloadFormatIndicatorStringID = "00"
	const payloadFormatIndicatorStringValue = "01"
	t2 := strconv.Itoa(utf8.RuneCountInString(payloadFormatIndicatorStringValue))
	var payloadFormatIndicatorStringCharLength string = zeroConst + t2
	var payloadFormatIndicatorString string = payloadFormatIndicatorStringID + payloadFormatIndicatorStringCharLength + payloadFormatIndicatorStringValue
	fmt.Println("payloadFormatIndicatorString : ", payloadFormatIndicatorString)

	//Point Of Initiation Method
	const pointOfInitiationMethodStringID = "01"
	//For The Value, 11 = Static, 12 = Dynamic
	const pointOfInitiationMethodStringValue = "12"
	t4 := strconv.Itoa(utf8.RuneCountInString(pointOfInitiationMethodStringValue))
	var pointOfInitiationMethodStringCharLength string = zeroConst + t4
	var pointOfInitiationMethodString string = pointOfInitiationMethodStringID + pointOfInitiationMethodStringCharLength + pointOfInitiationMethodStringValue
	fmt.Println("pointOfInitiationMethodString : ", pointOfInitiationMethodString)
	//Merchant Account Info Template Sub-Category : Electronic Fund Transfer Service
	const electronicFundTransferServiceStringID = "00"
	const electronicFundTransferServiceStringValue = "SG.PAYNOW"
	t6 := strconv.Itoa(utf8.RuneCountInString(electronicFundTransferServiceStringValue))
	var electronicFundTransferServiceStringCharLength string = zeroConst + t6
	var electronicFundTransferServiceString string = electronicFundTransferServiceStringID + electronicFundTransferServiceStringCharLength + electronicFundTransferServiceStringValue
	fmt.Println("electronicFundTransferServiceString : ", electronicFundTransferServiceString)
	//Merchant Account Info Template Sub-Category : Beneficiary Type Selected
	const beneficiaryTypeSelectedStringID = "01"
	//The Value 0 = Mobile, 1 = Unused, 2 = UEN
	var beneficiaryTypeStringValue string = beneficiaryType
	t8 := strconv.Itoa(utf8.RuneCountInString(beneficiaryTypeStringValue))
	var beneficiaryTypeStringCharLength string = zeroConst + t8
	var beneficiaryTypeString string = beneficiaryTypeSelectedStringID + beneficiaryTypeStringCharLength + beneficiaryTypeStringValue
	fmt.Println("beneficiaryTypeString : ", beneficiaryTypeString)
	//Merchant Account Info Template Sub-Category : UEN Value (Company Unique Entity Number)
	const beneficiaryValueStringID = "02"
	var beneficiaryValueStringValue string = beneficiary
	var beneficiaryValueStringValueCount = utf8.RuneCountInString(beneficiaryValueStringValue)

	var beneficiaryValueStringCharLength string

	if beneficiaryValueStringValueCount >= 10 {
		beneficiaryValueStringCharLength = strconv.Itoa(beneficiaryValueStringValueCount)
	} else if beneficiaryValueStringValueCount < 10 {
		beneficiaryValueStringCharLength = zeroConst + strconv.Itoa(beneficiaryValueStringValueCount)
	} else {
		fmt.Println("Failed To Return Beneficiary Value Character Length")
	}
	var beneficiaryValueString string = beneficiaryValueStringID + beneficiaryValueStringCharLength + beneficiaryValueStringValue
	fmt.Println("beneficiaryValueString : ", beneficiaryValueString)
	//Merchant Account Info Template Sub-Category : Payment Is Or Is Not Editable
	const isPaymentEditableStringID = "03"
	//The Value 0 = Payment Not Editable, 1 = Payment Is Editable
	var isPaymentEditableStringValue string
	if isEditable == true {
		isPaymentEditableStringValue = "1"
	} else {
		isPaymentEditableStringValue = "0"
	}
	t11 := strconv.Itoa(utf8.RuneCountInString(isPaymentEditableStringValue))
	var isPaymentEditableStringCharLength string = zeroConst + t11
	var isPaymentEditableString string = isPaymentEditableStringID + isPaymentEditableStringCharLength + isPaymentEditableStringValue
	fmt.Println("isPaymentEditableString : ", isPaymentEditableString)
	//Merchant Account Info Template Sub-Category : Expiry Date (YYYYMMDD Format) (This Is An Optional Category)
	var expiryDateString string

	if expiryDate == "none" {
		expiryDateString = ""
	} else {
		const expiryDateStringID = "04"
		var expiryDateStringValue string = expiryDate
		t12 := "0"
		t13 := strconv.Itoa(utf8.RuneCountInString(expiryDateStringValue))
		var expiryDateStringCharLength string = t12 + t13
		expiryDateString = expiryDateStringID + expiryDateStringCharLength + expiryDateStringValue
	}
	fmt.Println("expiryDateString : ", expiryDateString)
	//Merchant Account Info Template (ID-26)
	const merchantAccountInfoTemplateStringID = "26"
	t14 := utf8.RuneCountInString(electronicFundTransferServiceString)
	t15 := utf8.RuneCountInString(beneficiaryTypeString)
	t16 := utf8.RuneCountInString(beneficiaryValueString)
	t17 := utf8.RuneCountInString(isPaymentEditableString)
	t18 := utf8.RuneCountInString(expiryDateString)
	var merchantAccountInfoTemplateStringCount int = t14 + t15 + t16 + t17 + t18
	var merchantAccountInfoTemplateStringCharLength string = strconv.Itoa(merchantAccountInfoTemplateStringCount)
	var merchantAccountInfoTemplateString string = merchantAccountInfoTemplateStringID + merchantAccountInfoTemplateStringCharLength + electronicFundTransferServiceString + beneficiaryTypeString + beneficiaryValueString + isPaymentEditableString + expiryDateString
	fmt.Println("merchantAccountInfoTemplateString : ", merchantAccountInfoTemplateString)
	//Merchant Category Code
	const merchantCategoryCodeStringID = "52"
	//The Value For The Merchant Category Code = 0000 If It Is Unused
	const merchantCategoryCodeStringValue = "0000"
	t3245 := strconv.Itoa(utf8.RuneCountInString(merchantCategoryCodeStringValue))
	var merchantCategoryCodeStringCharLength string = zeroConst + t3245
	var merchantCategoryCodeString string = merchantCategoryCodeStringID + merchantCategoryCodeStringCharLength + merchantCategoryCodeStringValue
	fmt.Println("merchantCategoryCodeString : ", merchantCategoryCodeString)
	//Currency Code
	const currencyCodeStringID = "53"
	//The Currency Code Of Singapore Is 702
	const currencyCodeStringValue = "702"
	t20 := strconv.Itoa(utf8.RuneCountInString(currencyCodeStringValue))
	var currencyCodeStringCharLength string = zeroConst + t20
	var currencyCodeString string = currencyCodeStringID + currencyCodeStringValue + currencyCodeStringCharLength
	fmt.Println("currencyCodeString ", currencyCodeString)
	//The Transaction Amount In Dollars
	const transactionAmountStringID = "54"
	var transactionAmountStringValue string = transactionAmount
	var transactionAmountStringValueCount int = utf8.RuneCountInString(transactionAmountStringValue)
	var transactionAmountStringCharLength string

	if transactionAmountStringValueCount >= 10 {
		transactionAmountStringCharLength = strconv.Itoa(transactionAmountStringValueCount)
	} else if transactionAmountStringValueCount < 10 {
		t22 := strconv.Itoa(transactionAmountStringValueCount)
		transactionAmountStringCharLength = zeroConst + t22
	} else {
		fmt.Println("Failed To Return Transaction Amount Character Length")
	}

	var transactionAmountString string = transactionAmountStringID + transactionAmountStringCharLength + transactionAmountStringValue
	fmt.Println("transactionAmountString : ", transactionAmountString)
	//Country Code (2 Letters)
	const countryCodeStringID = "58"
	const countryCodeStringValue = "SG"
	t24 := strconv.Itoa(utf8.RuneCountInString(countryCodeStringValue))
	var countryCodeStringCharLength string = zeroConst + t24
	var countryCodeString string = countryCodeStringID + countryCodeStringCharLength + countryCodeStringValue
	fmt.Println("countryCodeString ", countryCodeString)
	//Company Name
	const beneficiaryNameStringID = "59"
	var beneficiaryNameStringValue string = beneficiaryName
	var beneficiaryNameStringCharLength string
	var beneficiaryNameStringValueCount int = utf8.RuneCountInString(beneficiaryNameStringValue)

	if beneficiaryNameStringValueCount >= 10 {
		beneficiaryNameStringCharLength = strconv.Itoa(beneficiaryNameStringValueCount)
	} else if beneficiaryNameStringValueCount < 10 {
		t26 := strconv.Itoa(beneficiaryNameStringValueCount)
		beneficiaryNameStringCharLength = zeroConst + t26
	} else {
		fmt.Println("Failed To Return Beneficiary Name Character Length")
	}

	var beneficiaryNameString string = beneficiaryNameStringID + beneficiaryNameStringCharLength + beneficiaryNameStringValue
	fmt.Println("beneficiaryNameString : ", beneficiaryNameString)
	//Merchant City
	const merchantCityStringID = "60"
	const merchantCityStringValue = "Singapore"
	var merchantCityStringValueCount int = utf8.RuneCountInString(merchantCityStringValue)
	var merchantCityStringCharLength string

	if merchantCityStringValueCount >= 10 {
		merchantCityStringCharLength = strconv.Itoa(merchantCityStringValueCount)
	} else if merchantCityStringValueCount < 10 {
		t28 := strconv.Itoa(merchantCityStringValueCount)
		merchantCityStringCharLength = zeroConst + t28
	} else {
		fmt.Println("Failed To Return Merchant City Character Length")
	}
	var merchantCityString string = merchantCityStringID + merchantCityStringCharLength + merchantCityStringValue
	fmt.Println("merchantCityString : ", merchantCityString)
	//Additional Data Fields Sub-Category : Actual Bill / Reference Number
	const referenceNumberStringID = "01"
	var referenceNumberStringValue string = referenceNumber
	var referenceNumberStringValueCount int = utf8.RuneCountInString(referenceNumberStringValue)
	var referenceNumberStringCharLength string
	if referenceNumberStringValueCount >= 10 {
		referenceNumberStringCharLength = strconv.Itoa(referenceNumberStringValueCount)
	} else if referenceNumberStringValueCount < 10 {
		t30 := strconv.Itoa(referenceNumberStringValueCount)
		referenceNumberStringCharLength = zeroConst + t30
	} else {
		fmt.Println("Failed To Return Reference Number String Character Length")
	}
	var referenceNumberString string = referenceNumberStringID + referenceNumberStringCharLength + referenceNumberStringValue
	fmt.Println("referenceNumberString : ", referenceNumberString)
	//Additional Data Fields (ID 62)
	const additionalDataFieldsStringID = "62"
	var referenceNumberStringCount int = utf8.RuneCountInString(referenceNumberString)
	var additionalDataFieldsStringCharLength string

	if referenceNumberStringCount >= 10 {
		additionalDataFieldsStringCharLength = strconv.Itoa(referenceNumberStringCount)
	} else if referenceNumberStringCount < 10 {
		t30 := strconv.Itoa(referenceNumberStringCount)
		additionalDataFieldsStringCharLength = zeroConst + t30
	} else {
		fmt.Println("Failed To Return Additional Data Fields Character Length")
	}

	var additionalDataFieldsString string = additionalDataFieldsStringID + additionalDataFieldsStringCharLength + referenceNumberString
	fmt.Println("additionalDataFieldsString : ", additionalDataFieldsString)
	//Checksum
	const checksumStringID = "63"
	const checksumStringCharLength = "04"
	var checksumString string = checksumStringID + checksumStringCharLength
	fmt.Println("checksumString : ", checksumString)
	//PayNow QR Code String Without The CRC-16 Checksum
	var payNowQRCodeStringWithoutChecksumCRC16 string = payloadFormatIndicatorString + pointOfInitiationMethodString + merchantAccountInfoTemplateString + merchantCategoryCodeString + currencyCodeString + transactionAmountString + countryCodeString + beneficiaryNameString + merchantCityString + additionalDataFieldsString + checksumString
	temp := []byte("payNowQRCodeStringWithoutChecksumCRC16")
	var checksum_CRC16_StringHex string = strings.ToUpper(fmt.Sprintf("%x", crc16(temp)))
	fmt.Println("checksum_CRC16_String : ", checksum_CRC16_StringHex)
	var finalPayNowQRString = payNowQRCodeStringWithoutChecksumCRC16 + checksum_CRC16_StringHex
	return finalPayNowQRString

}

func main() {
	fmt.Println("Welcome to a PayNowQRString Generator!")

	var beneficiaryType string = "2"
	var beneficiaryValue string = "S62SS0057G"
	var beneficiaryName string = "Singapore Children's Society"
	var amount string = "10"
	var reference string = "Donation"
	var expiryDate string = "20041129"
	var amounstIsEditable bool = false
	var withoutChecksumString string = generatePayNowQRString(beneficiaryType, beneficiaryValue, beneficiaryName, amount, reference, expiryDate, amounstIsEditable)
	fmt.Println("withoutChecksumString : ", withoutChecksumString)
}
