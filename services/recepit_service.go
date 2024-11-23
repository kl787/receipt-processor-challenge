package services

import (
	"github.com/google/uuid"
	"math"
	"receipt-processor-challenge/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GenerateID() string {
	return uuid.New().String()
}

func CalculatePoints(receipt models.Receipt) (int, error) {
	points := 0

	// One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if regexp.MustCompile(`^[a-zA-Z0-9]$`).MatchString(string(char)) {
			points += 1
		}
	}

	// 50 points if the total is a round dollar amount with no cents.
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.

	if math.Mod(float64(total), 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	itemTotals := float64(len(receipt.Items))
	points += 5 * int(itemTotals/2)

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		shortDescription := strings.TrimSpace(item.ShortDescription)
		if len(shortDescription)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd.
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if (purchaseTime.Hour() == 14 && purchaseTime.Minute() >= 0) || (purchaseTime.Hour() == 15 && purchaseTime.Minute() <= 59) {
		points += 10
	}

	return points, nil
}
