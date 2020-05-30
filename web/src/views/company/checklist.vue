<template >
    <div style="background-color: #ffffff;width: 100%;height: 100%;overflow: auto">
    <div class="checklist_bigDiv">
        <el-card class="checklist_top1">
            <div v-if="radio=='1'">
                <span style="text-align: center;font-size: 28px;width: 88%;float: left">供应商-已审核</span>
            </div>
            <div v-else>
                <span style="text-align: center;font-size: 28px;width: 88%;float: left">供应商-待审核</span>
                <el-button v-model="dialogform.company_status" label="-1" type="text" @click=" companystatu"
                           style="float: right;margin-right: 10px">驳回
                </el-button>
                <el-button v-model="dialogform.company_status" label="1" type="text" @click="shenHe"
                           style="float: right;margin-right: 10px">审核
                </el-button>

                <!--      表单      -->
                <el-dialog title="审核资质" :visible.sync="dialogFormVisible" width="550px">
                    <el-form :model="dialogform">
                        <!--<el-form-item label="公司性质" :label-width="formLabelWidth">
                            <el-input v-model="dialogform.company_nature" style="width: 70%"></el-input>
                        </el-form-item>-->
                        <el-form-item label="大客户认证" :label-width="formLabelWidth">
                            <el-select v-model="dialogform.auth_vip" placeholder="是否大客户" style="width: 70%">
                                <el-option label="是" :value=1></el-option>
                                <el-option label="否" :value=0></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="代理权限" :label-width="formLabelWidth">
                            <el-select v-model="dialogform.auth_agent" placeholder="是否有找代理权限" style="width: 70%">
                                <el-option label="开通" :value=1></el-option>
                                <el-option label="禁止" :value=0></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="调拨" :label-width="formLabelWidth">
                            <el-select v-model="dialogform.auth_allocation" placeholder="是否开启调拨" style="width: 70%">
                                <el-option label="开通" :value=1></el-option>
                                <el-option label="禁止" :value=0></el-option>
                            </el-select>
                        </el-form-item>
                        <!--<el-form-item label="公司简介" :label-width="formLabelWidth">
                            &lt;!&ndash;<textarea v-model="dialogform.company_profile" style="width:50% "></textarea>&ndash;&gt;
                            <el-input type="textarea" placeholder="请输入内容(不能超过600字)" style="width:70% "
                                      :rows="2" maxlength='600' v-model="dialogform.company_profile">

                            </el-input>
                        </el-form-item>-->
                    </el-form>

                    <div slot="footer" class="dialog-footer">
                        <el-button @click="dialogFormVisible = false">取 消</el-button>
                        <el-button type="primary" @click="companystatus">确 定</el-button>
                    </div>
                </el-dialog>
            </div>
        </el-card>

        <div v-model="form">

            <div class="checklist_top2_left">
                <el-card>
                    <div></div>
                    <h4 style="text-align: center">基本信息</h4>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">企业ID:</span><span
                            style="float: left">{{form.id}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">企业名称:</span><span style="float: left">{{form.company_name}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">法人姓名:</span><span style="float: left">{{form.legal_person_name}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">法人身份证:</span><span style="float: left">{{form.legal_person_id}}</span>
                    </div>

                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">联系人:</span><span style="float: left">{{form.contacts}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">联系电话:</span><span
                            style="float: left">{{form.tel}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">营业执照号:</span><span style="float: left">{{form.business_license}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">纳税识别号:</span><span style="float: left">{{form.taxpayer}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">企业状态:</span>
                        <span style="float: left" v-if="form.company_status==0">审核中</span>
                        <span style="float: left" v-else-if="form.company_status==1">已通过</span>
                        <span style="float: left" v-else></span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">企业类型:</span>
                        <span style="float: left" v-if="form.company_type==0">采购</span>
                        <span style="float: left" v-else-if="form.company_type==1">供应</span>
                        <span style="float: left" v-else-if="form.company_type==2">采购/供应</span>
                        <span style="float: left" v-else></span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">地址:</span><span
                            style="float: left">{{form.address}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">省份:</span><span style="float: left">{{form.province}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">城市:</span><span
                            style="float: left">{{form.city}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">区域:</span><span
                            style="float: left">{{form.area}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">经度:</span><span
                            style="float: left">{{form.lat}}</span>
                    </div>
                    <div class="checklist_top2_left_row">
                        <span class="checklist_top2_left_row_title">维度:</span><span
                            style="float: left">{{form.lag}}</span>
                    </div>

                    <div class="checklist_top2_left_rowPic">
                        <span style="margin-right: 5%;float: left">企业图标:</span>
                        <img :src="form.logo"
                             style="width:120px; height:140px; float: left;margin-top: 2%"/>

                    </div>
                </el-card>

            </div>

            <div class="checklist_top2_right">
                <el-card style="height: 100%">
                    <h4 style="text-align: center">审核文件</h4>
                    <div v-for="(item,index) in form.company_file" class="checklist_top2_left_rowPic">
                        <span style="margin-right: 5%;float: left">{{item.file_name}}:</span>
                        <el-popover
                                placement="left"
                                width="450"
                                offset="50"
                                trigger="click">
                            <img :src="item.file_url"
                                 style=" height:450px; "/>
                            <img :src="item.file_url"
                                 style="width:120px; height:140px; float: left;margin-top: 2%" slot="reference"/>
                        </el-popover>
                    </div>
                </el-card>
            </div>

        </div>
    </div>
    </div>
</template>
<script>
    import {FindByCompany, ModifyCompany} from '../../api/company'

    export default {
        name: 'checklist',
        data() {
            return {
                dialogFormVisible: false,
                url: '',
                statusForm: {
                    company_status: 0,
                },
                id: '',
                form: [],
                radio: '0',
                dialogform: {
                    company_nature: '',
                    auth_vip: '',
                    auth_agent: '',
                    auth_allocation: '',
                    company_profile: '',
                    company_status: 0,

                },
                formLabelWidth: '100px',
            }
        },
        mounted() {
            this.form = []
            this.id = this.$route.query.cid
            this.radio = this.$route.query.radio
            FindByCompany(this.id).then(res => {
                this.form = res.data
            })
        },
        watch: {

            '$route'(to, from) {
                this.form = []
                this.id = this.$route.query.cid
                this.radio = this.$route.query.radio
                FindByCompany(this.id).then(res => {
                    this.form = res.data
                })
            }
        },
        methods: {
            shenHe() {
                this.dialogform = []
                this.dialogFormVisible = true
            },
            companystatus: function () {
                this.dialogform.company_status = 1

                ModifyCompany(this.id, this.dialogform).then(res => {

                    this.dialogFormVisible = false
                    if (res.msg == "成功") {
                        this.$router.push({path: "/company"})
                    }
                })

            },
            companystatu: function () {
                this.dialogform.company_status = 2

                ModifyCompany(this.id, this.dialogform).then(res => {

                    this.$router.push({path: "/company"})
                })

            }
        }
    }
</script>
<style>


</style>
