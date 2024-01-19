// 运行时配置

// 全局初始化数据配置，用于 Layout 用户信息和权限初始化
// 更多信息见文档：https://umijs.org/docs/api/runtime-config#getinitialstate
import { RunTimeLayoutConfig } from '@@/plugin-layout/types';
import { RequestConfig } from '@@/plugin-request/request';
import FingerprintJS from '@fingerprintjs/fingerprintjs';
import NiceModal from '@ebay/nice-modal-react';
import { message } from 'antd';
import { getConfig } from '@/services/config';

export async function getInitialState(): Promise<{ userId: string; part_domain: string; public_domain: string }> {
  let userId = localStorage.getItem('userId');
  if (!userId) {
    const res = await (await FingerprintJS.load()).get();
    userId = 'b:' + res.visitorId;
    localStorage.setItem('userId', userId);
  }
  const res = await getConfig();

  return { userId, ...res.config };
}

export const layout: RunTimeLayoutConfig = () => {
  return {
    logo: false,
    title: false,
    // logo: 'https://img.alicdn.com/tfs/TB1YHEpwUT1gK0jSZFhXXaAtVXa-28-27.svg',
    menuRender: false,
    headerRender: false,
    menuHeaderRender: false,
    childrenRender: (children) => (
      <NiceModal.Provider>
        {children}
      </NiceModal.Provider>
    ),
  };
};

export const request: RequestConfig = {
  baseURL: '/api',
  requestInterceptors: [
    (config: any) => {
      config.headers['X-UserId'] = localStorage.getItem('userId');
      return config;
    },
  ],
  responseInterceptors: [
    [
      (resp) => {
        return resp;
      },
      (err: any) => {
        if (!err.config?.skipErrorHandler && err.response.data && err.response.data.error) {
          message.error(err.response.data.error);
          return Promise.reject(err.response.data.error);
        }
        return Promise.reject(err);
      },
    ],
  ],
};
