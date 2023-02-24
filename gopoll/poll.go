package main

import (
    // "database/sql"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "gopoll/handlers"
    // _ "github.com/mattn/go-sqlite3"
)

func main(){
    e:= echo.New()

    //middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    //definindo as rotas HTTP
    e.GET("/polls", func(c echo.Context) error{
        return c.JSON(200, "PUT Polls")
    })

    e.PUT("/polls", func(c echo.Context) error{
        return  c.JSON(200, "PUT Polls")
    })

    e.PUT("/polls", func(c echo.Context) error{
        return c.JSON(200, "UPDATE Poll " + c.Param("id"))
    })

    //startando servidor
    e.Logger.Fatal(e.Start(":9000"))
}

func initDB(filepath string) *ql.DB{
    db, err := sql.Open("sqlite3", filepath)

    if err !: nil{
        panic(err)
    }

    if db == nil {
        panic("db nil")
    }

    return db
}

func migrate(db *sql.DB) {
    sql := `
            CREATE TABLE IF NOT EXISTS polls(
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            name VARCHAR NOT NULL,
            topic VARCHAR NOT NULL,
            src VARCHAR NOT NULL,
            upvotes INTEGER NOT NULL,
            downvotes INTEGER NOT NULL,
            UNIQUE(name)
            );

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Angular','Awesome Angular', 'https://cdn.colorlib.com/wp/wp-content/uploads/sites/2/angular-logo.png', 1, 0);

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Vue', 'Voguish Vue','https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Vue.js_Logo.svg/400px-Vue.js_Logo.svg.png', 1, 0);

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('React','Remarkable React','https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/1200px-React-icon.svg.png', 1, 0);

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Ember','Excellent Ember','https://cdn-images-1.medium.com/max/741/1*9oD6P0dEfPYp3Vkk2UTzCg.png', 1, 0);

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Knockout','Knightly Knockout','https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_1489710848/knockout-js.png', 1, 0);
            `
    _, err := db.Exec(sql)

    if err != nil {
        panic(err)
    }
}
