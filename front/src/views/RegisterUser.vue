<template>
  <v-app id="app">
  <center>
    <v-alert v-if="isSuccess==='success'" type="success">登録に成功しました</v-alert>
    <v-alert v-if="isSuccess==='error'" type="error">登録に失敗しました</v-alert>
    <h1>ユーザ登録</h1>
    <v-container>
      <v-col cols="5">
        <v-text-field v-model="id" label="ID"></v-text-field>
        <v-text-field v-model="password" label="PASSWORD"></v-text-field>
        <v-btn @click="register">submit</v-btn>
      </v-col>
      <v-btn href="/">戻る</v-btn>
    </v-container>
  </center>
  </v-app>
</template>
<script>
import axios from "axios";
export default {
  auth: "guest",
  data() {
    return {
      id: "",
      password: "",
      isSuccess: "",
    };
  },
  methods: {
    register: async function () {
        try {
      await axios({
        method: "post",
        url: "/user/register",
        data: {
          id: this.id,
          pw: this.password,
        },
      });
      this.isSuccess="success";
      } catch(err) {
        this.isSuccess="error";
        console.log(err);
      }
    },
  },
};
</script>
