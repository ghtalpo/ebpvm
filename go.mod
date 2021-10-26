module github.com/ghtalpo/egb

go 1.16

require (
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20211024062804-40e447a793be // indirect
	github.com/hajimehoshi/bitmapfont/v2 v2.1.3
	github.com/hajimehoshi/ebiten/v2 v2.2.1
	github.com/juan-medina/goecs v1.5.2
	github.com/theckman/go-fsm v0.0.2
	github.com/tkanos/gonfig v0.0.0-20210106201359-53e13348de2f
	golang.org/x/exp v0.0.0-20211025140241-8418b01e8c3b // indirect
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d
	golang.org/x/mobile v0.0.0-20210924032853-1c027f395ef7 // indirect
	golang.org/x/sys v0.0.0-20211025201205-69cdffdb9359 // indirect
	golang.org/x/text v0.3.7
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace github.com/ghtalpo/egb/game v0.0.0 => ./game/ // for easy develop
