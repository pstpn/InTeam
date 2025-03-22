using TypeSpec.Http;
using TypeSpec.Rest;

model WithPaging {
  @query
  pattern?: string,
  @query
  field?: string,
  @query
  sort?: string
}

model WithRacketId {
  @path
  racket_id: integer
}

model WithOrderId {
  @path
  order_id: integer
}

model WithUserId {
  @path
  user_id: integer
}

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

model ResponseWithUsers {
  users: Array<User>
}

model ResponseWithUser {
  user: User
}

model ResponseWithOrders {
  orders: Array<Order>
}

model ResponseWithOrder {
  order: Order
}

model ResponseWithRackets {
  rackets: Array<Racket>
}

model ResponseWithRacket {
  racket: Racket
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

model OrderLine {
  quantity: integer,
  racket_id: integer
}

model Order {
  id: integer,
  user_id: integer,
  creation_date: plainTime,
  delivery_date: plainTime,
  address: string,
  recipient_name: string,
  status: string,
  lines: Array<OrderLine>,
  total_price: float32
}

model Racket {
  id: integer,
  brand: string,
  weight: float32,
  balance: float32,
  head_size: float32,
  available: boolean,
  quantity: integer,
  price: float32,
  image: bytes
}