package anagram

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sort"
)

var encoderCfg = zap.NewProductionEncoderConfig()
var atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
var logger = zap.New(zapcore.NewCore(
	zapcore.NewJSONEncoder(encoderCfg),
	zapcore.Lock(os.Stdout),
	atomicLevel,
))
var sugar = logger.Sugar()

func Solve(strings []string) [][]string {
	mp := make(map[string][]string)
	for _, s := range strings {
		sugar.Debugf("sorting function called for %v", s)
		sortedString := sortString(s)
		sugar.Debugf("value after sort, %v", sortedString)
		mp[sortedString] = append(mp[sortedString], s)
	}
	ans := make([][]string, 0)
	for _, v := range mp {
		sugar.Debugf("Inserting slice of string into ans: %v", v)
		ans = append(ans, v)
	}
	return ans
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func Changeloglevel(newlevel string){
	switch newlevel {
	case "Debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "Warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	default:
		atomicLevel.SetLevel(zap.InfoLevel)
	}
}