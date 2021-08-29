<template>
  <div>
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
  </div>
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
      categories: [{category_code:"burn",category_name:"浪費"}, {category_code:"food",category_name:"食費"}, {category_code:"car",category_name:"車"}],
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
};
</script>
