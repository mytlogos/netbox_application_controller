# Iris MVC Application Template

[![build status](https://img.shields.io/github/actions/workflow/status/mytlogos/netbox_application_controller/ci.yml?branch=main&style=for-the-badge)](https://github.com/mytlogos/netbox_application_controller/actions) [![report card](https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/mytlogos/netbox_application_controller)

This is a project template for [Iris](https://iris-go.com) MVC apps.

## Prerequisites

You will need to have:

- [Iris CLI](https://github.com/kataras/iris-cli) installed to run the `iris-cli` command.
- [Node.js](https://nodejs.org) to build the frontend application.

## Get started

Install the template with [Iris CLI](https://github.com/kataras/iris-cli):

```sh
iris-cli new --module=my-app mvc
```

Install the dependencies and build the MVC app manually with [Rollup](https://rollupjs.org)...

```sh
cd app
npm install
npm run build
```

...then **start the Iris web server**:

```sh
cd ../
go run main.go
```

Navigate to [localhost:8080](http://localhost:8080). You should see your app running.
