package main

import (
    "github.com/kataras/iris"
)

type baseTemplate struct {
    Title, Resources string
}

type user struct {
    Name, Email string
    Password string // hehehe
}

// App provides necessary data and functions to handle requests
type App struct {
    RegisteredUsers []user
}

// Static information about the site
const (
    ResourceURL string = "/res"
    Port string = ":8080"
)

func main() {
    app := App{}
    iris.OptionIsDevelopment(true)
    iris.Static(ResourceURL, "./templates/", 1)

    iris.Get("/", app.home)

    app.RegisteredUsers = append(app.RegisteredUsers, user{ "admin", "admin@blogexample.com", "1234!"})
    app.RegisteredUsers = append(app.RegisteredUsers, user{ "lebenasa", "lebenasa@hotmail.com", "1234" })

    iris.Get("/join-check/username/:name", app.checkUsername)

    iris.Listen(Port)
}

func (a App) home(ctx *iris.Context) {
    ctx.Render("index_template.html", baseTemplate{"Iris Play!", ResourceURL}, iris.RenderOptions{"gzip": true})
}

func (a App) checkUsername(ctx *iris.Context) {
    testname := ctx.Param("name")
    for _, i := range a.RegisteredUsers {
        if testname == i.Name {
            ctx.JSON(iris.StatusOK, map[string]string{"available": "false"})
            return
        }
    }
    ctx.JSON(iris.StatusOK, map[string]string{"available":"true"})
}
