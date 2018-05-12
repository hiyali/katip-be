package utils

import (
  "math/rand"
)

var Sources = map[uint]string{
  1: "abcdefghijklmnopqrstuvwxyz", // LowerLetters
  2: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", // UpperLetters
  3: "0123456789", // Digits
  4: "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./", // Symbols
}

const (
  LowerLetters = 1
  UpperLetters
  Digits
  Symbols
)

func getSource(sourceTypes []uint) (result string) {
  for i := 0; i < len(sourceTypes); i++ {
    result += Sources[sourceTypes[i]]
  }
  return
}

func GeneratePassword(len int) (result string) {
  sourceStr := getSource([]uint{1,2,3,4})
  for i := 0; i < len; i++ {
    result += randomItem(sourceStr)
  }
  return
}

func randomItem(str string) string {
  n := rand.Intn(len(str))
	return string(str[n])
}
