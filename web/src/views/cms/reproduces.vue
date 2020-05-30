<template>
    <!-- 热销商品开始 -->
    <div style="overflow: auto;width: 100%;height: 100%">
        <el-row>
            <el-card shadow="never" style="border-radius: 0px">
                <el-button type="primary" size="small" @click="chouti">选择商品</el-button>
                <span v-if="form.subject_id==1" style="float: right;font-size: 30px;margin-right: 10%">热销商品</span>
                <span v-else style="float: right;font-size: 30px;margin-right: 10%">商品推荐</span>
            </el-card>
        </el-row>
        <div style="overflow: auto; overflow-x: hidden;">
            <!--抽屉-->
            <el-drawer
                    style="height: auto; overflow-x: hidden;"
                    title="我是标题"
                    :visible.sync="drawer"
                    size="35%"
                    :with-header="false"
                    @open="">
                <!--滚动条-->
                <el-scrollbar style="height: 700px;overflow: auto">
                    <el-row>
                        <el-input v-model="brandinput" @change="findListByName" placeholder="请输入商铺名称" size="small"
                                  style="width: 23%;margin: 10px"
                        ></el-input>
                        <el-input v-model="brandinput" @change="findListByName" placeholder="请输入品牌名称" size="small"
                                  style="width: 23%;margin-top: 10px"
                        ></el-input>
                        <el-input v-model="form.input" @change="findListByName" size="small" placeholder="请输入商品名称"
                                  style="width: 25%;margin: 10px"></el-input>
                        <el-button type="primary" size="small" style="margin-left: 0%" @click="findListByName">搜索
                        </el-button>
                    </el-row>
                    <!--Radio单选框-->
                    <el-row>
                        <div style="float: left;margin-top: 1%;width: 100%">
                            排序:
                            <el-radio-group v-model="radio">
                                <el-radio :label="3" @change="findListByName">默认ID</el-radio>
                                <el-radio :label="6" @change="findListByName">点击量</el-radio>
                                <el-radio :label="9" @change="findListByName">销量</el-radio>
                                <el-radio :label="10" @change="findListByName">收藏量</el-radio>
                            </el-radio-group>
                        </div>
                    </el-row>
                    <div>
                        <el-row style="padding: 10px;"
                                v-infinite-scroll="load"
                                :infinite-scroll-disabled="disablea"
                        >
                            <el-col :span="7" v-for="(item, index) in addProduct" :key="item.ID" style="margin: 10px;"
                            >
                                <el-card :body-style="{ padding: '0px',height:'280px'}"
                                >
                                    <img :src="item.Pic" class="image1">
                                    <div style="padding: 14px">
                                        <span style="font-size:15px ">{{item.ID }}&nbsp;{{item.ProductName }}</span> <span
                                            style="font-size: 10px;float: right;color: gray;margin-top: 9px">
                                销量{{item.SellNum}}&nbsp;&nbsp;点击量{{item.ClickNum}}&nbsp;&nbsp;收藏量{{item.favourite}}
                             </span>
                                        <div class="bottom clearfix">
                                            <el-popconfirm
                                                    title="确定添加吗？" @onConfirm="submitForm(item.ID)" @onCancel=""
                                            >
                                                <el-button type="text" slot="reference">添加</el-button>
                                            </el-popconfirm>
                                        </div>
                                    </div>
                                </el-card>
                            </el-col>
                        </el-row>
                    </div>
                    <p v-if="loading">加载中...</p>
                    <p v-if="noMore">没有更多了</p>
                </el-scrollbar>
            </el-drawer>
        </div>
        <el-row style="padding: 10px">
            <el-col :span="6" v-for="(item, index) in hotGoodsList" :key="item.ID" style="margin: 10px">
                <el-card :body-style="{ padding: '0px',height: '330px' }">
                    <img :src="item.Pic" class="image">
                    <div style="padding: 14px;">
                        <span>{{item.ID }}&nbsp;{{item.ProductName }}</span> <span
                            style="font-size: 10px;float: right;color: gray;margin-top: 9px">
                                销量{{item.SellNum}}&nbsp;&nbsp;点击量{{item.ClickNum}}&nbsp;&nbsp;收藏量{{item.favourite}}
                             </span>
                        <div class="bottom clearfix">
                            <el-popconfirm
                                    title="确定删除产品吗？" @onConfirm="handleDelete(item.ID)" @onCancel=""
                            >
                                <el-button type="text" slot="reference">删除</el-button>
                            </el-popconfirm>
                        </div>
                    </div>
                </el-card>
            </el-col>
        </el-row>
    </div>
</template>
<script>
    import {AddSubject, FindBySubjectList, FindByUnSubjectList, RemoveSubject} from '../../api/company'

    export default {
        name: "reproduces",
        data() {
            return {
                count: 9,  //limit条数
                loading: false,
                shopinput: '',
                brandinput: '',
                form: {
                    subject_id: 0,
                    object_id: 0,
                    contentType: '',
                    input: '',
                },
                drawer: false,
                radio: 3,
                contentType: '',
                limit: 0,
                activeIndex: '1',
                hotGoodsList: [],
                addProduct: [],
                len: 0,
            }
        },
        computed: {
            noMore() {
                return this.count >= this.len
            },
            disablea() {
                return (this.loading || this.noMore)
            },
        },
        mounted() {
            this.form.subject_id = this.$route.query.sid
            this.form.contentType = this.$route.query.sname
            this.limit = this.$route.query.slimit
            FindBySubjectList(this.form.subject_id, this.form.contentType, this.limit).then(res => {
                console.log(res.data)
                this.hotGoodsList = res.data
            })
        },
        watch: {
            '$route'(to, from) {
                this.form.subject_id = this.$route.query.sid
                this.form.contentType = this.$route.query.sname
                this.limit = this.$route.query.slimit
                FindBySubjectList(this.form.subject_id, this.form.contentType, this.limit).then(res => {
                    console.log(res.data)
                    this.hotGoodsList = res.data
                })
            },

        },
        methods: {
            chouti() {
                this.count = 9
                this.form.input=''
                this.shopinput=''
                this.brandinput=''
                this.radio=3
                FindByUnSubjectList(this.form.subject_id, this.form.contentType, this.form.input, this.shopinput, this.brandinput, this.radio, this.count).then(res => {
                    this.len = res.data.count
                    this.addProduct = res.data.data
                    this.drawer = true
                })
            },
            load() {
                this.loading = true
                setTimeout(() => {
                    this.count += 6
                    FindByUnSubjectList(this.form.subject_id, this.form.contentType, this.form.input, this.shopinput, this.brandinput, this.radio, this.count).then(res => {
                        this.addProduct = res.data.data
                        this.loading = false
                    })
                }, 2000)
            },
            submitForm(object_id) {
                this.form.object_id = object_id
                AddSubject(this.form).then(res => {
                    FindBySubjectList(this.form.subject_id, this.form.contentType, this.limit).then(res => {
                        console.log(res.data)
                        this.hotGoodsList = res.data
                    })
                    this.drawer = false
                })
            },
            findListByName() {
                this.count = 9
                FindByUnSubjectList(this.form.subject_id, this.form.contentType, this.form.input, this.shopinput, this.brandinput, this.radio, this.count).then(res => {
                    this.addProduct = res.data.data
                    this.len = res.data.count
                })
            },
            handleDelete: function (index) {
                RemoveSubject(index, this.form.subject_id).then(res => {
                    console.log(res.data)
                    FindBySubjectList(this.form.subject_id, this.form.contentType, this.limit).then(res => {
                        console.log(res.data)
                        this.hotGoodsList = res.data
                    })
                })
            },
        }
    };
</script>
<style scoped>
    @import url("../../../public/static/css/index/index.css");

    .upload-demo {
        float: right;
    }

    body .el-scrollbar__wrap {
        overflow-x: hidden;
    }
</style>