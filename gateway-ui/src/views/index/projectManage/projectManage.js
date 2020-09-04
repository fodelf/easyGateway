/*
 * @Description: 首页
 * @Author: 吴文周
 * @Github: https://github.com/fodelf
 * @Date: 2020-03-16 21:55:11
 * @LastEditors: 吴文周
 * @LastEditTime: 2020-05-14 17:34:09
 */
import menuList from 'components/menuList/menuList.vue'
import tableBox from 'components/tableBox/tableBox.vue'
import proDialog from 'components/proDialog/proDialog.vue'
import compCard from './children/compCard.vue'
import {
  getProjectSum,
  getProjectList,
  deleteProject,
  action
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
        { name: '序号', code: 'index', width: 60 },
        { name: '服务名称', code: 'projectName' },
        { name: '服务类型', code: 'type' },
        { name: '服务关键字', code: 'keyword' },
        { name: '创建时间', code: 'createTime' },
        { name: '本地路径', code: 'pathUrl' }
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
      this.type = 'add'
      this.$refs.proDialog.show()
    },
    /**
     * @name: getPageNo
     * @description: 切换分页查询
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    getPageNo(val) {
      this.tablePag.pageNo = val
      this.selectMenu(this.itemObj)
    },
    getProList(item) {
      // console.log(item)
      this.tablePag.pageNo = 1
      this.selectMenu(item)
    },
    /**
     * @name: selectMenu
     * @description: 根据服务类型查询服务
     * @param {type}: 默认参数
     * @return {type}: 默认类型
     */
    selectMenu(item) {
      this.itemObj = item
      let params = {
        type: item.type,
        pageNum: this.tablePag.pageNo,
        pageSize: this.tablePag.pageSize,
        keyword: this.keyword
      }
      getProjectList(params).then((res) => {
        this.dataList = (res.list || []).map((item, index) => {
          item.index =
            index + (this.tablePag.pageNo - 1) * this.tablePag.pageSize + 1
          return item
        })
        this.tablePag.totalRecord = res.total || 0
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
    queryProList(flag) {
      getProjectSum({}).then(res => {
        // console.log(res)
        this.menuObj.total = res.total || 0
        this.menuObj.menuList = res.list || []
        this.menuObj.active = this.itemObj.type ? this.itemObj.type : this.menuObj.menuList[0].type
        if (this.menuObj.menuList.length !== 0 && flag) {
          this.$nextTick(() => {
            this.selectMenu(this.menuObj.menuList[0])
          })
        }
      })
    },
    deleteRow(data) {
      this.$confirm('确认删除此服务？')
        .then(() => {
          deleteProject({ projectId: data.projectId }).then(() => {
            this.$message({
              type: 'success',
              message: '删除成功'
            })
            this.getList(data.type)
          })
        })
        .catch(() => {})
    },
    editRow(data) {
      this.type = 'modify'
      this.itemObj = data
      this.$nextTick(() => {
        this.$refs.proDialog.show()
      })
    },
    // 新建分支
    newBranch(data){
      this.itemObj = data
      this.$nextTick(() => {
        this.$refs.proDialog.show()
      })
    },
    getList(type){
      this.itemObj.type = type;
      this.queryProList(true)
    },
    action(data){
      action(data).then(() => {
        this.$message({
          type: 'success',
          message: '脚本已经启动'
        })
      })
    }
  },
  mounted() {
    this.queryProList(true)
  },
  created() {
    
  }
}
