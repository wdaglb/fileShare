declare namespace API {
  /**
   * File
   */
  export interface File {
    /**
     * 创建时间
     */
    CreatedAt: string;
    DeletedAt: null;
    DirId: number;
    /**
     * 过期类型
     */
    ExpireType: number;
    /**
     * 过期值
     */
    ExpireValue: number;
    /**
     * 文件名
     */
    FileName: string;
    /**
     * 文件大小
     */
    FileSize: number;
    /**
     * 文件id
     */
    ID: number;
    /**
     * 媒体类型
     */
    MimeType: string;
    /**
     * 访问密码
     */
    Password: string;
    /**
     * 保存路径
     */
    SavePath: string;
    UpdatedAt: string;
    UserId: string;
    /**
     * 下载次数
     */
    DownloadNumber: number;

    [property: string]: any;
  }

}
