import request from "@/utils/request";
import { Auth } from "@/utils/auth";

const FileAPI = {
  /**
   * 上传文件
   *
   * @param formData
   */
  upload(formData: FormData) {
    return request<any, FileInfo>({
      url: "/api/v1/files",
      method: "post",
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  },

  /**
   * 上传文件
   */
  uploadFile(file: File) {
    const formData = new FormData();
    formData.append("file", file);
    return request<any, FileInfo>({
      url: "/api/v1/files",
      method: "post",
      data: formData,
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  },

  /**
   * 删除文件
   *
   * @param filePath 文件完整路径
   */
  delete(filePath?: string) {
    return request({
      url: "/api/v1/files",
      method: "delete",
      params: { filePath },
    });
  },

  /**
   * 下载文件
   * @param url
   * @param fileName
   */
  download(url: string, fileName?: string) {
    // 处理URL编码问题：确保文件名部分正确编码
    // 使用URL对象来正确构建编码后的URL，避免双重编码
    let processedUrl = url;
    
    if (url.includes("/api/v1/files/")) {
      try {
        const parts = url.split("/api/v1/files/");
        if (parts.length === 2) {
          const basePath = parts[0]; // 获取基础路径（可能是空字符串或完整URL）
          let fileNamePart = parts[1];
          
          // 如果文件名已经包含编码字符，先解码获取原始文件名
          if (fileNamePart.includes("%")) {
            try {
              fileNamePart = decodeURIComponent(fileNamePart);
            } catch (e) {
              // 解码失败，使用原始值
              console.warn("Failed to decode filename, using original:", fileNamePart);
            }
          }
          
          // 对文件名部分使用encodeURIComponent编码（这是标准做法）
          const encodedFileName = encodeURIComponent(fileNamePart);
          
          // 始终使用相对路径，这样可以通过Vite代理转发到后端
          // 如果原始URL是绝对URL，提取路径部分；否则直接使用
          if (url.startsWith("http://") || url.startsWith("https://")) {
            // 从完整URL中提取路径部分
            try {
              const urlObj = new URL(url);
              processedUrl = `${urlObj.pathname.split("/api/v1/files/")[0]}/api/v1/files/${encodedFileName}`;
            } catch (e) {
              // URL解析失败，使用相对路径
              processedUrl = `/api/v1/files/${encodedFileName}`;
            }
          } else if (url.startsWith("/")) {
            // 相对路径，直接构建
            processedUrl = `/api/v1/files/${encodedFileName}`;
          } else {
            // 其他情况，使用相对路径
            processedUrl = `/api/v1/files/${encodedFileName}`;
          }
        }
      } catch (error) {
        console.error("URL processing error:", error);
        processedUrl = url;
      }
    }
    
    // 使用fetch API来避免axios的URL编码问题
    // fetch对已编码的URL处理更可靠
    const accessToken = Auth.getAccessToken();
    const headers: HeadersInit = {};
    if (accessToken) {
      headers["Authorization"] = `Bearer ${accessToken}`;
    }
    
    return fetch(processedUrl, {
      method: "GET",
      headers,
    })
      .then(async (response) => {
        if (!response.ok) {
          const errorText = await response.text();
          console.error("Download failed:", response.status, errorText);
          throw new Error(`下载失败: ${response.status} ${errorText || 'No error message'}`);
        }
        return response.blob();
      })
      .then((blob) => {
        const downloadUrl = window.URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = downloadUrl;
        a.download = fileName || "下载文件";
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
        window.URL.revokeObjectURL(downloadUrl);
      })
      .catch((error) => {
        console.error("Download error:", error);
        throw error;
      });
  },
};

export default FileAPI;

/**
 * 文件API类型声明
 */
export interface FileInfo {
  /** 文件名 */
  name: string;
  /** 文件路径 */
  url: string;
  /** 文件ID（如果是已关联的文件） */
  id?: number;
}
