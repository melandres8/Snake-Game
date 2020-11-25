<template>
<v-container>
  <v-row
  class="pt-9"
  justify="center">
    <v-img
    src="../assets/snake.png"
    max-width="300"
    height="300"
    >
    </v-img>
  </v-row>
  <v-container class="mt-9">
    <v-row justify="center">
      <v-btn
      class="mr-5"
      color="primary"
      elevation="10"
      @click="getUsers"
      v-on:click="dialog = true"
      >
        SCORES
      </v-btn>

      <v-dialog
      v-model="dialog"
      max-width="400"
      >
        <v-card>
          <v-card-title primary-title class="justify-center">
            <h1><b>SCORES</b></h1>
          </v-card-title>
          <v-container>
            <v-row>
              <v-col align="center">
                <h2>Name:</h2>
              </v-col>
              <v-col align="center">
                <h2>Score:</h2>
              </v-col>
            </v-row>
            <v-row v-for="user in users" :key="user.score">
              <v-col align="center" class="pa-1">
                <h3>{{ user.username }}</h3>
              </v-col>
              <v-col align="center" class="pa-1">
                <h3>{{ user.score }}</h3>
              </v-col>
            </v-row>
          </v-container>
        </v-card>


        <v-card-actions>
          <v-btn text
          color="primary darken-1"
          @click="dialog = false">
            CLOSE
          </v-btn>
        </v-card-actions>
      </v-dialog>

      <UserName/>
    </v-row>
  </v-container>
</v-container>
</template>
<script>
import axios from "axios";
import UserName from "../components/UserName"
export default {
  name: "get-users",
  components: {
    UserName,
  },
  data() {
    return {
      dialog: false,
      users: []
    };
  },
  methods: {
    getUsers: async function () {
      await axios.get('http://localhost:8000/users')
      .then(res => {
        this.users = res.data;
        return this.users
      })
      .catch(err => {
        console.log(err);
      });
    }
  }
}
</script>
<style scoped>
</style>
