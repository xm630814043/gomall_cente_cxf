<template>
    <div style="background-color: #ffffff;width: 100%;height: 100%;overflow: auto">
    <div class="checklist_bigDiv">
        <el-card>
            <div class="lists_top1" v-if="form.user_type==0">
                <span class="listsspan">采购商-信息</span>
            </div>
            <div class="lists_top1" v-else>
                <span class="listsspan">供应商-信息</span>
            </div>
        </el-card>


        <div class="lists_top2" v-model="form">

                <div class="lists_top2_left">
                    <el-card style="width: 99%;height: 99%">
                    <div style="height: 10%;text-align: center">头像</div>
                        <hr style="width: 100%">
                    <img :src="form.avatar"
                         style="width:100%; height:90%; float: left;"/>
                    </el-card>
                </div>
            <div class="lists_top2_right">

                <el-card style="width: 99%;height: 99%">
                    <div style="height: 10%;text-align: center">基础信息</div>
                    <hr style="width: 100%">
                    <div class="lists_top2_right_row">
                        <span class="checklist_top2_left_row_title">
                            用户名称:
                        </span><span>{{form.username}}</span>
                    </div>
                    <div class="lists_top2_right_row">
                        <span class="checklist_top2_left_row_title">
                            性别:
                        </span><span>{{form.sex}}</span>
                    </div>
                    <div class="lists_top2_right_row">
                        <span class="checklist_top2_left_row_title">
                            手机号码:
                        </span><span>{{form.mobile}}</span>
                    </div>
                    <div class="lists_top2_right_row">
                        <span class="checklist_top2_left_row_title">
                            用户类型:
                        </span>
                        <span v-if="form.user_type==0">采购</span>
                        <span v-else-if="form.user_type==1">供应</span>
                    </div>
                    <div class="lists_top2_right_row">
                        <span class="checklist_top2_left_row_title">
                            email:
                        </span><span>{{form.email}}</span>
                    </div>
                </el-card>

            </div>

        </div>


    </div>
    </div>
</template>
<script>
    import {FindByAccount} from '../../api/company'

    export default {
        name: 'lists',
        data() {
            return {
                form: [],
            }
        },
        mounted() {
            this.id = this.$route.query.cid
            FindByAccount(this.id).then(res => {
                console.log(res.data)
                this.form = res.data
            })
        },
        watch: {
            '$route'(to, from) {
                this.id = this.$route.query.cid
                FindByAccount(this.id).then(res => {
                    console.log(res.data)
                    this.form = res.data
                })
            }
        }
    }
</script>
<style>

    .listsblock {
        margin-top: 10px;
    }

    .listsspan {
        margin-left: 10px;
        font-size: 28px;
    }

    .buttons {
        margin-top: 10px;
        margin-left: 40px;
        background-color: #ffffff;
    }

    .right {
        margin-left: 100px;
    }
</style>