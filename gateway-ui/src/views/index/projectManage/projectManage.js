/*
 * @Description: 首页
 * @Author: 吴文周
 * @Github: https://github.com/fodelf
 * @Date: 2020-03-16 21:55:11
 * @LastEditors: pym
 * @LastEditTime: 2020-09-06 21:40:28
 */
import menuList from 'components/menuList/menuList.vue'
import tableBox from 'components/tableBox/tableBox.vue'
import proDialog from 'components/proDialog/proDialog.vue'
import compCard from './children/compCard.vue'
import {
  getServiceSum,
  getServiceList,
  deleteService
} from '@/api/index/projectManage.js'
export default {
  name: 'projectManage',
  data() {
    return {
      type: 'add',
      menuObj: {
        title: '服务总计',
        total: 0,
        menuList: [],
        active:''
      },
      tablePag: {
        pageNo: 1,
        pageSize: 15,
        totalRecord: 0
      },
      dataList: [],
      headerList: [
        { name: '服务名称', code: 'serviceName' },
        { name: '服务地址', code: 'serviceAddress' },
        { name: '监控检查接口', code: 'useConsulCheckPath' },
        { name: '熔断阈值', code: 'serviceBreak' },
        { name: '限流阈值', code: 'serviceLimit' }
      ],
      keyword: '',
      itemObj: {},
      proFormObj: {}
    }
  },
  components: {
    menuList,
    tableBox,
    proDialog,
    compCard
  },
  methods: {
    addPro() {
      this.$router.push({
        name:'projectAdd',
        type:'add'
      })
    },
    /**
     * @name: selectMenu
     * @description: 根据服务类型查询服务
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    selectMenu(item) {
      let params = {}
      if(!item) {
        params = {
          type: 'all'
        }
      }else {
        this.itemObj = item
        params = {
          type: item.value
        }
      }
      getServiceList(params).then((res) => {
        this.dataList = res.serverList || []
        this.$forceUpdate()
        // this.$nextTick(() => {
        //   this.$refs.table.doLayout()
        // })
      })
    },
    /**
     * @name: queryProList
     * @description: 获取服务列表
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    queryProList() {
      getServiceSum({}).then(res => {
        this.menuObj.total = res.sum || 0
        this.menuObj.menuList = res.serverList || []
        // this.menuObj.active = this.itemObj.type ? this.itemObj.type : this.menuObj.menuList[0].type
        // if (this.menuObj.menuList.length !== 0 && flag) {
          this.$nextTick(() => {
            this.selectMenu()
          })
        // }
      })
    },
    deleteRow(data) {
      this.$confirm('确认删除此服务？')
        .then(() => {
          deleteService({ serviceId: data.serviceId }).then(() => {
            this.$message({
              type: 'success',
              message: '删除成功'
            })
            this.selectMenu(data)
          })
        })
        .catch(() => {})
    },
    editRow(row) {
      this.$router.push({
        name:'projectAdd',
        query:{
          id: row.serverId,
          type:'edit'
        }
      })
    },
    checkRow(row){
      this.$router.push({
        name:'projectAdd',
        query:{
          id: row.serverId,
          type:'check'
        }
      })
    }
  },
  mounted() {
    this.queryProList(true)
  },
  created() {
    
  }
}
