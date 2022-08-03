# it costs money to be alive
A short, non-fighty indie side scroller game. Run with the right / left arrow keys, jump with up arrow or space. Collect coins and bring them to the end of the level, but you lose coins over time. 

Currently in very preliminary state, there's one level, three or four minutes of gameplay. I've put a wasm build [right here](https://jcgraybill.github.io/it-costs-money/) that you can play in your web browser - try it out!

![screenshot](https://github.com/jcgraybill/it-costs-money/blob/main/screenshot.png)

Built in [golang](https://go.dev/) using the [ebiten](https://ebiten.org/) 2D game library. Tileset thanks to [ludicarts](https://ludicarts.itch.io/) *[license](https://www.ludicarts.com/license-2/)*, sounds [GameDev Market](https://www.gamedevmarket.net/) *[license](https://static.gamedevmarket.net/terms-conditions/#pro-licence)*, little running guy yoinked from the [ebiten animation demo](https://ebiten.org/examples/animation.html) (thanks!), typeface [Modak](https://github.com/EkType/Modak) *[license](https://github.com/EkType/Modak/blob/master/OFL.txt)*.

# Running the game locally
With go installed, download and run the game with:
```
go run -tags=deploy github.com/jcgraybill/it-costs-money@latest
```

To build on Ubuntu, install additional packages:
```
apt install libgl1-mesa-dev xorg-dev libasound2-dev
```
