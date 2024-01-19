import { Button, Form, Input, InputNumber, Select, Space, Upload } from 'antd';
import { InboxOutlined, UploadOutlined } from '@ant-design/icons';

import styles from '../style.less';
import React, { ChangeEvent, useRef, useState } from 'react';
import { useModel, useRequest, useSearchParams } from '@@/exports';
import { uploadFile } from '@/services/file';
import useUrlState from '@ahooksjs/use-url-state';

const FileUpload = () => {
  const { initialState } = useModel('@@initialState');
  const [, setParams] = useUrlState();
  const [saveType, setSaveType] = useState(0);
  const [saveValue, setSaveValue] = useState(1);

  return (
    <>
      <div className={styles.item}>
        储存方式：
        <Space>
          <Select
            style={{
              width: 80,
            }}
            value={saveType}
            onChange={value => setSaveType(value)}
            options={[
              {
                label: '时间',
                value: 0,
              },
              {
                label: '次数',
                value: 1,
              },
              {
                label: '永久',
                value: 2,
              },
            ]}
          />

          {saveType < 2 && (
            <InputNumber
              style={{ width: 120 }}
              addonAfter={<>{['天', '次'][saveType]}</>}
              min={1}
              value={saveValue}
              onChange={evt => setSaveValue(Number(evt))}
            />
          )}
        </Space>
      </div>

      <div>
        <Upload.Dragger
          name={'file'}
          multiple={false}
          headers={{
            ['X-UserId']: initialState?.userId ?? '',
          }}
          data={{
            ExpireType: saveType,
            ExpireValue: saveValue,
          }}
          customRequest={async ops => {
            const res = await uploadFile(ops.file as any, [
              { name: 'ExpireType', value: saveType },
              { name: 'ExpireValue', value: saveValue },
            ], {
              onProgress: (args) => {
                ops.onProgress?.({
                  percent: Math.ceil((args.loaded / args.total) * 100),
                });
              },
            });
            ops.onSuccess?.(res);
            setParams({
              code: res.code,
            });
          }}
        >
          <div
            className={styles.drop}>
            <p>
              <InboxOutlined className={styles.icon} />
            </p>
            <p>
              点击或拖拽文件到此上传
            </p>
          </div>
        </Upload.Dragger>
      </div>
    </>
  );
};

export default FileUpload;
