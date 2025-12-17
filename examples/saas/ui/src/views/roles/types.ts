export interface Category {
  id: number,
  code: string,
  name: string,
}

export type ElTagStyle = "info" | "primary" | "warning" | "danger" | "success";

export interface Tag {
  name: string,
  code: string,
  category: string,
  style?: ElTagStyle,
  projectCode?: string,
}

export type CategoryMapFunc = (category: string) => ElTagStyle;
export const categoryMapFuncWrapper = (a: () => void): CategoryMapFunc => {
  return (_) => {
    a();
    return "info";
  };
}

export type TagsModifyFunc = (index: string, done: () => void, current: Tag[]) => void;
export const tagsModifyFuncWrapper = (a: () => void): TagsModifyFunc => {
  return (_, done, _c) => {
    a();
    done();
  };
}
