package facebook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacebookDataWithoutPageAdress(t *testing.T) {
	facebookData := NewFacebookData()
	facebookData.FacebookPageAdress = ""
	facebookData.FacebookPageId = "1234"
	facebookData.FacebookToken = "Token"
	facebookData.FacebookUserToken = "Token"

	err := facebookData.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Invalid Facebook data", err.Error())
}

func TestFacebookDataWithoutFacebookPageId(t *testing.T) {
	facebookData := NewFacebookData()
	facebookData.FacebookPageAdress = "www.test.com"
	facebookData.FacebookPageId = ""
	facebookData.FacebookToken = "Token"
	facebookData.FacebookUserToken = "Token"

	err := facebookData.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Invalid Facebook data", err.Error())
}

func TestFacebookDataWithoutFacebookToken(t *testing.T) {
	facebookData := NewFacebookData()
	facebookData.FacebookPageAdress = "www.test.com"
	facebookData.FacebookPageId = "1234"
	facebookData.FacebookToken = ""
	facebookData.FacebookUserToken = "Token"

	err := facebookData.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Invalid Facebook data", err.Error())
}

func TestFacebookDataWithoutFacebookUserToken(t *testing.T) {
	facebookData := NewFacebookData()
	facebookData.FacebookPageAdress = "www.test.com"
	facebookData.FacebookPageId = "1234"
	facebookData.FacebookToken = "Token"
	facebookData.FacebookUserToken = ""

	err := facebookData.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Invalid Facebook data", err.Error())
}
