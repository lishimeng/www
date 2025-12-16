export interface Category {
  id: number,
  code: string,
  name: string,
}

export type ElTagStyle = "info" | "primary" | "warning" | "danger" | "success";

export interface Permission {
  name: string,
  code: string,
  category: string,
  style?: ElTagStyle,
}

export type CategoryMapFunc = (category: string) => ElTagStyle;
export const categoryMapFuncWrapper = (a: () => void): CategoryMapFunc => {
  return (_) => {
    a();
    return "info";
  };
}

export type PermissionModifyFunc = (role: string, done: () => void, current: Permission[]) => void;
export const permissionModifyFuncWrapper = (a: () => void): PermissionModifyFunc => {
  return (_, done, _c) => {
    a();
    done();
  };
}
