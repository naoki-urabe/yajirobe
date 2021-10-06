module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    port: process.env.VUE_APP_FRONT_PORT,
    disableHostCheck: true,
  }
}
