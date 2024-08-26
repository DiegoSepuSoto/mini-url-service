# Mini URL Service

## Overview

This Golang API will serve and return mini URLs created with [Mini URL Builder API](https://github.com/DiegoSepuSoto/mini-url-builder-api)

## Local workspace

First: clone the repository:

```bash
git clone https://github.com/DiegoSepuSoto/mini-url-service && cd mini-url-service
```

Then, download the dependencies:

```bash
go mod download
```

Now you can run the application using the Makefile:

```bash
make run
```

The available endpoint is the following:

```bash
curl --location --request GET 'localhost:8081/api/xyz789' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.8x2hIBGylPBtKnAoEP8wJqqXbXaQyOK0z8bjpasZGfo'
```

which will return the mini URL created.

You can even visit the following link: **localhost:8081/xyz789**
to be redirected to the original URL.

Also, you can access:

- Prometheus metrics at: **localhost:8081/metrics**
- Swagger documentation at: **localhost:8081/swagger/index.html**

### Tech Stack

- Golang library - Echo framework for http server
- Golang library - Logrus for application logs
- Golang library - Viper for application environment variables
- Golang library - Testify for unit testing
- Golang library - Testcontainer for integration testing
- Prometheus metrics
- Docker & Docker Compose
