package passwd

import (
  "crypto/sha256"
  "encoding/base64"
)

func GetHash(password string) string {
  h := sha256.New();
  h.Write([]byte(password))
  return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
