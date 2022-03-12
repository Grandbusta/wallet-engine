package utils

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func GenerateWallet(email string) int {
	total := 0

	for i := 0; i < int(math.Min(float64(len(email)), 50)); i++ {
		r, _ := utf8.DecodeLastRuneInString(string(email[i]))
		total += int(r)
	}
	sec := time.Now().Unix()
	gen := strconv.Itoa(total + int(sec))
	rand.Seed(time.Now().UnixNano())
	x := strings.Replace(gen, string(gen[(rand.Intn(9-0)+0)]), string(strconv.Itoa(rand.Intn(9-0)+0)), 1)
	y := strings.Replace(x, string(gen[0]), string(strconv.Itoa(rand.Intn(9-0)+0)), 1)
	wallet, _ := strconv.Atoi(y)
	return wallet
}
