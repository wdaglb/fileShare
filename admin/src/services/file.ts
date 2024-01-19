import { request } from '@umijs/max';

// 获取文件列表
export const getFileList = (): Promise<{ list: API.File[] }> => {
  return request('/file/list', {});
};

// 获取文件信息
export const getFileInfo = (code: string) => {
  return request('/file/get', {
    params: {
      code,
    },
  });
};

type UploadFileCallback = (args: any) => void;

interface UploadFileProps {
  onProgress: UploadFileCallback;
}

// 上传文件
export const uploadFile = (file: File, params: {
  name: string;
  value: any
}[], options: UploadFileProps) => {
  const fd = new FormData();
  fd.append('file', file);
  params.forEach(param => {
    fd.append(param.name, param.value);
  });
  return request('/upload', {
    method: 'post',
    data: fd,
    onUploadProgress: options?.onProgress,
  });
};

// 删除文件
export const deleteFile = (code: string) => {
  request('/file/delete', {
    method: 'post',
    data: {
      code,
    },
  });
};

// 下载文件
export const downloadFile = (code: string) => {
  return request('/file/download', {
    params: {
      code,
    },
    responseType: 'blob',
    skipErrorHandler: true,
  });
};
