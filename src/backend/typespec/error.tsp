enum Errors {
  BadRequest,
  ValidationError,
  NotFoundError,
  InternalError,
  NotModifiedError,
  ConflictError,
  Unauthorized
}

@error
model ErrorResponse {
  error: Errors;
}