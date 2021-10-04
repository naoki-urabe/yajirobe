<template>
  <v-app id="app">
  <p>{{ username }}さんの家計簿</p>
    <v-container>
      <Chart />
      <v-row>
        <v-col cols="2">
          <v-radio-group
            v-model="radios"
            row
          >
          <v-radio
            label="支出"
            value="payment"
          ></v-radio>
          <v-radio
            label="収入"
            value="income"
          ></v-radio>
        </v-radio-group>
      </v-col>
        <v-col cols="3">
          <v-text-field label="摘要" v-model="summary" />
        </v-col>
        <v-col cols="3">
          <v-text-field label="収支" v-model="income" />
        </v-col>
        <v-col cols="2">
          <v-select
            v-model="category"
            :items="categories"
            item-text="category_name"
            item-value="category_code"
            label="カテゴリ"
          ></v-select>
        </v-col>
        <v-col cols="2">
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
import dayjs from "dayjs"
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
      username: "",
      radios: "payment",
    };
  },
  methods: {
    registerIncome: async function() {
      const sign = this.radios === "payment" ? -1 : 1;
      const bodyParameter = {
        dt: new Date(),
        summary: this.summary,
        income: sign * parseInt(this.income),
        tag: this.category,
        user: this.username
      };
      await axios.post("/income/add", bodyParameter);
      const latestIncome = await this.getLatestIncome();
      this.updateIncomes(latestIncome);
      this.summary = "";
      this.income = null;
      this.category = "";
    },
    updateIncomes: function(latestIncome) {
      this.incomes.splice(this.incomes.length, 0, {
        dt: dayjs(latestIncome.dt).format("YYYY-MM-DD"),
        summary: latestIncome.summary,
        income: latestIncome.income,
        tag: latestIncome.tag,
      });
    },
    getAllCategories: async function() {
      const response = await axios.get(
        "/category/all"
      );
      return response.data;
    },
    getIncomes: async function() {
      const bodyParameter = {
        user: this.username
      };
      const response = (await axios.post("/income/all",bodyParameter)).data;
      for(let i=0;i<response.length;i++){
        this.incomes.splice(this.incomes.length,0, 
        {
          "dt": dayjs(response[i].dt).format("YYYY-MM-DD"),
          "summary": response[i].summary,
          "income": response[i].income,
          "tag": response[i].tag,
        })
      }
      return response.data;
    },
    getLatestIncome: async function() {
      const bodyParameter = {
        user: this.username
      };
      const response = await axios.post(
        "/income/latest",bodyParameter
      );
      return response.data;
    },
  },
  mounted: async function() {
    this.username = this.$auth.remember();
    this.categories = await this.getAllCategories();
    this.getIncomes()
  },
};
</script>
