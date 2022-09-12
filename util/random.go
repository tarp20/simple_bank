package util

import (
	"math/rand"
	"time"
	"strings"
)

const alfabet = "abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min +1)
}

func RandomString(n int) string{
    var sb strings.Builder
	k := len(alfabet)
	for i:= 0; i<n; i++{
		c:= alfabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
} 

//return random owner name
func RandomOwnerName() string{
	 return RandomString(6)
}


//return random amount
func RandomMoney() int64{
	return RandomInt(0,10000)
}

//return random currency

func RandomCurrency() string{
	
	currencies := []string{"USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]

}