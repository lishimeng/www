export interface TableColumns {
  prop: string;
  fk?: Fk;
  label: string;
  component?: 'text' | 'enum' | 'enum_tag' | 'enum_dict' | 'time' | 'custom';
  width?: string;
  showOverflowTooltip?: boolean;
  enums?: Enum[];
  dictCode?: string;
  isName?: boolean;
  fixed?: boolean; // 左侧固定显示
  sortable?: boolean | 'custom'; // 是否可排序
}

export interface SearchItem {
  prop: string;
  label: string;
  placeholder?: string;
  component: 'input' | 'select' | 'date-picker' | 'table-select' | 'dict-select';
  options?: { label: string; value: any }[];
  table?: {title: string, url: string, columns: TableColumns[], search: SearchItem[]};
  dictCode?: string;
}

interface Enum {
  value: any,
  label: string,
  type?: 'warning' | 'info' | 'success' | 'primary' | 'danger',
}

interface Fk {
  url: string,
  field: string,
}
