<template>
  <v-app>
    <v-toolbar app :dark="unavailable" :color="unavailable ? 'red':'default'">
      <v-toolbar-title class="headline text-uppercase">
        <span>dbcheck</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn flat medium fab @click="refresh()">
        <v-icon>refresh</v-icon>
      </v-btn>
    </v-toolbar>
    <v-content>
      <router-view/>
    </v-content>
    <v-snackbar
      :value="unavailable"
      color="error"
      :timeout="2000"
    >
      Server is not available
      <v-btn flat color="default" @click.native="value = false">Close</v-btn>
    </v-snackbar>
  </v-app>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'App',
  methods: {
    refresh(){
      window.location.reload()
    }
  },
  computed: {
    ...mapGetters({
      redis: 'redis/disabled',
      postgres: 'postgres/disabled',
      mongo: 'mongo/disabled',
      mysql: 'mysql/disabled',
    }),
    unavailable(){
      return this.redis && this.postgres && this.mongo && this.mysql
    }
  }
}
</script>