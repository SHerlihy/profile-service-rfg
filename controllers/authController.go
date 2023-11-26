package controllers

import (
	"context"
	"reflect"

	"github.com/SHerlihy/profile-service-rfg/database"
	"github.com/gofiber/fiber/v2"
)

type Profile struct {
	Theme string
}

func (i *Profile) setProperty(propName string, propValue string) *Profile {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

type Cookie struct {
	Value string `cookie:"value"`
}

func Create(c *fiber.Ctx) error {
	cookie := new(Cookie)

	if err := c.CookieParser(cookie); err != nil {
		return err
	}

    profile := Profile{
        Theme: "default",
    }

    profileVal := reflect.ValueOf(profile)

    profileValues := make([]interface{}, profileVal.NumField())

    for i:=0;i<profileVal.NumField();i++{
        profileValues[i] = profileVal.Field(i).Interface()
    }

	ctx := context.Background()
	for k, v := range profileValues {
		err := database.DBClient.HSet(ctx, cookie.Value, k, v).Err()
		if err != nil {
			panic(err)
		}
	}

	c.SendStatus(200)
	return c.JSON(profile)
}

func UpdateAccess(c *fiber.Ctx) error {
	cookie := new(Cookie)

	if err := c.CookieParser(cookie); err != nil {
		return err
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	ctx := context.Background()
	userProfile := database.DBClient.HGetAll(ctx, data["prevCookie"]).Val()

	for k, v := range userProfile {
		err := database.DBClient.HSet(ctx, cookie.Value, k, v).Err()
		if err != nil {
			panic(err)
		}
	}

	var cpProfile Profile
	for k, v := range userProfile {
		cpProfile.setProperty(k, v)
	}

	c.SendStatus(200)
	return c.JSON(cpProfile)
}
