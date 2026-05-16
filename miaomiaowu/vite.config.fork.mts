import path from 'path'
import { fileURLToPath } from 'url'
import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tailwindcss from '@tailwindcss/vite'
import { tanstackRouter } from '@tanstack/router-plugin/vite'

const __dirname = path.dirname(fileURLToPath(import.meta.url))
const repoRoot = path.resolve(__dirname, '..')

/** Fork 构建：子路径由 .env 中 VITE_BASE_PATH 控制，默认 /mmw/ */
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, repoRoot, '')
  const base = env.VITE_BASE_PATH || '/mmw/'
  const prefix = base.replace(/\/$/, '') || ''

  const apiProxy = prefix
    ? {
        [`${prefix}/api`]: {
          target: env.VITE_API_URL || 'http://localhost:8080',
          changeOrigin: true,
          rewrite: (p: string) => p.replace(new RegExp(`^${prefix}`), ''),
        },
        [`${prefix}/t`]: {
          target: env.VITE_API_URL || 'http://localhost:8080',
          changeOrigin: true,
          rewrite: (p: string) => p.replace(new RegExp(`^${prefix}`), ''),
        },
      }
    : {
        '/api': {
          target: env.VITE_API_URL || 'http://localhost:8080',
          changeOrigin: true,
        },
        '/t/': {
          target: env.VITE_API_URL || 'http://localhost:8080',
          changeOrigin: true,
        },
      }

  return {
    base,
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
      proxy: apiProxy,
    },
  }
})
