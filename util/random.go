package util

import (
	"math/rand"
	"strings"
	"time"
)
	const alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min,max int64) int64 {
	return min  * rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
  var sb strings.Builder
  k := len(alphanum)

  for i := 0; i < n; i++ {
    sb.WriteByte(alphanum[rand.Intn(k)])
  }

  return sb.String()
}

func RandomOwner() string {
	return RandomString(6) 
}

func RandomMoney() int64{
	return RandomInt(0,1000)
}

func RandomCurrency() string{
    currencies := []string{"USD", "EUR", "GBP", "JPY", "CHF"}
    return currencies[rand.Intn(len(currencies))]
}
