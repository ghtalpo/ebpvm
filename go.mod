module github.com/ghtalpo/egb

go 1.16

require (
	github.com/hajimehoshi/bitmapfont/v2 v2.2.3
	github.com/hajimehoshi/ebiten/v2 v2.5.8
	github.com/hajimehoshi/file2byteslice v0.0.0-20210813153925-5340248a8f41 // indirect
	github.com/hajimehoshi/oto v0.6.1 // indirect
	github.com/juan-medina/goecs v1.5.2
	github.com/theckman/go-fsm v0.0.2
	github.com/tkanos/gonfig v0.0.0-20210106201359-53e13348de2f
	golang.org/x/exp v0.0.0-20211025140241-8418b01e8c3b // indirect
	golang.org/x/exp/shiny v0.0.0-20230817173708-d852ddb80c63 // indirect
	golang.org/x/image v0.11.0
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mobile v0.0.0-20230818142238-7088062f872d // indirect
	golang.org/x/text v0.12.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace github.com/ghtalpo/egb/game v0.0.0 => ./game/ // for easy develop
