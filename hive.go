package hive

import (
	"github.com/yorkjs/hive-go/internal/constant"
	"github.com/yorkjs/hive-go/internal/is"
	"github.com/yorkjs/hive-go/internal/normalize"
)

var (
	AUTH_CODE_WECHAT = constant.AUTH_CODE_WECHAT
	AUTH_CODE_ALIPAY = constant.AUTH_CODE_ALIPAY

	DATE_YEAR_MONTH_DATE = constant.DATE_YEAR_MONTH_DATE
	DATE_YEAR_MONTH      = constant.DATE_YEAR_MONTH
	DATE_MONTH_DATE      = constant.DATE_MONTH_DATE

	DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND = constant.DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND
	DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE        = constant.DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE
	DATE_TIME_MONTH_DATE_HOUR_MINUTE             = constant.DATE_TIME_MONTH_DATE_HOUR_MINUTE

	MS_SECOND = constant.MS_SECOND
	MS_MINUTE = constant.MS_MINUTE
	MS_HOUR   = constant.MS_HOUR
	MS_DAY    = constant.MS_DAY
	MS_WEEK   = constant.MS_WEEK
	MS_YEAR   = constant.MS_YEAR

	MONEY_YUAN_TO_CENT              = constant.MONEY_YUAN_TO_CENT
	MONEY_TEN_THOUSAND_YUAN_TO_CENT = constant.MONEY_TEN_THOUSAND_YUAN_TO_CENT
	MONEY_YUAN_TO_PENNY             = constant.MONEY_YUAN_TO_PENNY

	SHELF_LIFE_DAY   = constant.SHELF_LIFE_DAY
	SHELF_LIFE_MONTH = constant.SHELF_LIFE_MONTH
	SHELF_LIFE_YEAR  = constant.SHELF_LIFE_YEAR

	SIZE_KB = constant.SIZE_KB
	SIZE_MB = constant.SIZE_MB
	SIZE_GB = constant.SIZE_GB

	NormalizeVersion  = normalize.Version
	IsStandardBarcode = is.StandardBarcode
	IsCustomBarcode   = is.CustomBarcode
	IsInteger         = is.Integer
)
