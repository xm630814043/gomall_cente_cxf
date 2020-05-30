<template>
    <div>
        <Split style="height: 100vh;">
            <SplitArea :size="35"
                       class="to-do-list">
                <div style="float: left">
                    <div class="list-title">
                        <el-input v-model="input" placeholder="请输入产品名称" @change="changeRadio"></el-input>
                        <el-radio v-model="radio" label="1" @change="changeRadio">待审核</el-radio>
                        <el-radio v-model="radio" label="4" @change="changeRadio">未通过</el-radio>
                        <hr/>
                        <!--表格-->
                        <el-table
                                :row-style="{height:'75px',textAlign:'center'}"

                                :data="tableData"
                                style="width: 90%"
                                @row-click="getRow"
                                height=450

                                :highlight-current-row="true">
                            <el-table-column
                                    prop="id"
                                    label="#ID"
                                    width=120>
                            </el-table-column>
                            <el-table-column
                                    prop="product_name"
                                    label="产品名称"
                                    width=120>
                            </el-table-column>
                            <el-table-column
                                    prop="producer"
                                    label="生产商"
                                    width=150>
                            </el-table-column>
                        </el-table>
                        <!--分页-->
                        <el-pagination style="margin-left: 80px"
                                       layout="prev, pager, next"
                                       @current-change="handleCurrentChange"
                                       :current-page="page.currentPage"
                                       :page-size="page.pageSize"
                                       :total="page.pageTotal">
                        </el-pagination>
                    </div>
                </div>
            </SplitArea>
            <SplitArea :size="65">
                <router-view></router-view>
            </SplitArea>
        </Split>
    </div>
</template>
<script>
    import {GetProductBasic, UpdateProductStatus} from "../../api/company"

    export default {
        name: 'goods',
        data() {
            return {
                data: [],
                page: {
                    pageNum: 1,
                    pageSize: 5,
                    pageTotal: 0,
                },
                radio: '1',
                activeIndex: '1',
                channels: [],
                input: '',
                tableData: [],
            };
        },
        mounted() {
            GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                this.tableData = res.data.data
                this.page.pageTotal = res.data.count
            })
        },
        watch: {
            '$route'(to, from) {
                GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                    this.tableData = res.data.data
                    this.page.pageTotal = res.data.count
                })
            }
        },
        methods: {
            noPassStatus() {
                if (this.data.id == undefined) {
                    alert("请选择商品")
                } else {
                    UpdateProductStatus(this.data.id, 4).then(res => {
                        alert(res.msg)
                        this.data = []
                        this.productImg = ''
                        GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                            this.tableData = res.data.data
                            this.page.pageTotal = res.data.count
                        })
                    })
                }
            },
            passStatus() {
                if (this.data.id == undefined) {
                    alert("请选择商品")
                    return
                } else {
                    UpdateProductStatus(this.data.id, 2).then(res => {
                        alert(res.msg)
                        this.data = []
                        this.productImg = ''
                        GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                            this.tableData = res.data.data
                            this.page.pageTotal = res.data.count
                        })
                    })
                }

            },
            handleCurrentChange(currentPage) {
                this.page.currentPage = currentPage
                GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                    this.tableData = res.data.data
                })
            },
            resize() {
                console.log('resize')
            },
            getRow(row, column, event) {
                this.$router.push({path: '/goods/detail', query: {id: row.id, status: this.radio}})
            },
            changeRadio(label) {
                this.data = []
                this.page.currentPage = 1
                GetProductBasic(this.page.currentPage, this.page.pageSize, this.input, this.radio).then(res => {
                    this.tableData = res.data.data
                    this.page.pageTotal = res.data.count
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

    .list-title {
        line-height: 50px;
        height: 50px;
        background-color: white;
        padding-left: 20px;

    }
</style>
