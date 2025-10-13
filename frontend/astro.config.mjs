import { defineConfig } from 'astro/config';
import vue from '@astrojs/vue';
import node from '@astrojs/node'; // Add this import

export default defineConfig({
  integrations: [vue()],
  output: 'server',
  adapter: node({
    mode: 'standalone'
  }),
  server: {
    host: true  // This tells Astro to listen on 0.0.0.0
  }
});