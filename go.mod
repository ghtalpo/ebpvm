module github.com/ghtalpo/ebpvm

go 1.13

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/hajimehoshi/bitmapfont/v2 v2.1.1
	github.com/hajimehoshi/ebiten/v2 v2.0.0
	github.com/juan-medina/goecs v1.5.2
	github.com/theckman/go-fsm v0.0.2
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5
	golang.org/x/text v0.3.4
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace github.com/ghtalpo/ebpvm/pvm v0.0.0 => ./pvm/ // for easy develop
