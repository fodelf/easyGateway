/*
 * @Descripttion: 
 * @version: 
 * @Author: pym
 * @Date: 2020-09-06 15:56:49
 * @LastEditors: 吴文周
 * @LastEditTime: 2020-09-15 19:46:35
 */
import {
  getServiceType,
  addService,
  serviceDetail,
  updateService
} from '@/api/index/projectManage.js'
export default {
  name: 'projectAdd',
  data() {
    const validateEn = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入服务名称'));
      } else {
        if((/[^A-Za-z]/g).test(value)){
          callback(new Error('只能输入英文字母'));
        }
        callback();
      }
    }
    return {
      ruleForm: {
        serviceName: '',
        serviceType: '',
        serviceAddress: '',
        servicePort: '',
        serviceRules: [],
        serviceBreak:"",
        serviceLimit:'',
        useConsulId:'',
        useConsulTag:'',
        useConsulCheckPath:'',
        useConsulPort:'',
        useConsulInterval:'',
        useConsulTimeout:'',
        dingdingAccessToken:'',
        dingdingSecret:'',
        dingdingList:[],
      },
      serviceRules:{
        serviceName:[
          { required: true, validator: validateEn, trigger: 'blur' },
        ],
        serviceType:[
          { required: true, message: '请选择服务类型', trigger: 'blur' },
        ],
        serviceAddress:[
          { required: true, message: '请输入服务地址', trigger: 'blur' },
        ]
      },
      rules:{
        apiUrl:[
          { required: true, message: '请输入拦截地址', trigger: 'blur' },
        ]
      },
      inputVisible:false,
      inputValue:'',
      baseInfo:'baseInfo',
      useConsul:'useConsul',
      messageWarn:'messageWarn',
      serverList:[]
    }
  },
  methods:{
    handleClose(tag) {
      this.ruleForm.dingdingList.splice(this.ruleForm.dynamicTags.indexOf(tag), 1);
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        this.ruleForm.dingdingList.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = '';
    },
    addRule() {
      this.ruleForm.serviceRules.push(
        {
          interceptLoc:'',
          locationReset:''
        }
      );
    },
    deleteRule(index) {
      this.ruleForm.serviceRules.splice(index,1)
    },
    cancel() {
      this.$router.push({
        name:'projectManage'
      })
    },
    queryProjectType() {
      getServiceType().then(res=>{
        this.serverList = res.serverTypeList || []
      })
    },
    saveRule() {
      this.$refs["ruleForm"].validate((valid) => {
        if (valid) {
          let params = JSON.parse(JSON.stringify(this.ruleForm))
          params.servicePort = params.servicePort*1
          params.serviceLimit = params.serviceLimit*1
          params.serviceBreak = params.serviceBreak*1
          params.useConsulPort = params.useConsulPort*1
          params.useConsulInterval = params.useConsulInterval*1
          params.useConsulTimeout = params.useConsulTimeout*1
          addService(params).then(res=>{
            this.$router.push({
              name:'projectManage'
            })
          })
        } else {
          return false
        }
      })
    },
    initDetail() {
      let id = this.$route.query.id
      serviceDetail(id).then(res=>{
        this.ruleForm = JSON.parse(JSON.stringify(res))
      })
    },
    updateRule() {
      let params = this.ruleForm
      updateService(params).then(res=>{
        this.$router.push({
          name:'projectManage'
        })
      })
    }
  },
  created() {
    this.queryProjectType()
    if(this.$route.query.id){
      this.initDetail()
    }else {
      this.ruleForm= {
        serviceName: '',
        serviceType: 'http',
        serviceAddress: '',
        servicePort: '',
        serviceRules: [],
        serviceBreak:"",
        serviceLimit:'',
        useConsulId:'',
        useConsulTag:'',
        useConsulCheckPath:'',
        useConsulPort:'',
        useConsulInterval:'',
        useConsulTimeout:'',
        dingdingAccessToken:'',
        dingdingSecret:'',
        dingdingList:[],
      }
    }
  }
}