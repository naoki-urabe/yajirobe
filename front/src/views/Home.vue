<template>
  <v-app id="app">
    <v-container>
      <Chart />
      <v-row>
        <v-col cols="3">
          <v-text-field label="摘要" v-model="summary" />
        </v-col>
        <v-col cols="3">
          <v-text-field label="収支" v-model="income" />
        </v-col>
        <v-col cols="3">
          <v-select
            v-model="category"
            :items="categories"
            item-text="category_name"
            item-value="category_code"
          ></v-select>
        </v-col>
        <v-col cols="3">
          <v-btn @click="registerIncome">submit</v-btn>
        </v-col>
      </v-row>
    </v-container>
    <v-card>
      <v-data-table :headers="headers" :items="incomes" :items-per-page="20">
      </v-data-table>
    </v-card>
  </v-app>
</template>
<script>
import axios from "axios";
import Chart from "@/components/Chart";
export default {
  components: {
    Chart,
  },
  name: "Home",
  data() {
    return {
      summary: "",
      income: null,
      category: "",
      categories: [],
      incomes: [],
      headers: [
        { text: "取引日", value: "dt" },
        { text: "名目", value: "summary" },
        { text: "取引額", value: "income" },
        { text: "カテゴリ", value: "tag" },
      ],
    };
  },
  methods: {
    registerIncome: async function() {
      const bodyParameter = {
        dt: new Date(),
        summary: this.summary,
        income: parseInt(this.income),
        tag: this.category,
      };
      await axios.post("http://localhost:8080/api/income/add", bodyParameter);
      const latestIncome = await this.getLatestIncome();
      this.updateIncomes(latestIncome);
      this.summary = "";
      this.income = null;
      this.category = "";
    },
    updateIncomes: async function(latestIncome) {
      this.incomes.splice(this.incomes.length, 0, {
        dt: latestIncome.dt,
        summary: latestIncome.summary,
        income: latestIncome.income,
        tag: latestIncome.tag,
      });
    },
    getAllCategories: async function() {
      const response = await axios.get(
        "http://localhost:8080/api/category/all"
      );
      return response.data;
    },
    getIncomes: async function() {
      const response = await axios.get("http://localhost:8080/api/income/all");
      return response.data;
    },
    getLatestIncome: async function() {
      const response = await axios.get(
        "http://localhost:8080/api/income/latest"
      );
      return response.data;
    },
  },
  mounted: async function() {
    this.categories = await this.getAllCategories();
    this.incomes = await this.getIncomes();
  },
};
</script>
