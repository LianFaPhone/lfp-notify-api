package models

import (
	. "LianFaPhone/lfp-base/log/zap"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"runtime/debug"
)

const (
	CONST_SMS_TP_Yzm = 0
	CONST_SMS_TP_Tz = 1
	CONST_SMS_TP_Hyyx = 2
)

func AdjustFloatAcc(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

type SqlPairCondition struct {
	Key   interface{}
	Value interface{}
}

func GetRandomString(n int) string {
	const letterBytes = "abcdefghijk012lmnopqrstuvwxy345zABCDEFGHIJKLMNOPQRSTUVWXY678Z90"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int()%(len(letterBytes))]
	}

	return string(b)
}

func PanicPrint() {
	if err := recover(); err != nil {
		ZapLog().With(zap.Any("error", err)).Error(string(debug.Stack()))
	}
}
