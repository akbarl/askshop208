package lazada_constants

const COUNTRY = "vn"
const BASE_REST_URL = "https://api.lazada." + COUNTRY
const AUTH_OAUTH_URL = "https://auth.lazada.com/oauth/authorize"
const PRODUCTS_REST_URL = BASE_REST_URL + "/products/get"
const BASE_AUTH_REST_URL = "https://auth.lazada.com/rest"
const GENERATE_ACCESS_TOKEN_REST_URL = "/auth/token/create"
const DEFAULT_SIGN_METHOD_NAME = "sha256"

// Request Params
const CODE = "code"
const SIGN = "sign"
