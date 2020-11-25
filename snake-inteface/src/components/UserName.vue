<template>
  <v-row class="flex-grow-0 flex-shrink-0">
    <v-btn
      class="ml-5"
      color="primary"
      elevation="10"
      width="95"
      @click.stop="dialog = true"
      >
        PLAY
    </v-btn>

    <v-dialog
      v-model="dialog"
      max-width="400"
    >
      <v-card>
        <v-card-title class="headline">
          ENTER A USERNAME
        </v-card-title>

        <v-text-field
          placeholder="Username"
          v-model="username"
          filled
          rounded
          dense
          required
          :rules="[() => !!username.length || 'Required field']"
        ></v-text-field>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            class="play-btn"
            color="primary"
            elevation="9"
            @click="dialog = false"
            v-on:click="postUser"
          >
            SUBMIT
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      user: {
        username: {
          default: "",
          required: true
        },
      },
      dialog: false,
    }
  },
  methods: {
    postUser: async function () {
      await axios.post('http://localhost:8000/user', {
        username: this.username,
        score: "0"
      })
      .then(function (response) {
        console.log(response)
      })
      .catch(function (error) {
        console.log(error)
      })
      if (this.username) {
        window.location.href = 'http://localhost:3000';
      }
    }
  }
}
</script>
