# 搭建vite项目 vue3+ts
## 包含：element-plus; axios; vue-router; pinia; sass

## 搭建过程
1. npm create vite@latest
2. npm i
3. npm run dev
4. 使用@引入路径，npm i @types/node -D
   在tsconfig.json中添加：
   ```
   "paths": {
      "@/*": ["./src/*"]
   }
   ```
   在vite.config.ts中添加并修改defineConfig：
   ```
   import { resolve } from 'path'
   export default defineConfig({
   plugins: [vue()],
   resolve: {
      alias: {
         '@': resolve(__dirname, './src')
      },
      //extensions: [".ts", ".js", ".vue", ".json", ".mjs"],
      extensions: [".mjs", ".js", ".ts", ".jsx", ".tsx", ".json", ".vue"]
   }
   })
   ```
5. 需要其他库，install即可
