# it costs money to be alive
A short, non-fighty indie side scroller game. Run with the right / left arrow keys, jump with up arrow or space. Collect coins and bring them to the end of the level, but you lose coins over time. 

Currently in very preliminary state. Builds and runs on Windows, *theoretically* on [all the platforms](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#readme-platforms) supported by the ebiten library.

![screenshot](https://github.com/jcgraybill/it-costs-money/blob/main/screenshot.png)

Built in [golang](https://go.dev/) using the [ebiten](https://ebiten.org/) 2D game library. Tileset thanks to [ludicarts](https://ludicarts.itch.io/) *[license](https://www.ludicarts.com/license-2/)*, sounds [GameDev Market](https://www.gamedevmarket.net/) *[license](https://static.gamedevmarket.net/terms-conditions/#pro-licence)*, little running guy yoinked from the [ebiten animation demo](https://ebiten.org/examples/animation.html) (thanks!).

# Running the game
```
git clone git@github.com:jcgraybill/it-costs-money.git
go mod tidy
go build
```