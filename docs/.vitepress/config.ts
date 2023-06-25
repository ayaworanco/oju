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
          //{ text: 'Using Qolu', link: '/using-qoju' }
        ]
      },
			{
				text: "Architecture",
				items: [
					{text: 'AWO Protocol', link: 'awo-protocol'},
          {text: 'Distributed Tracings', items: [
            {text: 'Tracer', link: 'distributed-tracing/tracer'},
            {text: 'Service Discovery', link: 'distributed-tracing/service-discovery'},
          ]},
					{text: 'Logs', items: [
            {text: 'Parser', link: 'logs/parser'},
					]},
					{text: 'Contributing', items: [
            {text: 'How to', link: 'contributing/how-to'},
            {text: 'Example', link: 'contributing/example'},
					]}
				]
			}
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/roqueando/oju' }
    ]
  }
})
