enum Types {
  BadRequest,
  ValidationError,
  NotFoundError,
  InternalError,
  NotModifiedError,
  ConflictError,
}

@error
model ErrorResponse {
  type: Types;
}