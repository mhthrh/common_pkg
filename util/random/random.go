package random

import (
	"math/rand"
	"strings"
	"time"
)

type Random struct {
	seed int64
	min  int64
	max  int64
}

const alphabet = "abcdfghijkmnopuwqxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}
func New(min, max, seed int64) Random {
	return Random{
		seed: seed,
		min:  min,
		max:  max,
	}
}

func (r Random) GetString(length int) string {
	var sb strings.Builder
	for range length {
		c := alphabet[rand.Intn(len(alphabet))]
		sb.WriteByte(c)
	}
	return sb.String()
}
func (r Random) GetInt() int64 {
	return r.min + rand.Int63n(r.max-r.min+1)
}
func (r Random) RandomOfList(list []string) string {
	return list[rand.Intn(len(list))]
}
