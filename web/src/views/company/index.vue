<template>
    <div>
        <Split style="height: 100vh;">
            <SplitArea :size="35" class="to-do-list">
                <div style="float: left">
                    <div class="list-title">
                        <el-input v-model="filterParams.KeyWord" placeholder="请输入公司名称" @change="changeRadio"></el-input>
                        <el-radio v-model="radio" label="0" @change="changeRadio">待审核</el-radio>
                        <el-radio v-model="radio" label="1" @change="changeRadio">已通过</el-radio>
                        <hr/>
                        <el-table
                                :data="tableData"
                                stripe
                                height=400
                                :highlight-current-row="true"
                                style="width: 100%"
                                :row-style="{height:'75px',textAlign:'center'}"
                                :default-sort="{prop: 'date', order: 'descending'}"
                                @row-click="handdle"
                        >
                            <el-table-column
                                    prop="id"
                                    label="ID"
                                    sortable
                                    min-width="8.19%">
                            </el-table-column>
                            <el-table-column
                                    prop="company_name"
                                    label="公司名称"
                                    sortable
                                    min-width="8.19%">
                            </el-table-column>
                            <el-table-column
                                    prop="legal_person_name"
                                    label="法人姓名"
                                    min-width="8.19%">
                            </el-table-column>

                        </el-table>
                        <div class="eltable">
                            <!-- <span class="demonstration">共{{total}}条</span>-->
                            <el-pagination class="fy"
                                           layout="prev, pager, next"
                                           @current-change="current_change"
                                           :total="total"
                                           :page-size="pageSize"
                                           :current-page="filterParams.PageNum"
                                           background
                            >
                            </el-pagination>
                        </div>
                    </div>

                    <div v-show="tableshows" style="margin-top:600px ">
                        <el-table
                                :data="tableDatas"
                                stripe
                                :highlight-current-row="true"
                                style="width: 100%"
                                :default-sort="{prop: 'date', order: 'descending'}"
                                @row-click="handdles"
                        >
                            <el-table-column
                                    prop="ID"
                                    label="ID"
                                    sortable
                                    min-width="8.19%">
                            </el-table-column>
                            <el-table-column
                                    prop="username"
                                    label="用户名称"
                                    sortable
                                    min-width="8.19%">
                            </el-table-column>

                           <el-table-column
                                    prop="mobile"
                                    label="手机"
                                    min-width="8.19%">
                            </el-table-column>

                        </el-table>
                        <div class="eltable">
                            <el-pagination class="fy"
                                           layout="prev, pager, next"
                                           @current-change="current_changes"
                                           :total="accountTotle"
                                           :page-size="pageSize"
                                           background
                            >
                            </el-pagination>
                        </div>
                    </div>
                </div>
            </SplitArea>
            <SplitArea :size="65" class="splitArea" >
                <router-view/>
            </SplitArea>
        </Split>
    </div>
</template>
<script>
    import {FindAccountList, FindCompanyList} from '../../api/company'
    export default {
        name: 'company',
        data: function () {
            return {
                statusId: '',
                ids: '',
                activeIndex: '1',
                PageNum: 1,//默认开始页面
                pageSize: 4,//每页的数据条数
                total: 100,//默认数据总数
                accountTotle: 1,
                tableshows: false,
                errors: false,
                tableData: [],
                tableDatas: [],
                filterParams: {
                    PageNum: 0,
                    PageSize: 4,
                    KeyWord: ''
                },
                radio: '0',
            };
        },
        mounted() {
            this.statusId = this.radio
            FindCompanyList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.statusId).then(res => {
                this.tableData = res.data.data
                this.total = res.data.count
            })
        },
        watch: {
            '$route'(to, from) {
                this.statusId = this.radio
                FindCompanyList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.statusId).then(res => {
                    this.tableData = res.data.data
                    this.total = res.data.count
                })
            }
        },
        methods: {
            changeRadio() {
                this.errors = false
                this.tableshows = false
                this.statusId = this.radio  //已审核未审核
                this.filterParams.PageNum = 1
                FindCompanyList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.statusId).then(res => {
                    this.tableData = res.data.data
                    this.total = res.data.count
                })
            },
            handdle: function (row) {
                this.errors = true
                this.tableshows = true

                    FindAccountList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, row.id).then(res => {
                        this.tableDatas = res.data.data
                        this.accountTotle = res.data.count
                    })
                    this.$router.push({path: '/company/checklist', query: {cid: row.id,radio:this.radio}})
                /*}*/
            },
            handdles: function (row) {
                this.$router.push({path: '/company/lists', query: {cid: row.ID}})
            },
            current_change: function (currentPage) {
                this.filterParams.PageNum = currentPage;
                FindCompanyList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.statusId).then(res => {
                    this.tableData = res.data.data
                    //this.total = res.data.count
                })
            },
            current_changes: function (currentPage) {
                this.filterParams.PageNum = currentPage;
                FindAccountList(this.filterParams.PageSize, this.filterParams.PageNum, this.filterParams.KeyWord, this.ids).then(res => {
                    this.tableDatas = res.data.data
                })
            },
        }
    }
</script>
<style scoped>
    .to-do-list {
        background-color: #fff;
        text-align: center;
        overflow: auto;
    }

    .eltable {
        margin-top: 20px;
        text-align: center;
    }
    .splitArea {

        height: auto;
        width: 100%;
        overflow:auto;
    }
    .list-title {
        line-height: 50px;
        height: 50px;
        background-color: white;
        padding-left: 20px;
        float: left;

    }
</style>
