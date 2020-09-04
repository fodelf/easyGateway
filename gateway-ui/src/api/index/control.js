/*
 * @Description:控制台
 * @Author: 吴文周
 * @Github: https://github.com/fodelf
 * @Date: 2019-06-05 17:27:29
 * @LastEditors: 吴文周
 * @LastEditTime: 2020-05-11 09:26:13
 */
import request from '@/utils/request'

export function queryTemplateList (params) {
  return request({
    url: '/api/drag/getPageList',
    method: 'GET',
    params: params
  })
}
export function query (params) {
  return request({
    url: '/api/dashboard/getVisitDetail',
    method: 'POST',
    params: params
  })
}
export function SOLIDER_33 (params) {
  return request({
    url: '/compcag/SOLIDER_33',
    method: 'POST',
    params: params
  })
}
export function createApp (params) {
  return request({
    url: '/app/createApp',
    method: 'POST',
    params: params
  })
}
