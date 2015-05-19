#Lets Kill Things
A tool for helping manage combat in a turn based environment. Allowing DM's/GM's to control stats such as AC, HP, and damage taken for a group.

##API Server
We currently use a mock API server (Drakov), whhich reads in [API Blueprint](https://apiblueprint.org/) to generate a mock server.

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
