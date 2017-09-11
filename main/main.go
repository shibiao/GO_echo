package main

import (
	"github.com/labstack/echo"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/labstack/gommon/log"
	"encoding/json"
)

type Cat struct {
	Name     string      `json:"name"`
	Type     string		 `json:"type"`
}
type Dog struct {
	Name     string      `json:"name"`
	Type     string		 `json:"type"`
}
type Hamster struct {
	Name     string      `json:"name"`
	Type     string		 `json:"type"`
}
func yallo(c echo.Context) error {
	return  c.String(http.StatusOK,"yallo1 from the web side")
}
func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is %s\n and his type is %s\n",catName,catType ))
	}
	if dataType == "json" {
		return  c.JSON(http.StatusOK,map[string]string{
			"name": catName,
			"type": catType,
		})
	}
	return c.JSON(http.StatusBadRequest,map[string]string {
		"error":"you need to lets us",
	})
}
func addCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()
	b,err := ioutil.ReadAll(c.Request().Body)
	if err!=nil {
		log.Printf("Failed reading the request body: %s",err)
		return c.String(http.StatusInternalServerError,"")
	}
	err = json.Unmarshal(b,&cat)
	if err != nil {
		log.Printf("Failed unmarshaling in addCats: %s",err)
		return c.String(http.StatusInternalServerError,"")
	}
	log.Printf("this is your cat: %#v",cat)
	return c.String(http.StatusOK,"we got your cat")
}
func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing addDog request: %s",err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this is your dog: %#v",dog)
	return c.String(http.StatusOK,"we got your dog")

}
func addHamster(c echo.Context) error {
	hamster := Hamster{}
	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed Processing addHamster request:%s",err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this is your hamster:%#v",hamster)
	return c.String(http.StatusOK,"we got your hamster!")
}
func main() {
	fmt.Println("Welcome to the server")

	e := echo.New()
	//GET请求你
	e.GET("/", yallo)
	e.GET("/cats/:data",getCats)
	//POST请求
	e.POST("cats",addCat)
	e.POST("dogs",addDog)
	e.POST("/hamsters",addHamster)

	e.Start(":8001")
}