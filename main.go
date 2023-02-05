package main

import (
	"net/http"

	"github.com/NagatsukiSep/asset-manager-server/model"
	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Asset struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Amount    int    `json:"amount" db:"amount"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	db, err := model.InitDB()
	if err != nil {
		panic(err)
	}

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello!!!")
		// return c.String(http.StatusOK, subpkg.Hello())
	})

	e.POST("/add", func(c echo.Context) error {
		var asset Asset
		if err := c.Bind(&asset); err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		if asset.Name == "" {
			return c.String(http.StatusBadRequest, "Name is required")
		}
		if asset.Amount <= 0 {
			return c.String(http.StatusBadRequest, "Amount is required")
		}

		sql := `INSERT INTO asset (id, name, amount) VALUES (UUID(), :name, :amount);`
		_, err = db.NamedExec(sql, asset)
		if err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/list", func(c echo.Context) error {
		var assets []Asset
		sql := `SELECT * FROM asset;`
		err = db.Select(&assets, sql)
		if err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.JSON(http.StatusOK, assets)
	})

	e.POST("/delete", func(c echo.Context) error {
		var asset Asset
		if err := c.Bind(&asset); err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		id := asset.ID
		if id == "" {
			return c.String(http.StatusBadRequest, "ID is invalid")
		}
		sql := `DELETE FROM asset WHERE id = ?;`
		_, err = db.Exec(sql, id)
		if err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.String(http.StatusOK, "OK")
	})

	e.POST("/update-amount", func(c echo.Context) error {
		var asset Asset
		if err := c.Bind(&asset); err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		id := asset.ID
		if id == "" {
			return c.String(http.StatusBadRequest, "ID is required")
		}
		amount := asset.Amount
		if amount <= 0 {
			return c.String(http.StatusBadRequest, "Amount is invalid")
		}
		sql := `UPDATE asset SET amount = ? WHERE id = ?;`
		_, err = db.Exec(sql, amount, id)
		if err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.String(http.StatusOK, "OK")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
