import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Frontend calls /auth/status; proxy rewrites it to GET / on the backend
      // so the browser session cookie is shared (same-origin) with no CORS needed.
      '/auth/status': {
        target: 'http://localhost:9000',
        changeOrigin: true,
        rewrite: () => '/',
      },
      '/api': {
        target: 'http://localhost:9000',
        changeOrigin: true,
      },
    },
  },
})
