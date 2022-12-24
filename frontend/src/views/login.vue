<template>
  <div class="center_div">
    <div class="login_card">
      <div class="login_txt center_div">登录</div>
      <div class="vertical_div">
        <el-input v-model="user" placeholder="账号" class="input_div" />
        <el-input
          v-model="password"
          type="password"
          placeholder="密码"
          show-password
          class="input_div"
        />
        <el-button
          type="primary"
          class="input_div"
          :loading="btnLoading"
          @click="tryLogin"
          >确定</el-button
        >
      </div>
    </div>
  </div>
</template>

<script>
import { login } from "../api/login";
import storage from "@/utils/storage"

export default {
  name: "login",
  data() {
    return {
      btnLoading: false,
      user: "",
      password: "",
    };
  },
  methods: {
    tryLogin() {
      this.btnLoading = true;
      login(this.user, this.password)
        .then((response) => {
          console.log(response);
          storage.set("token",response.data.token);
          storage.set("user",this.user);
          this.btnLoading = false;
          this.$router.push("/index/store");
        })
        .catch((error) => {
          this.btnLoading = false;
        });
    },
  },
};
</script>

<style scoped>
@import "../css/common.css";

.login_card {
  width: 500px;
  height: 300px;
  background-color: white;
  margin-top: 200px;
  border-radius: 10px;
  padding: 8px;
}

.login_txt {
  margin-top: 20px;
  font-size: 24px;
}

.input_div {
  height: 40px;
  width: 300px;
  margin-top: 30px;
}
</style>