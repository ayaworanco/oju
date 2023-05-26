import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Oju",
  description: "A streaming tracing and logger watcher for your applications",
  base: "/oju/",
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
          { text: 'Using Qolu', link: '/using-qoju' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/roqueando/oju' }
    ]
  }
})
