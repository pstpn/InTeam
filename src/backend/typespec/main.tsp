import "@typespec/http";
import "@typespec/rest";
import "@typespec/versioning";

import "./error.tsp";
import "./models.tsp";

@service(#{title: "RacketShop"})
@server("http://localhost:8080", "localhost")

@versioned(Versions)
namespace RacketShop;

enum Versions {
  v1
}

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.Versioning;

/**
  API Requests
**/
model RegisterRequest {
  @body
  registerBody: {
    email: string,
    name: string,
    password: string,
    surname: string
  }
};
model LoginRequest {
  @body
  loginBody: {
    email: string,
    password: string
  }
};
model AddRacketRequest {
  @body
  addRacketBody: {
    quantity: integer,
    racket_id: integer
  }
};
model DeleteRacketRequest {
  @path
  racket_id: integer
};
model UpdateRacketsCountRequest {
  @path
  racket_id: integer,
  @body
  updateRacketsCountBody: {
    quantity: integer
  }
};
model CreateFeedbackRequest {
  @body
  createFeedbackBody: {
    feedback: string,
    racket_id: integer,
    rating: integer
  }
}
model DeleteFeedbackRequest {
  @path
  racket_id: integer
};
model CreateOrderRequest {
  @body
  createOrderBody: {
    address: string,
    delivery_date: plainTime,
    recipient_name: string
  }
}

/**
  API Responses
**/
model RegisterResponse extends ResponseWithAccessToken {};
model LoginResponse extends ResponseWithAccessToken {};
model GetCartResponse extends ResponseWithCart {};
model AddRacketResponse extends ResponseWithCart {};
model DeleteRacketResponse extends ResponseWithCart {};
model UpdateRacketsCountResponse extends ResponseWithCart {};
model GetFeedbacksResponse extends ResponseWithFeedbacks {};
model CreateFeedbackResponse extends ResponseWithFeedback {};
model GetProfileResponse extends ResponseWithUser {};

/**
  API Definition
**/
@route("/api")
namespace Api {

  @tag("auth")
  @route("/auth")
  namespace Auth {
    @summary("register user")
    @route("/register")
    @post op register(...RegisterRequest): RegisterResponse | ErrorResponse;

    @summary("login user")
    @route("/login")
    @post op login(...LoginRequest): LoginResponse | ErrorResponse;
  }

  @tag("user")
  namespace User {
    @summary("get cart items")
    @route("/cart")
    @useAuth(BearerAuth)
    @get op getCart(): GetCartResponse | ErrorResponse;

    @summary("add racket in cart")
    @route("/cart")
    @useAuth(BearerAuth)
    @post op addRacket(...AddRacketRequest): AddRacketResponse | ErrorResponse;

    @summary("delete racket from cart")
    @route("/cart/rackets/{racket_id}")
    @useAuth(BearerAuth)
    @delete op deleteRacket(...DeleteRacketRequest): DeleteRacketResponse | ErrorResponse;

    @summary("update rackets count in cart")
    @route("/cart/rackets/{racket_id}")
    @useAuth(BearerAuth)
    @put op updateRacketsCount(...UpdateRacketsCountRequest): UpdateRacketsCountResponse | ErrorResponse;

    @summary("get user feedbacks")
    @route("/feedbacks")
    @useAuth(BearerAuth)
    @get op getFeedbacks(): GetFeedbacksResponse | ErrorResponse;

    @summary("create feedback")
    @route("/feedbacks")
    @useAuth(BearerAuth)
    @post op createFeedback(...CreateFeedbackRequest): CreateFeedbackResponse | ErrorResponse;

    @summary("delete feedback")
    @route("/feedbacks/{racket_id}")
    @useAuth(BearerAuth)
    @delete op deleteFeedback(...DeleteFeedbackRequest): OkResponse | ErrorResponse;

    @summary("create order")
    @route("/orders")
    @useAuth(BearerAuth)
    @post op createOrder(...CreateOrderRequest): OkResponse | ErrorResponse;

    @summary("get user profile")
    @route("/profile")
    @useAuth(BearerAuth)
    @get op getProfile(): GetProfileResponse | ErrorResponse;
  }
}