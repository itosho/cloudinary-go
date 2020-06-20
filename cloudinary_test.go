package imgix

import (
	"testing"
)

func testClient() Client {
	return NewClient("itosho", "key123", "secret123")
}

func TestSignUrlByRawTransformation(t *testing.T) {
	c := testClient()

	productId := "users/1.png"
	accessType := "upload"
	rawTransformation := "w_50,h_50"
	expected := "https://res.cloudinary.com/itosho/image/upload/s--yZ29lbgG--/w_50,h_50/users/1.png"
	actual := c.SignUrlByRawTransformation(productId, accessType, rawTransformation)

	if c.SignUrlByRawTransformation(productId, accessType, rawTransformation) != expected {
		t.Errorf("Actual=%q, Expected=%q", actual, expected)
	}
}
