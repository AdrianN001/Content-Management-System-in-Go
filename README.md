# Http webserver, written in Go

Inspired by [this video](https://youtu.be/cEH_ipqHbUw "A http webserver written in C language, and for some reason, I learned Go for this project")

## Routing 

Routes are handled by the file system. 
The "static" folder contains some examples. 
| File system            | Route              |
|------------------------|--------------------|
| static/home/main.html  | http://[...]/home  |
| static/about/main.html | http://[...]/about |

The default html of the route is "main.html". 
