# GoWebServer
Go-based web-server for a Boise State University team project

## Running (on linux)
I'm not sure how the following applies to other platforms, but it should be similar.
You need root access (able to use sudo) to run this, otherwise you won't have the correct
permissions to host a webserver.
* Be in the project root directory (containing the assets folder and testserver.go)
* Run ```go build``` to compile the program and generate an executable
* Run ```sudo ./testserver```
* Open up your favorite web-browser and type in 'localhost' to see the website
