package handlers

import (
	"io"
	"net/http"

	"github.com/Omotolani98/github-insights/config"
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": "🚀Server is Healthy!",
	})
}

func GetUserAccessToken(c *fiber.Ctx) error {
	code := c.Params("code")
	url := "https://github.com/login/oauth/access_token?client_id=" + config.LoadEnv().GITHUB_CLIENT_ID + "&client_secret=" + config.LoadEnv().GITHUB_CLIENT_SECRET + "&code=" + code

	log.Infof("URL ::>> %s", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("Could not set up Http Client")
		return c.JSON(&fiber.Map{
			"code": 500,
			"message": "Error on Http Client",
		})
	}
	
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("Error on API <::> %v", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"message": "Error on Github Auth",
		})
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error on body <::> %v", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"message": "Error on Github Auth",
		})
	}

	if body != nil {
		return c.JSON(&fiber.Map{
			"code": 200,
			"message": "Access Token Generated",
			"body": body,
		})
	}
	
	return nil
}
