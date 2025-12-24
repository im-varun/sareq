// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

import {themes as prismThemes} from "prism-react-renderer";

const config = {
  title: "SAReq",
  tagline: "A modern, open-source HTTP client for the command line",
  favicon: "img/logo.png",

  future: {
    v4: true,
  },

  url: "https://sareq.dev",
  baseUrl: "/",

  onBrokenLinks: "throw",

  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      ({
        docs: {
          sidebarPath: "./sidebars.js",
        },
        theme: {
          customCss: "./src/css/custom.css",
        },
      }),
    ],
  ],

  themeConfig:
    ({
      colorMode: {
        defaultMode: "dark",
        disableSwitch: true,
        respectPrefersColorScheme: false,
      },
      navbar: {
        title: "/>>/",
        items: [
          {
            type: "docSidebar",
            sidebarId: "docsSidebar",
            position: "left",
            label: "Docs",
          },
          {
            href: "https://github.com/im-varun/sareq",
            label: "GitHub",
            position: "right",
          },
        ],
      },
      footer: {
        style: "dark",
        copyright: `Copyright © ${new Date().getFullYear()} Varun Mulchandani.`,
      },
      prism: {
        theme: prismThemes.github,
        darkTheme: prismThemes.dracula,
        additionalLanguages: ["shell-session", "bash", "powershell"],
      },
    }),
};

export default config;
