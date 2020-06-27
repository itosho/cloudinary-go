package cloudinary

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

const (
	baseUrl             = "https://res.cloudinary.com"
	defaultResourceType = "image"
)

type Client struct {
	cloudName string
	secure    bool
	apiKey    string
	apiSecret string
}

func NewClient(cloudName string, apiKey string, apiSecret string) Client {
	return Client{
		cloudName: cloudName,
		secure:    true,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (c *Client) SignUrlByRawTransformation(productId string, accessType string, rawTransformation string) string {
	seed := rawTransformation + "/" + productId + c.apiSecret

	hash := sha1.New()
	hash.Write([]byte(seed))
	b := hash.Sum(nil)

	signature := base64.StdEncoding.EncodeToString(b)
	signature = strings.Replace(signature, "+", "-", -1)
	signature = strings.Replace(signature, "/", "_", -1)
	signature = string([]rune(signature)[:8])
	signature = "s--" + signature + "--"

	transformation := strings.Replace(rawTransformation, "\n", "%0A", -1)

	signUrl := baseUrl + "/" + c.cloudName + "/" + defaultResourceType + "/" + accessType
	signUrl = signUrl + "/" + signature + "/" + transformation + "/" + productId

	return signUrl
}
