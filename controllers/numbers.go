package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nanafox/simple-http-client/pkg/client"

	"github.com/nanafox/number-classifier-api-go/internal"
)

// GetNumberClassification returns a classification of the number received.
func GetNumberClassification(c *fiber.Ctx) error {
	// timeout after 200ms of no response
	client := client.ApiClient{Timeout: 200 * time.Millisecond}

	query := c.Query("number")

	number, err := strconv.Atoi(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"number": query, "error": true,
		})
	}

	url := fmt.Sprintf("http://numbersapi.com/%d/math/", number)

	client.Get(url, nil) // fetch the fun fact

	funFact := client.Body

	if funFact == "" {
		// use this as a fallback when the Numbers API fails to reply in time
		funFact = fmt.Sprintf("%d is cool number", number)
	}

	return c.JSON(fiber.Map{
		"number":     number,
		"is_prime":   internal.IsPrime(number),
		"is_perfect": internal.IsPerfect(number),
		"properties": internal.NumberProperties(number),
		"digit_sum":  internal.DigitSum(number),
		"fun_fact":   funFact,
	})
}
