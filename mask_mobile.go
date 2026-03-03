package hive

// 脱敏手机号
func MaskMobile(mobile string) string {
	if len(mobile) == 11 {
		return mobile[0:3] + "****" + mobile[7:]
	}
	return mobile
}
