import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
import { resolve } from 'path'

const IP = '' // 代理ip
export default defineConfig({
  server:{
    port: 8080,
    proxy: {
      '/api/': {
        headers: { Connection:'keep-alive' },
        target: IP,
        changeOrigin: true,
        rewrite: (path) => {
          return path.replace('/api/', '/')
        }
      }
    }
  },
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, './src')
    },
    //extensions: [".ts", ".js", ".vue", ".json", ".mjs"],
    extensions: [".mjs", ".js", ".ts", ".jsx", ".tsx", ".json", ".vue"]
  }
})

