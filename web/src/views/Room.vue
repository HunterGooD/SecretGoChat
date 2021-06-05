<template>
  <v-container fluid>
    <v-card dark>
      <v-toolbar flat dark>
        <v-toolbar-title>Name room</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon dark>
          <v-icon>mdi-logout</v-icon>
        </v-btn>
      </v-toolbar>

      <v-card-text class="grey darken-3">
        <v-container fluid class="overflow-y size-chat">
          <v-list-item v-for="item in items" :key="item.id">
            <v-list-item-avatar
              class="align-self-start mr-2"
              v-if="item.me == false"
            >
              <v-avatar size="40">
                <v-img src="@/assets/img/anon.png"></v-img>
              </v-avatar>
            </v-list-item-avatar>

            <v-list-item-content
              :class="{
                'received-message': item.me == false,
                'sent-message': item.me == true,
              }"
            >
              <v-card dark elevation="15">
                <v-card-title> {{ item.name }} </v-card-title>
                <v-card-text class="pa-2 d-flex flex-column">
                  {{ item.text }}
                  <span class="text-caption font-italic align-self-end"
                    >19.12</span
                  >
                </v-card-text>
              </v-card>
            </v-list-item-content>
            <v-list-item-avatar
              class="align-self-start mr-2"
              v-if="item.me == true"
            >
              <v-avatar size="40">
                <v-img src="@/assets/img/anon.png"></v-img>
              </v-avatar>
            </v-list-item-avatar>
          </v-list-item>
        </v-container>

        <v-divider></v-divider>
        <v-container fluid>
          <v-row>
            <v-col>
              <div class="d-flex flex-row align-center">
                <v-textarea
                  dark
                  name="input-7-1"
                  label="Сообщение"
                  rows="1"
                  @keypress.enter="() => ({})"
                ></v-textarea>
                <v-btn icon class="ml-4" @click="() => ({})"
                  ><v-icon>mdi-send</v-icon></v-btn
                >
              </div>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: "Room",
  data: () => ({
    messages: [
      {
        id: 1,
        name: "askdasd",
        text:
          "sadasdasdasdasdasdasdasdasdasdddd\ndddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
        me: true,
      },
      {
        id: 2,
        name: "1askdasd",
        text: "1sadasdasdasdasdasdasdasdasd",
        me: false,
      },
      {
        id: 3,
        name: "2askdasd",
        text: "2sadasdasdasdasdasdasdasdasd",
        me: false,
      },
      {
        id: 4,
        name: "askdasd",
        text: "sadasdasdasdasdasdasdasdasd",
        me: true,
      },
      {
        id: 5,
        name: "askdasd",
        text: "sadasdasdasdasdasdasdasdasd",
        me: true,
      },
      {
        id: 6,
        name: "askdasd",
        text: "sadasdasdasdasdasdasdasdasd",
        me: true,
      },
    ],
  }),
  computed: {
    items() {
      return this.messages;
    },
  },
  components: {},
};
</script>

<style scoped>
.overflow-y {
  overflow-y: auto;
}

.size-chat {
  height: 70vh;
}

.received-message::after {
  content: " ";
  position: absolute;
  width: 0;
  height: 0;
  left: 54px;
  right: auto;
  top: 12px;
  bottom: auto;
  border: 12px solid;
  border-color: #1e1e1e transparent transparent transparent;
}

.sent-message::after {
  content: " ";
  position: absolute;
  width: 0;
  height: 0;
  left: auto;
  right: 68px;
  top: 12px;
  bottom: auto;
  border: 12px solid;
  border-color: #1e1e1e transparent transparent transparent;
}

@media screen and (max-width: 730px) {
  .size-chat {
    height: 63vh;
  }
}
</style>