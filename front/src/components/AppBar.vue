<template>
  <v-app-bar color="deep-purple accent-4" dark>
    <v-toolbar-title>YAJIROBE</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn @click="logout">ログアウト</v-btn>
  </v-app-bar>
</template>
<script>
export default {
  name: 'AppBar',
  methods: {
    async logout() {
      this.$auth.logout({redirect: "/login"})
    }
  }
};
</script>
