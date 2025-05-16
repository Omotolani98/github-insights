package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Omotolani98/github-insights/config"
	"github.com/Omotolani98/github-insights/models"
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": "🚀Server is Healthy!",
	})
}

func GetUserAccessToken(c *fiber.Ctx) error {
	code := c.Query("code")
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

	var tokenResponse models.TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	
	log.Infof("TokenResponse :: %s", tokenResponse)
	if err != nil {
		log.Errorf("Unmarshal could not be completed ::> %v", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"message": "Could not parse response data",
		})
	}
		
	return c.JSON(&fiber.Map{
		"code": 200,
		"message": "Access Token Generated",
		"body": tokenResponse,
	})	
}

func GetUserDetails(c *fiber.Ctx) error {
	authHeader := c.GetReqHeaders()
	bearer := authHeader["Authorization"][0]
	token := bearer[7:]
	log.Infof("Token :: %s", token)

	url := "https://api.github.com/user"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("Could not set up Http Client")
		return c.JSON(&fiber.Map{
			"code": 500,
			"message": "Error on Http Client",
		})
	}
	
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Authorization", "Bearer " + token)

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
	
	var userDetailsResponse models.UserDetailsResponse
	err = json.Unmarshal(body, &userDetailsResponse)

	log.Infof("Data Response :: %s", body)
	if err != nil {
		log.Errorf("Unmarshal could not be completed ::> %v", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"message": "Could not parse response data",
		})
	}
		
	return c.JSON(&fiber.Map{
		"code": 200,
		"message": "User Details Fetched",
		"body": userDetailsResponse,
	})
}
