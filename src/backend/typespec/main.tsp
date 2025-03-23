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
  ...WithRacketId
};
model UpdateRacketsCountRequest {
  ...WithRacketId,
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
  ...WithRacketId
};
model CreateOrderRequest {
  @body
  createOrderBody: {
    address: string,
    delivery_date: plainTime,
    recipient_name: string
  }
}
model GetOrderRequest {
  ...WithOrderId
}
model UpdateOrderStatusRequest {
  ...WithOrderId,
  @body
  updateStatusBody: {
    status: string
  }
}
model CreateRacketRequest {
  @header
  contentType: "multipart/form-data",
  @multipartBody
  createRacketBody: {
    balance: HttpPart<float32>,
    brand: HttpPart<string>,
    head_size: HttpPart<float32>,
    price: HttpPart<float32>,
    // Float32 for fix ogen encode error.
    // Original value - integer. Need to convert on server side
    quantity: HttpPart<float32>,
    weight: HttpPart<float32>,
    image: HttpPart<bytes>
  }
}
model UpdateRacketQuantityRequest {
  ...WithRacketId,
  @body
  updateQuantityBody: {
    quantity: integer
  }
}
model GetUserRequest {
  ...WithUserId
}
model ChangeUserRole {
  ...WithUserId,
  @body
  changeRoleBody: {
    role: string
  }
}
model GetOrdersRequest {
  ...WithPaging
}
model GetRacketFeedbacksRequest {
  ...WithRacketId
}
model GetRacketsRequest {
  ...WithPaging
}
model GetRacketRequest {
  ...WithRacketId
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
model GetOrderResponse extends ResponseWithOrder {};
model CreateRacketResponse extends ResponseWithRacket {};
model GetUsersResponse extends ResponseWithUsers {};
model GetUserResponse extends ResponseWithUser {};
model GetOrdersResponse extends ResponseWithOrders {};
model GetRacketFeedbacksResponse extends ResponseWithFeedbacks {};
model GetRacketsResponse extends ResponseWithRackets {};
model GetRacketResponse extends ResponseWithRacket {};

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
  @useAuth(BearerAuth)
  namespace User {
    @route("/cart")
    namespace Cart {
      @summary("get cart items")
      @get op getCart(): GetCartResponse | ErrorResponse;

      @summary("add racket in cart")
      @post op addRacket(...AddRacketRequest): AddRacketResponse | ErrorResponse;

      @route("/rackets")
      namespace Rackets {
        @summary("delete racket from cart")
        @delete op deleteRacket(...DeleteRacketRequest): DeleteRacketResponse | ErrorResponse;

        @summary("update rackets count in cart")
        @put op updateRacketsCount(...UpdateRacketsCountRequest): UpdateRacketsCountResponse | ErrorResponse;
      }
    }

    @route("/feedbacks")
    namespace Feedbacks {
      @summary("get user feedbacks")
      @get op getFeedbacks(): GetFeedbacksResponse | ErrorResponse;

      @summary("create feedback")
      @post op createFeedback(...CreateFeedbackRequest): CreateFeedbackResponse | ErrorResponse;

      @summary("delete feedback")
      @delete op deleteFeedback(...DeleteFeedbackRequest): OkResponse | ErrorResponse;
    }

    @route("/orders")
    namespace Orders {
      @summary("create order")
      @post op createOrder(...CreateOrderRequest): OkResponse | ErrorResponse;
    }

    @route("/profile")
    namespace Profile {
      @summary("get user profile")
      @get op getProfile(): GetProfileResponse | ErrorResponse;
    }
  }

  @tag("admin")
  @useAuth(BearerAuth)
  namespace Admin {
    @route("/orders")
    namespace Orders {
      @summary("get order")
      @get op getOrder(...GetOrderRequest): GetOrderResponse | ErrorResponse;

      @summary("update order status")
      @patch op updateOrderStatus(...UpdateOrderStatusRequest): OkResponse | ErrorResponse;
    }

    @route("/rackets")
    namespace Rackets {
      @summary("create racket in shop")
      @post op createRacket(...CreateRacketRequest): CreateRacketResponse | ErrorResponse;

      @summary("update racket quantity")
      @patch op updateRacketQuantity(...UpdateRacketQuantityRequest): OkResponse | ErrorResponse;
    }

    @route("/users")
    namespace Users {
      @summary("get all users")
      @get op getUsers(): GetUsersResponse | ErrorResponse;

      @summary("get user")
      @get op getUser(...GetUserRequest): GetUserResponse | ErrorResponse;

      @summary("change user role")
      @patch op changeUserRole(...ChangeUserRole): OkResponse | ErrorResponse;
    }
  }

  @tag("authorized users")
  @useAuth(BearerAuth)
  namespace AuthorizedUsers {
    @route("/orders")
    namespace Orders {
      @summary("get orders")
      @get op getOrders(...GetOrdersRequest): GetOrdersResponse | ErrorResponse;
    }
  }

  @tag("all users")
  namespace AllUsers {
    @route("/feedbacks")
    namespace Feedbacks {
      @route("/rackets")
      namespace Rackets {
        @summary("get feedbacks for racket")
        @get op getRacketFeedbacks(...GetRacketFeedbacksRequest): GetRacketFeedbacksResponse | ErrorResponse;
      }
    }

    @route("/rackets")
    namespace Rackets {
      @summary("get all rackets")
      @get op getRackets(...GetRacketsRequest): GetRacketsResponse | ErrorResponse;

      @summary("get racket")
      @get op getRacket(...GetRacketRequest): GetRacketResponse | ErrorResponse;
    }
  }
}