# GoWebServer
Go-based web-server for a Boise State University team project

## Read before cloning!
To make sure this works correctly, make sure to put it in your Go src directory, where all your Go projects are located. Naturally, this means you'll want to make sure you have Go installed and working on your computer before cloning into this project.

## Running (on linux)
I'm not sure how the following applies to other platforms, but it should be similar.
You need root access (able to use sudo) to run this, otherwise you won't have the correct
permissions to host a webserver.
* run ```go get -u github.com/gorilla/mux``` to get a third party dependancy
* Be in the project root directory (containing the assets folder and testserver.go)
* Run ```go build``` to compile the program and generate an executable
* Run ```sudo ./testserver```
* Open up your favorite web-browser and type in 'localhost' to see the website
