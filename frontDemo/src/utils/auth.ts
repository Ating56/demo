export const TOKENKEY = 'my-blog-session'

export function getToken() {
  return sessionStorage.getItem(TOKENKEY)
}

export function setToken(token: string) {
  return sessionStorage.setItem(TOKENKEY, token)
}
