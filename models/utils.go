package models

import (
	"math"
	"math/rand"
	"go.uber.org/zap"
	"runtime/debug"
	. "LianFaPhone/lfp-base/log/zap"
)

const(
	CONST_NOTIFY_MSG_TYPE_ROB = 1

)

func AdjustFloatAcc(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

type SqlPairCondition struct{
	Key  interface{}
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