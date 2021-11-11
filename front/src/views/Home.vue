<template>
  <v-app id="app">
  <p>{{ username }}さんの家計簿</p>
    <v-container>
      <v-row>
      <Chart />
      </v-row>
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
      <v-col cols="2">
        <v-menu
          v-model="menu"
          :close-on-content-click="false"
          :nudge-right="40"
          transition="scale-transition"
          offset-y
          min-width="auto"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
              v-model="date"
              label="取引日"
              prepend-icon="mdi-calendar"
              readonly
              v-bind="attrs"
              v-on="on"
            ></v-text-field>
          </template>
          <v-date-picker
            v-model="date"
            @input="menu = false"
          ></v-date-picker>
        </v-menu>
      </v-col>
        <v-col cols="2">
          <v-text-field label="摘要" v-model="summary" />
        </v-col>
        <v-col cols="2">
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
      <v-row>
        <v-select
            v-model="month"
            :items="months"
            label="対象月"
            @change="getIncomes"
          ></v-select>
      </v-row>
      <v-row>
        <v-col cols="4">
        <h2>収入:{{ totalIncome }}</h2>
        </v-col>
        <v-col cols="4">
        <h2>支出:{{ totalPayment }}</h2>
        </v-col>
        <v-col cols="4">
        <h2>収支:{{ balanceOfPayments }}</h2>
        </v-col>
      </v-row>
    </v-container>
    <v-card>
      <v-data-table :headers="headers" :items="incomes" :items-per-page="20">
        <template #[`item.actions`]="{ item }">
          <v-icon
            small
            @click="editItem(item)"
          >
            mdi-pencil
          </v-icon>
          <v-icon
            small
            @click="deleteItem(item)"
          >
            mdi-delete
          </v-icon>
        </template>
      </v-data-table>
      <v-dialog
        v-model="dialogEdit"
        max-width="500px"
      >
        <v-card>
          <v-card-title class="justify-center">
            編集しますか?
          </v-card-title>
          <v-card-actions class="justify-center">
            <v-container>
            <v-row>
                <v-col cols="6">
                <v-text-field v-model="editIncome.summary" label="摘要"/>
                </v-col>
                <v-col cols="6">
                <v-text-field v-model="editIncome.income" label="取引額"/>
                </v-col>
            </v-row>
            <v-row>
                <v-col cols="6">
                <v-text-field v-model="editTime" label="取引日"/>
                </v-col>
             <v-col cols="6">
                <v-select
                  v-model="editIncome.tag"
                  :items="categories"
                  item-text="category_name"
                  item-value="category_code"
                  label="タグ名"
                  no-data-text="タグを設定してください"
                  :menu-props="{
                    closeOnClick: true,
                    closeOnContentClick: true,
                  }"
                  :error="isError"
                  :error-messages="errMsg"
                />
              </v-col>
            </v-row>
            <v-row>
                <v-col>
                  <v-btn @click="editConfirm">
                    OK
                  </v-btn>
                  <v-btn @click="dialogEdit=false">
                    キャンセル
                  </v-btn>
                </v-col>
              </v-row>
            </v-container>
          </v-card-actions>
        </v-card>
      </v-dialog>
    <v-dialog
        v-model="dialogDelete"
        max-width="500px"
      >
        <v-card>
          <v-card-title class="justify-center">
            削除しますか?
          </v-card-title>
          <v-card-actions class="justify-center">
            <v-btn @click="deleteConfirm">
              はい
            </v-btn>
            <v-btn @click="dialogDelete=false">
              いいえ
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
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
      dialogEdit: false,
      dialogDelete: false,
      menu: false,
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
        { text: "編集/削除", value: "actions"}
      ],
      editIncome: {summary: "", income: 0},
      editDate: null,
      username: "",
      radios: "payment",
      date: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
      month: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 7),
      totalPayment: 0,
      totalIncome: 0,
      balanceOfPayments: 0
    };
  },
  methods: {
    editItem: async function(item) {
      this.dialogEdit=true;
      this.editId=item.id;
      this.editIncome=item;
      this.editTime=dayjs(this.editIncome.dt).format("YYYY-MM-DD");
      console.log(item);
    },
    deleteItem: async function(item) {
      this.dialogDelete=true;
      this.deleteId=item.id;
    },
    editConfirm: async function() {
        const bodyParameters = {dt: dayjs(this.editTime), summary: this.editIncome.summary, income: parseInt(this.editIncome.income), tag: this.editIncome.tag};
        await axios.post(`/income/edit/${this.editId}`, bodyParameters);
        this.dialogEdit=false;
        this.getIncomes();
    },
    deleteConfirm: async function() {
      await axios.get(`/income/delete/${this.deleteId}`);
      this.dialogDelete=false;
      this.getIncomes();
    },
    registerIncome: async function() {
      const sign = this.radios === "payment" ? -1 : 1;
      const bodyParameter = {
        dt: new Date(this.date),
        summary: this.summary,
        income: sign * parseInt(this.income),
        tag: this.category,
        user: this.username
      };
      await axios.post("/income/add", bodyParameter);
      const latestIncome = await this.getLatestIncome();
      this.updateIncomes(latestIncome);
      this.months = await this.getAllMonths();
      this.months.sort().reverse()
      this.summary = "";
      this.income = null;
      this.category = "";
    },
    updateIncomes: function(latestIncome) {
      if(latestIncome.month === this.month){
        if(latestIncome.income > 0){
          this.totalIncome += latestIncome.income;
        } else {
          this.totalPayment += latestIncome.income;
        }
        this.incomes.splice(this.incomes.length, 0, {
          dt: dayjs(latestIncome.dt).format("YYYY-MM-DD"),
          summary: latestIncome.summary,
          income: latestIncome.income,
          tag: latestIncome.tag,
        });
        this.balanceOfPayments = this.totalIncome + this.totalPayment;
      }
    },
    getAllCategories: async function() {
      const response = await axios.get(
        "/category/all"
      );
      return response.data;
    },
    getIncomes: async function() {
      const bodyParameter = {
        user: this.username,
        month: this.month
      };
      this.incomes = []
      const response = (await axios.post("/income/monthly-incomes",bodyParameter)).data;
      let tmpIncome = 0, tmpPayment = 0;      
      for(let i=0;i<response.length;i++){
        if(response[i].income > 0){
          tmpIncome += response[i].income;
        } else {
          tmpPayment += response[i].income;
        }
        this.incomes.splice(this.incomes.length,0, 
        {
          "id": response[i].id,
          "dt": dayjs(response[i].dt).format("YYYY-MM-DD"),
          "summary": response[i].summary,
          "income": response[i].income,
          "tag": response[i].tag,
        })
      }
      this.totalIncome = tmpIncome;
      this.totalPayment = tmpPayment;
      this.balanceOfPayments = tmpIncome + tmpPayment;
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
    getAllMonths: async function() {
      this.months = []
      const response = await axios.post(
        "/income/months"
      );
      return response.data;
    }
  },
  mounted: async function() {
    this.username = this.$auth.remember();
    this.categories = await this.getAllCategories();
    this.months = await this.getAllMonths();
    this.months.sort().reverse();
    this.getIncomes();
  },
};
</script>
