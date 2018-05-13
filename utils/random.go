package utils

import "math/rand"

var Sources = map[string]string{
  "LowerLetters": "abcdefghijklmnopqrstuvwxyz", // LowerLetters
  "UpperLetters": "ABCDEFGHIJKLMNOPQRSTUVWXYZ", // UpperLetters
  "Digits": "0123456789", // Digits
  "Symbols": "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./", // Symbols
}

type SourceTypes struct {
  All           bool
  LowerLetters  bool
  UpperLetters  bool
  Digits        bool
  Symbols       bool
}

func getSource(st SourceTypes) (result string) {
  if st.All {
    for _, value := range Sources {
      result += value
    }
  } else {
    if st.UpperLetters {
      result += Sources["UpperLetters"]
    }
    if st.LowerLetters {
      result += Sources["LowerLetters"]
    }
    if st.Digits {
      result += Sources["Digits"]
    }
    if st.Symbols {
      result += Sources["Symbols"]
    }
  }
  return
}

func GenerateRandomStr(len int, st SourceTypes) (result string) {
  sourceStr := getSource(st)
  for i := 0; i < len; i++ {
    result += randomItem(sourceStr)
  }
  return
}

func randomItem(str string) string {
  n := rand.Intn(len(str))
	return string(str[n])
}
