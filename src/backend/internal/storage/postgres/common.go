package mypostgres

import "fmt"

const (
	userTable        = "\"user\""
	racketTable      = "racket"
	orderTable       = "\"order\""
	cartTable        = "cart"
	feedbackTable    = "feedback"
	deliveryTable    = "delivery"
	cartRacketTable  = "cart_racket"
	orderRacketTable = "order_racket"

	idField            = "id"
	userIDField        = "user_id"
	orderIDField       = "order_id"
	cartIDField        = "cart_id"
	racketIDField      = "racket_id"
	feedbackField      = "feedback"
	nameField          = "name"
	surnameField       = "surname"
	totalPriceField    = "total_price"
	emailField         = "email"
	passwordField      = "password"
	roleField          = "role"
	statusField        = "status"
	quantityField      = "quantity"
	priceField         = "price"
	brandField         = "brand"
	weightField        = "weight"
	balanceField       = "balance"
	headSizeField      = "head_size"
	addressField       = "address"
	recipientNameField = "recipient_name"
	creationDateField  = "creation_date"
	deliveryDateField  = "delivery_date"
	ratingField        = "rating"
	dateField          = "date"
	townField          = "town"
	phoneField         = "phone"
	subscriptionField  = "subscription"
	availableField     = "available"
	imageField         = "image"
)

func on(baseTable string, targetTable string, baseColumn string, targetColumn string) string {
	return fmt.Sprintf("%s on %s.%s=%s.%s",
		targetTable,
		baseTable,
		baseColumn,
		targetTable,
		targetColumn,
	)
}
