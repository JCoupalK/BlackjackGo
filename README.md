
<p align="center">
 <img src="https://github.com/JCoupalK/BlackjackGo/assets/108779415/789a025c-89ec-41f6-a7fb-23dc21eb544a"
</p>

# BlackjackGo

Welcome to BlackjackGo! This project is a simple yet exciting implementation of the classic Blackjack card game, written in Go. It features a text-based user interface and an intuitive gameplay experience.

## Features

- Text-based Blackjack game playable in the terminal.
- Simple and intuitive user interface with color styling.
- Automatic handling of game rules including scoring, hitting, standing, and dealing.
- Soft hand score detection and display.
- Dealer play logic according to standard Blackjack rules.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Gameplay](#gameplay)
- [Build](#build)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Download the binary with wget:

    ```shell
    wget https://github.com/JCoupalK/blackjackgo/releases/download/1.0/blackjackgo_linux_amd64_1.0.tar.gz
    ```

2. Unpack it with tar

    ```shell
    tar -xf blackjackgo_linux_amd64_1.0.tar.gz
    ```

3. To start the game, run:

    ```shell
    ./blackjackgo
    ```

You'll be presented with the game interface in your terminal. Follow the on-screen instructions to play.

## Gameplay

- Press 'H' to hit (draw a new card).
- Press 'S' to stand (end your turn).
- Press 'R' to restart the game after it ends.
- Press 'Q' at any time to quit.

The game follows standard Blackjack rules. Your goal is to beat the dealer's hand without going over 21.

![image](https://github.com/JCoupalK/BlackjackGo/assets/108779415/134b90d7-4e3d-41ad-9834-f73d5558a33e)

## Build

To build  BlackjackGo, you need to have Go installed on your machine. If you don't have Go installed, you can download it from [the official Go website](https://golang.org/dl/).

Once Go is installed, clone this repository:

```shell
git clone https://github.com/JCoupalK/blackjackgo
cd blackjackgo
go build .
```

## Contributing

Contributions to BlackjackGo are welcome! Feel free to fork the repository and submit pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
