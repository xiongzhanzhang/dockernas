import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import { VitePWA } from 'vite-plugin-pwa'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    VitePWA({
      // devOptions: {
      //   enabled: true
      //   /* other options */
      // },
      manifest: {
        name: 'DockerNAS',
        short_name: 'DockerNAS',
        icons: [
          // {
          //   src: 'pwa-192x192.png',
          //   sizes: '192x192',
          //   type: 'image/png',
          // },
          {
            src: 'pwa-512x512.png',
            sizes: '591x591',
            type: 'image/png',
          },
        ],
      },
      workbox: {
        // globPatterns: [],
        runtimeCaching: [
          {
            urlPattern: /\.(?:png|jpg|jpeg|svg|ico)$/,
            handler: 'NetworkFirst'
          },
          {
            urlPattern:  /.*\.js.*/,
            handler: 'NetworkFirst'
          },
          {
            urlPattern: /.*\.css.*/,
            handler: 'NetworkFirst'
          },
          {
            urlPattern: /.*\/index\/.*/,
            handler: 'NetworkFirst'
          },
          {
            urlPattern: /.*\/api\/(?:instance|app)$/,
            handler: 'NetworkFirst'
          },
        ]
      }
    })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  base: "/",
  server: {
    port: 8081,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        ws: true,
        changeOrigin: true
      },
      '/apps': 'http://localhost:8080',
      '/extra/apps': 'http://localhost:8080',
      '/terminal': {
        target: 'http://localhost:8080',
        ws: true,
        changeOrigin: true
      }
    }
  },
  define: {
    // "global": {},
  },
})
