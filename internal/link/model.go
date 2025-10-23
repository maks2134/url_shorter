package link

import (
	"crypto/rand"
	"math/big"
	"shorter-url/internal/stat"

	"gorm.io/gorm"
)

const (
	hashLength  = 10
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	err := link.generateHash()
	if err != nil {
		panic("failed to generate hash: " + err.Error())
	}
	return link
}

func (link *Link) generateHash() error {
	hash, err := generateRandomHash(hashLength)
	if err != nil {
		return err
	}
	link.Hash = hash
	return nil
}

func generateRandomHash(n int) (string, error) {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		b[i] = letterBytes[num.Int64()]
	}
	return string(b), nil
}
