package rtx

type Texture interface {
	D(uv Point2) Spectrum
}
