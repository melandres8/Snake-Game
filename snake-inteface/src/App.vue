<template>
  <v-app style="background-color: black;">
      <GetUsers/>
  </v-app>
</template>

<script>
import GetUsers from './components/GetUsers';
export default {
  name: 'App',

  components: {
    GetUsers,
  },

  data: () => ({
    //
  }),
};
</script>
