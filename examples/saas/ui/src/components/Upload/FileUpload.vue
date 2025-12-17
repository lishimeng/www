<!-- 文件上传组件 -->
<template>
  <div>
    <el-upload
      v-model:file-list="fileList"
      :style="props.style"
      :before-upload="handleBeforeUpload"
      :http-request="handleUpload"
      :on-progress="handleProgress"
      :on-success="handleSuccess"
      :on-error="handleError"
      :accept="props.accept"
      :limit="props.limit"
      multiple
    >
      <!-- 上传文件按钮 -->
      <el-button type="primary" :disabled="fileList.length >= props.limit">
        {{ props.uploadBtnText }}
      </el-button>

      <!-- 文件列表 -->
      <template #file="{ file }">
        <div class="el-upload-list__item-info">
          <div class="file-item-container">
            <a class="el-upload-list__item-name" @click="handleView(file)">
              <el-icon><Document /></el-icon>
              <el-tooltip :content="file.name" placement="top" :disabled="!file.name || file.name.length <= 30">
                <span class="el-upload-list__item-file-name">{{ truncateFileName(file.name) }}</span>
              </el-tooltip>
            </a>
            <div class="file-actions">
              <!-- 预览按钮（仅图片和PDF显示） -->
              <span 
                v-if="isPreviewable(file)" 
                class="file-action-btn" 
                title="预览"
                @click.stop="handlePreview(file)"
              >
                <el-icon><View /></el-icon>
              </span>
              <!-- 下载按钮 -->
              <span 
                class="file-action-btn" 
                title="下载"
                @click.stop="handleDownload(file)"
              >
                <el-icon><Download /></el-icon>
              </span>
              <!-- 删除按钮 -->
              <span 
                class="file-action-btn" 
                title="删除"
                @click.stop="handleRemove(file.url!)"
              >
                <el-icon><Close /></el-icon>
              </span>
            </div>
          </div>
        </div>
      </template>
    </el-upload>

    <!-- 图片预览 -->
    <ElImageViewer
      v-if="previewVisible && previewImageUrl"
      :url-list="[previewImageUrl]"
      @close="handlePreviewClose"
    />

    <el-progress
      :style="{
        display: showProgress ? 'inline-flex' : 'none',
        width: '100%',
      }"
      :percentage="progressPercent"
    />
  </div>
</template>
<script lang="ts" setup>
import { ref, watch } from 'vue';
import type { PropType } from 'vue';
import {
  UploadRawFile,
  UploadUserFile,
  UploadFile,
  UploadFiles,
  UploadProgressEvent,
  UploadRequestOptions,
} from "element-plus";
import { Document, Close, View, Download } from '@element-plus/icons-vue';
import { ElImageViewer } from 'element-plus';

import FileAPI, { FileInfo } from "@/api/file.api";
import { ElMessage } from "element-plus";

const props = defineProps({
  /**
   * 请求携带的额外参数
   */
  data: {
    type: Object,
    default: () => {
      return {};
    },
  },
  /**
   * 上传文件的参数名
   */
  name: {
    type: String,
    default: "file",
  },
  /**
   * 文件上传数量限制
   */
  limit: {
    type: Number,
    default: 10,
  },
  /**
   * 单个文件上传大小限制(单位MB)
   */
  maxFileSize: {
    type: Number,
    default: 10,
  },
  /**
   * 上传文件类型
   */
  accept: {
    type: String,
    default: "*",
  },
  /**
   * 上传按钮文本
   */
  uploadBtnText: {
    type: String,
    default: "上传文件",
  },

  /**
   * 样式
   */
  style: {
    type: Object,
    default: () => {
      return {
        width: "300px",
      };
    },
  },
});

const modelValue = defineModel("modelValue", {
  type: [Array] as PropType<FileInfo[]>,
  required: false,
  default: () => [],
});

const fileList = ref([] as UploadFile[]);

const showProgress = ref(false);
const progressPercent = ref(0);

// 预览相关
const previewVisible = ref(false);
const previewImageUrl = ref('');

// 监听 modelValue 转换用于显示的 fileList
watch(
  modelValue,
  (value) => {
    // 确保 value 是数组，避免 null/undefined 错误
    if (!value || !Array.isArray(value)) {
      fileList.value = [];
      return;
    }
    fileList.value = value.map((item) => {
      const name = item.name ? item.name : item.url?.substring(item.url.lastIndexOf("/") + 1);
      return {
        name,
        url: item.url,
        status: "success",
        uid: getUid(),
      } as UploadFile;
    });
  },
  {
    immediate: true,
  }
);

/**
 * 上传前校验
 */
function handleBeforeUpload(file: UploadRawFile) {
  // 检查文件类型和大小
  const isImage = file.type.startsWith('image/') || 
    ['.jpg', '.jpeg', '.png'].some(ext => file.name.toLowerCase().endsWith(ext))
  const isFile = ['.pdf', '.docx'].some(ext => file.name.toLowerCase().endsWith(ext))
  
  // 图片限制5MB，其他文件限制20MB
  const maxSize = isImage ? 5 : (isFile ? 20 : props.maxFileSize)
  
  if (file.size > maxSize * 1024 * 1024) {
    ElMessage.warning(`上传${isImage ? '图片' : '文件'}不能大于${maxSize}M`);
    return false;
  }
  return true;
}

/*
 * 上传文件
 */
function handleUpload(options: UploadRequestOptions) {
  return new Promise((resolve, reject) => {
    const file = options.file;

    const formData = new FormData();
    formData.append(props.name, file);

    // 处理附加参数
    Object.keys(props.data).forEach((key) => {
      formData.append(key, props.data[key]);
    });

    FileAPI.upload(formData)
      .then((data) => {
        resolve(data);
      })
      .catch((error) => {
        reject(error);
      });
  });
}

/**
 * 上传进度
 *
 * @param event
 */
const handleProgress = (event: UploadProgressEvent) => {
  progressPercent.value = event.percent;
};

/**
 * 上传成功
 */
const handleSuccess = (_response: any, uploadFile: UploadFile, files: UploadFiles) => {
  ElMessage.success("上传成功");
  //只有当状态为success或者fail，代表文件上传全部完成了，失败也算完成
  if (
    files.every((file: UploadFile) => {
      return file.status === "success" || file.status === "fail";
    })
  ) {
    const fileInfos = [] as FileInfo[];
    files.map((file: UploadFile) => {
      if (file.status === "success") {
        //只取携带response的才是刚上传的
        const res = file.response as FileInfo;
        if (res) {
          fileInfos.push({ name: res.name, url: res.url } as FileInfo);
        }
      } else {
        //失败上传 从fileList删掉，不展示
        fileList.value.splice(
          fileList.value.findIndex((e) => e.uid === file.uid),
          1
        );
      }
    });
    if (fileInfos.length > 0) {
      const currentFiles = Array.isArray(modelValue.value) ? modelValue.value : [];
      modelValue.value = [...currentFiles, ...fileInfos];
    }
  }
};

/**
 * 上传失败
 */
const handleError = (_error: any) => {
  console.error(_error);
  ElMessage.error("上传失败");
};

/**
 * 删除文件
 */
async function handleRemove(fileUrl: string) {
  // 查找要删除的文件信息
  const fileToRemove = modelValue.value?.find((file) => file.url === fileUrl);
  
  // 如果文件有ID，说明是已关联的文件，需要删除数据库关联
  if (fileToRemove?.id) {
    try {
      // 先删除数据库关联
      await fetch(`/api/quality/inspection-item/files/${fileToRemove.id}`, {
        method: 'DELETE'
      });
      // 无论数据库删除是否成功，都尝试删除文件系统中的文件
      // 如果文件不存在，忽略错误
      try {
        await FileAPI.delete(fileUrl);
      } catch (err) {
        // 文件可能已经不存在，这是正常的，不抛出错误
        console.warn('删除文件系统中的文件失败（可能文件已不存在）:', err);
      }
      // 从列表中移除
      if (!modelValue.value || !Array.isArray(modelValue.value)) {
        modelValue.value = [];
        return;
      }
      modelValue.value = modelValue.value.filter((file) => file.url !== fileUrl);
      ElMessage.success('删除成功');
    } catch (error) {
      console.error('删除文件失败:', error);
      ElMessage.error('删除文件关联失败');
    }
  } else {
    // 如果是新上传但未关联的文件，直接删除文件
    try {
      await FileAPI.delete(fileUrl);
      if (!modelValue.value || !Array.isArray(modelValue.value)) {
        modelValue.value = [];
        return;
      }
      modelValue.value = modelValue.value.filter((file) => file.url !== fileUrl);
      ElMessage.success('删除成功');
    } catch (error) {
      // 文件可能已经不存在，仍然从列表中移除
      if (!modelValue.value || !Array.isArray(modelValue.value)) {
        modelValue.value = [];
        return;
      }
      modelValue.value = modelValue.value.filter((file) => file.url !== fileUrl);
      // 不显示错误，因为文件可能已经不存在了
      console.warn('删除文件失败（可能文件已不存在）:', error);
    }
  }
}

/**
 * 判断文件是否可预览（图片和PDF）
 */
function isPreviewable(file: UploadUserFile): boolean {
  if (!file.name && !file.url) return false;
  const fileName = (file.name || file.url || '').toLowerCase();
  return fileName.match(/\.(jpg|jpeg|png|gif|bmp|webp|pdf)$/i) !== null;
}

/**
 * 预览文件
 */
function handlePreview(file: UploadUserFile) {
  if (!file.url) return;
  
  const fileName = (file.name || file.url || '').toLowerCase();
  
  // 如果是PDF，在新窗口打开
  if (fileName.endsWith('.pdf')) {
    window.open(file.url, '_blank');
    return;
  }
  
  // 如果是图片，使用图片预览器
  if (fileName.match(/\.(jpg|jpeg|png|gif|bmp|webp)$/i)) {
    previewImageUrl.value = file.url;
    previewVisible.value = true;
  }
}

/**
 * 查看文件（预览或下载）
 */
function handleView(file: UploadUserFile) {
  if (isPreviewable(file)) {
    handlePreview(file);
  } else {
    handleDownload(file);
  }
}

/**
 * 关闭预览
 */
function handlePreviewClose() {
  previewVisible.value = false;
  previewImageUrl.value = '';
}

/**
 * 下载文件
 */
function handleDownload(file: UploadUserFile) {
  const { url, name } = file;
  if (url) {
    FileAPI.download(url, name);
  }
}

/** 获取一个不重复的id */
function getUid(): number {
  // 时间戳左移13位（相当于乘以8192） + 4位随机数
  return (Date.now() << 13) | Math.floor(Math.random() * 8192);
}

/**
 * 截断文件名，如果过长则显示省略号
 */
function truncateFileName(fileName: string | undefined): string {
  if (!fileName) return '';
  const maxLength = 30; // 最大显示长度
  if (fileName.length <= maxLength) {
    return fileName;
  }
  // 保留文件扩展名，截断文件名主体部分
  const lastDotIndex = fileName.lastIndexOf('.');
  if (lastDotIndex > 0) {
    const name = fileName.substring(0, lastDotIndex);
    const ext = fileName.substring(lastDotIndex);
    const truncatedName = name.substring(0, maxLength - ext.length - 3) + '...';
    return truncatedName + ext;
  }
  return fileName.substring(0, maxLength - 3) + '...';
}
</script>
<style lang="scss" scoped>

.file-item-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.file-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.file-action-btn {
  cursor: pointer;
  color: var(--el-text-color-regular);
  transition: color var(--el-transition-duration);
  
  &:hover {
    color: var(--el-color-primary);
  }
}

.el-upload-list__item-name {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 4px;
  min-width: 0; /* 允许flex子元素收缩 */
}

.el-upload-list__item-file-name {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 200px; /* 限制最大宽度 */
}
</style>
