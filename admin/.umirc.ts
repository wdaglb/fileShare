import { defineConfig } from '@umijs/max';

export default defineConfig({
  antd: {},
  access: {},
  model: {},
  initialState: {},
  request: {},
  layout: {
    title: '@umijs/max',
    locale: false,
  },
  routes: [
    {
      name: '文件快运',
      path: '/',
      component: './Home',
    },
  ],

  npmClient: 'pnpm',

  proxy: {
    '/api': {
      'target': 'http://127.0.0.1:5555/',
      'pathRewrite': { '^/api': '' },
    },
  },
});
