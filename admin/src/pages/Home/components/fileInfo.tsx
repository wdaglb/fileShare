import { Button, Card, Col, Descriptions, Modal, QRCode, Row, Space, Typography } from 'antd';
import { convertFileSize } from '@/utils/format';
import dayjs from 'dayjs';
import React, { useMemo, useState } from 'react';
import styles from '../style.less';
import { useRequest } from 'ahooks';
import { downloadFile } from '@/services/file';
import { useModel, useNavigate } from '@@/exports';
import useUrlState from '@ahooksjs/use-url-state';

interface Props {
  file: API.File;
}

const FileInfo = (props: Props) => {
  const { initialState } = useModel('@@initialState');
  const [, setParams] = useUrlState({});
  const [qrcode, setQrcode] = useState(false);
  const navigate = useNavigate();

  const file = useMemo(() => {
    return props.file;
  }, [props.file]);

  const { run: download } = useRequest(downloadFile, {
    manual: true,
    onSuccess: (res) => {
      if (!file) {
        return;
      }
      const reader = new FileReader();
      reader.onload = (evt) => {
        const a = document.createElement('a');
        a.setAttribute('href', evt.target?.result as any);
        a.setAttribute('download', file.FileName);
        a.click();
      };
      reader.readAsDataURL(res);
    },
  });

  // 内网地址
  const partUrl = useMemo(() => {
    return `${initialState?.part_domain}/?code=${file.Password}`;
  }, [initialState, file]);

  // 公网地址
  const publicUrl = useMemo(() => {
    return `${initialState?.public_domain}/?code=${file.Password}`;
  }, [initialState, file]);

  return (
    <>
      <Descriptions className={styles.fileBody}>
        <Descriptions.Item span={3} label={'文件名'}>
          {file?.FileName}
        </Descriptions.Item>

        <Descriptions.Item span={3} label={'文件大小'}>
          {convertFileSize(file?.FileSize ?? 0)}
        </Descriptions.Item>

        <Descriptions.Item span={3} label={'下载次数'}>
          {file?.DownloadNumber ?? 0}
        </Descriptions.Item>

        <Descriptions.Item span={3} label={'上传时间'}>
          {dayjs(file?.CreatedAt).format('YYYY-MM-DD HH:mm:ss')}
        </Descriptions.Item>

      </Descriptions>

      <Space>
        <Button type={'primary'} onClick={() => download(file?.Password)}>直接下载</Button>

        <Button onClick={() => navigate('/')}>我要上传</Button>

        {!qrcode && (
          <Button onClick={() => setQrcode(true)}>手机下载</Button>
        )}
      </Space>

      {qrcode && (
        <div className={styles.qrcode}>
          <QRCode value={publicUrl} size={200} />
        </div>
      )}
    </>
  );
};

export default FileInfo;
