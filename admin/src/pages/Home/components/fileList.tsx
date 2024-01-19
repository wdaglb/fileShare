import styles from '@/pages/Home/style.less';
import { FileImageFilled, FileWordFilled, FolderFilled } from '@ant-design/icons';
import React from 'react';
import { useRequest } from 'ahooks';
import { getFileList } from '@/services/file';

const FileList = () => {
  const { loading, data } = useRequest(() => {
    return getFileList();
  }, {});
  return (
    <div className={styles.items}>
      <div className={styles.item}>
        <FolderFilled className={styles.icon} />
        <div className={styles.text}>默认目录</div>
      </div>
      <div className={styles.item}>
        <FileImageFilled className={styles.icon} />
        <div className={styles.text}>默认目录</div>
      </div>

      {data?.list.map(item => (
        <div className={styles.item} key={item.ID}>
          <FileWordFilled className={styles.icon} />
          <div className={styles.text}>{item.FileName}</div>
        </div>
      ))}
    </div>
  );
};

export default FileList;
