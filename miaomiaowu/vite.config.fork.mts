import path from 'path'
import { fileURLToPath } from 'url'
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tailwindcss from '@tailwindcss/vite'
import { tanstackRouter } from '@tanstack/router-plugin/vite'

const __dirname = path.dirname(fileURLToPath(import.meta.url))

/** Fork 构建：子路径 /mmw/，与 extensions/basepath 默认 BASE_PATH 一致 */
export default defineConfig({
  base: '/mmw/',
  plugins: [
    tanstackRouter({
      target: 'react',
      autoCodeSplitting: true,
    }),
    react(),
    tailwindcss(),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  build: {
    outDir: path.resolve(__dirname, '../internal/web/dist'),
    emptyOutDir: true,
    sourcemap: true,
  },
  css: {
    devSourcemap: true,
  },
  server: {
    proxy: {
      '/mmw/api': {
        target: process.env.VITE_API_URL || 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (p) => p.replace(/^\/mmw/, ''),
      },
      '/mmw/t': {
        target: process.env.VITE_API_URL || 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (p) => p.replace(/^\/mmw/, ''),
      },
    },
  },
})
