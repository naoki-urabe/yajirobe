<template>
  <center>
    <h1>カテゴリ登録</h1>
    <v-container>
      <v-col cols="5">
        <v-text-field v-model="categoryName" label="カテゴリ名"></v-text-field>
        <v-text-field
          v-model="categoryCode"
          label="カテゴリコード"
        ></v-text-field>
        <v-btn @click="register">submit</v-btn>
      </v-col>
    </v-container>
  </center>
</template>
<script>
import axios from "axios"
export default{
    name: "RegisterCategory",
    data() {
        return {
            categoryName: "",
            categoryCode: "",
        };
    },
    methods: {
        register: async function() {
            const bodyParameters = {
                category_name: this.categoryName,
                category_code: this.categoryCode
            }
            const response = await axios.post(
                "http://localhost:8080/api/category/add",
                bodyParameters
            )
            console.log(response);
            this.categoryName = "";
            this.categoryCode = "";
        }
    }
}
</script>
