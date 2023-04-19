import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Oluwoye",
  description: "A rule-based network and log monitoring system",
  base: "/oluwoye/",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Get Started', link: '/installing' }
    ],

    sidebar: [
      {
        text: 'Get Started',
        items: [
          { text: 'Installing', link: '/installing' },
          { text: 'Running', link: '/running' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/roqueando/oluwoye' }
    ]
  }
})
