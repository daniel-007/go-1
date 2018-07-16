package pubnub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeRequestHTTP2(t *testing.T) {
	assert := assert.New(t)

	config := NewConfig()
	config.Origin = "ssp.pubnub.com"
	config.UseHTTP2 = true
	pn := NewPubNub(config)

	_, s, err := pn.Time().Execute()

	assert.Nil(err)
	assert.Equal(200, s.StatusCode)
}
