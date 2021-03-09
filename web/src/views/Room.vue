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
        <v-virtual-scroll height="500" item-height="20" :items="messages">
          <template v-slot:default="{ message }">
            <!-- <v-list-item>
              <v-list-item-avatar
                class="align-self-start mr-2"
                v-if="message.me == false"
              >
                <v-avatar size="40">
                  <v-img src="@/assets/img/anon.png"></v-img>
                </v-avatar>
              </v-list-item-avatar>

              <v-list-item-content :class="{'received-message': message.me == false, 'sent-message': message.me == true}">
                <v-card dark elevation="15">
                  <v-card-title> {{ message.name }} </v-card-title>
                  <v-card-text class="pa-2 d-flex flex-column">
                    {{ message.text }}
                    <span class="text-caption font-italic align-self-end"
                      >19.12</span
                    >
                  </v-card-text>
                </v-card>
              </v-list-item-content>
              <v-list-item-avatar
                class="align-self-start mr-2"
                v-if="message.me == true"
              >
                <v-avatar size="40">
                  <v-img src="@/assets/img/anon.png"></v-img>
                </v-avatar>
              </v-list-item-avatar>
            </v-list-item> -->
          </template>
        </v-virtual-scroll>

        <v-divider></v-divider>
        <v-container fluid>
          <v-row>
            <v-col>
              <div class="d-flex flex-row align-center">
                <v-textarea
                  dark
                  name="input-7-1"
                  filled
                  label="Сообщение"
                  auto-grow
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
        name: "askdasd",
        text: "sadasdasdasdasdasdasdasdasd",
        me: true,
      },
      {
        name: "1askdasd",
        text: "1sadasdasdasdasdasdasdasdasd",
        me: false,
      },
      {
        name: "2askdasd",
        text: "2sadasdasdasdasdasdasdasdasd",
        me: false,
      },
      {
        name: "askdasd",
        text: "sadasdasdasdasdasdasdasdasd",
        me: true,
      },
    ],
  }),
  components: {},
};
</script>

<style scoped>
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
</style>