package handlers

import (
	"backend/internal/model"
	api "backend/internal/router/ogen"
)

func modelToAPICartLines(cartLines []*model.CartLine) []api.CartLine {
	apiCartLines := make([]api.CartLine, 0, len(cartLines))
	for _, line := range cartLines {
		apiCartLines = append(apiCartLines, api.CartLine{
			RacketID: line.RacketID,
			Price:    line.Price,
			Quantity: line.Quantity,
		})
	}
	return apiCartLines
}

func modelToAPIFeedbacks(feedbacks []*model.Feedback) []api.Feedback {
	apiFeedbacks := make([]api.Feedback, 0, len(feedbacks))
	for _, feedback := range feedbacks {
		apiFeedbacks = append(apiFeedbacks, api.Feedback{
			Date:     feedback.Date,
			Feedback: feedback.Text,
			RacketID: feedback.RacketID,
			Rating:   feedback.Rating,
		})
	}
	return apiFeedbacks
}

func modelToAPIFeedbacksWithUsername(feedbacks []*model.FeedbackWithUsername) []api.FeedbackWithUsername {
	apiFeedbacks := make([]api.FeedbackWithUsername, 0, len(feedbacks))
	for _, feedback := range feedbacks {
		apiFeedbacks = append(apiFeedbacks, api.FeedbackWithUsername{
			Date:     feedback.Date,
			Feedback: feedback.Text,
			RacketID: feedback.RacketID,
			Rating:   feedback.Rating,
			UserID:   feedback.UserID,
			Username: feedback.Username,
		})
	}
	return apiFeedbacks
}

func modelToAPIOrderLines(orderLines []*model.OrderLine) []api.OrderLine {
	apiOrderLines := make([]api.OrderLine, 0, len(orderLines))
	for _, line := range orderLines {
		apiOrderLines = append(apiOrderLines, api.OrderLine{
			RacketID: line.RacketID,
			Quantity: line.Quantity,
		})
	}
	return apiOrderLines
}

func modelToAPIOrder(order *model.Order) api.Order {
	return api.Order{
		ID:            order.ID,
		UserID:        order.UserID,
		CreationDate:  order.CreationDate,
		DeliveryDate:  order.DeliveryDate,
		Address:       order.Address,
		RecipientName: order.RecipientName,
		Status:        string(order.Status),
		Lines:         modelToAPIOrderLines(order.Lines),
		TotalPrice:    order.TotalPrice,
	}
}

func modelToAPIOrders(orders []*model.Order) []api.Order {
	apiOrders := make([]api.Order, 0, len(orders))
	for _, order := range orders {
		apiOrders = append(apiOrders, modelToAPIOrder(order))
	}
	return apiOrders
}

func modelToAPIRacket(racket *model.Racket) api.Racket {
	return api.Racket{
		ID:        racket.ID,
		Brand:     racket.Brand,
		Weight:    racket.Weight,
		Balance:   racket.Balance,
		HeadSize:  racket.HeadSize,
		Available: racket.Available,
		Quantity:  racket.Quantity,
		Price:     racket.Price,
		Image:     racket.Image,
	}
}

func modelToAPIRackets(rackets []*model.Racket) []api.Racket {
	apiRackets := make([]api.Racket, 0, len(rackets))
	for _, racket := range rackets {
		apiRackets = append(apiRackets, modelToAPIRacket(racket))
	}
	return apiRackets
}

func modelToAPIUser(user *model.User) api.User {
	return api.User{
		Email:   user.Email,
		Name:    user.Name,
		Role:    string(user.Role),
		Surname: user.Surname,
	}
}

func modelToAPIUsers(users []*model.User) []api.User {
	apiUsers := make([]api.User, 0, len(users))
	for _, user := range users {
		apiUsers = append(apiUsers, modelToAPIUser(user))
	}
	return apiUsers
}
