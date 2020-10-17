<div>
  <div style="float: left;">
    <img width="200" src="https://golang.org/lib/godoc/images/go-logo-blue.svg" alt="Go Language icon">
  </div>
  <div style="float: left;">
    <img width="200" src="https://www.vernier.com/wp-content/uploads/2020/05/Arduino-Loop-logo.png"   alt="Arduino icon">
  </div>
</div>

# Goduino
This is a hello world starting point for using golang with an arduino

## Setup

### Install Dependencies
* `go` [golang install](https://golang.org/doc/install)
* `platformio` [PlatformIO Core (CLI) install](http://docs.platformio.org/en/latest/installation.html)

### Install firmata on an arduino
* `cd platformio`
* compile: `platformio run -e uno`
* burn to uno: `platformio run -e uno --target upload`
* See the below link if different firmata versions for different functions are desired
  * [Firmata Examples](https://platformio.org/lib/show/307/Firmata/examples)
* over time operator may need to update the "platforms" that platformio supports
  * one use case is if an old platformio installation was using python2 code
  * fix with `platformio platform update`

### gobot.io
* `go get -d -u gobot.io/x/gobot/...`