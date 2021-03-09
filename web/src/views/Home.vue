<template>
  <v-layout justify-center align-center>
    <v-card
      dark
      elevation="24"
      class="align-center"
      max-width="400"
      width="100%"
    >
      <v-card-title> Создание комнаты </v-card-title>
      <v-card-text>
        <v-text-field
          dark
          label="Наименование комнаты"
          outlined
          clearable
          :loading="isLoad"
        ></v-text-field>
      </v-card-text>

      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn text dark outlined @click="click">
          Создать
        </v-btn></v-card-actions
      >
    </v-card>
  </v-layout>
</template>

<script>
export default {
  name: "Home",
  data: () => ({
    isLoad: false,
  }),
  components: {},
  methods: {
    click() {
      console.log();
    },
  },
};
</script>

<style scoped>
.layout {
  height: 60vh;
}
</style>