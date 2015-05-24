[![Coverage Status](https://coveralls.io/repos/fear-the-dice/api/badge.svg)](https://coveralls.io/r/fear-the-dice/api)
[![Build Status](https://travis-ci.org/fear-the-dice/api.svg)](https://travis-ci.org/fear-the-dice/api)

#Fear the Dice
A tool for helping manage combat in a turn based environment. Allowing DM's/GM's to control stats such as AC, HP, and damage taken for a group.

##API Server
Our API server is built in Go.

* **Use:**
> To quickly run the server you can use the following command:
    ```
    $ go run server.go
    ```

* **Installation:**
> To install the server (which is required if you are using supervisor) you have to run:
    ```
    $ go get
    ```
> Then you can run:
    ```
    $ api 
    ```

##Mock API
We have a mock API server (Drakov), whhich reads in [API Blueprint](https://apiblueprint.org/) to generate an API server.

* **Installation:**

    ```
    $ npm install -g drakov
    ```
* **Use:**

    ```
    $ drakov -f api/server.MD
    ```

##License
This tool is protected by the [GNU General Public License v2](http://www.gnu.org/licenses/gpl-2.0.html).

Copyright [Jeffrey Hann](http://jeffreyhann.ca/) 2015
