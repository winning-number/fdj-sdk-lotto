# FDJ-SDK-LOTTO

<!-- Badge section [godoc / code quality / ci / last Tag / License ]-->

<!-- Description section -->
This package is a sdk of the lotto (FRANCAISE DES JEUX) and could get the full draws history. It get the draws history provided from [any csv files](https://www.fdj.fr/jeux-de-tirage/loto/historique) and convert them in a standardized Draw format.

## Installation

Download fdj-sdk-lotto:

```sh
$ go get -u github.com/winning-number/fdj-sdk-lotto@latest
# go: downloading github.com/winning-number/fdj-sdk-lotto ...
# go: added github.com/winning-number/fdj-sdk-lotto ...
```

## Usage

They are two way for use this sdk. You can get directly the history from the `FRANCAISE DES JEUX` or directly by providing a CSV file (in this case, you need to know [all type of Draw](#type-of-draw))

* From the `FRANCAISE DES JEUX` api

```golang
import github.com/winning-number/fdj-sdk-lotto

func main() {
    var err error
    var driver lotto.Lotto

    if driver, err = lotto.New(); err != nil {
        panic(err)
    }
    // get all by default
    if err = driver.LoadAPI(lotto.LoadAPIOption{}); err != nil {
        panic(err)
    }
    draws := driver.Draws(lotto.Filter{
        SuperLotto:   true
        GrandLotto:   true
        XmasLotto:    true
        ClassicLotto: true
        OldLotto:     true
    })
}
```

* From your csv files

```golang
import github.com/winning-number/fdj-sdk-lotto

func main() {
    var err error
    var driver lotto.Lotto
    var file *os.File

    if file, err = os.Open("your_filepath.csv"); err != nil {
        panic(err)
    }
    defer file.Close()
    if driver, err = lotto.New(); err != nil {
        panic(err)
    }
    // get all by default
    if err = driver.LoadCSV(file, lotto.DrawLottoType, lotto.DrawV1); err != nil {
        panic(err)
    }
    draws := driver.Draws(lotto.Filter{
        SuperLotto:   true
        GrandLotto:   true
        XmasLotto:    true
        ClassicLotto: true
        OldLotto:     true
    })
}
```

## Type of DRAW

From the begin of the lotto, the rules are been updated any time. So, to exploit the full history, each version are interpreting like a global `Draw` type.

### DrawV0

* Concern only the `super lotto` type `before 2008 october`. They was `6` balls between `1` and `49` included and one lucky ball between `1` and `49` included. All balls are added inside the global slice of ball inside the Draw type. Result is provinding in a pick order and number order.
* Only one a `joker+` picked by draw.
* `7 Winners Rank` by Draw.

### DrawV1

* Concern only the `classic lotto` type `before 2008 october`. They was `6` balls between `1` and `49` included and one lucky ball between `1` and `49` included. All balls are added inside the global slice of ball inside the Draw type. Result is provinding in a pick order and number order.
* They was one a `joker+` and one `joker number` picked by draw.
* `7` Winners Rank by Draw.

### DrawV2

* Concern the `classic lotto` and the `super lotto` type `between 2008 october and 2017 march`. They was `5` balls between `1` and `49` included and one lucky ball between `1` and `9` included. these 5 balls are added inside the global slice of ball inside the Draw type and the lucky ball is added inside a LuckyBall field (int32).
* Only one a `joker+` picked by draw.
* `6` Winners Rank by Draw.

### DrawV3

* Concern the `classic lotto`, the `super lotto`, the `grand lotto` and the `grand lotto (xmas)` type `between 2017 march and 2019 november`. They was `5` balls between `1` and `49` included and one lucky ball between `1` and `9` included. these 5 balls are added inside the global slice of ball inside the Draw type and the lucky ball is added inside a LuckyBall field (int32). Result is provinding in a pick order and number order.
* Only one a `joker+` picked by draw.
* `9` Winners Rank by Draw.
* Add `Winning number` (any by draw depends of type of draw).

### DrawV4

* Concern the `classic lotto` type `from 2019 november`. They was `5` balls between `1` and `49` included and one lucky ball between `1` and `9` included. these 5 balls are added inside the global slice of ball inside the Draw type and the lucky ball is added inside a LuckyBall field (int32). Result is provinding in a pick order and number order.
* Only one a `joker+` picked by draw.
* `9` Winners Rank by Draw.
* Add `Winning number` (any by draw depends of type of draw).
* Add a `Second Roll` with `5` balls between `1` and `49` included (no lucky ball). Result is providing in only in a number order.
* `4` Winners Rank for the second Roll

## Project Structure

The package structure is classic with a flat mode for the lotto driver. They are just a quick separation for:

* `csvparser` which is a particular parser of csvfile with a strict decoder (column and header) to avoid to parse a old type like a new for example and lost data(s).
* `httpclient` which is a basic http client to get the csv file from the `FDJ` api. Probably this package will be move in another repository in the future.
* `reader` which read zip archive and csv file with a grace reader closer. It could record the downloaded files inside a folder if you want.

## Contributing

Pull requests are welcome. Before changes, please `open an issue first` to discuss what you would like to change.

Please make sure to update tests as appropriate and that the tests and the linters pass:

```sh
$ make test
# ...
# PASS
# ...
# tests executed !
$ make lint
# ... 
# done
```

## License

[GNU GPL v3](https://choosealicense.com/licenses/gpl-3.0/)
