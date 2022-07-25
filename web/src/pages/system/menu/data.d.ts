export interface MenuListItem {
  id: number;
  menuName: string;
  parentId: number;
  menuUrl: string;
  menuType: number;
  icon: string;
  sort: number;
  remark: string;
  statusId: number;
  apiUrl: string;
}

export interface MenuListPagination {
  total: number;
  pageSize: number;
  current: number;
}

export interface MenuListData {
  list: MenuListItem[];
  pagination: Partial<MenuListPagination>;
}

export interface MenuListParams {
  pageSize?: number;
  currentPage?: number;
  filter?: { [key: string]: any[] };
  sorter?: { [key: string]: any };

}
