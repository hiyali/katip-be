package handlers

import (
  "net/http"

	"github.com/labstack/echo"
  "github.com/dgrijalva/jwt-go"

  "github.com/hiyali/katip-be/config"
)

// Authorization: Bearer {TOKEN_HERE}
// return errors.New("failed to connect database")

// get positive integer from form value
/*
func getPosInt(c echo.Context, name string) (int, string) {
	intParam, err := strconv.Atoi(c.FormValue(name))

  if err != nil {
    return -1, fmt.Sprintf("Param %v is not a integer", name)
  }

  if intParam < 1 {
    return -1, fmt.Sprintf("Param %v can't be a negative integer", name)
  }

  return intParam, ""
} // */

/*
  @page  positive integer
  @limit positive integer

  return [record...]
*/
func RecordGetAllPageable(c echo.Context) (err error) {
  pageable := new(config.PageableParam)
  if err = c.Bind(pageable); err != nil {
    return
  }
  // if err = c.Validate(pageable); err != nil {
  //   return
  // }

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  var records []config.Record
  db.Where("creator_id = ?", claims.ID).Limit(pageable.Limit).Offset(pageable.Limit * (pageable.Page - 1)).Find(&records)

  return c.JSON(http.StatusOK, records)
}

func RecordCreateOne(c echo.Context) (err error) {
  record := new(config.Record)
  if err = c.Bind(record); err != nil {
    return
  }
  // if err = c.Validate(record); err != nil {
  //   return
  // }

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
  record.CreatorId = claims.ID

  db := config.GetDB()
  defer db.Close()

  if err := db.Create(&record).Error; err != nil {
    return c.JSON(http.StatusOK, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, record)
  }
}

func RecordGetOne(c echo.Context) (err error) {
  id := c.Param("id")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  var record config.Record
  if err := db.Where("creator_id = ? AND id = ?", claims.ID, id).First(&record).Error; err != nil {
    return c.JSON(http.StatusNotFound, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, record)
  }
}

func RecordUpdateOne(c echo.Context) (err error) {
  id := c.Param("id")

  record := new(config.Record)
  if err = c.Bind(record); err != nil {
    return
  }
  // if err = c.Validate(record); err != nil {
  //   return
  // }

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  var recordModel config.Record
  if err := db.Model(&recordModel).Where("creator_id = ? AND id = ?", claims.ID, id).Updates(record).Error; err != nil {
    return c.JSON(http.StatusOK, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, echo.Map{})
  }
}

func RecordDeleteOne(c echo.Context) (err error) {
  id := c.Param("id")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)

  db := config.GetDB()
  defer db.Close()

  var record config.Record
  if err := db.Where("creator_id = ? AND id = ?", claims.ID, id).Delete(&record).Error; err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": err,
    })
  } else {
    return c.JSON(http.StatusOK, echo.Map{})
  }
}
