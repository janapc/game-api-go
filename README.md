<div align="center">
  <h1>Game API</h1>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/janapc/game-api-go"/>
  <img alt="Language top" src="https://img.shields.io/github/languages/top/janapc/game-api-go"/>
  <img alt="Repo size" src="https://img.shields.io/github/repo-size/janapc/game-api-go"/>

<a href="#project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#run-project">Run Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#request-api">Request API</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#technologies">Technologies</a>

</div>

## Project

Api to register games(name, description, release date) and manager those games.Used only golang and mysql.

## Requirement

To this project your need:

- golang v1.21 [Golang](https://go.dev/)
- docker [Docker](https://www.docker.com/)

Inside _.docker_ create a first file **mysql_password.txt** and put your password mysql and create a second file **mysql_root_password** and put your password root mysql.This configuration wil be used by _docker-compose.yaml_.

In root folder create a file **.env** with:

```env
MYSQL_URL="" // mysql url connection
```

## Run Project

Start Docker in your machine and run this commands in your terminal:

```sh
## up mongodb, kafka and mysql
❯ docker compose up -d

## run this command to start api(localhost:8080):
❯ go run main.go

```

## Technologies

- golang
- mysql
- docker

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
