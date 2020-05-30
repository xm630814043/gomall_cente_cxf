<template>
    <div style="height: 100%;width: 100%;background-color: #fefff9;overflow: auto">
        <div style="margin-left: 9%;height: 100%;width: 750px;">
            <div v-if="status==1 ? true : false"
                 style="margin-right: 5%;border: 0px solid black">
                <div style="float: right;margin-right: 20%;margin-top: 1%">
                    <el-button v-model="status1" label="4" @click="noPassStatus" size="small">驳回</el-button>
                </div>
                <div style="float: right;margin-top: 1%">
                    <el-button v-model="status1" label="2" @click="passStatus" size="small">通过</el-button>
                </div>
            </div>
            <div v-else></div>
            <br/>
            <div style="margin-top: 10%;border: 0px solid black;height: auto">
                <div style="height:40px "><span><b>商品名称:</b></span> <span>{{ productdata.product_name}}</span></div>
                <div style="height:210px "><b>商品图片:</b> <br/> <img :src="productdata.pic"
                                                                   style="width:130px; height:175px; margin-left:11%"/>
                </div>
                <div style="height:85px;width: 100%;">
                    <div style="float: left;width: 10%;"><b>商品描述:</b></div>
                    <span style="width:85%;height: auto;border: 0px solid black">{{ productdata.description }}</span></div>
                <div style="height:40px "><b>产品链接:</b><a
                        v-bind:href="['http://47.93.162.54/goodsDetail.html?goodsid='+productdata.id]"
                        target="view_window">{{ "http://47.93.162.54/goodsDetail.html?goodsid="+productdata.id}}</a>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
    import {GetProductAuditInfo, UpdateProductStatus} from "../../api/company";

    export default {
        name: "detail",
        data() {
            return {
                status: 1,
                status1: 1,
                productdata: [],
            };
        },
        mounted() {
            this.id = this.$route.query.id
            this.dstatus = this.$route.query.status
            GetProductAuditInfo(this.id).then(res => {
                this.productdata = res.data
            })
        },
        watch: {
            '$route'(to, from) {
                this.id = this.$route.query.id
                this.dstatus = this.$route.query.status
                GetProductAuditInfo(this.id).then(res => {
                    this.productdata = res.data
                })
            }
        },
        methods: {
            noPassStatus() {

                UpdateProductStatus(this.productdata.id, 4).then(res => {
                    this.$router.push({
                        path: '/goods',
                    })
                })
            },
            passStatus() {
                UpdateProductStatus(this.productdata.id, 2).then(res => {
                    this.$router.push({
                        path: '/goods',
                    })
                })
            },
        }
    }

</script>

<style scoped>
    .to-do-list1 {
        background-color: #fefff9
    }
</style>