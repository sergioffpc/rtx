package rtx_test

import (
	"sergioffpc/rtx/pkg/rtx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpectrumAdd(t *testing.T) {
	a, b := rtx.Spectrum{0.9, 0.6, 0.75}, rtx.Spectrum{0.7, 0.1, 0.25}
	assert.True(t, rtx.Spectrum.Eq(rtx.Spectrum{1.6, 0.7, 1.0}, rtx.Spectrum.Add(a, b)))
}

func TestSpectrumSub(t *testing.T) {
	a, b := rtx.Spectrum{0.9, 0.6, 0.75}, rtx.Spectrum{0.7, 0.1, 0.25}
	assert.True(t, rtx.Spectrum.Eq(rtx.Spectrum{0.2, 0.5, 0.5}, rtx.Spectrum.Sub(a, b)))
}

func TestSpectrumMul(t *testing.T) {
	a, b := rtx.Spectrum{1, 0.2, 0.4}, rtx.Spectrum{0.9, 1, 0.1}
	assert.True(t, rtx.Spectrum.Eq(rtx.Spectrum{0.9, 0.2, 0.04}, rtx.Spectrum.Mul(a, b)))
}

func TestSpectrumMulFloat(t *testing.T) {
	assert.True(t, rtx.Spectrum.Eq(rtx.Spectrum{0.4, 0.6, 0.8}, rtx.Spectrum{0.2, 0.3, 0.4}.MulFloat(2)))
}
