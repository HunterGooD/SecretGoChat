<template>
  <v-app>
    <v-app-bar
      app
      dark
      v-if="
        $router.currentRoute.name == 'Home' ||
        $router.currentRoute.name == 'About'
      "
    >
      <v-img
        class="mx-2"
        src="@/assets/img/11-512.webp"
        max-height="40"
        max-width="40"
        contain
        style="filter: invert(100%)"
      ></v-img>
      <v-toolbar-title><router-link to="/">SecretChat</router-link> </v-toolbar-title>
    </v-app-bar>
    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>
    <v-footer app dark padless>
      <v-card class="flex" flat tile>
        <v-card-title
          class=""
          v-if="
            $router.currentRoute.name == 'Home' ||
            $router.currentRoute.name == 'About'
          "
        >
          <i class="subheading">По всем проблемам обращайтесь в соц сети.</i>
          <i class="subheading">Прочитать о нас <b><router-link to="/about" id="about">тут</router-link></b></i>
          <v-spacer></v-spacer>

          <v-btn v-for="icon in icons" :key="icon.icon" class="mx-4" dark icon>
            <v-icon size="24px" tag="a" :href="icon.link" target="_blank">
              {{ icon.icon }}
            </v-icon>
          </v-btn>
        </v-card-title>

        <v-divider></v-divider>

        <v-card-text class="py-2 white--text text-center">
          {{ new Date().getFullYear() }} — <strong>Vuetify</strong>
        </v-card-text>
      </v-card>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: "App",
  components: {},
  data: () => ({
    icons: [
      { icon: "mdi-vk", link: "https://vk.com/1proger1" },
      { icon: "mdi-linkedin", link: "/" },
      { icon: "mdi-github", link: "https://github.com/HunterGooD" },
    ],
  }),
};
</script>

<style>
a {
  text-decoration: none;
  color: #fff !important;
}
.v-main {
  background: #333336;
}
#about {
  color: red !important;
}
</style>