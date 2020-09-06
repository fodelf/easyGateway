/*
 * @Description: 描述
 * @Author: 吴文周
 * @Github: https://github.com/fodelf
 * @Date: 2020-03-21 19:52:07
 * @LastEditors: pym
 * @LastEditTime: 2020-09-06 19:29:15
 */
import request from '@/utils/request'

export function getIndexCount() {
  return request({
    url: '/uiApi/v1/index/sum',
    method: 'GET',
  })
}

export function queryIndexTrend(id) {
  return request({
    url: '/uiApi/v1/index/charts/{'+id+'}',
    method: 'GET',
  })
}

export function queryActualTime() {
  return request({
    url: '/uiApi/v1/index/actualTime',
    method: 'GET',
  })
}

export function queryWarningList() {
  return request({
    url: '/uiApi/v1/index/warningList',
    method: 'GET',
  })
}

export function getPersonActive() {
  return request({
    url: '/api/home/getPersonActive',
    method: 'GET',
  })
}
