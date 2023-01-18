# Restraunt Service

## _Below are the required information what tools/packages used and how to run this project_.

### Developed using these tools
[![golang](https://emojis.slackmojis.com/emojis/images/1643514073/291/golang.png?1643514073)](https://go.dev/)&emsp;&emsp;&emsp;[![mysql](https://emojis.slackmojis.com/emojis/images/1643514462/4439/mysql.png?1643514462)](https://www.mysql.com/)&emsp;&emsp;&emsp;&emsp;[![docker](https://emojis.slackmojis.com/emojis/images/1643514952/9754/docker.png?1643514952)](https://www.docker.com/)


### Packages used
- [go-chi](https://github.com/go-chi/chi) - Router used instead of default one
- [sqlx](https://github.com/jmoiron/sqlx) - Provides useful options on top of standard database/sql package
- [go-chi render](https://github.com/go-chi/render) -The render package helps manage HTTP request / response payloads.
- [gotenv](https://github.com/subosito/gotenv) - Loading variable from .env

### How to run this project
>Note: Before running this project you need to install [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/install/) in your system. As this project relies on docker to run mysql database.

Change the dockerfile name in docker-compose.yaml in dockerfile path.
>Note: For development use devapi.Dockerfile for production use api.Dockerfile

Then run the command
```sh
docker-compose up -d
```
Docker will download the required images and run the containers.

### check the api docs
This project uses swaggerui for api documentation, Navigate to [http://localhost:8080/swaggerui/](http://localhost:8080/swaggerui/)


### How db migration works
This project uses golang-migrate directly inside docker so at the time of initailly running the project for very first time the migrate tool will directly create the required tables inside the db.