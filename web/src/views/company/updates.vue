<template>
    <div class="title">
        <div class="container_table">
            <el-table
                    :data="tableData"
                    stripe
                    style="width: 100%"
                    :default-sort="{prop: 'date', order: 'descending'}"
            >
                <el-table-column
                        prop="ID"
                        sortable
                        label="待审核文件ID"
                        min-width="8.19%">
                </el-table-column>
                <el-table-column prop="file_url" label="经营许可" min-width="8.19%">
                    <template slot-scope="scope">
                        <el-popover
                                placement="right"
                                title=""
                                trigger="click">
                            <el-image slot="reference" :src="scope.row.file_url" :alt="scope.row.file_name"
                                      style="max-height: 100px;max-width: 100px"></el-image>
                            <el-image :src="scope.row.file_url"></el-image>
                        </el-popover>
                        <!--                <img :src="scope.row.file_name" width="40" height="40"/>-->
                    </template>
                </el-table-column>
                <el-table-column prop="file_url" label="企业资质" min-width="8.19%">
                    <template slot-scope="scope">
                        <!--                <a :href="scope.row.file_url" target="_blank">"企业资质"</a>-->
                        <router-link :to="{path:scope.row.file_url}">企业资质</router-link>
                    </template>
                </el-table-column>
                <el-table-column label="审核文件" min-width="8.19%" prop="file_name">

                </el-table-column>
            </el-table>
            <el-button v-model="form.company_status" label="1" type="text" @click="companystatus">通过</el-button>
            &nbsp;&nbsp;&nbsp;
            <el-button v-model="form.company_status" label="2" type="text" @click="companystatu ">驳回</el-button>
            <!--<el-button v-model="form.company_status" label="1" @click="passStatus" size="small">通过</el-button>
            <el-button v-model="form.company_status" label="2" @click="noPassStatus" size="small">驳回</el-button>-->
        </div>
    </div>
</template>
<script>
    import {FindByCompany, ModifyCompany} from '../../api/company'

    export default {
        name: 'updates',
                data() {
            return {
                form: {
                    company_status: 0,
                },
                tableData: [],

            }
        },
        mounted() {

            this.id = this.$route.query.cid
            //alert(this.id)
            FindByCompany(this.id).then(res => {
                // console.log(res.data)
                console.log("查找", res)
                this.tableData = res.data['company_file']
            })
        },
        watch:{
            '$route' (to, from) {
                /*if(this.id==this.$route.query.cid){
                    return
                }*/

                this.id = this.$route.query.cid

                FindByCompany(this.id).then(res => {
                    // console.log(res.data)
                    console.log("查找", res)
                    this.tableData = res.data['company_file']
                })
            }
        },
        methods: {


                    companystatus: function () {
                        this.form.company_status = 1
                        ModifyCompany(this.id, this.form).then(res => {
                            console.log(res.data)
                            console.log("审核批准已通过")
                            if(res.msg=="成功"){
                                this.tableData=[]
                                this.$router.push({path:"/company"})
                            }
                        })
                    }
                ,
                    companystatu: function () {
                        this.form.company_status = 2

                        ModifyCompany(this.id, this.form).then(res => {
                            console.log(res.data)
                            console.log("审核批准驳回")
                           // this.tableData=[]
                            this.$router.push({path:"/company"})
                        })

                    }
                }
            }
</script>
<style>
    .checklists {
        background-color: #ffffff;
        text-align: center;
    }

    .listsblock {
        margin-top: 10px;

    }

    .listsspan {
        margin-left: 10px;
        font-size: 28px;
    }

    .buttons {
        margin-top: 20px;
        margin-left: 40px;
        background-color: #ffffff;
    }

    .right {
        margin-left: 70px;
    }
</style>