<template>
    <div>
        <el-row>
            <el-card shadow="never" style="border-radius: 0px" >
                <el-button type="primary" size="small" @click="drawer = true">选择商品</el-button>
            </el-card>
<!--抽屉-->
            <el-drawer
                    title="我是标题"
                    :visible.sync="drawer"
                    :with-header="false">
                <span>我来啦!</span>
            </el-drawer>
        </el-row>
        <el-row style="padding: 10px">
            <el-col :span="6" v-for="(o, index) in 5" :key="o" style="margin: 10px">
                <el-card :body-style="{ padding: '0px' }">
                    <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png" >
                    <div style="padding: 14px;">
                        <span>好吃的汉堡</span>
                        <div class="bottom clearfix">
                            <time class="time">{{ currentDate }}</time>
                            <el-button type="text" class="button">删除</el-button>
                        </div>
                    </div>
                </el-card>
            </el-col>
        </el-row>
    </div>

</template>

<script>
    export default {
        name: "HotProduct.vue",
        data() {
            return {
                drawer: false,
            };
        }
    }
</script>

<style scoped>


</style>