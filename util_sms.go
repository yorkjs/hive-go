package hive

// 计算短信的预估字符数和短信条数
func CalculateSmsCount(signatureName, templateContent string, perSmsCharCount int) (int, int) {

	charCount := GetStringLength(templateContent)

	if signatureName != "" {
		// 【签名】
		charCount += GetStringLength(signatureName) + 2
	}

	if charCount == 0 || perSmsCharCount <= 0 {
		return charCount, 0
	}

	smsCount := (charCount + perSmsCharCount - 1) / perSmsCharCount
	return charCount, smsCount

}
