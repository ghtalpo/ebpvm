module github.com/ghtalpo/egb

go 1.16

require (
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20210715014612-ab6297867137 // indirect
	github.com/hajimehoshi/bitmapfont/v2 v2.1.3
	github.com/hajimehoshi/ebiten/v2 v2.1.3
	github.com/juan-medina/goecs v1.5.2
	github.com/theckman/go-fsm v0.0.2
	github.com/tkanos/gonfig v0.0.0-20210106201359-53e13348de2f
	golang.org/x/exp v0.0.0-20210714144626-1041f73d31d8 // indirect
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d
	golang.org/x/mobile v0.0.0-20210710064935-76c259c465ba // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/text v0.3.6
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace github.com/ghtalpo/egb/game v0.0.0 => ./game/ // for easy develop
