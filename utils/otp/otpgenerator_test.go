package otp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOtpGenerator_Generate(t *testing.T) {
	otpGen := New()
	otp := otpGen.Generate()
	assert.NotEmpty(t, otp)
	assert.Equal(t, 6, len(otp))
}

