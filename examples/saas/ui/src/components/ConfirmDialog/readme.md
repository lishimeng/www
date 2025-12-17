## Usage
直接调用全局注册的函数`withConfirm`。

```ts
const result = await withConfirm({
  message: "删除的数据不可恢复，是否继续？",
  title: "删除数据",
} as ConfirmOptions);
if (result) {
  // actions...
}
```

### 参数
param: ConfirmOptions | string

```ts
interface ConfirmOptions {
  title?: string; // default: "提示"
  message: string;
  confirmText?: string; // default: "确认"
  cancelText?: string; // default: "取消"
}
```
