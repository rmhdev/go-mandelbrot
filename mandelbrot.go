package mandelbrot

//             |
//             |
//             |
// ------------------------ real
//             |
//             |
//             | imaginary (i)

const realMin float64 = -2.0
const realMax float64 = 2.0
const imaginaryMin float64 = -2
const imaginaryMax float64 = 1

type Config struct {
	width         int
	height        int
	maxIterations int
	realMin       float64
	realMax       float64
	imaginaryMin  float64
	imaginaryMax  float64
}

func NewConfig(width int, height int, maxIterations int) Config {
	config := Config{}
	config.width = width
	config.height = height
	config.maxIterations = maxIterations
	config.realMin = realMin
	config.realMax = realMax
	config.imaginaryMin = imaginaryMin
	config.imaginaryMax = config.imaginaryMin + (config.realMax-config.realMin)*float64(config.height)/float64(config.width)

	return config
}

type Coord struct {
	x int
	y int
}

func (coord Coord) complex(config Config) complex128 {
	return complex(
		config.realMin+float64(coord.x)*((config.realMax-config.realMin)/float64(config.width-1)),
		config.imaginaryMax-float64(coord.y)*((config.imaginaryMax-config.imaginaryMin)/float64(config.height-1)))
}
