# Projector

*Projector* is a program that allows the user to remotely control the contents of an info screen. It is possible to play videos, display images or marquees. I built this program for a stage show where multiple performances were announced and/or accompanied by information on a large screen. The contents of the screen can be controlled by any device in the same network using a web browser.

It works like this:

* You prepare a JSON configuration file where you define a number of so-called *Elements*. Each *Element* represents a single piece of content, for example a video or a marquee with some pre-defined text. Related *Elements* can be organized in *Groups*.
* When *Projector* starts, it displays an URL and a corresponding QR code on the screen. Behind this URL is a controller web interface which is generated from the JSON configuration. Each *Element* defined in the JSON file is represented by a button. When a button is clicked, the screen immediately switches to the view defined by the *Element*.

For more information on how to write a configuration, have a look at the [README file in the `conf` directory](conf/README.md).

## Screenshots

[![Screenshot: Welcome View](https://i.imgur.com/EhTpjSLl.png)](https://i.imgur.com/EhTpjSL.png)

[![Screenshot: Controller View](https://i.imgur.com/g2YZya7l.png)](https://i.imgur.com/g2YZya7.png)

## Currently supported **Element** types

* `video`
  * Parameter: Path to a video file relative to the `config` directory.
  * Presents a video file. When the video has finished, it stops on the last frame.
* `video-loop`
  * Parameter: Path to a video file relative to the `config` directory.
  * Same as `video`, but repeats the video after playback has finished.
* `image`
  * Parameter: Path to an image file relative to the `config` directory.
  * Presents an image file.
* `marquee`
  * Parameter: Text
  * Displays the specified string as a moving text.
* `black` 
  * Parameter: *None*.
  * Just a black screen.
* `welcome`
  * Parameter: *None*.
  * Displays the URL and QR code of the controller web interface. This is the view that is automatically opened when *Projector* starts.

## Requirements

* Go *(version 1.12.15 worked for me)*
* [Go Dep](https://github.com/golang/dep) *(version 0.5.1 worked for me)*
* Chrome/Chromium *(version >= 70, required by [Lorca](https://github.com/zserge/lorca))*
* I used Fedora and Ubuntu to develop and execute *Projector*. Windows and MacOS may work as well, but I didn't test it.

## Security (read this, it's important)

Please be aware that I created this program essentialy on a single Saturday evening and for the needs of a very specific event. I could trust all devices sharing the network with the PC that executed *Projector*. The web interface to control the screen is not protected in any way. Furthermore, it is possible to call the AJAX endpoints with arbitraty parameters. This implies that any user in the same network can make the screen display arbitrary information. For me, this was not an issue: I simply restricted network access and saved the effort to implement protection means in *Projector*. If this could be a problem for you, don't use *Projector* or add some protection to the web interface and the AJAX endpoints (pull requests are welcome!).

## How to compile

* Install the dependencies mentioned above. For Ubuntu, the packages `go-dep` and `chromium-browser` are only present in the universe repositories, so you have to enable those first:

  ```
  # add-apt-repository universe
  # apt install git golang make go-dep chromium-browser
  ```
* Clone this repository into your `GOPATH`:

  ```
  $ go get github.com/rfde/projector
  ```
* Change into the cloned directory (adjust if your `GOPATH` is different):

  ```
  $ cd ~/go/src/github.com/rfde/projector
  ```
* Run `make` in the root directory of this repository:

  ```
  $ make
  ```
  `make` will take care of the dependencies (using Go Dep and the files in `vendor/`) and produce the binary `bin/projector`.

## How to execute

Just run the binary and supply the path to a JSON configuration file as a command line parameter, for example:

```
$ bin/projector conf/my_config.json
```

You can close *Projector* by pressing `Alt`+`F4` or switch to another window with `Alt`+`Tab`.

## Thank you

I would like to thank all the developers who did the preparatory work that made it possible to implement *Projector* in no time. *Projector* makes use of the following open source projects:

* [Lorca](https://github.com/zserge/lorca) (for the integration with Chrome/Chromium)
* [TentCSS](https://css.sitetent.com/) (for the clear and responsive design in the web interface)
* [QRCode.js](https://davidshimjs.github.io/qrcodejs/) (for the QR code on the `welcome` screen)
* All the tutorials on how to implement a smooth marquee in HTML/CSS/JS, especially [this one](https://stackoverflow.com/a/38118591) and [this series](https://smil-control.de/category/smilcontrol-academy/lauftexte/).

## License

I release the portions of this repository which I created myself under the MIT license.
