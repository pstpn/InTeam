import "@typespec/http";
import "@typespec/rest";
import "@typespec/versioning";

import "./error.tsp";
import "./models.tsp";

@service({
  title: "RacketShop",
})
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
    surname: string,
  }
};
model LoginRequest {
  @body
  loginBody: {
    email: string,
    password: string,
  }
};

/**
  API Responses
**/
model RegisterResponse extends ResponseWithAccessToken {};
model LoginResponse extends ResponseWithAccessToken {};

/**
  API Definition
**/
@route("/api")
namespace Api {

  @tag("auth")
  @route("/auth")
  namespace Auth {
    @summary("register user")
    @doc("Register user")
    @route("/register")
    @post op register(...RegisterRequest): RegisterResponse | ErrorResponse;

    @summary("login user")
    @doc("Login user")
    @route("/login")
    @post op login(...LoginRequest): LoginResponse | ErrorResponse;
  }
}