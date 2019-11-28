package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
  "net/http"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
) 
func GetName(ctx *gin.Context) {
  name := ctx.Param("name")

  ctx.JSON(http.StatusOK, gin.H{
      "message": "hello " + name,
  })
}
func main () { 

  db, err := sql.Open("mysql", "root:root@/sample_database")
  if err!= nil {
    panic(err.Error())
  }

  fmt.Println(db)

  defer db.Close()

	router := gin.Default()

	router.GET("/cars/:name", GetName)

	if err := router.Run(); err != nil {
		fmt.Println("error")
	}

}



/*func AddCar(ctx *gin.Context) {
    body, err := ioutil.ReadAll(ctx.Request.Body)
    if err != nil {
		fmt.Println("error")
		return 
    }

    var car Auto
    if err = json.Unmarshal(body, &car); err != nil {
		fmt.Println("error")
		return
    }

    a.carDatabase.agenciaX.AgregarAuto(car)

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "car added",
    })
}*/