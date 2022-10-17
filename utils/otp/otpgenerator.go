package otp

import (
	"fmt"
	"math/rand"
	"time"
)

//go:generate mockgen -source=otpgenerator.go -destination=../../mocks/otpgenerator_mock.go -package=mocks
type OtpGenerator interface {
	Generate() string
}

type otpGenerator struct{}

func New() OtpGenerator {
	rand.Seed(time.Now().UnixNano())
	return &otpGenerator{}
}

func (*otpGenerator) Generate() string {
	return fmt.Sprintf("%06d", rand.Intn(999999))
}
