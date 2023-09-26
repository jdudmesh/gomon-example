// vite.config.js
import { resolve } from "path"
/** @type {import('vite').UserConfig} */
export default {
  root: resolve(__dirname),
  base: '/assets/dist',
  build: {
    outDir: '../static/dist',
    rollupOptions: {
      input: 'src/index.ts',
      output: {
        entryFileNames: 'index.js',
        assetFileNames: '[name].[ext]',
      }
    },
  },
  server: {
    watch: {
      ignored: ["**/views/**"]
    }
  },
  plugins: [],
}