package utils

import (
  "crypto/sha256"
  "fmt"
  "math/big"
  "net/url"
  "strings"
)

const (
  base62Charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  shortcodeLength = 6
)

func PrepareURL(urlLink string) (string, error) {
  parsedURL, err := url.ParseRequestURI(urlLink)
  if err != nil {
    return "", fmt.Errorf("invalid URL: %v", err)
  }

  if parsedURL.Scheme == "" {
    parsedURL, err = url.Parse("http://" + urlLink)
    if err != nil {
      return "", fmt.Errorf("invalid URL after adding scheme: %v", err)
    }
  }

  if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
    return "", fmt.Errorf("unsupported URL scheme: %s", parsedURL.Scheme)
  }
  if parsedURL.Host == "" {
    return "", fmt.Errorf("missing host in URL")
  }

  return parsedURL.String(), nil
}

func CreateShortcode(url []byte) string {
  var result strings.Builder
  hash := sha256.Sum256(url)
  num := new(big.Int).SetBytes(hash[:])
  base := big.NewInt(62)

  for result.Len() < shortcodeLength {
    rem := new(big.Int)
    num.DivMod(num, base, rem)
    result.WriteByte(base62Charset[rem.Int64()])
  }

  return result.String()
}
