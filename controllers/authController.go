package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SHerlihy/profile-service-rfg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Profile struct {
	Theme string
}

func Create(c *fiber.Ctx) error {
    fmt.Println("CREATE CALLED")
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	profile := Profile{
		Theme: "default",
	}

    profBytes, err := json.Marshal(profile)
    profileStr := string(profBytes)

	expiration, _ := time.ParseDuration("10s")

	err = database.DBClient.Set(c.Context(), data["profileKey"], profileStr, expiration).Err()
	if err != nil {
        fmt.Println("set failed")
        fmt.Println(err)
		panic(err)
	}

	val, err := database.DBClient.Get(c.Context(), data["profileKey"]).Result()
	switch {
	case err == redis.Nil:
		c.SendStatus(204)
		return c.JSON(fiber.Map{
			"message": "key does not exist",
		})
	case err != nil:
		c.SendStatus(204)
		return c.JSON(fiber.Map{
			"message": err,
		})
	case val == "":
		c.SendStatus(204)
		return c.JSON(fiber.Map{
			"message": "value is empty",
		})
	}

	c.SendStatus(200)
	return c.JSON(val)
}

func UpdateAccess(c *fiber.Ctx) error {
	c.SendStatus(200)
    return nil
}

//
//type Profile struct {
//	Theme string
//}
//
//func (i *Profile) setProperty(propName string, propValue string) *Profile {
//	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
//	return i
//}
//
//type Cookie struct {
//	Value string `cookie:"value"`
//}
//
//func Create(c *fiber.Ctx) error {
//	cookie := new(Cookie)
//
//	if err := c.CookieParser(cookie); err != nil {
//		return err
//	}
//
//    profile := Profile{
//        Theme: "default",
//    }
//
//    profileVal := reflect.ValueOf(profile)
//
//    profileValues := make([]interface{}, profileVal.NumField())
//
//    for i:=0;i<profileVal.NumField();i++{
//        profileValues[i] = profileVal.Field(i).Interface()
//    }
//
//	ctx := context.Background()
//	for k, v := range profileValues {
//		err := database.DBClient.HSet(ctx, cookie.Value, k, v).Err()
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	c.SendStatus(200)
//	return c.JSON(profile)
//}
//
//func UpdateAccess(c *fiber.Ctx) error {
//	cookie := new(Cookie)
//
//	if err := c.CookieParser(cookie); err != nil {
//		return err
//	}
//
//	var data map[string]string
//
//	if err := c.BodyParser(&data); err != nil {
//		return err
//	}
//
//	ctx := context.Background()
//	userProfile := database.DBClient.HGetAll(ctx, data["prevCookie"]).Val()
//
//	for k, v := range userProfile {
//		err := database.DBClient.HSet(ctx, cookie.Value, k, v).Err()
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	var cpProfile Profile
//	for k, v := range userProfile {
//		cpProfile.setProperty(k, v)
//	}
//
//	c.SendStatus(200)
//	return c.JSON(cpProfile)
//}
