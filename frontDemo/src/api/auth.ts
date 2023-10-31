import { $post } from '@/utils/request'

export const ajaxLog = {
  login: (data: any) => $post('/login', data),
  logout: () => $post('/logout')
}

export const logout = () => {
  ajaxLog.logout().finally(() => {
    sessionStorage.clear()
    localStorage.clear()
    location.reload()
  })
}
