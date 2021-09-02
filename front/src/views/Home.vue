<template>
  <v-app id="app">
    <v-container>
      <v-row>
        <v-col cols="3">
          <v-text-field label="摘要" v-model="summary" />
        </v-col>
        <v-col cols="3">
          <v-text-field label="収支" v-model="income" />
        </v-col>
        <v-col cols="3">
          <v-select v-model="category" :items="categories" item-text="category_name" item-value="category_code"></v-select>
        </v-col>
        <v-col cols="3">
          <v-btn @click="registerIncome">submit</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>
<script>
import axios from "axios"
export default {
  name: "Home",
  data() {
    return {
      summary: "",
      income: null,
      category: "",
      categories: [],
    };
  },
  methods: {
    registerIncome: async function() {
      const bodyParameter = {
        dt: new Date(),
        summary: this.summary,
        income: parseInt(this.income),
        tag: this.category
      }
      const response = await axios.post(
        "http://localhost:8080/api/income/add",
        bodyParameter
      );
      console.log(response);
      this.summary = "";
      this.income = null;
      this.category = "";
    },
  },
  mounted: async function() {
    const response = await axios.get("http://localhost:8080/api/category/all")
    this.categories = response.data;
    console.log(this.categories);
  }
};
</script>
