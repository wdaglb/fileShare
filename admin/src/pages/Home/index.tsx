import React, { useEffect, useMemo, useState } from 'react';
import { useRequest, useSearchParams } from '@@/exports';
import { getFileInfo } from '@/services/file';
import {
  Card,
} from 'antd';
import styles from './style.less';
import FileInfo from '@/pages/Home/components/fileInfo';
import FileUpload from '@/pages/Home/components/fileUpload';
import useUrlState from '@ahooksjs/use-url-state';

const Btn = (props: { text: string; icon: React.ReactElement; onClick?: () => void }) => {
  return (
    <div
      onClick={props.onClick}
      className={styles.btn}>
      {React.cloneElement(props.icon, {
        className: styles.icon,
        style: {
          fontSize: 24,
        },
      })}
      <div className={styles.text}>{props.text}</div>
    </div>
  );
};

const HomePage: React.FC = () => {
  const [params] = useUrlState({ code: '' });
  const [error, setError] = useState('');
  const [file, setFile] = useState<API.File | null>(null);
  const { run: getFileInfoReq } = useRequest<API.File>(async (code) => {
    try {
      const res = await getFileInfo(code);
      setFile(res.file);
    } catch (e: any) {
      setError(e.response.data.error);
    }
  }, {
    manual: true,
  });

  useEffect(() => {
    if (!params.code) {
      setFile(null);
      setError('');
      return;
    }
    if (params.code && (!file || file.Password !== params.code)) {
      getFileInfoReq(params.code);
    }
  }, [params.code]);

  return (
    <div>
      <Card title={'文件快运'} className={styles.container}>
        {file ? <FileInfo file={file} /> : <FileUpload />}
      </Card>
    </div>
  );
};

export default HomePage;
