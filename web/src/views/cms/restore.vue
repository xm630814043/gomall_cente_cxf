<template>
    <!-- 优选商家开始 -->
    <div>
        <el-row>
            <el-card shadow="never" style="border-radius: 0px">
                <el-button type="primary" size="small" @click="chouti()">选择商品</el-button>
                <span  style="float: right;font-size: 30px;margin-right: 10%">优选店铺</span>
            </el-card>

        </el-row>
        <!--抽屉-->
        <div style="overflow: auto; overflow-x: hidden;">
            <el-drawer
                    style="height: auto; overflow-x: hidden;"
                    title="我是标题"
                    :visible.sync="drawer"
                    size="35%"
                    :with-header="false">
                <el-scrollbar style="height: 700px;overflow: auto">  <!--滚动条-->

                    <el-row>
                        <el-input v-model="form.input" size="small" placeholder="请输入店铺名称"
                                  style="width: 60%;margin: 10px" @change="findListByName"></el-input>

                        <el-button type="primary" size="small" style="margin-left: 0%" @click="findListByName">搜索</el-button>
                    </el-row>
                    <!--新增Radio单选框-->
                    <el-row>
                        <div style="float: left;margin-top: 1%;width: 100%">
                            排序:
                            <el-radio-group v-model="radio">
                                <el-radio :label="3" @change="findListByName">默认ID</el-radio>
                                <el-radio :label="6" @change="findListByName">热度</el-radio>
                                <el-radio :label="9" @change="findListByName">收藏量</el-radio>
                            </el-radio-group>
                        </div>
                    </el-row>
                    <div>
                        <el-row style="padding: 10px"
                                v-infinite-scroll="load"
                                :infinite-scroll-disabled="disablea">
                            <el-col :span="6" v-for="(item, index) in addShop" :key="item.ID" style="margin: 10px">
                                <el-card :body-style="{ padding: '0px',height:'250px'}">
                                    <img :src="item.logo" class="image1">
                                    <div style="padding: 14px;">
                                        <span style="font-size:15px ">{{item.ID }}&nbsp;{{item.shop_name }}</span> <span
                                            style="font-size: 10px;float: right;color: gray;margin-top: 7px">
                                热度{{item.hot}}&nbsp;&nbsp;收藏量{{item.favourite}}
                             </span>
                                        <div class="bottom clearfix">
                                            <el-popconfirm
                                                    title="确定添加这家推荐店铺吗？" @onConfirm="submitForm(item.ID)" @onCancel=""
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
            <el-col :span="6" v-for="(item, index) in goodShopList" :key="item.ID" style="margin: 10px">
                <el-card :body-style="{ padding: '0px',height: '330px' }">
                    <img :src="item.logo" class="image">
                    <div style="padding: 14px;">
                        <span>{{item.ID }}&nbsp;{{item.shop_name }}</span> <span
                            style="font-size: 10px;float: right;color: gray;margin-top: 7px">
                                热度{{item.hot}}&nbsp;&nbsp;收藏量{{item.favourite}}
                             </span>
                        <div class="bottom clearfix">
                            <el-popconfirm
                                    title="确定删除这家推荐店铺吗？" @onConfirm="handleDelete(item.ID)" @onCancel=""
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
        name: "restore",
        data() {
            return {
                loading: false,
                drawer: false,
                id: 0,
                contentType: '',
                limit: 0,
                activeIndex: '1',
                goodShopList: [],
                addShop: [],
                form: {
                    subject_id: 0,
                    object_id: 0,
                    contentType: '',
                    input: '',
                },
                radio: 3,  //排序
                count: 0,
                shopinput: '',
                brandinput: '',
                len:0,
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

                this.goodShopList = res.data
            })
        },
        methods: {
            load() {
                this.loading = true
                setTimeout(() => {
                    this.count += 6
                    FindByUnSubjectList(this.form.subject_id, this.form.contentType, this.form.input, this.shopinput, this.brandinput, this.radio, this.count).then(res => {
                        this.addShop = res.data.data
                        this.loading = false
                    })
                }, 2000)

            },
            chouti() {
                this.count = 9
                this.shopinput=''
                this.radio=3
                FindByUnSubjectList(this.form.subject_id, this.form.contentType, this.form.input, this.shopinput, this.brandinput, this.radio, this.count).then(res => {
                    this.addShop = res.data.data
                    this.len = res.data.count
                    this.drawer = true
                })
            },
            findListByName() {
                FindByUnSubjectList(this.form.subject_id, this.form.contentType, this.form.input, this.shopinput, this.brandinput, this.radio, this.count).then(res => {
                    this.addShop = res.data.data
                    this.len = res.data.count

                })
            },
            handleDelete: function (index) {
                // console.log(index,this.id)
                RemoveSubject(index, this.form.subject_id).then(res => {
                    console.log(res.data)
                    FindBySubjectList(this.form.subject_id, this.form.contentType, this.limit).then(res => {
                        console.log(res.data)
                        this.goodShopList = res.data
                    })

                })
            },
            submitForm(object_id) {

                this.form.object_id = object_id
                console.log(this.form)
                AddSubject(this.form).then(res => {
                    FindBySubjectList(this.form.subject_id, this.form.contentType, this.limit).then(res => {
                        console.log(res.data)
                        this.goodShopList = res.data
                    })
                    this.drawer=false
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


</style>