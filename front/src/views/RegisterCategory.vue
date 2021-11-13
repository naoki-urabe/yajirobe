<template>
  <v-app id="app">
    <v-container>
        <v-row>
        <v-col cols="4">
        <v-text-field v-model="categoryName" label="カテゴリ名"></v-text-field>
        </v-col>
        <v-col cols="4">
        <v-text-field
          v-model="categoryCode"
          label="カテゴリコード"
        ></v-text-field>
        </v-col>
        <v-col cols="4">
        <v-btn @click="register">submit</v-btn>
        </v-col>
      </v-row>
    </v-container>
    <v-card>
        <v-data-table :headers="headers" :items="categories" :items-per-page="20">
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
            編集しますか？
          </v-card-title>
          <v-card-actions class="justify-center">
            <v-container>
              <v-row>
                <v-col cols="6">
                  <v-text-field
                    v-model="editCategoryCode"
                    label="カテゴリコード"
                  />
                </v-col>
                <v-col cols="6">
                  <v-text-field
                    v-model="editCategoryName"
                    label="カテゴリ名"
                  />
                </v-col>
              </v-row>
              <v-btn @click="editConfirm">
                はい
              </v-btn>
              <v-btn @click="dialogEdit=false">
                いいえ
              </v-btn>
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
            削除しますか？
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
import axios from "axios"
export default{
    name: "RegisterCategory",
    data() {
        return {
            headers: [
                {text: "カテゴリ名", value: "category_name", align: "center"},
                {text: "カテゴリコード", value: "category_code"},
                {text: "編集/削除", value: "actions"}
            ],
            categoryName: "",
            categoryCode: "",
            editCategoryName: "",
            editCategoryCode: "",
            categories: [],
            dialogEdit: false,
            dialogDelete: false,
        };
    },
    methods: {
        editItem: async function(item) {
          this.dialogEdit=true;
          this.editId=item.category_code;
          this.editCategoryCode = item.category_code;
          this.editCategoryName = item.category_name;
        },
        editConfirm: async function() {
          const bodyParameter = {category_code: this.editCategoryCode, category_name: this.editCategoryName};
          await axios.post(`/category/edit/${this.editId}`,bodyParameter);
          this.dialogEdit=false;
          await this.getAllCategories();
        },
        deleteItem: async function(item) {
          this.dialogDelete=true;
          this.deleteId=item.category_code;
          console.log(this.deleteId);
        },
        deleteConfirm: async function() {
          await axios.post(`/category/delete/${this.deleteId}`);
          await this.getAllCategories();
          this.dialogDelete=false;
        },
        register: async function() {
            const bodyParameters = {
                category_name: this.categoryName,
                category_code: this.categoryCode
            }
            await axios.post(
                "/category/add",
                bodyParameters
            )
            this.categoryName = "";
            this.categoryCode = "";
            await this.getAllCategories();
        },
        getAllCategories: async function() {
          this.categories = [];
          const response = await axios.get(
            "/category/all"
          );
          this.categories = response.data;
          if(this.categories == null) {
            this.categories = [];
            return;
          }
          this.categories.splice(0,0);
        },
    },
    mounted: async function() {
        await this.getAllCategories();
    }
}
</script>
