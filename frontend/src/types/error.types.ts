type ErrorData = {
  [key: string]: any;
};

export class ApiError extends Error {
  constructor(
    public message: string,
    public data?: ErrorData
  ) {
    super(message);
    this.name = 'ApiError';
  }
}
