/*
 * @Descripttion: 
 * @version: 
 * @Author: pym
 * @Date: 2020-08-11 10:20:11
 * @LastEditors: 吴文周
 * @LastEditTime: 2020-09-04 09:37:24
 */
const menuTree = [
  { name: 'home', path: '/home', icon: 'icon-shouye', label: '首页' },
  {
    name: 'projectManage',
    path: '/project/projectManage',
    icon: 'icon-xiangmu',
    label: '服务管理'
  },
  {
    name: 'systemManage',
    path: '/system/userManage',
    icon: 'icon-shezhi',
    label: '系统设置',
    children:[
      {
        name: 'serviceSet',
        path: '/system/serviceSet',
        label: '服务设置'
      }
    ]
  }
]

export default menuTree