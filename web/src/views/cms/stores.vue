<template>
    <div>
        <!--<div class="good-goodshops-title">
            店铺列表
            <span>大牌聚集 源头直供</span>
            <el-button slot="reference" class="upload-demo" @click="handleCreate()">添加</el-button>
        </div>-->
        <el-row>
            <el-card shadow="never" style="border-radius: 0px">
                <el-button type="primary" size="small" @click="handleCreate">添加</el-button>
                <span  style="float: right;font-size: 30px;margin-right: 10%">轮播图</span>
            </el-card>
        </el-row>

        <el-row style="padding: 10px">
            <el-col :span="6" v-for="(item, index) in articleList" :key="item.ID" style="margin: 10px">
                <el-card :body-style="{ padding: '0px',height: '330px' }">
                    <img :src="item.image_url" class="image">
                    <div style="padding: 14px;">
                        <span>{{item.ID }}&nbsp;{{item.title }}</span>
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
        <!--<div class="good-goodshops" v-model="activeIndex">
            <div class="good-goodshops-item" v-for="(item,index) in goodShopsList" :key="index">
                <img :src="item.image_url" @click="clickImg(item.ID)" class="goodsimg"/>
                <p style="color: black;font-size: 20px">{{item.ID}} &nbsp;{{ item.title }}</p>
                <hr style="height:1px;border-width:0;color:gray;background-color:gray">
                <el-popconfirm
                        title="这家店铺确定删除吗？" class="buttons" @onConfirm="handleDelete(item.ID)" onCancel=""
                >
                    <el-button slot="reference">删除</el-button>
                </el-popconfirm>
            </div>
        </div>-->
    </div>
</template>

<script>
    import {FindBySubject, FindBySubjectList, RemoveSubject} from '../../api/company'

    export default {
        name: "stores",
        data() {
            return {
                id: 0,
                contentType: '',
                limit: 0,
                activeIndex: '1',
                articleList: []
            }
        },
        mounted() {
            this.id = this.$route.query.sid
            this.contentType = this.$route.query.sname
            this.limit = this.$route.query.slimit
            FindBySubjectList(this.id, this.contentType, this.limit).then(res => {
                console.log(res.data)
                this.articleList = res.data
            })
        },
        methods: {
            handleDelete: function (index) {
                RemoveSubject(index, this.id).then(res => {
                    console.log(res.data)
                    FindBySubjectList(this.id, this.contentType, this.limit).then(res => {
                        this.articleList = res.data
                    })
                })
            },
            handleCreate: function () {
                this.$router.push({
                    path: '/cms/create',
                    query: {cid: this.id, sname: this.contentType, slimit: this.limit}
                })
            },
            /*clickImg: function (id) {
                // console.log(id)
                FindBySubject(id, this.contentType).then(res => {
                    console.log(res.data)
                })
            }*/
        }
    }
</script>

<style scoped>
    .buttons {
        margin-top: 1px;
    }

    .upload-demo {
        float: right;
    }
</style>