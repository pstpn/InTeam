using TypeSpec.Http;
using TypeSpec.Rest;

model ResponseWithAccessToken {
  access_token: string,
}

model ResponseWithCart {
  cart: Cart
}

model ResponseWithFeedbacks {
  feedbacks: Array<Feedback>
}

model ResponseWithFeedback {
  feedback: Feedback
}

model ResponseWithUser {
  user: User
}

model CartLine {
  racket_id: integer,
  price: float32,
  quantity: integer
}

model Cart {
  lines: Array<CartLine>,
  quantity: integer,
  total_price: float32
}

model Feedback {
  date: plainTime,
  feedback: string,
  racket_id: integer,
  rating: integer
}

model User {
  email: string,
  name: string,
  role: string,
  surname: string
}