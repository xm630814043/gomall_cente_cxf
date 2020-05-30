<template>
    <!--轮播-->
    <div class="goodsShop">

            <el-card class="box-card" >
            <el-form ref="form" :model="addform" label-width="100px"
                     style="margin-left: 10%;margin-top: 7%">
                <el-form-item label="标题" >
                    <el-input v-model="addform.title" style="width:100%"></el-input>
                </el-form-item>
                <el-form-item label="上传图片"><!--https://jsonplaceholder.typicode.com/posts/-->
                    <el-upload
                            class="avatar-uploader"
                            action="http://upload.qiniu.com/"
                            :data="qn"
                            :show-file-list="false"
                            :on-success="handleAvatarSuccess"
                            :before-upload="beforeAvatarUpload">
                        <img v-if="addform.image_url" :src="addform.image_url" class="avatar">
                        <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                    </el-upload>
                </el-form-item>
                <el-form-item label="轮播地址">
                    <el-input v-model="addform.contents"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="onSubmit">立即创建</el-button>
                    <el-button @click="fanHui">取消</el-button>
                </el-form-item>
            </el-form>
            </el-card>
    </div>
</template>
<script>
    import {AddArticle, AddSubject, FindByUnSubjectList} from '../../api/company'

    export default {
        name: "create",
        data() {
            return {
                qn: {
                    token: "生成凭证中生成的token",
                    key: ""
                },
                restaurants: [],
                state: '',
                timeout: null,

                radio: 3,
                addform: {
                    title: '',
                    contents: '',
                    image_url: '',
                },
                activeIndex: '1',
                form: {
                    subject_id: 0,
                    object_id: 0,
                    contentType: '',
                    input: '',
                },
                shopinput: '',
                brandinput: '',
                datalist: [],
                contentType: '',
                id: 0,
            };
        },
        mounted() {
            /*this.restaurants = this.loadAll();*/

            this.id = this.$route.query.cid
            this.contentType = this.$route.query.sname
            this.limit = this.$route.query.slimit
            FindByUnSubjectList(this.id, this.contentType, this.form.input, this.shopinput, this.brandinput, this.radio).then(res => {
                this.datalist = res.data
                if (res.data.length == 0) {
                    /* alert("暂时没有数据")*/
                }
            })
        },
        methods: {
            /*loadAll() {
                return [
                    {"value": "三全鲜食（北新泾店）", "address": "长宁区新渔路144号"},
                    {"value": "Hot honey 首尔炸鸡（仙霞路）", "address": "上海市长宁区淞虹路661号"},
                    {"value": "新旺角茶餐厅", "address": "上海市普陀区真北路988号创邑金沙谷6号楼113"},
                ]
            },*/
            querySearchAsync(queryString, cb) {
                var restaurants = this.restaurants;
                var results = queryString ? restaurants.filter(this.createStateFilter(queryString)) : restaurants;

                clearTimeout(this.timeout);
                this.timeout = setTimeout(() => {
                    cb(results);
                }, 3000 * Math.random());
            },
            createStateFilter(queryString) {
                return (state) => {
                    return (state.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0);
                };
            },
            handleSelect(item) {
                console.log(item);
            },
            shopTiShi(value) {
                /* FindNameByLike(value).then(res=>{
                 })*/
            },
            findListByName() {
                FindByUnSubjectList(this.id, this.contentType, this.form.input, this.shopinput, this.brandinput, this.radio).then(res => {
                    this.datalist = res.data
                })
            },
            submitForm(object_id) {
                this.id = this.$route.query.cid
                this.form.subject_id = this.id
                this.form.object_id = object_id
                AddSubject(this.form).then(res => {
                    this.fanHui()
                })
            },
            fanHui() {
                if (this.id == 1) {
                    this.$router.push({
                        path: '/cms/reproduces',
                        query: {sid: this.id, sname: this.contentType, slimit: this.limit}
                    })
                }
                if (this.id == 2) {
                    this.$router.push({
                        path: '/cms/restore',
                        query: {sid: this.id, sname: this.contentType, slimit: this.limit}
                    })
                }
                if (this.id == 3) {

                    this.$router.push({
                        path: '/cms/reproduces',
                        query: {sid: this.id, sname: this.contentType, slimit: this.limit}
                    })
                }
                if (this.id == 4) {
                    this.$router.push({
                        path: '/cms/stores',
                        query: {sid: this.id, sname: this.contentType, slimit: this.limit}
                    })
                }
            },
            onSubmit() {
                AddArticle(this.addform)
            },
            handleAvatarSuccess(res, file) {
                this.addform.image_url = "http://qiniu.zt.snpit.cn/" + res.key;
                console.log(this.addform.image_url)
                console.log(file, res)
            },
            beforeAvatarUpload(file) {
                this.qn.token = genUpToken(AK, SK, policy)
                this.qn.key = file.name; // 通过设置key为文件名，上传到七牛中会显示对应的图片名
                const isJPG = file.type === "image/jpeg" || file.type === "image/png";
                const isLt2M = file.size / 1024 / 1024 < 2;

                if (!isJPG) {
                    this.$message.error('上传头像只能是图片格式!');
                }
                if (!isLt2M) {
                    this.$message.error('上传头像图片大小不能超过 2MB!');
                }
                return isJPG && isLt2M;
            },
        },
    }
</script>
<style>
    .avatar-uploader .el-upload {
        border: 1px dashed #d9d9d9;
        border-radius: 6px;
        cursor: pointer;
        position: relative;
        overflow: hidden;
    }

    .avatar-uploader .el-upload:hover {
        border-color: #409EFF;
    }

    .avatar-uploader-icon {
        font-size: 28px;
        color: #8c939d;
        width: 178px;
        height: 178px;
        line-height: 178px;
        text-align: center;
    }

    .avatar {
        width: 178px;
        height: 178px;
        display: block;
    }
    .box-card {
        width: 480px;
        margin: 50px;
    }
</style>