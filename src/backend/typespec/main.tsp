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
model LoginResponse {
  access_token: string,
};

/**
  API Definition
**/
@route("/api")
namespace Api {

  @tag("auth")
  @route("/auth")
  namespace Auth {
    @summary("login user")
    @doc("Login user")
    @post op login(...LoginRequest): LoginResponse | ErrorResponse;
  }
}