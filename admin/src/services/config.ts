import { request } from '@@/exports';

export const getConfig = (): Promise<{ config: API.Config }> => {
  return request('/config', {});
};
