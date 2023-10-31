import { defineStore } from 'pinia'

function changeThemeMode(theme: string, val: any) {
  const body = document.documentElement as HTMLElement;
  if (theme === 'day') {
    body.setAttribute("data-themeMode", "")
    changeTheme(val)
  } else {
    body.setAttribute("data-themeMode", "dark")
    changeTheme(val)
  }
}
function changeTheme(val: any) {
  document.documentElement.style.backgroundColor = val
}
export const useAppStore = defineStore('app', {
  state() {
    const currTheme = localStorage.getItem('theme') || '#fff'
    changeTheme(currTheme)
    return {
      themeMode: 'day',
      theme: currTheme
    };
  },
  actions: {
    changeTheme(color: string) {
      localStorage.setItem('theme', color)
      changeTheme(color)
    },
    toggleThemeMode(val:'day'|'dark', color: string | undefined) {
      // 切换主题
      changeThemeMode(val, color)
    }
  }
})
